package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	substrate "github.com/threefoldtech/tfchain/clients/tfchain-client-go"
	"github.com/threefoldtech/tfgrid-sdk-go/rmb-sdk-go/direct"
	"github.com/threefoldtech/zos/pkg/gridtypes"
)

type rmbCmdArgs map[string]interface{}

func deploymentChanges(client *direct.DirectClient, dst uint32, contractID uint64) ([]gridtypes.Workload, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var changes []gridtypes.Workload
	args := rmbCmdArgs{
		"contract_id": contractID,
	}
	err := client.Call(ctx, dst, "zos.deployment.changes", args, &changes)
	if err != nil {
		return nil, fmt.Errorf("failed to get deployment changes after deploy %w", err)
	}
	fmt.Println(changes)
	return changes, nil
}

func deploymentDeploy(mnemonics string, substrate_url string, relay_url string, dst uint32, data string) error {
	subManager := substrate.NewManager(substrate_url)
	sub, err := subManager.Substrate()
	if err != nil {
		return fmt.Errorf("failed to connect to substrate: %w", err)
	}

	defer sub.Close()
	client, err := direct.NewClient(context.Background(), direct.KeyTypeSr25519, mnemonics, relay_url, "tfgrid-vclient", sub, true)
	if err != nil {
		return fmt.Errorf("failed to create direct client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var dl gridtypes.Deployment
	err = json.Unmarshal([]byte(data), &dl)
	if err != nil {
		return fmt.Errorf("failed to marshal deployment %w", err)
	}

	if err := client.Call(ctx, dst, "zos.deployment.deploy", dl, nil); err != nil {
		return fmt.Errorf("failed to deploy deployment %w", err)
	}

	dlChanges, err := deploymentChanges(client, dst, dl.ContractID)
	if err != nil {
		return err
	}
	json, err := json.Marshal(dlChanges)
	if err != nil {
		return fmt.Errorf("failed to marshal deployment changes %w", err)
	}

	fmt.Println(string(json))
	return nil
}

func deploymentGet(mnemonics string, substrate_url string, relay_url string, dst uint32, data string) error {
	subManager := substrate.NewManager(substrate_url)
	sub, err := subManager.Substrate()
	if err != nil {
		return fmt.Errorf("failed to connect to substrate: %w", err)
	}

	defer sub.Close()
	client, err := direct.NewClient(context.Background(), direct.KeyTypeSr25519, mnemonics, relay_url, "tfgrid-vclient", sub, true)
	if err != nil {
		return fmt.Errorf("failed to create direct client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var args rmbCmdArgs
	err = json.Unmarshal([]byte(data), &args)
	if err != nil {
		return fmt.Errorf("failed to marshal data to get deployment %w", err)
	}
	var dl gridtypes.Deployment

	if err := client.Call(ctx, dst, "zos.deployment.get", args, &dl); err != nil {
		return fmt.Errorf("failed to get deployment %w", err)
	}
	json, err := json.Marshal(dl)
	if err != nil {
		return fmt.Errorf("failed to marshal deployment %w", err)
	}

	fmt.Println(string(json))

	return nil
}
