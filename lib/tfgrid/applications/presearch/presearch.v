module presearch

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

[params]
pub struct Deploy {
pub:
	name              string [required] // identifier for the instance, must be unique
	farm_id           u64    // farm id to deploy on, if 0, a random eligible node on a random farm will be selected
	disk_size         u32    // size of disk to mount on instance. must be in GB
	ssh_key           string // public ssh key to access the instance in a later stage
	registration_code string [required] // You need to sign up on Presearch in order to get your Presearch Registration Code.
	public_ipv4       bool // if true, a public ipv4 will be added to the instance
	public_ipv6       bool // if true, a public ipv6 will be added to the instance
	// presearch config for restoring old nodes
	public_restore_key  string
	private_restore_key string
}

// PresearchClient is a client containig an RpcWsClient instance, and provides access for presearch applications on tfgrid
[openrpc: exclude]
pub struct PresearchClient {
mut:
	client  &RpcWsClient
	timeout int
}

// Deploys a presearch instance
pub fn (mut t PresearchClient) deploy(args Deploy) !PresearchResult {
	return t.client.send_json_rpc[[]Deploy, PresearchResult]('tfgrid.DeployPresearch',
		[args], t.timeout)!
}

// Gets a deployed presearch instance
pub fn (mut t PresearchClient) get(presearch_name string) !PresearchResult {
	return t.client.send_json_rpc[[]string, PresearchResult]('tfgrid.GetPresearch', [
		presearch_name,
	], t.timeout)!
}

// Deletes a deployed presearch instance.
pub fn (mut t PresearchClient) delete(presearch_name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.DeletePresearch', [
		presearch_name,
	], t.timeout)!
}
