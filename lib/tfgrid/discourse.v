module tfgrid

[params]
pub struct Discourse {
pub:
	name            string
	farm_id         u64
	capacity        string
	disk_size       u32 // in giga bytes
	ssh_key         string
	developer_email string
	smtp_username   string
	smtp_password   string
	smtp_address    string = 'smtp.gmail.com'
	smtp_enable_tls bool   = true
	smtp_port       u32    = 587
	public_ipv6     bool
}

// Deploys a discourse instance
pub fn (mut t TFGridClient) deploy_discourse(discourse Discourse) !DiscourseResult {
	return t.client.send_json_rpc[[]Discourse, DiscourseResult]('tfgrid.DeployDiscourse',
		[discourse], default_timeout)!
}

// Gets a deployed discourse instance
pub fn (mut t TFGridClient) get_discourse(discourse_name string) !DiscourseResult {
	return t.client.send_json_rpc[[]string, DiscourseResult]('tfgrid.GetDiscourse', [
		discourse_name,
	], default_timeout)!
}

// Deletes a deployed discourse instance.
pub fn (mut t TFGridClient) delete_discourse(discourse_name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.DeleteDiscourse', [
		discourse_name,
	], default_timeout)!
}
