module tfchain

import freeflowuniverse.crystallib.actionsparser{Action}
import freeflowuniverse.crystallib.rpcwebsocket {RpcWsClient}

pub struct TFChainHandler{
	
}

pub fn new(mut rpc_client &RpcWsClient) TFChainHandler{
	return TFChainHandler{}
}

pub fn(mut t TFChainHandler) handle_action(action Action) !{
	
}