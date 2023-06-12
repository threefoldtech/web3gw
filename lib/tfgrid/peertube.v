module tfgrid

[params]
pub struct Peertube {
pub:
	name          string
	farm_id       u64
	capacity      string
	ssh_key       string
	db_username   string
	db_password   string
	admin_email   string
	smtp_hostname string
	smtp_username string
	smtp_password string
	public_ipv6   bool
}

// Deploys a peertube instance
pub fn (mut t TFGridClient) deploy_peertube(peertube Peertube) !PeertubeResult {
	return t.client.send_json_rpc[[]Peertube, PeertubeResult]('tfgrid.DeployPeertube',
		[peertube], default_timeout)!
}

// Gets a deployed peertube instance
pub fn (mut t TFGridClient) get_peertube(peertube_name string) !PeertubeResult {
	return t.client.send_json_rpc[[]string, PeertubeResult]('tfgrid.GetPeertube', [
		peertube_name,
	], default_timeout)!
}

// Deletes a deployed peertube instance.
pub fn (mut t TFGridClient) delete_peertube(peertube_name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.DeletePeertube', [
		peertube_name,
	], default_timeout)!
}
