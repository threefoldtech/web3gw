module main

import cli
import os
import rand
import log
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution
import threefoldtech.threebot.explorer
import freeflowuniverse.crystallib.rpcwebsocket

fn main() {
	mut app := cli.Command{
		name: 'tfgrid-cli'
		description: 'is a command line interface for the ThreeFold Grid.'
		usage: 'tfgrid-cli [module] [operation] [flags]'
		disable_help: true
		disable_man: true
		execute: fn (cmd cli.Command) ! {
			print_manual(cmd, 'root')
			return
		}
		commands: get_grid_modules_commands()
		flags: [
			cli.Flag{
				flag: .string
				name: 'grid'
				description: 'Grid Network (dev, qa, test, main)'
				required: false
				global: true
			},
			cli.Flag{
				flag: .string
				name: 'mnemonic'
				description: 'Secret mnemonic to use to access the grid'
				required: true
				global: true
			},
		]
	}
	app.setup()
	app.parse(os.args)
}

fn get_grid_modules_commands() []cli.Command {
	return [
		cli.Command{
			name: 'vms'
			description: 'machines module on the grid'
			usage: 'tfgrid-cli vms [operation] [flags]'
			disable_help: true
			disable_man: true
			execute: fn (cmd cli.Command) ! {
				print_manual(cmd, 'module')
				return
			}
			commands: get_vms_commands()
		},
		cli.Command{
			name: 'k8s'
			description: 'kubernetes module on the grid'
			usage: 'tfgrid-cli k8s [operation] [flags]'
			disable_help: true
			disable_man: true
			execute: fn (cmd cli.Command) ! {
				print_manual(cmd, 'module')
				return
			}
			commands: get_k8s_commands()
		},
		cli.Command{
			name: 'gws_fqdn'
			description: 'FQDN Gateway module on the grid'
			usage: 'tfgrid-cli gws_fqdn [operation] [flags]'
			disable_help: true
			disable_man: true
			execute: fn (cmd cli.Command) ! {
				print_manual(cmd, 'module')
				return
			}
			commands: get_gws_fqdn_commands()
		},
		cli.Command{
			name: 'gws_name'
			description: 'Name Gateway module on the grid'
			usage: 'tfgrid-cli gws_name [operation] [flags]'
			disable_help: true
			disable_man: true
			execute: fn (cmd cli.Command) ! {
				print_manual(cmd, 'module')
				return
			}
			commands: get_gws_name_commands()
		},
		cli.Command{
			name: 'zdb'
			description: 'ZDB module on the grid'
			usage: 'tfgrid-cli zdb [operation] [flags]'
			disable_help: true
			disable_man: true
			execute: fn (cmd cli.Command) ! {
				print_manual(cmd, 'module')
				return
			}
			commands: get_zdb_commands()
		},
		cli.Command{
			name: 'discourse'
			description: 'Discourse solution on the grid'
			usage: 'tfgrid-cli discourse [operation] [flags]'
			disable_help: true
			disable_man: true
			execute: fn (cmd cli.Command) ! {
				print_manual(cmd, 'module')
				return
			}
			commands: get_discourse_commands()
		},
		cli.Command{
			name: 'funkwhale'
			description: 'funkwhale solution on the grid'
			usage: 'tfgrid-cli funkwhale [operation] [flags]'
			disable_help: true
			disable_man: true
			execute: fn (cmd cli.Command) ! {
				print_manual(cmd, 'module')
				return
			}
			// commands: get_funkwhale_commands()
		},
		cli.Command{
			name: 'peertube'
			description: 'peertube solution on the grid'
			usage: 'tfgrid-cli peertube [operation] [flags]'
			disable_help: true
			disable_man: true
			execute: fn (cmd cli.Command) ! {
				print_manual(cmd, 'module')
				return
			}
			// commands: get_peertube_commands()
		},
		cli.Command{
			name: 'presearch'
			description: 'presearch solution on the grid'
			usage: 'tfgrid-cli presearch [operation] [flags]'
			disable_help: true
			disable_man: true
			execute: fn (cmd cli.Command) ! {
				print_manual(cmd, 'module')
				return
			}
			// commands: get_presearch_commands()
		},
	]
}
