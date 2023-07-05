module applications

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import threefoldtech.threebot.tfgrid.applications.peertube
import threefoldtech.threebot.tfgrid.applications.discourse
import threefoldtech.threebot.tfgrid.applications.funkwhale
import threefoldtech.threebot.tfgrid.applications.presearch
import threefoldtech.threebot.tfgrid.applications.taiga

// ApplicationsClient is a client containig an RpcWsClient instance, and provides access to tfgrid application clients
[openrpc: exclude]
pub struct ApplicationsClient {
mut:
	client  &RpcWsClient
	timeout int
}

// peertube returns a client that provides access to peertube applications for the current tfgrid instance
[openrpc: exclude]
pub fn (t ApplicationsClient) peertube() peertube.PeerTubeClient {
	return peertube.PeerTubeClient{
		client: t.client
		timeout: t.timeout
	}
}

// discourse returns a client that provides access to discourse applications for the current tfgrid instance
[openrpc: exclude]
pub fn (t ApplicationsClient) discourse() discourse.DiscourseClient {
	return discourse.DiscourseClient{
		client: t.client
		timeout: t.timeout
	}
}

// funkwhale returns a client that provides access to funkwhale applications for the current tfgrid instance
[openrpc: exclude]
pub fn (t ApplicationsClient) funkwhale() funkwhale.FunkwhaleClient {
	return funkwhale.FunkwhaleClient{
		client: t.client
		timeout: t.timeout
	}
}

// presearch returns a client that provides access to presearch applications for the current tfgrid instance
[openrpc: exclude]
pub fn (t ApplicationsClient) presearch() presearch.PresearchClient {
	return presearch.PresearchClient{
		client: t.client
		timeout: t.timeout
	}
}

// taiga returns a client that provides access to taiga applications for the current tfgrid instance
[openrpc: exclude]
pub fn (t ApplicationsClient) taiga() taiga.TaigaClient {
	return taiga.TaigaClient{
		client: t.client
		timeout: t.timeout
	}
}
