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

[params]
pub struct BridgeTransfer {
	amount string
	destination string
}

[params]
pub struct TfchainBridgeTransfer {
	amount string
	destination string
	twin_id u32
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
	_ := s.client.send_json_rpc[[]Load, string]('stellar.Load', [args], default_timeout)!
}

// Transer an amount of TFT from the loaded account to the destination.
pub fn (mut s StellarClient) transer(args Transfer) ! {
	_ := s.client.send_json_rpc[[]Transfer, string]('stellar.Transfer', [args], default_timeout)!
}

// Balance of an account for TFT on stellar.
pub fn (mut s StellarClient) balance(address string) ! i64 {
	return s.client.send_json_rpc[[]string, i64]('stellar.Balance', [address], default_timeout)!
}

// bridge_to_eth bridge to eth from stellar
pub fn (mut s StellarClient) bridge_to_eth(args BridgeTransfer) ! {
	_ := s.client.send_json_rpc[[]BridgeTransfer, string]('stellar.BridgeToEth', [args], default_timeout)!
}

// Reinstate later

// // bridge_to_bsc bridge to bsc from stellar
// pub fn (mut s StellarClient) bridge_to_bsc(args BridgeTransfer) ! {
// 	_ := s.client.send_json_rpc[[]BridgeTransfer, string]('stellar.BridgeToBsc', [args], default_timeout)!
// }

// bridge_to_tfchain bridge to tfchain from stellar
pub fn (mut s StellarClient) bridge_to_tfchain(args TfchainBridgeTransfer) ! {
 	_ := s.client.send_json_rpc[[]TfchainBridgeTransfer, string]('stellar.BridgeToTfchain', [args], default_timeout)!
}