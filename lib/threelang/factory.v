module threelang

import log

import freeflowuniverse.crystallib.actionsparser
import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }


import threefoldtech.threebot.tfgrid as tfgrid_client
import threefoldtech.threebot.tfchain as tfchain_client
import threefoldtech.threebot.stellar as stellar_client
import threefoldtech.threebot.eth as eth_client
import threefoldtech.threebot.btc as btc_client

import threefoldtech.threebot.threelang.tfgrid { TFGridHandler }
import threefoldtech.threebot.threelang.web3gw { Web3GWHandler }
import threefoldtech.threebot.threelang.clients { Clients }

const (
	tfgrid_book = 'tfgrid'
	web3gw_book  = 'web3gw'
)

pub struct Runner {
pub mut:
	path string
	clients &Clients
	tfgrid_handler TFGridHandler
	web3gw_handler Web3GWHandler
}

[params]
pub struct RunnerArgs {
pub mut:
	name    string
	path    string
	address string
}

pub fn new(args RunnerArgs, debug_log bool) !Runner {
	mut ap := actionsparser.new(path: args.path, defaultbook: 'aaa')!

	mut logger := log.Logger(&log.Log{
		level: if debug_log { .debug } else { .info }
	})

	mut rpc_client := rpcwebsocket.new_rpcwsclient(args.address, &logger) or {
		return error('Failed creating rpc websocket client: ${err}')
	}
	_ := spawn rpc_client.run()

	gw_clients := get_clients(mut rpc_client)!

	tfgrid_handler := tfgrid.new(mut rpc_client, logger)
	web3gw_handler := web3gw.new(mut rpc_client, logger, &gw_clients)

	mut runner := Runner{
		path: args.path
		tfgrid_handler: tfgrid_handler
		web3gw_handler: web3gw_handler
		clients: &gw_clients
	}

	runner.run(mut ap)!
	return runner
}

pub fn (mut r Runner) run(mut action_parser actionsparser.ActionsParser) ! {
	for action in action_parser.actions {
		match action.book {
			tfgrid_book {
				r.tfgrid_handler.handle_action(action)!
			}
			web3gw_book {
				r.web3gw_handler.handle(action)!
			}
			else {
				return error('module ${action.book} is invalid')
			}
		}
	}
}

pub fn get_clients(mut rpc_client &RpcWsClient) !Clients {
	return Clients{
		tfg_client: tfgrid_client.new(mut rpc_client)
		tfc_client: tfchain_client.new(mut rpc_client)
		btc_client: btc_client.new(mut rpc_client)
		eth_client: eth_client.new(mut rpc_client)
		str_client: stellar_client.new(mut rpc_client)
	}
}