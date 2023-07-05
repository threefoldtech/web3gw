module discourse

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

[params]
pub struct Deploy {
pub:
	name            string // identifier for the instance, must be unique
	farm_id         u64    // farm id to deploy on, if 0, a random eligible node on a random farm will be selected
	capacity        string // capacity of the instance. one of small, medium, large, extra-large
	disk_size       u32    // size of disk to mount on instance. must be in GB
	ssh_key         string // public ssh key to access the instance in a later stage
	developer_email string // developer email to get development emails, only works if smtp is configured
	public_ipv6     bool   // if true, a public ipv6 will be added to the instance
	// smtp server configurations
	smtp_username   string // smtp server username
	smtp_password   string // smtp server password
	smtp_address    string = 'smtp.gmail.com' // smtp server domain address
	smtp_port       u32    = 587 // smtp server port
	smtp_enable_tls bool   // if true, tls encryption will be used in the smtp server
}

// DiscourseClient is a client containig an RpcWsClient instance, and provides access for discourse applications on tfgrid
[openrpc: exclude]
pub struct DiscourseClient {
mut:
	client  &RpcWsClient
	timeout int
}

// Deploys a discourse instance
pub fn (mut t DiscourseClient) deploy(args Deploy) !DiscourseResult {
	return t.client.send_json_rpc[[]Deploy, DiscourseResult]('tfgrid.DeployDiscourse',
		[args], t.timeout)!
}

// Gets a deployed discourse instance
pub fn (mut t DiscourseClient) get(discourse_name string) !DiscourseResult {
	return t.client.send_json_rpc[[]string, DiscourseResult]('tfgrid.GetDiscourse', [
		discourse_name,
	], t.timeout)!
}

// Deletes a deployed discourse instance.
pub fn (mut t DiscourseClient) delete(discourse_name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.DeleteDiscourse', [
		discourse_name,
	], t.timeout)!
}
