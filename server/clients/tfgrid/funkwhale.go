package tfgrid

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-proxy/pkg/types"
	"github.com/threefoldtech/zos/pkg/gridtypes/zos"
)

var funkwhaleCapacity = map[string]capacityPackage{
	"small": {
		cru: 1,
		mru: 1024,
		sru: 51200,
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

type Funkwhale struct {
	Name          string `json:"name"`
	FarmID        uint64 `json:"farm_id"`
	Capacity      string `json:"capacity"`
	SSHKey        string `json:"ssh_key"`
	AdminEmail    string `json:"admin_email"`
	AdminUsername string `json:"admin_username"`
	AdminPassword string `json:"admin_password"`
	PublicIPv6    bool   `json:"public_ipv6"`
}

type FunkwhaleResult struct {
	Name         string `json:"name"`
	MachineYGGIP string `json:"machine_ygg_ip"`
	MachineIPv6  string `json:"machine_ipv6"`
	FQDN         string `json:"fqdn"`
}

func (c *Client) Deployfunkwhale(ctx context.Context, funkwhale Funkwhale) (FunkwhaleResult, error) {
	if err := c.validateProjectName(ctx, funkwhale.Name); err != nil {
		return FunkwhaleResult{}, err
	}

	gwNode, err := c.findfunkwhaleGWNode(uint32(funkwhale.FarmID))
	if err != nil {
		return FunkwhaleResult{}, err
	}

	machinesModel, err := funkwhale.generateMachinesModel(gwNode)
	if err != nil {
		return FunkwhaleResult{}, err
	}

	machinesModel, err = c.MachinesDeploy(ctx, machinesModel)
	if err != nil {
		return FunkwhaleResult{}, err
	}

	yggIP := machinesModel.Machines[0].YggIP
	ipv6 := machinesModel.Machines[0].ComputedIP6

	gwModel := funkwhale.generateGWModel(gwNode, yggIP)
	gw, err := c.GatewayNameDeploy(ctx, gwModel)
	if err != nil {
		return FunkwhaleResult{}, err
	}

	return FunkwhaleResult{
		Name:         funkwhale.Name,
		MachineYGGIP: yggIP,
		MachineIPv6:  ipv6,
		FQDN:         gw.FQDN,
	}, nil
}

func (f *Funkwhale) generateMachinesModel(gwNode types.Node) (MachinesModel, error) {
	cap, ok := funkwhaleCapacity[f.Capacity]
	if !ok {
		return MachinesModel{}, fmt.Errorf("capacity %s is invalid", f.Capacity)
	}

	model := MachinesModel{
		Name: generatefunkwhaleModelName(f.Name),
		Network: Network{
			IPRange: "10.1.0.0/16",
		},
		Machines: []Machine{
			{
				Name:       fmt.Sprintf("%sVM", f.Name),
				Flist:      "https://hub.grid.tf/tf-official-apps/funkwhale-dec21.flist",
				CPU:        int(cap.cru),
				Memory:     int(cap.mru),
				RootfsSize: int(cap.sru),
				EnvVars: map[string]string{
					"SSH_KEY":                   f.SSHKey,
					"FUNKWHALE_HOSTNAME":        fmt.Sprintf("%s.%s", f.Name, gwNode.PublicConfig.Domain),
					"DJANGO_SUPERUSER_EMAIL":    f.AdminEmail,
					"DJANGO_SUPERUSER_USERNAME": f.AdminUsername,
					"DJANGO_SUPERUSER_PASSWORD": f.AdminPassword,
				},
				PublicIP6:  f.PublicIPv6,
				Entrypoint: "/init.sh",
				Planetary:  true,
				FarmID:     uint32(f.FarmID),
			},
		},
	}

	return model, nil
}

func (d *Funkwhale) generateGWModel(gwNode types.Node, yggIP string) GatewayNameModel {
	gw := GatewayNameModel{
		NodeID:   uint32(gwNode.NodeID),
		Name:     d.Name,
		Backends: []zos.Backend{zos.Backend(fmt.Sprintf("http://[%s]:9000", yggIP))},
	}

	return gw
}

func (c *Client) Getfunkwhale(ctx context.Context, name string) (FunkwhaleResult, error) {
	machinesModel, err := c.MachinesGet(ctx, generatefunkwhaleModelName(name))
	if err != nil {
		return FunkwhaleResult{}, err
	}

	gw, err := c.GatewayNameGet(ctx, name)
	if err != nil {
		return FunkwhaleResult{}, err
	}

	yggIP := machinesModel.Machines[0].YggIP
	ipv6 := machinesModel.Machines[0].ComputedIP6

	return FunkwhaleResult{
		Name:         name,
		MachineYGGIP: yggIP,
		MachineIPv6:  ipv6,
		FQDN:         gw.FQDN,
	}, nil
}

func (c *Client) Deletefunkwhale(ctx context.Context, name string) error {
	if err := c.cancelModel(ctx, generatefunkwhaleModelName(name)); err != nil {
		return err
	}

	if err := c.cancelModel(ctx, name); err != nil {
		return err
	}

	return nil
}

func (c *Client) findfunkwhaleGWNode(farmID uint32) (types.Node, error) {
	filter := BuildGridProxyFilters(FilterOptions{
		FarmID:       farmID,
		PublicConfig: true,
	}, uint64(c.TwinID))

	res, _, err := c.GridClient.FilterNodes(filter, types.Limit{Size: 1})
	if err != nil {
		return types.Node{}, err
	}

	if len(res) == 0 {
		return types.Node{}, errors.New("failed to find an elibile gateway for the funkwhale instance")
	}

	return res[0], nil
}

func generatefunkwhaleModelName(name string) string {
	return fmt.Sprintf("%sFunkwhale", name)
}
