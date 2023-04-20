module tfgrid

// Deploys a zdb workload on the grid given a model. This call returns the zdb model with extra computed 
// data from the grid upon success. 
pub fn (mut t TFGridClient) zdb_deploy(model ZDB) !ZDBResult {
	return t.client.send_json_rpc[[]ZDB, ZDBResult]('tfgrid.ZDBDeploy', [model], default_timeout)!
}

// Deletes a deployed zdb workload given its name. Returns an error if it does not succeed. 
pub fn (mut t TFGridClient) zdb_delete(model_name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.ZDBDelete', [model_name], default_timeout)!
}

// Gets a deployed zdb deployment info given its configuration name. Returns the zdb deployment data 
// upon success.
pub fn (mut t TFGridClient) zdb_get(model_name string) !ZDBResult {
	return t.client.send_json_rpc[[]string, ZDBResult]('tfgrid.ZDBGet', [model_name], default_timeout)!
}
