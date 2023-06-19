module tfchain

import threefoldtech.threebot.tfchain { TfChainClient }
import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import freeflowuniverse.crystallib.actionsparser { Action }
import log { Logger } 

pub struct TFChainHandler {
pub mut:
	tfchain TfChainClient
	logger Logger
}

pub fn new(mut rpc_client RpcWsClient, logger Logger, mut chain_client TfChainClient) TFChainHandler {
	return TFChainHandler{
		tfchain: chain_client
		logger: logger
	}
}

pub fn (mut t TFChainHandler) handle_action(action Action) ! {
	match action.actor {
		'account' {
			t.account(action)!
		}
		'money' {
			t.money(action)!
		}
		'contracts' {
			t.contracts(action)!
		}
		'service_contract' {
			t.service_contract(action)!
		}
		'metadata' {
			t.metadata(action)!
		}
		'client' {
			t.client(action)!
		}
		'farms' {
			t.farms(action)!
		}
		'nodes' {
			t.nodes(action)!
		}
		'twins' {
			t.twins(action)!
		} else {
			return error('actor ${action.actor} is invalid')
		}
	}
}