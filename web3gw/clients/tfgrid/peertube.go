package tfgrid

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-proxy/pkg/types"
	"github.com/threefoldtech/zos/pkg/gridtypes/zos"
)

var peertubeCapacity = map[string]capacityPackage{
	"small": {
		cru: 1,
		mru: 2048,
		sru: 10240,
	},
	"medium": {
		cru: 2,
		mru: 2048,
		sru: 102400,
	},
	"large": {
		cru: 4,
		mru: 4096,
		sru: 256000,
	},
	"extra-large": {
		cru: 4,
		mru: 8192,
		sru: 409600,
	},
}

type Peertube struct {
	Name       string `json:"name"`
	FarmID     uint64 `json:"farm_id"`
	Capacity   string `json:"capacity"`
	SSHKey     string `json:"ssh_key"`
	DBUserName string `json:"db_username"`
	DBPassword string `json:"db_password"`
	AdminEmail string `json:"admin_email"`

	PublicIPv6 bool `json:"public_ipv6"`
}

type PeertubeResult struct {
	Name         string `json:"name"`
	MachineYGGIP string `json:"ygg_ip"`
	MachineIPv6  string `json:"ipv6"`
	FQDN         string `json:"fqdn"`
}

func (c *Client) DeployPeertube(ctx context.Context, peertube Peertube) (PeertubeResult, error) {
	if err := c.validateProjectName(ctx, peertube.Name); err != nil {
		return PeertubeResult{}, err
	}

	gwNode, err := c.findPeertubeGWNode(uint32(peertube.FarmID))
	if err != nil {
		return PeertubeResult{}, err
	}

	machinesModel, err := peertube.generateMachinesModel(gwNode)
	if err != nil {
		return PeertubeResult{}, err
	}

	machinesModel, err = c.DeployNetwork(ctx, machinesModel)
	if err != nil {
		return PeertubeResult{}, err
	}

	yggIP := machinesModel.VMs[0].YggIP
	ipv6 := machinesModel.VMs[0].ComputedIP6

	gwModel := peertube.generateGWModel(gwNode, yggIP)
	gw, err := c.GatewayNameDeploy(ctx, gwModel)
	if err != nil {
		return PeertubeResult{}, err
	}

	return PeertubeResult{
		Name:         peertube.Name,
		MachineYGGIP: yggIP,
		MachineIPv6:  ipv6,
		FQDN:         gw.FQDN,
	}, nil
}

func (p *Peertube) generateMachinesModel(gwNode types.Node) (NetworkDeployment, error) {
	cap, ok := peertubeCapacity[p.Capacity]
	if !ok {
		return NetworkDeployment{}, fmt.Errorf("capacity %s is invalid", p.Capacity)
	}

	model := NetworkDeployment{
		Name: generatePeertubeModelName(p.Name),
		Network: NetworkConfiguration{
			IPRange: "10.1.0.0/16",
		},
		VMs: []VMConfiguration{
			{
				Name:       fmt.Sprintf("%sVM", p.Name),
				Flist:      "https://hub.grid.tf/tf-official-apps/peertube-v3.1.1.flist",
				CPU:        int(cap.cru),
				Memory:     int(cap.mru),
				RootfsSize: int(cap.sru),
				EnvVars: map[string]string{
					"SSH_KEY":                     p.SSHKey,
					"PEERTUBE_WEBSERVER_HOSTNAME": fmt.Sprintf("%s.%s", p.Name, gwNode.PublicConfig.Domain),
					"PEERTUBE_DB_SUFFIX":          "_prod",
					"PEERTUBE_DB_USERNAME":        p.DBUserName,
					"PEERTUBE_DB_PASSWORD":        p.DBPassword,
					"PEERTUBE_ADMIN_EMAIL":        p.AdminEmail,
					"PEERTUBE_WEBSERVER_PORT":     "443",
					"PEERTUBE_BIND_ADDRESS":       "::",
				},
				Entrypoint: "/sbin/zinit init",
				Planetary:  true,
				FarmID:     uint32(p.FarmID),
			},
		},
	}

	return model, nil
}

func (d *Peertube) generateGWModel(gwNode types.Node, yggIP string) GatewayNameModel {
	gw := GatewayNameModel{
		NodeID:   uint32(gwNode.NodeID),
		Name:     d.Name,
		Backends: []zos.Backend{zos.Backend(fmt.Sprintf("http://[%s]:9000", yggIP))},
	}

	return gw
}

func (c *Client) GetPeertube(ctx context.Context, name string) (PeertubeResult, error) {
	machinesModel, err := c.GetNetworkDeployment(ctx, generatePeertubeModelName(name))
	if err != nil {
		return PeertubeResult{}, err
	}

	gw, err := c.GatewayNameGet(ctx, name)
	if err != nil {
		return PeertubeResult{}, err
	}

	yggIP := machinesModel.VMs[0].YggIP
	ipv6 := machinesModel.VMs[0].ComputedIP6

	return PeertubeResult{
		Name:         name,
		MachineYGGIP: yggIP,
		MachineIPv6:  ipv6,
		FQDN:         gw.FQDN,
	}, nil
}

func (c *Client) DeletePeertube(ctx context.Context, name string) error {
	if err := c.cancelModel(ctx, generatePeertubeModelName(name)); err != nil {
		return err
	}

	if err := c.cancelModel(ctx, name); err != nil {
		return err
	}

	return nil
}

func (c *Client) findPeertubeGWNode(farmID uint32) (types.Node, error) {
	filter := BuildGridProxyNodeFilters(NodeFilterOptions{
		FarmID:       farmID,
		PublicConfig: true,
	}, uint64(c.TwinID))

	res, _, err := c.GridClient.FilterNodes(filter, types.Limit{Size: 1})
	if err != nil {
		return types.Node{}, err
	}

	if len(res) == 0 {
		return types.Node{}, errors.New("failed to find an elibile gateway for the peertube instance")
	}

	return res[0], nil
}

func generatePeertubeModelName(name string) string {
	return fmt.Sprintf("%sPeertube", name)
}
