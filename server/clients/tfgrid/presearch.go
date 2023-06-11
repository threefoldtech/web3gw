package tfgrid

import (
	"context"
	"fmt"
)

type Presearch struct {
	Name     string `json:"name"`
	FarmID   uint64 `json:"farm_id"`
	SSHKey   string `json:"ssh_key"`
	DiskSize uint32 `json:"disk_size"`
	PublicIP bool   `json:"public_ipv4"`
}

type PresearchResult struct {
	Name         string `json:"name"`
	MachineYGGIP string `json:"machine_ygg_ip"`
	MachineIPV4  string `json:"machine_ipv4"`
}

func (c *Client) DeployPresearch(ctx context.Context, presearch Presearch) (PresearchResult, error) {
	if err := c.validateProjectName(ctx, presearch.Name); err != nil {
		return PresearchResult{}, err
	}

	machinesModel := presearch.generateMachinesModel()

	machinesModel, err := c.MachinesDeploy(ctx, machinesModel)
	if err != nil {
		return PresearchResult{}, err
	}

	yggIP := machinesModel.Machines[0].YggIP
	publicIP := machinesModel.Machines[0].ComputedIP4

	return PresearchResult{
		Name:         presearch.Name,
		MachineYGGIP: yggIP,
		MachineIPV4:  publicIP,
	}, nil
}

func (p *Presearch) generateMachinesModel() MachinesModel {
	model := MachinesModel{
		Name: p.Name,
		Network: Network{
			IPRange: "10.1.0.0/16",
		},
		Machines: []Machine{
			{
				Name:   fmt.Sprintf("%sVM", p.Name),
				Flist:  "https://hub.grid.tf/tf-official-apps/presearch-v2.2.flist",
				CPU:    1,
				Memory: 512,
				EnvVars: map[string]string{
					"SSH_KEY":                     p.SSHKey,
					"PRESEARCH_REGISTRATION_CODE": "",
				},
				Entrypoint: "/sbin/zinit init",
				Planetary:  true,
				FarmID:     uint32(p.FarmID),
				PublicIP:   p.PublicIP,
			},
		},
	}

	return model
}

func (c *Client) GetPresearch(ctx context.Context, name string) (PresearchResult, error) {
	machinesModel, err := c.MachinesGet(ctx, name)
	if err != nil {
		return PresearchResult{}, err
	}

	yggIP := machinesModel.Machines[0].YggIP
	publicIP := machinesModel.Machines[0].ComputedIP4

	return PresearchResult{
		Name:         name,
		MachineYGGIP: yggIP,
		MachineIPV4:  publicIP,
	}, nil
}

func (c *Client) DeletePresearch(ctx context.Context, name string) error {
	if err := c.cancelModel(ctx, name); err != nil {
		return err
	}

	return nil
}
