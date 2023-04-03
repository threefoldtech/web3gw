module stellar

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

const (
	default_timeout = 500000
)

pub fn load(mut client RpcWsClient, secret string, network string) ! {
	_ := client.send_json_rpc[[]string, string]('stellar.Load', [network, secret], default_timeout)!
}

pub fn transer(mut client RpcWsClient, amount string, destination string, memo string) ! {
	_ := client.send_json_rpc[[]string, string]('stellar.Transfer', [amount, destination, memo], default_timeout)!
}

pub fn balance(mut client RpcWsClient, address string) ! i64 {
	return client.send_json_rpc[[]string, i64]('stellar.Balance', [address], default_timeout)!
}