module funkwhale

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

[params]
pub struct Deploy {
pub:
	name        string // identifier for the instance, must be unique
	farm_id     u64    // farm id to deploy on, if 0, a random eligible node on a random farm will be selected
	capacity    string // capacity of the instance. one of small, medium, large, extra-large
	ssh_key     string // public ssh key to access the instance in a later stage
	public_ipv6 bool   // if true, a public ipv6 will be added to the instance
	// admin configuration
	admin_email    string [required] // admin email to access admin dashboard
	admin_username string // admin username to access admin dashboard
	admin_password string // admin password to access admin dashboard
}

// PeerTubeClient is a client containig an RpcWsClient instance, and provides access for peertube applications on tfgrid
[openrpc: exclude]
pub struct FunkwhaleClient {
mut:
	client  &RpcWsClient
	timeout int
}

// Deploys a funkwhale instance
pub fn (mut t FunkwhaleClient) deploy(args Deploy) !FunkwhaleResult {
	return t.client.send_json_rpc[[]Deploy, FunkwhaleResult]('tfgrid.DeployFunkwhale',
		[args], t.timeout)!
}

// Gets a deployed funkwhale instance
pub fn (mut t FunkwhaleClient) get(funkwhale_name string) !FunkwhaleResult {
	return t.client.send_json_rpc[[]string, FunkwhaleResult]('tfgrid.GetFunkwhale', [
		funkwhale_name,
	], t.timeout)!
}

// Deletes a deployed funkwhale instance.
pub fn (mut t FunkwhaleClient) delete(funkwhale_name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.DeleteFunkwhale', [
		funkwhale_name,
	], t.timeout)!
}
