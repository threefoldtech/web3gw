module peertube

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

[params]
pub struct Deploy {
pub:
	name        string // identifier for the instance, must be unique
	farm_id     u64    // farm id to deploy on, if 0, a random eligible farm will be selected
	capacity    string // capacity of the instance. one of small, medium, large, extra-large
	ssh_key     string // public ssh key to access the instance in a later stage
	public_ipv6 bool   // if true, a public ipv6 will be added to the instance
	// database configs
	db_username string // database username
	db_password string // database password

	admin_email string // admin email
}

// PeerTubeClient is a client containig an RpcWsClient instance, and provides access for peertube applications on tfgrid
[openrpc: exclude]
pub struct PeerTubeClient {
mut:
	client  &RpcWsClient
	timeout int
}

// Deploys a peertube instance
pub fn (mut t PeerTubeClient) deploy(args Deploy) !PeertubeResult {
	return t.client.send_json_rpc[[]Deploy, PeertubeResult]('tfgrid.DeployPeertube', [
		args,
	], t.timeout)!
}

// Gets a deployed peertube instance
pub fn (mut t PeerTubeClient) get(peertube_name string) !PeertubeResult {
	return t.client.send_json_rpc[[]string, PeertubeResult]('tfgrid.GetPeertube', [
		peertube_name,
	], t.timeout)!
}

// Deletes a deployed peertube instance.
pub fn (mut t PeerTubeClient) delete(peertube_name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.DeletePeertube', [
		peertube_name,
	], t.timeout)!
}
