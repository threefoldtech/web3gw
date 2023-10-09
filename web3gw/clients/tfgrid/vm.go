package tfgrid

import (
	"context"
	"fmt"
	"math/rand"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/threefoldtech/zos/pkg/gridtypes/zos"
)

const gwNameEnvVar = "WEB3PROXY_DOMAIN_NAME"
const letters = "abcdefghijklmnopqrstuvwxyz"

var vmCapacity = map[string]capacityPackage{
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

type DeployVM struct {
	VMConfiguration    `json:"VMConfiguration"`
	AddWireguardAccess bool `json:"add_wireguard_access"`
}

type VMDeployment struct {
	VMConfiguration `json:"VMConfiguration"`

	Network         string `json:"network"`
	WireguardConfig string `json:"wireguard_config"`
	GatewayName     string `json:"gateway_name"`
}

type GatewayedMachines struct {
	Machine VMConfiguration  `json:"machine"`
	Gateway GatewayNameModel `json:"gateway"`
}

type RemoveVM struct {
	Network string `json:"network"`
	VMName  string `json:"vm_name"`
}

func (c *Client) DeployVM(ctx context.Context, args DeployVM) (VMDeployment, error) {
	// TODO generate network name
	// TODO return error if vm already exists!
	_, err := c.GetNetworkDeployment(ctx, args.Name)
	if err != nil {
		if strings.Contains(err.Error(), "found 0 contracts for model") {
			// this is a new network
			return c.createNetworkAndAddVM(ctx, args)
		}
		log.Error().Msgf("error: %+v", err)
		return VMDeployment{}, err
	}

	return VMDeployment{}, fmt.Errorf("there already is a vm with the name %s", args.Name)
}

func (c *Client) createNetworkAndAddVM(ctx context.Context, args DeployVM) (VMDeployment, error) {
	networkDeployment := NetworkDeployment{
		Name: args.Name,
		Network: NetworkConfiguration{
			AddWireguardAccess: args.AddWireguardAccess,
			IPRange:            "10.1.0.0/16",
		},
		VMs: []VMConfiguration{args.VMConfiguration},
		// todo check other arguments
	}

	networkDeployment, err := c.DeployNetwork(ctx, networkDeployment)
	if err != nil {
		return VMDeployment{}, err
	}
	vm := networkDeployment.VMs[0]
	gws := map[string]GatewayNameModel{}
	gwName, ok := vm.EnvVars[gwNameEnvVar]
	if ok {
		gw := GatewayNameModel{
			Name:     gwName,
			Backends: []zos.Backend{zos.Backend(fmt.Sprintf("http://[%s]:9000", vm.YggIP))},
		}

		gw, err := c.GatewayNameDeploy(ctx, gw)
		if err != nil {
			return VMDeployment{}, err
		}

		gws[vm.Name] = gw
	}

	return VMDeployment{
		Network:         networkDeployment.Name,
		WireguardConfig: networkDeployment.Network.WireguardConfig,
		VMConfiguration: networkDeployment.VMs[0],
		GatewayName:     gws[vm.Name].Name,
	}, nil
}

func (c *Client) GetVMDeployment(ctx context.Context, name string) (VMDeployment, error) {
	networkDeployment, err := c.GetNetworkDeployment(ctx, name)
	if err != nil {
		return VMDeployment{}, err
	}

	res := VMDeployment{
		Network:         name,
		WireguardConfig: networkDeployment.Network.WireguardConfig,
	}

	for _, vm := range networkDeployment.VMs {
		if vm.Name == name {
			res.VMConfiguration = vm
			gwName, ok := vm.EnvVars[gwNameEnvVar]
			if !ok {
				continue
			}

			gw, err := c.GatewayNameGet(ctx, gwName)
			if err != nil {
				return VMDeployment{}, err
			}
			res.GatewayName = gw.Name
		}
	}

	return res, nil
}

func (c *Client) CancelVMDeployment(ctx context.Context, name string) error {
	networkDeployment, err := c.GetNetworkDeployment(ctx, name)
	if err != nil {
		return err
	}

	for _, m := range networkDeployment.VMs {
		gwName, ok := m.EnvVars[gwNameEnvVar]
		if !ok {
			continue
		}

		if err := c.cancelModel(ctx, gwName); err != nil {
			return err
		}
	}

	if err := c.cancelModel(ctx, name); err != nil {
		return err
	}

	return nil
}

func generateRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
