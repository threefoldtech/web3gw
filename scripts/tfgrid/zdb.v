module tfgrid

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

pub fn zdb_deploy(mut client RpcWsClient, model ZDB) !ZDBResult {
	return client.send_json_rpc[[]ZDB, ZDBResult]('tfgrid.ZDBDeploy', [model], default_timeout)!
}

pub fn zdb_delete(mut client RpcWsClient, model_name string) ! {
	_ := client.send_json_rpc[[]string, string]('tfgrid.ZDBDelete', [model_name], default_timeout)!
}

pub fn zdb_get(mut client RpcWsClient, model_name string) !ZDBResult {
	return client.send_json_rpc[[]string, ZDBResult]('tfgrid.ZDBGet', [model_name], default_timeout)!
}