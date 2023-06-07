module main

import cli
import os
import rand
import log
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { SolutionHandler }
import threefoldtech.threebot.explorer
import freeflowuniverse.crystallib.rpcwebsocket

fn print_manual(cmd cli.Command, level string) string {
	levels := {
		'root':      'modules'
		'module':    'operations'
		'operation': 'flags'
	}

	mut msg := ''

	msg += cmd.name + ' ' + cmd.description
	msg += '\n\n'
	msg += 'Usage: ' + cmd.usage
	msg += '\n\n'

	msg += 'Available ' + levels[level] or { 'commands' } + ':'

	for sub in cmd.commands {
		msg += '\n\t- ' + sub.name + '\t: ' + sub.description
	}

	msg += '\n\n'
	msg += 'Available flags:'
	for flag in cmd.flags {
		msg += '\n\t- ' + flag.name + '\t: ' + flag.description
	}

	println(msg)
	return msg
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
	]
}

fn get_vms_commands() []cli.Command {
	return [
		cli.Command{
			name: 'create'
			description: 'Create a machine module on the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'network'
					description: 'Network name to deploy the machine on'
					required: false
				},
				cli.Flag{
					flag: .string
					name: 'capacity'
					description: 'Capacity package (small, medium, large, extra-large)'
					required: false
				},
				cli.Flag{
					flag: .int
					name: 'times'
					description: 'Number of machines to deploy'
					required: false
				},
				cli.Flag{
					flag: .bool
					name: 'gateway'
					description: 'Either or not to deploy a gateway'
					required: false
				},
				cli.Flag{
					flag: .bool
					name: 'wg'
					description: 'Either or not to add a wireguard access to the machine'
					required: false
				},
				cli.Flag{
					flag: .string
					name: 'ssh_key'
					description: 'Public ssh key to add to the machine'
					required: false
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_vms_create(cmd)!
				return
			}
		},
		cli.Command{
			name: 'delete'
			description: 'Delete a machine module on the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'network'
					description: 'Network name to delete the machines on'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_vms_delete(cmd)!
				return
			}
		},
		cli.Command{
			name: 'add'
			description: 'Add a machine to deployed network'
			flags: [
				cli.Flag{
					flag: .string
					name: 'network'
					description: 'Network name to deploy the machine on'
					required: false
				},
				cli.Flag{
					flag: .string
					name: 'capacity'
					description: 'Capacity package (small, medium, large, extra-large)'
					required: false
				},
				cli.Flag{
					flag: .int
					name: 'times'
					description: 'Number of machines to deploy'
					required: false
				},
				cli.Flag{
					flag: .bool
					name: 'gateway'
					description: 'Either or not to deploy a gateway'
					required: false
				},
				cli.Flag{
					flag: .bool
					name: 'wg'
					description: 'Either or not to add a wireguard access to the machine'
					required: false
				},
				cli.Flag{
					flag: .string
					name: 'ssh_key'
					description: 'Public ssh key to add to the machine'
					required: false
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_vms_create(cmd)!
				return
			}
		},
		cli.Command{
			name: 'remove'
			description: 'Remove a single machine from a deployed network'
			flags: [
				cli.Flag{
					flag: .string
					name: 'network'
					description: 'Network name where the machine is deployed'
					required: true
				},
				cli.Flag{
					flag: .string
					name: 'machine_name'
					description: 'Machine name to remove'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_vms_remove(cmd)!
				return
			}
		},
		cli.Command{
			name: 'get'
			description: 'Get a deployed machine from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'network'
					description: 'Network name where the machine is deployed'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_vms_get(cmd)!
				return
			}
		},
	]
}

fn get_solution_handler(cmd cli.Command) !(SolutionHandler, log.Logger) {
	mut logger := log.Log{
		level: .info
	}

	mut rpc_client := rpcwebsocket.new_rpcwsclient('ws://127.0.0.1:8080', &logger)!
	_ := spawn rpc_client.run()

	mut grid_network := cmd.flags.get_string('grid')!
	mnemonic := cmd.flags.get_string('mnemonic')!

	if grid_network == '' {
		grid_network = 'dev'
	}

	mut tfgrid_client := tfgrid.new(mut rpc_client)
	mut explorer_client := explorer.new(mut rpc_client)

	tfgrid_client.load(tfgrid.Credentials{
		mnemonic: mnemonic
		network: grid_network
	})!
	explorer_client.load(grid_network)!

	mut sl := SolutionHandler{
		tfclient: &tfgrid_client
		explorer: &explorer_client
	}

	return sl, logger
}

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

fn execute_vms_create(cmd cli.Command) ! {
	// get flags
	mut network := cmd.flags.get_string('network')!
	mut capacity := cmd.flags.get_string('capacity')!
	mut times := cmd.flags.get_int('times')!
	mut gateway := cmd.flags.get_bool('gateway')!
	mut wg := cmd.flags.get_bool('wg')!
	mut ssh_key := cmd.flags.get_string('ssh_key')!

	// assign default values
	if network == '' {
		network = rand.string(8)
	}
	if capacity == '' {
		capacity = 'small'
	}
	if times == 0 {
		times = 1
	}
	if ssh_key == '' {
		default_ssh_key := os.execute('cat ~/.ssh/id_rsa.pub')
		ssh_key = default_ssh_key.output
	}

	// execute
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	res := solution_handler.create_vm(solution.VM{
		network: network
		capacity: capacity
		times: u32(times)
		gateway: gateway
		add_wireguard_access: wg
		ssh_key: ssh_key
	})!
	logger.info('${res}')

	return
}

fn execute_vms_get(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	mut network := cmd.flags.get_string('network')!

	if network == '' {
		logger.info('Please provide a network name')
	}

	res := solution_handler.get_vm(network)!
	logger.info('${res}')
}

fn execute_vms_remove(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	mut network := cmd.flags.get_string('network')!
	mut machine_name := cmd.flags.get_string('machine_name')!

	if network == '' {
		logger.info('Please provide a network name')
	}

	if machine_name == '' {
		logger.info('Please provide a machine name')
	}

	res := solution_handler.remove_vm(network, machine_name)!
	logger.info('${res}')
}

fn execute_vms_delete(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	mut network := cmd.flags.get_string('network')!

	if network == '' {
		logger.info('Please provide a network name')
	}

	solution_handler.delete_vm(network)!
	logger.info('deleted network ${network}')
}

fn get_k8s_commands() []cli.Command {
	return [
		cli.Command{
			name: 'create'
			description: 'ops bardo'
			execute: fn (cmd cli.Command) ! {
				println('Creating k8s')
				return
			}
		},
		cli.Command{
			name: 'delete'
			description: 'k8s bardo'
			execute: fn (cmd cli.Command) ! {
				println('Deleting k8s')
				return
			}
		},
	]
}

fn get_gws_fqdn_commands() []cli.Command {
	return [
		cli.Command{
			name: 'create'
			description: 'ops bardo'
			execute: fn (cmd cli.Command) ! {
				println('Creating gw')
				return
			}
		},
		cli.Command{
			name: 'delete'
			description: 'gw bardo'
			execute: fn (cmd cli.Command) ! {
				println('Deleting gw')
				return
			}
		},
	]
}

fn get_gws_name_commands() []cli.Command {
	return [
		cli.Command{
			name: 'create'
			description: 'ops bardo'
			execute: fn (cmd cli.Command) ! {
				println('Creating gw')
				return
			}
		},
		cli.Command{
			name: 'delete'
			description: 'gw bardo'
			execute: fn (cmd cli.Command) ! {
				println('Deleting gw')
				return
			}
		},
	]
}

fn get_zdb_commands() []cli.Command {
	return [
		cli.Command{
			name: 'create'
			description: 'ops bardo'
			execute: fn (cmd cli.Command) ! {
				println('Creating gw')
				return
			}
		},
		cli.Command{
			name: 'delete'
			description: 'gw bardo'
			execute: fn (cmd cli.Command) ! {
				println('Deleting gw')
				return
			}
		},
	]
}
