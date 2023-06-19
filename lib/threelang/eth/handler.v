module eth

import threefoldtech.threebot.eth as eth_client { EthClient }
import freeflowuniverse.crystallib.actionsparser { Action }
import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import log { Logger }

pub struct EthHandler {
pub mut:
	client EthClient
	logger Logger
}

pub fn new(mut rpc_client RpcWsClient, logger Logger) EthHandler {
	mut client := eth_client.new(mut rpc_client)

	return EthHandler{
		client: client
		logger: logger
	}
}

pub fn (mut h EthHandler) handle_action(action Action) ! {
	match action.actor {
		'core' {
			h.core(action)!
		}
		'account' {
			h.account(action)!
		}
		'transfer' {
			h.transfer(action)!
		}
		'balance' {
			h.balance(action)!
		}
		'block' {
			h.block(action)!
		}
		'multisig' {
			h.multisig(action)!
		}
		'swap' {
			h.swap(action)!
		}
		'bridge' {
			h.bridge(action)!
		}
		'fungible' {
			h.fungible(action)!
		}
		'token' {
			h.token(action)!
		}
		'tft' {
			h.tft(action)!
		}
		else {
			return error('actor ${action.actor} is invalid')
		}
	}
}
