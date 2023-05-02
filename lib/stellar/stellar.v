module stellar

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

const (
	default_timeout = 500000
)

[params]
pub struct Load {
	network string
	secret string
}

[params]
pub struct Transfer {
	amount string
	destination string
	memo string
}


[noinit; openrpc: exclude]
pub struct StellarClient {
mut:
	client &RpcWsClient
}

[openrpc: exclude]
pub fn new(mut client RpcWsClient) StellarClient {
	return StellarClient{
		client: &client
	}
}

// Load a client, connecting to the rpc endpoint at the given network and loading a keypair from the given secret.
pub fn (mut s StellarClient) load(args Load) ! {
	_ := client.send_json_rpc[[]Load, string]('stellar.Load', [args], default_timeout)!
}

// Transer an amount of TFT from the loaded account to the destination.
pub fn (mut s StellarClient) transer(args Transfer) ! {
	_ := client.send_json_rpc[[]Transfer, string]('stellar.Transfer', [args], default_timeout)!
}

// Balance of an account for TFT on stellar.
pub fn (mut s StellarClient) balance(address string) ! i64 {
	return client.send_json_rpc[[]string, i64]('stellar.Balance', [address], default_timeout)!
}