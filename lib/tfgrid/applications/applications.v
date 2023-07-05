module applications

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import threefoldtech.threebot.tfgrid.applications.peertube

// ApplicationsClient is a client containig an RpcWsClient instance, and provides access to tfgrid application clients
[openrpc: exclude]
pub struct ApplicationsClient {
mut:
	client  &RpcWsClient
	timeout int
}

// applications returns a client that provides access to applications for the current tfgrid instance
[openrpc: exclude]
pub fn (t ApplicationsClient) peertube() peertube.PeerTubeClient {
	return peertube.PeerTubeClient{
		client: t.client
		timeout: t.timeout
	}
}
