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

type VM struct {
	Name               string `json:"name"`
	FarmID             uint64 `json:"farm_id"`
	Network            string `json:"network"`
	Capacity           string `json:"capacity"`
	Times              uint32 `json:"times"`
	DiskSize           uint32 `json:"disk_size"`
	SSHKey             string `json:"ssh_key"`
	Gateway            bool   `json:"gateway"`
	AddWireguardAccess bool   `json:"add_wireguard_access"`
	AddPublicIPv4      bool   `json:"add_public_ipv4"`
	AddPublicIPv6      bool   `json:"add_public_ipv6"`
}

type VMResult struct {
	Network         string              `json:"network"`
	WireguardConfig string              `json:"wireguard_config"`
	VMs             []GatewayedMachines `json:"vms"`
}

type GatewayedMachines struct {
	Machine Machine          `json:"machine"`
	Gateway GatewayNameModel `json:"gateway"`
}

type RemoveVM struct {
	Network string `json:"network"`
	VMName  string `json:"vm_name"`
}

func (c *Client) DeployVM(ctx context.Context, vm VM) (VMResult, error) {
	_, err := c.MachinesGet(ctx, vm.Network)
	if err != nil {
		log.Error().Msgf("error: %+v", err)
		if strings.Contains(err.Error(), "found 0 contracts for model") {
			// this is a new network
			return c.deployVM(ctx, vm)
		}

		return VMResult{}, err
	}

	return c.addVM(ctx, vm)
}

func (c *Client) deployVM(ctx context.Context, vm VM) (VMResult, error) {
	machinesModel := MachinesModel{
		Name: vm.Network,
		Network: Network{
			AddWireguardAccess: vm.AddWireguardAccess,
			IPRange:            "10.1.0.0/16",
		},
		Machines: []Machine{},
	}

	machines, err := vm.generateMachines()
	if err != nil {
		return VMResult{}, err
	}

	machinesModel.Machines = machines

	machinesModel, err = c.MachinesDeploy(ctx, machinesModel)
	if err != nil {
		return VMResult{}, err
	}

	gws := map[string]GatewayNameModel{}
	for _, m := range machinesModel.Machines {
		gwName, ok := m.EnvVars[gwNameEnvVar]
		if !ok {
			continue
		}

		gw := GatewayNameModel{
			Name:     gwName,
			Backends: []zos.Backend{zos.Backend(fmt.Sprintf("http://[%s]:9000", m.YggIP))},
		}

		gw, err := c.GatewayNameDeploy(ctx, gw)
		if err != nil {
			return VMResult{}, err
		}

		gws[m.Name] = gw
	}

	return newVMResult(machinesModel, gws), nil
}

func (c *Client) addVM(ctx context.Context, vm VM) (VMResult, error) {
	gws := map[string]GatewayNameModel{}

	machines, err := vm.generateMachines()
	if err != nil {
		return VMResult{}, err
	}

	machinesModel := MachinesModel{}
	for _, m := range machines {
		res, err := c.MachineAdd(ctx, AddMachineParams{
			ModelName: vm.Network,
			Machine:   m,
		})
		if err != nil {
			return VMResult{}, err
		}

		machinesModel = res

		gwName, ok := m.EnvVars[gwNameEnvVar]
		if !ok {
			continue
		}

		gws[m.Name] = GatewayNameModel{
			Name: gwName,
		}
	}

	for _, m := range machinesModel.Machines {
		gw, ok := gws[m.Name]
		if !ok {
			continue
		}

		gw.Backends = []zos.Backend{zos.Backend(fmt.Sprintf("http://[%s]:9000", m.YggIP))}

		gw, err := c.GatewayNameDeploy(ctx, gw)
		if err != nil {
			return VMResult{}, err
		}

		gws[m.Name] = gw
	}

	return newVMResult(machinesModel, gws), nil
}

func (c *Client) GetVM(ctx context.Context, networkName string) (VMResult, error) {
	gws := map[string]GatewayNameModel{}

	machinesModel, err := c.MachinesGet(ctx, networkName)
	if err != nil {
		return VMResult{}, err
	}

	for _, m := range machinesModel.Machines {
		gwName, ok := m.EnvVars[gwNameEnvVar]
		if !ok {
			continue
		}

		gw, err := c.GatewayNameGet(ctx, gwName)
		if err != nil {
			return VMResult{}, err
		}

		gws[m.Name] = gw
	}

	res := VMResult{
		Network:         networkName,
		WireguardConfig: machinesModel.Network.WireguardConfig,
		VMs:             []GatewayedMachines{},
	}

	for _, m := range machinesModel.Machines {
		res.VMs = append(res.VMs, GatewayedMachines{
			Machine: m,
			Gateway: gws[m.Name],
		})
	}

	return res, nil
}

func (c *Client) DeleteVM(ctx context.Context, networkName string) error {
	machinesModel, err := c.MachinesGet(ctx, networkName)
	if err != nil {
		return err
	}

	for _, m := range machinesModel.Machines {
		gwName, ok := m.EnvVars[gwNameEnvVar]
		if !ok {
			continue
		}

		if err := c.cancelModel(ctx, gwName); err != nil {
			return err
		}
	}

	if err := c.cancelModel(ctx, networkName); err != nil {
		return err
	}

	return nil
}

func (c *Client) RemoveVM(ctx context.Context, args RemoveVM) (VMResult, error) {
	machinesModel, err := c.MachinesGet(ctx, args.Network)
	if err != nil {
		return VMResult{}, err
	}

	for _, m := range machinesModel.Machines {
		if m.Name == args.VMName {
			gwName, ok := m.EnvVars[gwNameEnvVar]
			if ok {
				if err := c.cancelModel(ctx, gwName); err != nil {
					return VMResult{}, err
				}
			}

			if _, err := c.MachineRemove(ctx, RemoveMachineParams{
				ModelName:   args.Network,
				MachineName: args.VMName,
			}); err != nil {
				return VMResult{}, err
			}

			break
		}
	}

	return c.GetVM(ctx, args.Network)
}

func (vm *VM) generateMachines() ([]Machine, error) {
	machines := []Machine{}

	vmName := "vm"
	if vm.Name != "" {
		vmName = vm.Name
	}

	cap, ok := vmCapacity[vm.Capacity]
	if !ok {
		return nil, fmt.Errorf("capacity %s is invalid", vm.Capacity)
	}

	for i := 0; i < int(vm.Times); i++ {
		name := vmName
		if vm.Times > 0 {
			name = fmt.Sprintf("%s%d", name, i)
		}

		m := Machine{
			Name:       name,
			FarmID:     uint32(vm.FarmID),
			Flist:      "https://hub.grid.tf/tf-official-apps/base:latest.flist",
			Planetary:  true,
			PublicIP:   vm.AddPublicIPv4,
			CPU:        int(cap.cru),
			Memory:     int(cap.mru),
			RootfsSize: int(cap.sru),
			Entrypoint: "/sbin/zint init",
			EnvVars: map[string]string{
				"SSH_KEY": vm.SSHKey,
			},
		}

		if vm.DiskSize > 0 {
			m.Disks = append(m.Disks, Disk{
				MountPoint: "/mnt/disk",
				SizeGB:     int(vm.DiskSize),
			})
		}

		if vm.Gateway {
			gwName := generateRandomString(8)
			m.EnvVars[gwNameEnvVar] = gwName
		}

		machines = append(machines, m)
	}

	return machines, nil
}

func newVMResult(model MachinesModel, gws map[string]GatewayNameModel) VMResult {
	res := VMResult{
		Network:         model.Name,
		WireguardConfig: model.Network.WireguardConfig,
		VMs:             []GatewayedMachines{},
	}

	for _, m := range model.Machines {
		res.VMs = append(res.VMs, GatewayedMachines{
			Machine: m,
			Gateway: gws[m.Name],
		})
	}

	return res
}

func generateRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
