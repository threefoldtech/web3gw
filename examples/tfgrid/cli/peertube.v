module main

import cli
import os
import rand
import log
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { Peertube }
import threefoldtech.threebot.explorer
import freeflowuniverse.crystallib.rpcwebsocket

const (
	peertube_packages = {
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

fn get_peertube_commands() []cli.Command {
	return [
		cli.Command{
			name: 'create'
			description: 'Create a peertube instance on the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Name of peertube instance'
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
				cli.Flag{
					flag: .string
					name: 'smtp_username'
					description: 'SMTP username'
				},
				cli.Flag{
					flag: .string
					name: 'smtp_password'
					description: 'SMTP password'
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_peertube_create(cmd)!
				return
			}
		},
		cli.Command{
			name: 'get'
			description: 'Get a peertube instance from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Name of peertube instance'
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_peertube_get(cmd)!
				return
			}
		},
		cli.Command{
			name: 'delete'
			description: 'delete a peertube instance from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Name of peertube instance'
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_peertube_delete(cmd)!
				return
			}
		},
	]
}

fn execute_peertube_create(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!
	mut capacity := cmd.flags.get_string('capacity')!
	mut ssh_key := cmd.flags.get_string('ssh_key')!
	mut admin_email := cmd.flags.get_string('admin_email')!
	mut smtp_username := cmd.flags.get_string('smtp_username')!
	mut smtp_password := cmd.flags.get_string('smtp_password')!

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
	if smtp_username == '' {
		smtp_username = rand.string(8)
	}
	if smtp_password == '' {
		smtp_password = rand.string(8)
	}
	if admin_email == '' {
		logger.error('admin_email is required')
	}

	// packagee := peertube_packages[capacity]!

	// execute

	res := solution_handler.deploy_peertube(Peertube{
		name: name
		cpu: u32(peertube_packages[capacity]!['cpu'])
		memory: u32(peertube_packages[capacity]!['memory']) * 1024
		rootfs_size: u32(peertube_packages[capacity]!['rootfs']) * 1024
		ssh_key: ssh_key
		admin_email: admin_email
		smtp_username: smtp_username
		smtp_password: smtp_password
	})!
	logger.info('${res}')

	return
}

fn execute_peertube_get(cmd cli.Command) ! {
	// get flags
	mut name := cmd.flags.get_string('name')!

	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// assign default values
	if name == '' {
		logger.error('name is required')
	}

	res := solution_handler.get_peertube(name)!
	logger.info('${res}')

	return
}

fn execute_peertube_delete(cmd cli.Command) ! {
	// get flags
	mut name := cmd.flags.get_string('name')!

	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// assign default values
	if name == '' {
		logger.error('name is required')
	}

	solution_handler.delete_peertube(name)!
	logger.info('deleted peertube instance ${name}')

	return
}
