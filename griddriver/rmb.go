package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	substrate "github.com/threefoldtech/tfchain/clients/tfchain-client-go"
	"github.com/threefoldtech/tfgrid-sdk-go/rmb-sdk-go/direct"
	"github.com/threefoldtech/zos/pkg"
	"github.com/threefoldtech/zos/pkg/gridtypes"
	"github.com/urfave/cli"
)

type rmbCmdArgs map[string]interface{}

func rmbDecorator(action func(c *cli.Context, client *direct.RpcCLient) (interface{}, error)) cli.ActionFunc {
	return func(c *cli.Context) error {
		substrate_url := c.String("substrate")
		mnemonics := c.String("mnemonics")
		relay_url := c.String("relay")

		subManager := substrate.NewManager(substrate_url)
		sub, err := subManager.Substrate()
		if err != nil {
			return fmt.Errorf("failed to connect to substrate: %w", err)
		}
		defer sub.Close()
		client, err := direct.NewRpcClient(context.Background(), direct.KeyTypeSr25519, mnemonics, relay_url, "tfgrid-vclient", sub, true)

		if err != nil {
			return fmt.Errorf("failed to create direct client: %w", err)
		}

		res, err := action(c, client)

		if err != nil {
			return err
		}
		fmt.Printf("%v", res)
		return nil

	}
}

func rmbCall(c *cli.Context, client *direct.RpcCLient) (interface{}, error) {
	dst := uint32(c.Uint("dst"))
	cmd := c.String("cmd")
	payload := c.String("payload")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var pl interface{}
	if err := json.Unmarshal([]byte(payload), &pl); err != nil {
		return nil, err
	}

	var res interface{}
	if err := client.Call(ctx, dst, cmd, pl, &res); err != nil {
		return nil, err
	}

	b, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	return string(b), nil
}

func deploymentChanges(c *cli.Context, client *direct.RpcCLient) (interface{}, error) {
	dst := uint32(c.Uint("dst"))
	contractID := c.Uint64("contract_id")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	var changes []gridtypes.Workload
	args := rmbCmdArgs{
		"contract_id": contractID,
	}
	err := client.Call(ctx, dst, "zos.deployment.changes", args, &changes)
	if err != nil {
		return nil, fmt.Errorf("failed to get deployment changes after deploy: %w, contractID: %d", err, contractID)
	}
	res, err := json.Marshal(changes)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal deployment changes%w", err)
	}
	return string(res), nil
}

func deploymentDeploy(c *cli.Context, client *direct.RpcCLient) (interface{}, error) {
	dst := uint32(c.Uint("dst"))
	data := c.String("data")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var dl gridtypes.Deployment
	err := json.Unmarshal([]byte(data), &dl)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal deployment %w", err)
	}

	if err := client.Call(ctx, dst, "zos.deployment.deploy", dl, nil); err != nil {
		return nil, fmt.Errorf("failed to deploy deployment %w", err)
	}

	return nil, nil
}

func deploymentGet(c *cli.Context, client *direct.RpcCLient) (interface{}, error) {
	dst := uint32(c.Uint("dst"))
	data := c.String("data")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var args rmbCmdArgs
	err := json.Unmarshal([]byte(data), &args)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal data to get deployment %w", err)
	}
	var dl gridtypes.Deployment

	if err := client.Call(ctx, dst, "zos.deployment.get", args, &dl); err != nil {
		return nil, fmt.Errorf("failed to get deployment %w", err)
	}
	json, err := json.Marshal(dl)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal deployment %w", err)
	}

	return string(json), nil
}

func nodeTakenPorts(c *cli.Context, client *direct.RpcCLient) (interface{}, error) {
	dst := uint32(c.Uint("dst"))
	var takenPorts []uint16

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := client.Call(ctx, dst, "zos.network.list_wg_ports", nil, &takenPorts); err != nil {
		return nil, fmt.Errorf("failed to get node taken ports %w", err)
	}
	json, err := json.Marshal(takenPorts)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal taken ports %w", err)
	}

	return string(json), nil
}

func getNodePublicConfig(c *cli.Context, client *direct.RpcCLient) (interface{}, error) {
	dst := uint32(c.Uint("dst"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var pubConfig pkg.PublicConfig

	if err := client.Call(ctx, dst, "zos.network.public_config_get", nil, &pubConfig); err != nil {
		return nil, fmt.Errorf("failed to get node public configuration: %w", err)
	}
	json, err := json.Marshal(pubConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal public configuration: %w", err)
	}
	fmt.Println(string(json))
	return string(json), nil
}
