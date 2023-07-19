package tfgrid

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-proxy/pkg/types"
	"github.com/threefoldtech/zos/pkg/gridtypes/zos"
)

var discourseCapacity = map[string]capacityPackage{
	"small": {
		cru: 1,
		mru: 2048,
		sru: 10240,
	},
	"medium": {
		cru: 2,
		mru: 2048,
		sru: 51200,
	},
	"large": {
		cru: 1,
		mru: 4096,
		sru: 102400,
	},
	"extra-large": {
		cru: 4,
		mru: 8192,
		sru: 153600,
	},
}

type Discourse struct {
	Name           string `json:"name"`
	FarmID         uint64 `json:"farm_id"`
	Capacity       string `json:"capacity"`
	DiskSize       uint32 `json:"disk_size"`
	SSHKey         string `json:"ssh_key"`
	DeveloperEmail string `json:"developer_email"`
	SMTPUsername   string `json:"smtp_username"`
	SMTPPassword   string `json:"smtp_password"`
	SMTPAddress    string `json:"smtp_address"`
	SMTPEnableTLS  bool   `json:"smtp_enable_tls"`
	SMTPPort       uint32 `json:"smtp_port"`
	PublicIPv6     bool   `json:"public_ipv6"`
}

type DiscourseResult struct {
	Name         string `json:"name"`
	MachineYGGIP string `json:"ygg_ip"`
	MachineIPv6  string `json:"ipv6"`
	FQDN         string `json:"fqdn"`
}

func (c *Client) DeployDiscourse(ctx context.Context, discourse Discourse) (DiscourseResult, error) {
	if err := c.validateProjectName(ctx, discourse.Name); err != nil {
		return DiscourseResult{}, err
	}

	gwNode, err := c.findDiscourseGWNode(uint32(discourse.FarmID))
	if err != nil {
		return DiscourseResult{}, err
	}

	machinesModel, err := discourse.generateMachinesModel(gwNode)
	if err != nil {
		return DiscourseResult{}, err
	}

	machinesModel, err = c.DeployNetwork(ctx, machinesModel)
	if err != nil {
		return DiscourseResult{}, err
	}

	yggIP := machinesModel.VMs[0].YggIP
	ipv6 := machinesModel.VMs[0].ComputedIP6

	gwModel := discourse.generateGWModel(gwNode, yggIP)
	gw, err := c.GatewayNameDeploy(ctx, gwModel)
	if err != nil {
		return DiscourseResult{}, err
	}

	return DiscourseResult{
		Name:         discourse.Name,
		MachineYGGIP: yggIP,
		MachineIPv6:  ipv6,
		FQDN:         gw.FQDN,
	}, nil
}

func (d *Discourse) generateMachinesModel(gwNode types.Node) (NetworkDeployment, error) {
	var disks []Disk
	if d.DiskSize > 0 {
		disks = []Disk{
			{
				MountPoint: "/var/lib/docker",
				SizeGB:     int(d.DiskSize),
				Name:       "disk1",
			},
		}
	}

	cap, ok := discourseCapacity[d.Capacity]
	if !ok {
		return NetworkDeployment{}, fmt.Errorf("capacity %s is invalid", d.Capacity)
	}

	model := NetworkDeployment{
		Name: generateDiscourseModelName(d.Name),
		Network: NetworkConfiguration{
			IPRange: "10.1.0.0/16",
		},
		VMs: []VMConfiguration{
			{
				Name:       fmt.Sprintf("%sVM", d.Name),
				Flist:      "https://hub.grid.tf/tf-official-apps/forum-docker-v3.1.2.flist",
				CPU:        int(cap.cru),
				Memory:     int(cap.mru),
				RootfsSize: int(cap.sru),
				EnvVars: map[string]string{
					"SSH_KEY":                         d.SSHKey,
					"DISCOURSE_HOSTNAME":              fmt.Sprintf("%s.%s", d.Name, gwNode.PublicConfig.Domain),
					"DISCOURSE_DEVELOPER_EMAILS":      d.DeveloperEmail,
					"DISCOURSE_SMTP_ADDRESS":          d.SMTPAddress,
					"DISCOURSE_SMTP_PORT":             fmt.Sprint(d.SMTPPort),
					"DISCOURSE_SMTP_ENABLE_START_TLS": fmt.Sprint(d.SMTPEnableTLS),
					"DISCOURSE_SMTP_USER_NAME":        d.SMTPUsername,
					"DISCOURSE_SMTP_PASSWORD":         d.SMTPPassword,
					"THREEBOT_PRIVATE_KEY":            generateRandomString(16),
					"FLASK_SECRET_KEY":                generateRandomString(16),
				},
				PublicIP6:  d.PublicIPv6,
				Entrypoint: "/sbin/zinit init",
				Planetary:  true,
				FarmID:     uint32(d.FarmID),
				Disks:      disks,
			},
		},
	}

	return model, nil
}

func (d *Discourse) generateGWModel(gwNode types.Node, yggIP string) GatewayNameModel {
	gw := GatewayNameModel{
		NodeID:   uint32(gwNode.NodeID),
		Name:     d.Name,
		Backends: []zos.Backend{zos.Backend(fmt.Sprintf("http://[%s]:9000", yggIP))},
	}

	return gw
}

func (c *Client) GetDiscourse(ctx context.Context, name string) (DiscourseResult, error) {
	machinesModel, err := c.GetNetworkDeployment(ctx, generateDiscourseModelName(name))
	if err != nil {
		return DiscourseResult{}, err
	}

	gw, err := c.GatewayNameGet(ctx, name)
	if err != nil {
		return DiscourseResult{}, err
	}

	yggIP := machinesModel.VMs[0].YggIP
	ipv6 := machinesModel.VMs[0].ComputedIP6

	return DiscourseResult{
		Name:         name,
		MachineYGGIP: yggIP,
		MachineIPv6:  ipv6,
		FQDN:         gw.FQDN,
	}, nil
}

func (c *Client) DeleteDiscourse(ctx context.Context, name string) error {
	if err := c.cancelModel(ctx, generateDiscourseModelName(name)); err != nil {
		return err
	}

	if err := c.cancelModel(ctx, name); err != nil {
		return err
	}

	return nil
}

func (c *Client) findDiscourseGWNode(farmID uint32) (types.Node, error) {
	filter := BuildGridProxyFilters(FilterOptions{
		FarmID:       farmID,
		PublicConfig: true,
	}, uint64(c.TwinID))

	res, _, err := c.GridClient.FilterNodes(filter, types.Limit{Size: 1})
	if err != nil {
		return types.Node{}, err
	}

	if len(res) == 0 {
		return types.Node{}, errors.New("failed to find an elibile gateway for the discourse instance")
	}

	return res[0], nil
}

func generateDiscourseModelName(name string) string {
	return fmt.Sprintf("%sDiscourse", name)
}
