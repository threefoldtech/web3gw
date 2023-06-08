module main

import cli
import os
import rand
import log
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { Discourse }
import threefoldtech.threebot.explorer
import freeflowuniverse.crystallib.rpcwebsocket

const (
	discourse_packages = {
		'small':       {
			'cpu':    1
			'memory': 2
			'rootfs': 10
			'disk':   50
		}
		'medium':      {
			'cpu':    2
			'memory': 4
			'rootfs': 20
			'disk':   100
		}
		'large':       {
			'cpu':    4
			'memory': 8
			'rootfs': 40
			'disk':   200
		}
		'extra-large': {
			'cpu':    8
			'memory': 16
			'rootfs': 80
			'disk':   400
		}
	}
)

fn get_discourse_commands() []cli.Command {
	return [
		cli.Command{
			name: 'create'
			description: 'Create a discourse instance on the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Name of discourse instance'
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
					name: 'developer_email'
					description: 'Email of the developer'
					required: true
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
				execute_discourse_create(cmd)!
				return
			}
		},
		cli.Command{
			name: 'get'
			description: 'Get a discourse instance from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Name of discourse instance'
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_discourse_get(cmd)!
				return
			}
		},
		cli.Command{
			name: 'delete'
			description: 'delete a discourse instance from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Name of discourse instance'
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_discourse_delete(cmd)!
				return
			}
		},
	]
}

fn execute_discourse_create(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!
	mut capacity := cmd.flags.get_string('capacity')!
	mut ssh_key := cmd.flags.get_string('ssh_key')!
	mut developer_email := cmd.flags.get_string('developer_email')!
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
	if developer_email == '' {
		logger.error('developer_email is required')
	}

	// packagee := discourse_packages[capacity]!

	// execute

	res := solution_handler.deploy_discourse(Discourse{
		name: name
		cpu: u32(discourse_packages[capacity]!['cpu'])
		memory: u32(discourse_packages[capacity]!['memory']) * 1024
		rootfs_size: u32(discourse_packages[capacity]!['rootfs']) * 1024
		disk_size: u32(discourse_packages[capacity]!['disk'])
		ssh_key: ssh_key
		developer_email: developer_email
		smtp_username: smtp_username
		smtp_password: smtp_password
		threebot_private_key: rand.string(16)
		flask_secret_key: rand.string(16)
	})!
	logger.info('${res}')

	return
}

fn execute_discourse_get(cmd cli.Command) ! {
	// get flags
	mut name := cmd.flags.get_string('name')!

	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// assign default values
	if name == '' {
		logger.error('name is required')
	}

	res := solution_handler.get_discourse(name)!
	logger.info('${res}')

	return
}

fn execute_discourse_delete(cmd cli.Command) ! {
	// get flags
	mut name := cmd.flags.get_string('name')!

	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// assign default values
	if name == '' {
		logger.error('name is required')
	}

	solution_handler.delete_discourse(name)!
	logger.info('deleted discourse instance ${name}')

	return
}
