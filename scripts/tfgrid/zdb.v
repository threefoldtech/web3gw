module tfgrid

// zdb_deploy deploys a zdb workload on the grid
// - model: a zdb model with the required info
// returns the deployed zdb model with the computed fileds from the grid
pub fn (mut client TFGridClient) zdb_deploy(model ZDB) !ZDBResult {
	return client.send_json_rpc[[]ZDB, ZDBResult]('tfgrid.ZDBDeploy', [model], default_timeout)!
}

// zdb_delete delete a deployed zdb workload
// - model_name: the zdb deployment name
pub fn (mut client TFGridClient) zdb_delete(model_name string) ! {
	_ := client.send_json_rpc[[]string, string]('tfgrid.ZDBDelete', [model_name], default_timeout)!
}

// zdb_get get a deployed zdb deployment info
// - model_name: the zdb deployment name
// returns the deployed zdb deployemnt
pub fn (mut client TFGridClient) zdb_get(model_name string) !ZDBResult {
	return client.send_json_rpc[[]string, ZDBResult]('tfgrid.ZDBGet', [model_name], default_timeout)!
}
