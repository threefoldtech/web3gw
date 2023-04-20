module explorer

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

[noinit]
pub struct ExplorerClient {
mut:
	client &RpcWsClient
}

pub fn new(mut client RpcWsClient) ExplorerClient {
	return ExplorerClient{
		client: &client
	}
}
