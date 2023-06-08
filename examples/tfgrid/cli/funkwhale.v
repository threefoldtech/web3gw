module main

import cli
import os
import rand
import log
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { Funkwhale }
import threefoldtech.threebot.explorer
import freeflowuniverse.crystallib.rpcwebsocket

const (
	funkwhale_packages = {
		'small':       {
			'cpu':    1
			'memory': 2
			'rootfs': 10
		}
		'medium':      {
			'cpu':    2
			'memory': 4
			'rootfs': 20
		}
		'large':       {
			'cpu':    4
			'memory': 8
			'rootfs': 40
		}
		'extra-large': {
			'cpu':    8
			'memory': 16
			'rootfs': 80
		}
	}
)

fn get_funkwhale_commands() []cli.Command {
	return [
		cli.Command{
			name: 'create'
			description: 'Create a funkwhale instance on the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Name of funkwhale instance'
				},
				cli.Flag{
					flag: .string
					name: 'capacity'
					description: 'Capacity of the machine'
				},
				cli.Flag{
					flag: .string
					name: 'ssh_key'
					description: 'Public ssh key to use for the machine'
				},
				cli.Flag{
					flag: .string
					name: 'admin_email'
					description: 'Email of the admin'
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_funkwhale_create(cmd)!
				return
			}
		},
		cli.Command{
			name: 'get'
			description: 'Get a funkwhale instance from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Name of funkwhale instance'
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_funkwhale_get(cmd)!
				return
			}
		},
		cli.Command{
			name: 'delete'
			description: 'delete a funkwhale instance from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Name of funkwhale instance'
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_funkwhale_delete(cmd)!
				return
			}
		},
	]
}

fn execute_funkwhale_create(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!
	mut capacity := cmd.flags.get_string('capacity')!
	mut ssh_key := cmd.flags.get_string('ssh_key')!
	mut admin_email := cmd.flags.get_string('admin_email')!

	// assign default values
	if name == '' {
		name = rand.string(8)
	}
	if capacity == '' {
		capacity = 'small'
	}
	if ssh_key == '' {
		default_ssh_key := os.execute('cat ~/.ssh/id_rsa.pub')
		ssh_key = default_ssh_key.output
	}
	if admin_email == '' {
		logger.error('admin_email is required')
	}
	res := solution_handler.deploy_funkwhale(Funkwhale{
		name: name
		cpu: u32(funkwhale_packages[capacity]!['cpu'])
		memory: u32(funkwhale_packages[capacity]!['memory']) * 1024
		rootfs_size: u32(funkwhale_packages[capacity]!['rootfs']) * 1024
		admin_email: admin_email
	})!
	logger.info('${res}')

	return
}

fn execute_funkwhale_get(cmd cli.Command) ! {
	// get flags
	mut name := cmd.flags.get_string('name')!

	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// assign default values
	if name == '' {
		logger.error('name is required')
	}

	res := solution_handler.get_funkwhale(name)!
	logger.info('${res}')

	return
}

fn execute_funkwhale_delete(cmd cli.Command) ! {
	// get flags
	mut name := cmd.flags.get_string('name')!

	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// assign default values
	if name == '' {
		logger.error('name is required')
	}

	solution_handler.delete_funkwhale(name)!
	logger.info('deleted funkwhale instance ${name}')

	return
}
