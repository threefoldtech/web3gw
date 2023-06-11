module tfgrid

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

// TFGridClient is a client containig an RpcWsClient instance, and implements all tfgrid functionality
[openrpc: exclude]
[noinit]
pub struct TFGridClient {
mut:
	client &RpcWsClient
}

[openrpc: exclude]
pub fn new(mut client RpcWsClient) TFGridClient {
	return TFGridClient{
		client: &client
	}
}
