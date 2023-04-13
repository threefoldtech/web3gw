module tfgrid

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

pub fn zdb_deploy(mut client RpcWsClient, params ZDB) !ZDBResult {
	return client.send_json_rpc[[]ZDB, ZDBResult]('tfgrid.ZDBDeploy', [params], default_timeout)!
}

pub fn zdb_delete(mut client RpcWsClient, params string) ! {
	_ := client.send_json_rpc[[]string, string]('tfgrid.ZDBDelete', [params], default_timeout)!
}

pub fn zdb_get(mut client RpcWsClient, params string) !ZDBResult {
	return client.send_json_rpc[[]string, ZDBResult]('tfgrid.ZDBGet', [params], default_timeout)!
}