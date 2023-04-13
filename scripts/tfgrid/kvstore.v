module tfgrid

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

// NOTE: this is related to tfchain not tfgrid + not implemented

// Set a new record in my kvstore as key and value, if success return account_id
pub fn kvstore_set(mut client RpcWsClient, key string, value string) !string {
	return client.send_json_rpc[[]string, string]('tfgrid.KvstoreSet', [key, value], default_timeout)!
}

// Get a record from my kvstore using key
pub fn kvstore_get(mut client RpcWsClient, params string) !string {
	return client.send_json_rpc[[]string, string]('tfgrid.KvstoreGet', [params], default_timeout)!
}

// List all keys in my kvstore
pub fn kvstore_list(mut client RpcWsClient) ![]string {
	return client.send_json_rpc[string, []string]('tfgrid.KvstoreList', '', default_timeout)!
}

// Remove a record from my kvstore using key, if success return account_id
pub fn kvstore_remove(mut client RpcWsClient, params string) !string {
	return client.send_json_rpc[[]string, string]('tfgrid.KvstoreRemove', [params], default_timeout)!
}

// Remove all my records in my kvstore, if success return deleted Keys
pub fn kvstore_remove_all(mut client RpcWsClient) ![]string {
	return client.send_json_rpc[string, []string]('tfgrid.KvstoreRemoveall', '', default_timeout)!
}
