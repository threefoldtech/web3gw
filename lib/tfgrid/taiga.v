module tfgrid

[params]
pub struct Taiga {
pub:
	name           string
	farm_id        u64
	capacity       string
	disk_size      u32 // in giga bytes
	ssh_key        string
	admin_username string
	admin_password string
	admin_email    string
	public_ipv6    bool
}

// Deploys a taiga instance
pub fn (mut t TFGridClient) deploy_taiga(taiga Taiga) !TaigaResult {
	return t.client.send_json_rpc[[]Taiga, TaigaResult]('tfgrid.DeployTaiga', [taiga],
		default_timeout)!
}

// Gets a deployed taiga instance
pub fn (mut t TFGridClient) get_taiga(taiga_name string) !TaigaResult {
	return t.client.send_json_rpc[[]string, TaigaResult]('tfgrid.GetTaiga', [
		taiga_name,
	], default_timeout)!
}

// Deletes a deployed taiga instance.
pub fn (mut t TFGridClient) delete_taiga(taiga_name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.DeleteTaiga', [
		taiga_name,
	], default_timeout)!
}
