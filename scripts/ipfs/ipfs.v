module ipfs

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

const (
	default_timeout = 500000
)

[noinit]
pub struct IpfsClient {
mut:
	client &RpcWsClient
}

pub fn new(mut client RpcWsClient) IpfsClient {
	return IpfsClient{
		client: &client
	}
}

pub fn (mut e IpfsClient) store_file(content []byte) !string {
	return e.client.send_json_rpc[[][]byte, string]('ipfs.StoreFile', [content], ipfs.default_timeout)!
}

pub fn (mut e IpfsClient) get_file(cid string) ![]byte {
	return e.client.send_json_rpc[[]string, []byte]('ipfs.GetFile', [cid], ipfs.default_timeout)!
}

// remove file from ipfs
pub fn (mut e IpfsClient) remove_file(cid string) !bool {
	return e.client.send_json_rpc[[]string, bool]('ipfs.RemoveFile', [cid], ipfs.default_timeout)!
}