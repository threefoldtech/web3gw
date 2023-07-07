module taiga

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

[params]
pub struct Deploy {
pub:
	name        string // identifier for the instance, must be unique
	farm_id     u64    // farm id to deploy on, if 0, a random eligible node on a random farm will be selected
	capacity    string // capacity of the instance. one of small, medium, large, extra-large
	disk_size   u32    // size of disk to mount on instance. must be in GB
	ssh_key     string // public ssh key to access the instance in a later stage
	public_ipv6 bool   // if true, a public ipv6 will be added to the instance
	// admin configs
	admin_username string // admin username
	admin_password string // admin password
	admin_email    string // admin email
}

// PeerTubeClient is a client containig an RpcWsClient instance, and provides access for peertube applications on tfgrid
[openrpc: exclude]
pub struct TaigaClient {
mut:
	client  &RpcWsClient
	timeout int
}

// Deploys a taiga instance
pub fn (mut t TaigaClient) deploy(args Deploy) !TaigaResult {
	return t.client.send_json_rpc[[]Deploy, TaigaResult]('tfgrid.DeployTaiga', [args],
		t.timeout)!
}

// Gets a deployed taiga instance
pub fn (mut t TaigaClient) get(taiga_name string) !TaigaResult {
	return t.client.send_json_rpc[[]string, TaigaResult]('tfgrid.GetTaiga', [
		taiga_name,
	], t.timeout)!
}

// Deletes a deployed taiga instance.
pub fn (mut t TaigaClient) delete(taiga_name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.DeleteTaiga', [
		taiga_name,
	], t.timeout)!
}
