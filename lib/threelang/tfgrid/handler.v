module tfgrid

import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.explorer
import threefoldtech.threebot.tfgrid.solution { SolutionHandler }
import freeflowuniverse.crystallib.actionsparser { Action }
import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import log {Logger}

pub struct TFGridHandler {
pub mut:
	solution_handler SolutionHandler
	ssh_keys         map[string]string
	logger Logger
}

pub fn new(mut rpc_client RpcWsClient, logger Logger) TFGridHandler {
	mut tfgrid_client := tfgrid.new(mut rpc_client)
	mut exp := explorer.new(mut rpc_client)

	solution_handler := SolutionHandler{
		tfclient: tfgrid_client
		explorer: exp
	}

	return TFGridHandler{
		solution_handler: solution_handler
		logger: logger
	}
}

pub fn (mut t TFGridHandler) handle_action(action Action) ! {
	match action.actor {
		'core' {
			t.core(action)!
		}
		'gateway_fqdn' {
			t.gateway_fqdn(action)!
		}
		'gateway_name' {
			t.gateway_name(action)!
		}
		'kubernetes' {
			t.k8s(action)!
		}
		'machines' {
			t.vm(action)!
		}
		'zdbs' {
			t.zdb(action)!
		}
		'discourse'{
			t.discourse(action)!
		}
		'funkwhale'{
			t.funkwhale(action)!
		}
		'peertube'{
			t.peertube(action)!
		}
		'taiga'{
			t.taiga(action)!
		}
		'presearch'{
			t.presearch(action)!
		}
		else {
			t.helper(action)!
		}
	}
}
