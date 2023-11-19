package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/rs/zerolog"
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
						Name:  "substrate",
						Value: "wss://tfchain.grid.tf/ws",
						Usage: "substrate URL",
					},
					cli.StringFlag{
						Name:     "hash",
						Required: true,
						Usage:    "deployment hash",
					},
				},
				Action: substrateDecorator(signDeployment),
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
					nodeId := c.Uint("node_id")
					return getNodeTwin(c, substrateURL, uint32(nodeId))
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
				Action: rmbDecorator(deploymentDeploy),
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
				Action: rmbDecorator(deploymentGet),
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
				Action: rmbDecorator(deploymentChanges),
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
				Action: rmbDecorator(nodeTakenPorts),
			},
			{
				Name:  "rmb-node-pubConfig",
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
				Action: rmbDecorator(getNodePublicConfig),
			},
			{
				Name:  "deploy-single",
				Usage: "Get node public configuration",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "env",
						Value:    "main",
						Usage:    "network env",
						Required: true,
					},
					cli.StringFlag{
						Name:     "mnemonics",
						Value:    "",
						Usage:    "user mnemonics",
						Required: true,
					},
					cli.StringFlag{
						Name:     "data",
						Value:    "",
						Usage:    "vm data",
						Required: true,
					},
					cli.StringFlag{
						Name:     "solution_type",
						Value:    "",
						Usage:    "solution type",
						Required: true,
					},
					cli.UintFlag{
						Name:     "node",
						Value:    0,
						Usage:    "node id",
						Required: true,
					},
				},
				Action: deployVM(),
			},
			{
				Name:  "generate-wg-key",
				Usage: "Generates wireguard private key",
				Flags: []cli.Flag{},
				Action: func(c *cli.Context) error {
					return generateWgPrivKey()
				},
			},
			{
				Name:  "rmb",
				Usage: "Make RMB call",
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
						Value:    "wss://relay.grid.tf/ws",
						Usage:    "relay URL",
						Required: true,
					},
					cli.UintFlag{
						Name:     "dst",
						Value:    0,
						Usage:    "destination node",
						Required: true,
					},
					cli.StringFlag{
						Name:     "cmd",
						Usage:    "rmb command",
						Required: true,
					},
					cli.StringFlag{
						Name:  "payload",
						Value: "",
						Usage: "command payload",
					},
				},
				Action: rmbDecorator(rmbCall),
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
