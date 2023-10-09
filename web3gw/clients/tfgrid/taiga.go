package tfgrid

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-proxy/pkg/types"
	"github.com/threefoldtech/zos/pkg/gridtypes/zos"
)

var taigaCapacity = map[string]capacityPackage{
	"small": {
		cru: 1,
		mru: 2048,
		sru: 4096,
	},
	"medium": {
		cru: 2,
		mru: 4096,
		sru: 8192,
	},
	"large": {
		cru: 4,
		mru: 8192,
		sru: 16384,
	},
	"extra-large": {
		cru: 8,
		mru: 16384,
		sru: 32768,
	},
}

type Taiga struct {
	Name          string `json:"name"`
	FarmID        uint64 `json:"farm_id"`
	Capacity      string `json:"capacity"`
	SSHKey        string `json:"ssh_key"`
	DiskSize      uint32 `json:"disk_size"`
	AdminEmail    string `json:"admin_email"`
	AdminUsername string `json:"admin_username"`
	AdminPassword string `json:"admin_password"`
	PublicIPv6    bool   `json:"public_ipv6"`
}

type TaigaResult struct {
	Name         string `json:"name"`
	MachineYGGIP string `json:"ygg_ip"`
	MachineIPv6  string `json:"ipv6"`
	FQDN         string `json:"fqdn"`
}

func (c *Client) DeployTaiga(ctx context.Context, taiga Taiga) (TaigaResult, error) {
	if err := c.validateProjectName(ctx, taiga.Name); err != nil {
		return TaigaResult{}, err
	}

	gwNode, err := c.findTaigaGWNode(uint32(taiga.FarmID))
	if err != nil {
		return TaigaResult{}, err
	}

	machinesModel, err := taiga.generateMachinesModel(gwNode)
	if err != nil {
		return TaigaResult{}, err
	}

	machinesModel, err = c.DeployNetwork(ctx, machinesModel)
	if err != nil {
		return TaigaResult{}, err
	}

	yggIP := machinesModel.VMs[0].YggIP
	ipv6 := machinesModel.VMs[0].ComputedIP6

	gwModel := taiga.generateGWModel(gwNode, yggIP)
	gw, err := c.GatewayNameDeploy(ctx, gwModel)
	if err != nil {
		return TaigaResult{}, err
	}

	return TaigaResult{
		Name:         taiga.Name,
		MachineYGGIP: yggIP,
		MachineIPv6:  ipv6,
		FQDN:         gw.FQDN,
	}, nil
}

func (t *Taiga) generateMachinesModel(gwNode types.Node) (NetworkDeployment, error) {
	cap, ok := taigaCapacity[t.Capacity]
	if !ok {
		return NetworkDeployment{}, fmt.Errorf("capacity %s is invalid", t.Capacity)
	}

	model := NetworkDeployment{
		Name: generateTaigaModelName(t.Name),
		Network: NetworkConfiguration{
			IPRange: "10.1.0.0/16",
		},
		VMs: []VMConfiguration{
			{
				Name:       fmt.Sprintf("%sVM", t.Name),
				Flist:      "https://hub.grid.tf/tf-official-apps/grid3_taiga_docker-latest.flist",
				CPU:        int(cap.cru),
				Memory:     int(cap.mru),
				RootfsSize: int(cap.sru),
				EnvVars: map[string]string{
					"SSH_KEY":             t.SSHKey,
					"DOMAIN_NAME":         fmt.Sprintf("%s.%s", t.Name, gwNode.PublicConfig.Domain),
					"ADMIN_USERNAME":      t.AdminUsername,
					"ADMIN_PASSWORD":      t.AdminPassword,
					"ADMIN_EMAIL":         t.AdminEmail,
					"DEFAULT_FROM_EMAIL":  "",
					"EMAIL_USE_TLS":       "",
					"EMAIL_USE_SSL":       "",
					"EMAIL_HOST":          "",
					"EMAIL_PORT":          "",
					"EMAIL_HOST_USER":     "",
					"EMAIL_HOST_PASSWORD": "",
				},
				Entrypoint: "/sbin/zinit init",
				Planetary:  true,
				FarmID:     uint32(t.FarmID),
			},
		},
	}

	return model, nil
}

func (d *Taiga) generateGWModel(gwNode types.Node, yggIP string) GatewayNameModel {
	gw := GatewayNameModel{
		NodeID:   uint32(gwNode.NodeID),
		Name:     d.Name,
		Backends: []zos.Backend{zos.Backend(fmt.Sprintf("http://[%s]:9000", yggIP))},
	}

	return gw
}

func (c *Client) GetTaiga(ctx context.Context, name string) (TaigaResult, error) {
	machinesModel, err := c.GetNetworkDeployment(ctx, generateTaigaModelName(name))
	if err != nil {
		return TaigaResult{}, err
	}

	gw, err := c.GatewayNameGet(ctx, name)
	if err != nil {
		return TaigaResult{}, err
	}

	yggIP := machinesModel.VMs[0].YggIP
	ipv6 := machinesModel.VMs[0].ComputedIP6

	return TaigaResult{
		Name:         name,
		MachineYGGIP: yggIP,
		MachineIPv6:  ipv6,
		FQDN:         gw.FQDN,
	}, nil
}

func (c *Client) DeleteTaiga(ctx context.Context, name string) error {
	if err := c.cancelModel(ctx, generateTaigaModelName(name)); err != nil {
		return err
	}

	if err := c.cancelModel(ctx, name); err != nil {
		return err
	}

	return nil
}

func (c *Client) findTaigaGWNode(farmID uint32) (types.Node, error) {
	filter := BuildGridProxyNodeFilters(NodeFilterOptions{
		FarmID:       farmID,
		PublicConfig: true,
	}, uint64(c.TwinID))

	res, _, err := c.GridClient.FilterNodes(filter, types.Limit{Size: 1})
	if err != nil {
		return types.Node{}, err
	}

	if len(res) == 0 {
		return types.Node{}, errors.New("failed to find an elibile gateway for the Taiga instance")
	}

	return res[0], nil
}

func generateTaigaModelName(name string) string {
	return fmt.Sprintf("%sTaiga", name)
}
