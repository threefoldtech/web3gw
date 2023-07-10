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

pub fn new(mut rpc_client RpcWsClient, logger Logger, mut client EthClient) EthHandler {
	return EthHandler{
		client: client
		logger: logger
	}
}

pub fn (mut h EthHandler) handle_action(action Action) ! {
	match action.actor {
		'account' {
			h.account(action)!
		}
		'transfer' {
			h.transfer(action)!
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
		'fungible' {
			h.fungible(action)!
		}
		'token' {
			h.token(action)!
		}
		else {
			return error('actor ${action.actor} is invalid')
		}
	}
}
