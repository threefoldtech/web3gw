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
	Name               string `json:"name"`
	FarmID             uint64 `json:"farm_id"`
	Capacity           string `json:"capacity"`
	DiskSize           uint32 `json:"disk_size"`
	SSHKey             string `json:"ssh_key"`
	DeveloperEmail     string `json:"developer_email"`
	SMTPUsername       string `json:"smtp_username"`
	SMTPPassword       string `json:"smtp_password"`
	SMTPAddress        string `json:"smtp_address"`
	SMTPEnableTLS      bool   `json:"smtp_enable_tls"`
	SMTPPort           uint32 `json:"smtp_port"`
	ThreebotPrivateKey string `json:"threebot_private_key"`
	FlaskSecretKey     string `json:"flask_secret_key"`
}

type DiscourseResult struct {
	Name         string `json:"name"`
	MachineYGGIP string `json:"machine_ygg_ip"`
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

	machinesModel, err = c.MachinesDeploy(ctx, machinesModel)
	if err != nil {
		return DiscourseResult{}, err
	}

	yggIP := machinesModel.Machines[0].YggIP

	gwModel := discourse.generateGWModel(gwNode, yggIP)
	gw, err := c.GatewayNameDeploy(ctx, gwModel)
	if err != nil {
		return DiscourseResult{}, err
	}

	return DiscourseResult{
		Name:         discourse.Name,
		MachineYGGIP: yggIP,
		FQDN:         gw.FQDN,
	}, nil
}

func (d *Discourse) generateMachinesModel(gwNode types.Node) (MachinesModel, error) {
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
		return MachinesModel{}, fmt.Errorf("capacity %s is invalid", d.Capacity)
	}

	model := MachinesModel{
		Name: generateDiscourseModelName(d.Name),
		Network: Network{
			IPRange: "10.1.0.0/16",
		},
		Machines: []Machine{
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
					"THREEBOT_PRIVATE_KEY":            d.ThreebotPrivateKey,
					"FLASK_SECRET_KEY":                d.FlaskSecretKey,
				},
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
	machinesModel, err := c.MachinesGet(ctx, generateDiscourseModelName(name))
	if err != nil {
		return DiscourseResult{}, err
	}

	gw, err := c.GatewayNameGet(ctx, name)
	if err != nil {
		return DiscourseResult{}, err
	}

	yggIP := machinesModel.Machines[0].YggIP

	return DiscourseResult{
		Name:         name,
		MachineYGGIP: yggIP,
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
