module tfgrid

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

// TFGridClient is a client containig an RpcWsClient instance, and implements all tfgrid functionality
pub struct TFGridClient {
	RpcWsClient // RpcWsClient instance to talk to the rpc server
}
