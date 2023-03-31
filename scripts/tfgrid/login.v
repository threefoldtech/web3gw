module tfgrid

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

pub struct Credentials {
	mnemonic string
	network   string
}

pub fn load(mut client RpcWsClient, params Credentials) ! {
	_ := client.send_json_rpc[[]string, string]('tfgrid.Load', [params.mnemonic, params.network], default_timeout)!
}
