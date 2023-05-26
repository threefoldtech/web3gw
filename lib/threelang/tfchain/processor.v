module tfchain

import freeflowuniverse.crystallib.params { Params }
import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

// ChainProcessor should handle all tfchain related actions
pub struct ChainProcessor {
}

fn (c ChainProcessor) add_action(ns string, op string, action_params Params) ! {
}

fn (c ChainProcessor) execute(rpc_client &RpcWsClient) ! {
}
