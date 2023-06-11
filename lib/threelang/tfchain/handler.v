module tfchain

import freeflowuniverse.crystallib.actionsparser { Action }
import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import log { Logger }

pub struct TFChainHandler {
}

pub fn new(mut rpc_client RpcWsClient, logger Logger) TFChainHandler {
	return TFChainHandler{}
}

pub fn (mut t TFChainHandler) handle_action(action Action) ! {
}
