module explorer

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import freeflowuniverse.crystallib.actionparser { Action }
import threefoldtech.threebot.explorer { ExplorerClient }
import log { Logger }


pub struct ExplorerHandler {
pub mut:
	explorer ExplorerClient
	logger Logger
}

pub fn new(explorer ExplorerClient, logger Logger) ExplorerHandler {
	ExplorerHandler {
		explorer: explorer,
		logger: logger,
	}
}

pub fn (mut h ExplorerHandler) handle_action(action Action) ! {
	match action.actor {
		'nodes' {
			t.nodes(action)!
		}
		'farms' {
			t.farms(action)!
		}
		'contracts' {
			t.contracts(action)!
		}
		'twins' {
			t.twins(action)!
		}
		'counters' {
			t.counters(action)!
		}
		else {
			return error!("unknown actor: ${action.actor}")
		}
	}
}