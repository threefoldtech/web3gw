package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	substrate "github.com/threefoldtech/tfchain/clients/tfchain-client-go"
	"github.com/urfave/cli"
)

func main() {
	log.SetOutput(io.Discard)
	zerolog.SetGlobalLevel(zerolog.Disabled)

	app := &cli.App{
		Name:  "grid",
		Usage: "Example: grid [COMMAND]",
		Commands: []cli.Command{
			{
				Name:  "new-node-cn",
				Usage: "create node contract",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "substrate",
						Value: "wss://tfchain.grid.tf/ws",
						Usage: "substrate URL",
					},
					cli.StringFlag{
						Name:     "mnemonics",
						Usage:    "user mnemonics",
						Required: true,
					},
					cli.UintFlag{
						Name:     "node_id",
						Required: true,
						Usage:    "node id to create the contract on",
					},
					cli.StringFlag{
						Name:  "body",
						Value: "",
						Usage: "contract body",
					},
					cli.StringFlag{
						Name:     "hash",
						Required: true,
						Usage:    "deployment hash",
					},
					cli.UintFlag{
						Name:  "public_ips",
						Value: 0,
						Usage: "number of reserved public ips for this deployment",
					},
					cli.Uint64Flag{
						Name:  "solution_provider",
						Value: 0,
						Usage: "twin id for the solution provider",
					},
				},
				Action: substrateDecorator(createNodeContract),
			},
			{
				Name:  "update-cn",
				Usage: "update node contract",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "mnemonics",
						Value:    "",
						Usage:    "user mnemonics",
						Required: true,
					},
					cli.StringFlag{
						Name:  "substrate",
						Value: "wss://tfchain.grid.tf/ws",
						Usage: "substrate URL",
					},
					cli.UintFlag{
						Name:     "contract_id",
						Required: true,
						Usage:    "id of contract to update",
					},
					cli.StringFlag{
						Name:  "body",
						Value: "",
						Usage: "contract body",
					},
					cli.StringFlag{
						Name:     "hash",
						Required: true,
						Usage:    "deployment hash",
					},
				},
				Action: substrateDecorator(updateNodeContract),
			},
			{
				Name:  "cancel-cn",
				Usage: "cancel any type of contract",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "mnemonics",
						Value:    "",
						Usage:    "user mnemonics",
						Required: true,
					},
					cli.StringFlag{
						Name:  "substrate",
						Value: "wss://tfchain.grid.tf/ws",
						Usage: "substrate URL",
					},
					cli.UintFlag{
						Name:     "contract_id",
						Required: true,
						Usage:    "id of contract to delete",
					},
				},
				Action: substrateDecorator(cancelContract),
			},
			{
				Name:  "new-name-cn",
				Usage: "create name contract",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "mnemonics",
						Value:    "",
						Usage:    "user mnemonics",
						Required: true,
					},
					cli.StringFlag{
						Name:  "substrate",
						Value: "wss://tfchain.grid.tf/ws",
						Usage: "substrate URL",
					},
					cli.StringFlag{
						Name:     "name",
						Required: true,
						Usage:    "contract name",
					},
				},
				Action: substrateDecorator(createNameContract),
			},
			{
				Name:  "new-rent-cn",
				Usage: "create rent contract",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "mnemonics",
						Value:    "",
						Usage:    "user mnemonics",
						Required: true,
					},
					cli.StringFlag{
						Name:  "substrate",
						Value: "wss://tfchain.grid.tf/ws",
						Usage: "substrate URL",
					},
					cli.UintFlag{
						Name:     "node_id",
						Required: true,
						Usage:    "id of node to rent",
					},
					cli.UintFlag{
						Name:  "solution_provider",
						Value: 0,
						Usage: "solution provider twin id",
					},
				},
				Action: substrateDecorator(createRentContract),
			},
			{
				Name:  "sign",
				Usage: "sign a deployment",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "mnemonics",
						Value:    "",
						Usage:    "user mnemonics",
						Required: true,
					},
					cli.StringFlag{
						Name:     "hash",
						Required: true,
						Usage:    "deployment hash",
					},
				},
				Action: func(c *cli.Context) error {
					mnemonics := c.String("mnemonics")
					identity, err := substrate.NewIdentityFromSr25519Phrase(mnemonics)
					if err != nil {
						return errors.Wrap(err, "failed to create identity from provided mnemonics")
					}

					hashHex := c.String("hash")
					hashByets, err := hex.DecodeString(hashHex)
					if err != nil {
						return errors.Wrap(err, "failed to decode deployment hash")
					}
					signatureBytes, err := identity.Sign(hashByets)
					if err != nil {
						return errors.Wrap(err, "failed to sign deployment hash")
					}

					sig := hex.EncodeToString(signatureBytes)
					fmt.Printf("%s", sig)

					return nil
				},
			},
			{
				Name:  "node-twin",
				Usage: "get node twin id",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "substrate",
						Value: "wss://tfchain.grid.tf/ws",
						Usage: "substrate URL",
					},
					cli.StringFlag{
						Name:     "node_id",
						Required: true,
						Usage:    "node id",
					},
				},
				Action: func(c *cli.Context) error {
					substrateURL := c.String("substrate")
					manager := substrate.NewManager(substrateURL)
					sub, err := manager.Substrate()
					if err != nil {
						return errors.Wrap(err, "failed to create substrate connection")
					}
					defer sub.Close()
					nodeId := c.Uint("node_id")
					node, err := sub.GetNode(uint32(nodeId))
					if err != nil {
						return errors.Wrapf(err, "failed to get node data for Id: %d", nodeId)
					}
					fmt.Printf("%d", node.TwinID)
					return nil
				},
			},
			{
				Name:  "user-twin",
				Usage: "get user twin id",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "mnemonics",
						Value:    "",
						Usage:    "user mnemonics",
						Required: true,
					},
					cli.StringFlag{
						Name:  "substrate",
						Value: "wss://tfchain.grid.tf/ws",
						Usage: "substrate URL",
					},
				},
				Action: substrateDecorator(getUserTwin),
			},
			{
				Name:  "rmb-dl-deploy",
				Usage: "call rmb func",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "mnemonics",
						Value:    "",
						Usage:    "user mnemonics",
						Required: true,
					},
					cli.StringFlag{
						Name:     "substrate",
						Value:    "wss://tfchain.grid.tf/ws",
						Usage:    "substrate URL",
						Required: true,
					},
					cli.StringFlag{
						Name:     "relay",
						Value:    "wss://tfchain.grid.tf/ws",
						Usage:    "relay URL",
						Required: true,
					},
					cli.UintFlag{
						Name:     "dst",
						Value:    0,
						Usage:    "destination",
						Required: true,
					},

					cli.StringFlag{
						Name:     "data",
						Value:    "wss://tfchain.grid.tf/ws",
						Usage:    "Data to be sent",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					mnemonics := c.String("mnemonics")
					substrate_url := c.String("substrate")
					relay := c.String("relay")
					dst := uint32(c.Uint("dst"))
					data := c.String("data")
					return deploymentDeploy(mnemonics, substrate_url, relay, dst, data)

				},
			},
			{
				Name:  "rmb-dl-get",
				Usage: "call rmb func",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "mnemonics",
						Value:    "",
						Usage:    "user mnemonics",
						Required: true,
					},
					cli.StringFlag{
						Name:     "substrate",
						Value:    "wss://tfchain.grid.tf/ws",
						Usage:    "substrate URL",
						Required: true,
					},
					cli.StringFlag{
						Name:     "relay",
						Value:    "wss://tfchain.grid.tf/ws",
						Usage:    "relay URL",
						Required: true,
					},
					cli.UintFlag{
						Name:     "dst",
						Value:    0,
						Usage:    "destination",
						Required: true,
					},

					cli.StringFlag{
						Name:     "data",
						Value:    "wss://tfchain.grid.tf/ws",
						Usage:    "Data to be sent",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					mnemonics := c.String("mnemonics")
					substrate_url := c.String("substrate")
					relay := c.String("relay")
					dst := uint32(c.Uint("dst"))
					data := c.String("data")
					return deploymentGet(mnemonics, substrate_url, relay, dst, data)

				},
			},
			{
				Name:  "rmb-dl-changes",
				Usage: "call rmb func",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "mnemonics",
						Value:    "",
						Usage:    "user mnemonics",
						Required: true,
					},
					cli.StringFlag{
						Name:     "substrate",
						Value:    "wss://tfchain.grid.tf/ws",
						Usage:    "substrate URL",
						Required: true,
					},
					cli.StringFlag{
						Name:     "relay",
						Value:    "wss://tfchain.grid.tf/ws",
						Usage:    "relay URL",
						Required: true,
					},
					cli.UintFlag{
						Name:     "dst",
						Value:    0,
						Usage:    "destination",
						Required: true,
					},

					cli.StringFlag{
						Name:     "contract_id",
						Value:    "0",
						Usage:    "contract id to get changes for",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					mnemonics := c.String("mnemonics")
					substrate_url := c.String("substrate")
					relay := c.String("relay")
					dst := uint32(c.Uint("dst"))
					contract_id := uint64(c.Uint("contract_id"))
					return deploymentChanges(mnemonics, substrate_url, relay, dst, contract_id)

				},
			},
			{
				Name:  "rmb-taken-ports",
				Usage: "call rmb func",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "substrate",
						Value:    "wss://tfchain.grid.tf/ws",
						Usage:    "substrate URL",
						Required: true,
					},
					cli.StringFlag{
						Name:     "mnemonics",
						Value:    "",
						Usage:    "user mnemonics",
						Required: true,
					},
					cli.StringFlag{
						Name:     "relay",
						Value:    "wss://tfchain.grid.tf/ws",
						Usage:    "relay URL",
						Required: true,
					},
					cli.UintFlag{
						Name:     "dst",
						Value:    0,
						Usage:    "destination node",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					mnemonics := c.String("mnemonics")
					substrate_url := c.String("substrate")
					relay := c.String("relay")
					dst := uint32(c.Uint("dst"))
					return nodeTakenPorts(mnemonics, substrate_url, relay, dst)

				},
			},
			{
				Name: "rmb-node-pubConfig",
				Usage: "Get node public configuration",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "substrate",
						Value:    "wss://tfchain.grid.tf/ws",
						Usage:    "substrate URL",
						Required: true,
					},
					cli.StringFlag{
						Name:     "mnemonics",
						Value:    "",
						Usage:    "user mnemonics",
						Required: true,
					},
					cli.StringFlag{
						Name:     "relay",
						Value:    "wss://tfchain.grid.tf/ws",
						Usage:    "relay URL",
						Required: true,
					},
					cli.UintFlag{
						Name:     "dst",
						Value:    0,
						Usage:    "destination node",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					mnemonics := c.String("mnemonics")
					substrate_url := c.String("substrate")
					relay := c.String("relay")
					dst := uint32(c.Uint("dst"))
					return getNodePublicConfig(mnemonics, substrate_url, relay, dst)

				},

			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

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

