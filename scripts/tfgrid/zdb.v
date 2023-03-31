module tfgrid

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

pub fn zdb_deploy(mut client RpcWsClient, params ZDB) !ZDBResult {
	return client.send_json_rpc[ZDB, ZDBResult]('tfgrid.zdb.deploy', params, default_timeout)!
}

pub fn zdb_delete(mut client RpcWsClient, params string) ! {
	_ := client.send_json_rpc[string, string]('tfgrid.zdb.delete', params, default_timeout)!
}

pub fn zdb_get(mut client RpcWsClient, params string) !ZDBResult {
	return client.send_json_rpc[string, ZDBResult]('tfgrid.zdb.get', params, default_timeout)!
}