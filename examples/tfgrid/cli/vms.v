module main

import cli
import os
import rand
import log
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution
import threefoldtech.threebot.explorer
import freeflowuniverse.crystallib.rpcwebsocket

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
