module tfgrid

pub fn (mut client TFGridClient) zdb_deploy(model ZDB) !ZDBResult {
	return client.send_json_rpc[[]ZDB, ZDBResult]('tfgrid.ZDBDeploy', [model], default_timeout)!
}

pub fn (mut client TFGridClient) zdb_delete(model_name string) ! {
	_ := client.send_json_rpc[[]string, string]('tfgrid.ZDBDelete', [model_name], default_timeout)!
}

pub fn (mut client TFGridClient) zdb_get(model_name string) !ZDBResult {
	return client.send_json_rpc[[]string, ZDBResult]('tfgrid.ZDBGet', [model_name], default_timeout)!
}
