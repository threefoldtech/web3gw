module eth

import threefoldtech.threebot.eth { EthClient }
import freeflowuniverse.crystallib.actionsparser { Action }
import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import log { Logger }

pub struct EthHandler {
pub mut:
	client EthClient
	logger Logger
}

pub fn new(mut rpc_client RpcWsClient, logger Logger) EthHandler {
	mut eth_client := eth.new(mut rpc_client)

	return EthHandler{
		client: eth_client
		logger: logger
	}
}

pub fn (mut h EthHandler) handle_action(action Action) ! {
	match action.actor {
		'core' {
			h.core(action)!
		}
		else {
			return error('actor ${action.actor} is invalid')
		}
	}
}
