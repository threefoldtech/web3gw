package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/pkg/errors"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/deployer"
	"github.com/threefoldtech/tfgrid-sdk-go/grid-client/workloads"
	"github.com/threefoldtech/zos/pkg/gridtypes"
	"github.com/urfave/cli"
)

func deployNetwork() cli.ActionFunc {
	return func(ctx *cli.Context) error {

		mnemonics := ctx.String("mnemonics")

		if mnemonics == "" {
			return fmt.Errorf("must provide mnemonics")
		}
		env := ctx.String("env")

		t, err := deployer.NewTFPluginClient(mnemonics, "sr25519", env, "", "", "", 100, false)
		if err != nil {
			return err
		}

		networkName := ctx.String("network_name")
		vm_name := ctx.String("vm_name")
		entrypoint := ctx.String("entryPoint")
		flist := ctx.String("flist")
		sshKey := ctx.String("sshKey")
		node := uint32(ctx.Uint("node"))
		rootfs := ctx.Int("rootfs")
		cpu := ctx.Int("cpu")
		memory := ctx.Int("memory")

		vm := workloads.VM{
			Name:    vm_name,
			EnvVars: map[string]string{"SSH_KEY": string(sshKey)},
			CPU:     cpu,
			Memory:  memory * 1024,

			RootfsSize: rootfs * 1024,
			Flist:      flist,
			Entrypoint: entrypoint,
		}

		network := buildNetwork(networkName, vm.Name, []uint32{node})

		mounts := []workloads.Disk{}

		vm.NetworkName = networkName
		dl := workloads.NewDeployment(vm.Name, node, vm.Name, nil, networkName, mounts, nil, []workloads.VM{vm}, nil)

		c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err = t.NetworkDeployer.Deploy(c, &network)
		if err != nil {
			return errors.Wrapf(err, "failed to deploy network on node %d", node)
		}
		err = t.DeploymentDeployer.Deploy(c, &dl)
		if err != nil {
			return errors.Wrapf(err, "failed to deploy vm on node %d", node)
		}
		resVM, err := t.State.LoadVMFromGrid(node, vm.Name, dl.Name)
		if err != nil {
			return errors.Wrapf(err, "failed to load vm from node %d", node)
		}
		jsonRes, err := json.Marshal(resVM)
		if err != nil {
			return errors.Wrapf(err, "failed to load vm from node %d", node)
		}
		fmt.Println(string(jsonRes))
		return nil
	}
}
func buildNetwork(name, projectName string, nodes []uint32) workloads.ZNet {
	return workloads.ZNet{
		Name:  name,
		Nodes: nodes,
		IPRange: gridtypes.NewIPNet(net.IPNet{
			IP:   net.IPv4(10, 20, 0, 0),
			Mask: net.CIDRMask(16, 32),
		}),
		SolutionType: projectName,
	}
}
