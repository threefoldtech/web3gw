module main

import cli
import os
import rand
import log
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution
import threefoldtech.threebot.explorer
import freeflowuniverse.crystallib.rpcwebsocket

fn get_gws_fqdn_commands() []cli.Command {
	return [
		cli.Command{
			name: 'create'
			description: 'Get a deployed cluster from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'FQDN Gateway name'
				},
				cli.Flag{
					flag: .string
					name: 'fqdn'
					description: 'The full domain name'
					required: true
				},
				cli.Flag{
					flag: .string
					name: 'backend'
					description: 'backend of your deployment (ex: "http://1.1.1.1:9000")'
					required: true
				},
				cli.Flag{
					flag: .int
					name: 'node_id'
					description: 'grid node to deploy the fqdn Gateway on'
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_gws_fqdn_get(cmd)!
				return
			}
		},
		cli.Command{
			name: 'delete'
			description: 'delete a deployed fqdn Gateway from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'FQDN Gateway name'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_gws_fqdn_get(cmd)!
				return
			}
		},
		cli.Command{
			name: 'get'
			description: 'Get a deployed fqdn Gateway from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'FQDN Gateway name'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_gws_fqdn_get(cmd)!
				return
			}
		},
	]
}

fn execute_gws_fqdn_create(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!
	mut fqdn := cmd.flags.get_string('fqdn')!
	mut backend := cmd.flags.get_string('backend')!
	mut node_id := cmd.flags.get_int('node_id')!

	// assign default values
	if fqdn == '' {
		logger.error('Please provide a fqdn')
	}
	if backend == '' {
		logger.error('Please provide a backend')
	}
	if name == '' {
		name = rand.string(8)
	}

	res := solution_handler.tfclient.gateways_deploy_fqdn(tfgrid.GatewayFQDN{
		name: name
		node_id: u32(node_id)
		backends: [backend]
		fqdn: fqdn
	})!

	logger.info('${res}')
}

fn execute_gws_fqdn_delete(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!

	// assign default values
	if name == '' {
		logger.error('Please provide a name')
	}

	solution_handler.tfclient.gateways_delete_fqdn(name)!
	logger.info('deleted fqdn ${name}')
}

fn execute_gws_fqdn_get(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!

	// assign default values
	if name == '' {
		logger.error('Please provide a name')
	}

	res := solution_handler.tfclient.gateways_get_fqdn(name)!

	logger.info('${res}')
}

fn get_gws_name_commands() []cli.Command {
	return [
		cli.Command{
			name: 'create'
			description: 'Get a deployed cluster from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Name Gateway name'
				},
				cli.Flag{
					flag: .string
					name: 'backend'
					description: 'backend of your deployment (ex: "http://1.1.1.1:9000")'
					required: true
				},
				cli.Flag{
					flag: .int
					name: 'node_id'
					description: 'grid node to deploy the name Gateway on'
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_gws_name_get(cmd)!
				return
			}
		},
		cli.Command{
			name: 'delete'
			description: 'delete a deployed name Gateway from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Name Gateway name'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_gws_name_get(cmd)!
				return
			}
		},
		cli.Command{
			name: 'get'
			description: 'Get a deployed name Gateway from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Name Gateway name'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_gws_name_get(cmd)!
				return
			}
		},
	]
}

fn execute_gws_name_create(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!
	mut backend := cmd.flags.get_string('backend')!
	mut node_id := cmd.flags.get_int('node_id')!

	// assign default values
	if backend == '' {
		logger.error('Please provide a backend')
	}
	if name == '' {
		name = rand.string(8)
	}

	res := solution_handler.tfclient.gateways_deploy_name(tfgrid.GatewayName{
		name: name
		node_id: u32(node_id)
		backends: [backend]
	})!

	logger.info('${res}')
}

fn execute_gws_name_delete(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!

	// assign default values
	if name == '' {
		logger.error('Please provide a name')
	}

	solution_handler.tfclient.gateways_delete_name(name)!
	logger.info('deleted name ${name}')
}

fn execute_gws_name_get(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!

	// assign default values
	if name == '' {
		logger.error('Please provide a name')
	}

	res := solution_handler.tfclient.gateways_get_name(name)!

	logger.info('${res}')
}
