module tfgrid

[params]
pub struct Funkwhale {
pub:
	name           string
	farm_id        u64
	capacity       string
	ssh_key        string
	admin_email    string
	admin_username string
	admin_password string
	public_ipv6    bool
}

// Deploys a funkwhale instance
pub fn (mut t TFGridClient) deploy_funkwhale(funkwhale Funkwhale) !FunkwhaleResult {
	return t.client.send_json_rpc[[]Funkwhale, FunkwhaleResult]('tfgrid.DeployFunkwhale',
		[funkwhale], default_timeout)!
}

// Gets a deployed funkwhale instance
pub fn (mut t TFGridClient) get_funkwhale(funkwhale_name string) !FunkwhaleResult {
	return t.client.send_json_rpc[[]string, FunkwhaleResult]('tfgrid.GetFunkwhale', [
		funkwhale_name,
	], default_timeout)!
}

// Deletes a deployed funkwhale instance.
pub fn (mut t TFGridClient) delete_funkwhale(funkwhale_name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.DeleteFunkwhale', [
		funkwhale_name,
	], default_timeout)!
}
