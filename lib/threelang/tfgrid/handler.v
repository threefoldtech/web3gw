module tfgrid

import threefoldtech.threebot.tfgrid { TFGridClient }
import freeflowuniverse.crystallib.actionsparser { Action }
import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import log { Logger }

pub struct TFGridHandler {
pub mut:
	tfgrid TFGridClient
	ssh_keys map[string]string
	logger   Logger
}

pub fn new(mut rpc_client RpcWsClient, logger Logger) TFGridHandler {
	mut tfgrid_client := tfgrid.new(mut rpc_client)

	return TFGridHandler{
		tfgrid: tfgrid_client
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
		'discourse' {
			t.discourse(action)!
		}
		'funkwhale' {
			t.funkwhale(action)!
		}
		'peertube' {
			t.peertube(action)!
		}
		'taiga' {
			t.taiga(action)!
		}
		'presearch' {
			t.presearch(action)!
		}
		else {
			t.helper(action)!
		}
	}
}
