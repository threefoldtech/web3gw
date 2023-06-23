module stellar

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

const (
	default_timeout = 500000
)

[params]
pub struct Load {
	network string = 'public'
	secret  string
}

[params]
pub struct Swap {
	amount string
	source_asset string = "xlm"
	destination_asset string
}

[params]
pub struct Transfer {
	amount      string
	destination string
	memo        string
}

[params]
pub struct BridgeTransfer {
	amount      string
	destination string
}

[params]
pub struct TfchainBridgeTransfer {
	amount  string
	twin_id u32
}

[params]
pub struct Transactions {
	account string  // filter the transactions on the account with the address from this argument, leave empty for your account
	limit u32 // limit the amount of transactions to gather with this argument, this is 10 by default
	include_failed bool // include the failed arguments
	cursor string // list the last transactions starting from this cursor, leave empty to start from the top
	ascending bool // order the transactions in ascending order
}

[openrpc: exclude]
[noinit]
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
	_ := s.client.send_json_rpc[[]Load, string]('stellar.Load', [args], stellar.default_timeout)!
}

// Creates an account on the provided network and returns the seed. Consecutive calls will be using the newly created account.
pub fn (mut s StellarClient) create_account(network string) !string {
	return s.client.send_json_rpc[[]string, string]('stellar.CreateAccount', [network], stellar.default_timeout)!
}

// Get the public address of the loaded stellar secret
pub fn (mut s StellarClient) address() !string {
	return s.client.send_json_rpc[[]string, string]('stellar.Address', []string{}, default_timeout)!
}

// Swap tokens from one asset type to the other (for example from tft to xlm)
pub fn (mut s StellarClient) swap(args Swap) !string {
	return s.client.send_json_rpc[[]Swap, string]('stellar.Swap', [args], default_timeout)!
}

// Transfer an amount of TFT from the loaded account to the destination.
pub fn (mut s StellarClient) transfer(args Transfer) !string {
	return s.client.send_json_rpc[[]Transfer, string]('stellar.Transfer', [args], default_timeout)!
}

// Balance of an account for TFT on stellar.
pub fn (mut s StellarClient) balance(address string) !string {
	return s.client.send_json_rpc[[]string, string]('stellar.Balance', [address], default_timeout)!
}

// bridge_to_eth bridge to eth from stellar
pub fn (mut s StellarClient) bridge_to_eth(args BridgeTransfer) !string {
	return s.client.send_json_rpc[[]BridgeTransfer, string]('stellar.BridgeToEth', [args], default_timeout)!
}

// Reinstate later

// // bridge_to_bsc bridge to bsc from stellar
// pub fn (mut s StellarClient) bridge_to_bsc(args BridgeTransfer) ! {
// 	_ := s.client.send_json_rpc[[]BridgeTransfer, string]('stellar.BridgeToBsc', [args], default_timeout)!
// }

// bridge_to_tfchain bridge to tfchain from stellar
pub fn (mut s StellarClient) bridge_to_tfchain(args TfchainBridgeTransfer) !string {
 	return s.client.send_json_rpc[[]TfchainBridgeTransfer, string]('stellar.BridgeToTfchain', [args], default_timeout)!
}

// Await till a transaction is processed on ethereum bridge that contains a specific memo
pub fn (mut s StellarClient) await_transaction_on_eth_bridge(memo string) ! {
	_ := s.client.send_json_rpc[[]string, string]('stellar.AwaitTransactionOnEthBridge', [memo], default_timeout)!
}

// Return a limited amount of transactions bound to a specific account
pub fn (mut s StellarClient) transactions(args Transactions) ![]Transaction {
	return s.client.send_json_rpc[[]Transactions, []Transaction]('stellar.Transactions', [args], default_timeout)!
}

// Return the data that is related to an account
pub fn (mut s StellarClient) account_data(account string) !AccountData {
	return s.client.send_json_rpc[[]string, AccountData]('stellar.AccountData', [account], default_timeout)!
}