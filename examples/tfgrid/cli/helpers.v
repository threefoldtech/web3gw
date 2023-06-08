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
