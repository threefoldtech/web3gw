module tfgrid

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

pub struct Credentials {
	mnemonic string
	network  string
}

pub fn load(mut client RpcWsClient, credentials Credentials) ! {
	_ := client.send_json_rpc[[]string, string]('tfgrid.Load', [credentials.mnemonic, credentials.network],
		default_timeout)!
}
