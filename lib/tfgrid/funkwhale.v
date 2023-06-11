module tfgrid

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
