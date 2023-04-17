module tfgrid

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

pub struct Credentials {
	mnemonic string
	network  string
}

struct Login {
	RpcWsClient
}

pub fn (mut client Login) load(credentials Credentials) ! {
	_ := client.send_json_rpc[[]string, string]('tfgrid.Load', [credentials.mnemonic, credentials.network],
		default_timeout)!
}
