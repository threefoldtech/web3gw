module main

import cli
import os
import rand
import log
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { Presearch }
import threefoldtech.threebot.explorer
import freeflowuniverse.crystallib.rpcwebsocket

fn get_presearch_commands() []cli.Command {
	return [
		cli.Command{
			name: 'create'
			description: 'Create a presearch instance on the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Name of presearch instance'
				},
				cli.Flag{
					flag: .string
					name: 'ssh_key'
					description: 'Public ssh key to use for the machine'
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_presearch_create(cmd)!
				return
			}
		},
		cli.Command{
			name: 'get'
			description: 'Get a presearch instance from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Name of presearch instance'
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_presearch_get(cmd)!
				return
			}
		},
		cli.Command{
			name: 'delete'
			description: 'delete a presearch instance from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Name of presearch instance'
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_presearch_delete(cmd)!
				return
			}
		},
	]
}

fn execute_presearch_create(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!
	mut ssh_key := cmd.flags.get_string('ssh_key')!

	// assign default values
	if name == '' {
		name = rand.string(8)
	}
	if ssh_key == '' {
		default_ssh_key := os.execute('cat ~/.ssh/id_rsa.pub')
		ssh_key = default_ssh_key.output
	}

	res := solution_handler.deploy_presearch(Presearch{
		name: name
		cpu: 1
		memory: 512
		rootfs_size: 10
		ssh_key: ssh_key
	})!
	logger.info('${res}')

	return
}

fn execute_presearch_get(cmd cli.Command) ! {
	// get flags
	mut name := cmd.flags.get_string('name')!

	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// assign default values
	if name == '' {
		logger.error('name is required')
	}

	res := solution_handler.get_presearch(name)!
	logger.info('${res}')

	return
}

fn execute_presearch_delete(cmd cli.Command) ! {
	// get flags
	mut name := cmd.flags.get_string('name')!

	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// assign default values
	if name == '' {
		logger.error('name is required')
	}

	solution_handler.delete_presearch(name)!
	logger.info('deleted presearch instance ${name}')

	return
}
