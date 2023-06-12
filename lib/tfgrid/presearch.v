module tfgrid

[params]
pub struct Presearch {
pub:
	name        string
	farm_id     u64
	disk_size   u32 // in giga bytes
	ssh_key     string
	public_ipv4 bool
	public_ipv6 bool
}

// Deploys a presearch instance
pub fn (mut t TFGridClient) deploy_presearch(presearch Presearch) !PresearchResult {
	return t.client.send_json_rpc[[]Presearch, PresearchResult]('tfgrid.DeployPresearch',
		[presearch], default_timeout)!
}

// Gets a deployed presearch instance
pub fn (mut t TFGridClient) get_presearch(presearch_name string) !PresearchResult {
	return t.client.send_json_rpc[[]string, PresearchResult]('tfgrid.GetPresearch', [
		presearch_name,
	], default_timeout)!
}

// Deletes a deployed presearch instance.
pub fn (mut t TFGridClient) delete_presearch(presearch_name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.DeletePresearch', [
		presearch_name,
	], default_timeout)!
}
