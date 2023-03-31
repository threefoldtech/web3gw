module tfgrid

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

pub struct Credentials {
	mnemonic string
	network   string
}

pub fn login(mut client RpcWsClient, params Credentials) ! {
	_ := client.send_json_rpc[Credentials, string]('tfgrid.login', params, default_timeout)!
}
