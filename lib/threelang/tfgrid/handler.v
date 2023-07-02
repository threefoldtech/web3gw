module tfgrid

import threefoldtech.threebot.tfgrid as tfgrid_client { TFGridClient }
import threefoldtech.threebot.explorer { ExplorerClient }
import freeflowuniverse.crystallib.actionsparser { Action }
import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import log { Logger }

pub struct TFGridHandler {
pub mut:
	tfgrid   TFGridClient
	explorer ExplorerClient
	ssh_keys map[string]string
	logger   Logger
}

pub fn new(mut rpc_client RpcWsClient, logger Logger) TFGridHandler {
	mut grid_client := tfgrid_client.new(mut rpc_client)
	mut explorer_client := explorer.new(mut rpc_client)

	return TFGridHandler{
		tfgrid: grid_client
		explorer: explorer_client
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
		'machine' {
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
		'explorer'{
			t.explorer(action)!
		}
		else {
			t.helper(action)!
		}
	}
}
