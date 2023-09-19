package main

import (
	"encoding/hex"
	"fmt"

	"github.com/pkg/errors"
	substrate "github.com/threefoldtech/tfchain/clients/tfchain-client-go"
	"github.com/urfave/cli"
)

func substrateDecorator(action func(ctx *cli.Context, sub *substrate.Substrate, identity substrate.Identity) (interface{}, error)) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		substrateURL := ctx.String("substrate")
		manager := substrate.NewManager(substrateURL)
		sub, err := manager.Substrate()
		if err != nil {
			return errors.Wrap(err, "failed to create substrate connection")
		}
		defer sub.Close()

		mnemonics := ctx.String("mnemonics")
		if mnemonics == "" {
			return errors.Wrap(err, "failed to create identity please provide mnemonics")
		}
		identity, err := substrate.NewIdentityFromSr25519Phrase(mnemonics)
		if err != nil {
			return errors.Wrap(err, "failed to create identity from provided mnemonics")
		}

		ret, err := action(ctx, sub, identity)
		if err != nil {
			return err
		}
		fmt.Printf("%v", ret)
		return nil
	}
}

func createNameContract(ctx *cli.Context, sub *substrate.Substrate, identity substrate.Identity) (interface{}, error) {
	name := ctx.String("name")

	contractID, err := sub.CreateNameContract(identity, name)
	if err != nil {
		return nil, err
	}

	return contractID, nil
}

func createRentContract(ctx *cli.Context, sub *substrate.Substrate, identity substrate.Identity) (interface{}, error) {
	nodeID := ctx.Uint("node_id")
	solutionProvider := ctx.Uint64("solution_provider")
	spp := &solutionProvider
	if solutionProvider == 0 {
		spp = nil
	}

	contractID, err := sub.CreateRentContract(identity, uint32(nodeID), spp)
	if err != nil {
		return nil, err
	}

	return contractID, nil
}

func cancelContract(ctx *cli.Context, sub *substrate.Substrate, identity substrate.Identity) (interface{}, error) {
	contractID := ctx.Uint64("contract_id")

	if err := sub.CancelContract(identity, contractID); err != nil {
		return nil, err
	}

	return "", nil
}

func createNodeContract(ctx *cli.Context, sub *substrate.Substrate, identity substrate.Identity) (interface{}, error) {
	nodeID := ctx.Uint("node_id")
	body := ctx.String("body")
	hash := ctx.String("hash")
	publicIPs := ctx.Uint("public_ips")
	solutionProvider := ctx.Uint64("solution_provider")
	spp := &solutionProvider
	if solutionProvider == 0 {
		spp = nil
	}

	contractID, err := sub.CreateNodeContract(identity, uint32(nodeID), body, hash, uint32(publicIPs), spp)
	if err != nil {
		return nil, err
	}

	return contractID, nil
}

func updateNodeContract(ctx *cli.Context, sub *substrate.Substrate, identity substrate.Identity) (interface{}, error) {
	contractID := ctx.Uint64("contract_id")
	body := ctx.String("body")
	hash := ctx.String("hash")

	_, err := sub.UpdateNodeContract(identity, contractID, body, hash)
	if err != nil {
		return nil, err
	}

	return "", nil
}

func getUserTwin(ctx *cli.Context, sub *substrate.Substrate, identity substrate.Identity) (interface{}, error) {
	keypair, err := identity.KeyPair()
	if err != nil {
		return nil, err
	}
	twin, err := sub.GetTwinByPubKey(keypair.Public())
	if err != nil {
		return nil, err
	}

	return twin, nil
}

func getNodeTwin(ctx *cli.Context, substrateURL string, nodeId uint32) error {
	manager := substrate.NewManager(substrateURL)
	sub, err := manager.Substrate()
	if err != nil {
		return errors.Wrap(err, "failed to create substrate connection to get node twin")
	}
	defer sub.Close()
	node, err := sub.GetNode(uint32(nodeId))
	if err != nil {
		return errors.Wrapf(err, "failed to get node data for Id: %d", nodeId)
	}
	fmt.Printf("%d", node.TwinID)
	return nil
}

func signDeployment(ctx *cli.Context, sub *substrate.Substrate, identity substrate.Identity) (interface{}, error) {

	hashHex := ctx.String("hash")
	hashByets, err := hex.DecodeString(hashHex)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode deployment hash")
	}
	signatureBytes, err := identity.Sign(hashByets)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign deployment hash")
	}

	sig := hex.EncodeToString(signatureBytes)
	return sig, nil
}
