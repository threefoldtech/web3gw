module main

import cli
import os
import rand
import log
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution
import threefoldtech.threebot.explorer
import freeflowuniverse.crystallib.rpcwebsocket

fn get_zdb_commands() []cli.Command {
	return [
		cli.Command{
			name: 'create'
			description: 'deploy a zdb on the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Zdb name'
					required: true
				},
				cli.Flag{
					flag: .int
					name: 'node_id'
					description: 'Node id to deploy the zdb on'
					required: true
				},
				cli.Flag{
					flag: .int
					name: 'size'
					description: 'size of zdb in GB'
					required: true
				},
				cli.Flag{
					flag: .string
					name: 'password'
					description: 'secure password for zdb'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_zdb_create(cmd)!
				return
			}
		},
		cli.Command{
			name: 'get'
			description: 'get a zdb from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Zdb name'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_zdb_get(cmd)!
				return
			}
		},
		cli.Command{
			name: 'delete'
			description: 'delete a zdb from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Zdb name'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_zdb_delete(cmd)!
				return
			}
		},
	]
}

fn execute_zdb_create(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!
	mut node_id := cmd.flags.get_int('node_id')!
	mut password := cmd.flags.get_string('password')!
	mut size := cmd.flags.get_int('size')!

	// assign default values
	if name == '' {
		name = rand.string(8)
	}
	if password == '' {
		password = rand.string(8)
	}
	if size == 0 {
		size = 10
	}

	res := solution_handler.tfclient.zdb_deploy(tfgrid.ZDB{
		name: name
		node_id: u32(node_id)
		password: password
		size: u32(size)
	})!

	logger.info('${res}')
}

fn execute_zdb_get(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!

	// assign default values
	if name == '' {
		logger.error('Please provide a name')
	}
	res := solution_handler.tfclient.zdb_get(name)!

	logger.info('${res}')
}

fn execute_zdb_delete(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!

	// assign default values
	if name == '' {
		logger.error('Please provide a name')
	}
	solution_handler.tfclient.zdb_delete(name)!

	logger.info('deleted zdb ${name}')
}
