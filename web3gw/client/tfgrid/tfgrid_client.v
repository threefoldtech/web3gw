module tfgrid

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import threefoldtech.web3gw.tfgrid.applications

// TFGridClient is a client containig an RpcWsClient instance, and implements all tfgrid functionality
[openrpc: exclude]
[noinit]
pub struct TFGridClient {
mut:
	client  &RpcWsClient
	timeout int
}

[openrpc: exclude]
pub fn new(mut client RpcWsClient) TFGridClient {
	return TFGridClient{
		client: &client
		timeout: 500000
	}
}

// applications returns a client that provides access to applications for the current tfgrid instance
[openrpc: exclude]
pub fn (t TFGridClient) applications() applications.ApplicationsClient {
	return applications.ApplicationsClient{
		client: t.client
		timeout: t.timeout
	}
}
