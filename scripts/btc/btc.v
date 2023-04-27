module btc

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

const (
	default_timeout = 500000
)

[noinit]
pub struct BtcClient {
mut:
	client &RpcWsClient
}

pub fn new(mut client RpcWsClient) BtcClient {
	return BtcClient{
		client: &client
	}
}

// Needed calls
// Load(ctx context.Context, conState jsonrpc.State, args Load)
// CreateNewAccount(ctx context.Context, conState jsonrpc.State, account string) 
// CreateEncryptedWallet(ctx context.Context, conState jsonrpc.State, passphrase string) 
// ImportAddress(ctx context.Context, conState jsonrpc.State, address string) 
// ImportAddressRescan(ctx context.Context, conState jsonrpc.State, args ImportAddressRescan) 
// ImportPrivKey(ctx context.Context, conState jsonrpc.State, wif string) 
// ImportPrivKeyLabel(ctx context.Context, conState jsonrpc.State, args ImportPrivKeyLabel)
// ImportPrivKeyRescan(ctx context.Context, conState jsonrpc.State, args ImportPrivKeyRescan) 

// ImportPubKey(ctx context.Context, conState jsonrpc.State, pubKey string) 
// ImportPubKeyRescan(ctx context.Context, conState jsonrpc.State, args ImportPubKeyRescan) 
// InvalidateBlock(ctx context.Context, conState jsonrpc.State, hash string) 
// RenameAccount(ctx context.Context, conState jsonrpc.State, args RenameAccount)
// SubmitBlock(ctx context.Context, conState jsonrpc.State, args SubmitBlock)
// SendToAddress(ctx context.Context, conState jsonrpc.State, args SendToAddress)
// SendToAddressComment(ctx context.Context, conState jsonrpc.State, args SendToAddress) 


pub fn (mut e BtcClient) load(params Load) !string {
	return e.client.send_json_rpc[[]Load, string]('btc.Load', [
		params,
	], explorer.default_timeout)!
}

pub fn (mut e BtcClient) create_new_account(account string) ! {
	return e.client.send_json_rpc[[]string, string]('btc.CreateNewAccount', [
		account,
	], explorer.default_timeout)!
}

pub fn (mut e BtcClient) create_encrypted_wallet(passphrase string) ! {
	return e.client.send_json_rpc[[]string, string]('btc.CreateEncryptedWallet', [
		passphrase,
	], explorer.default_timeout)!
}

pub fn (mut e BtcClient) import_address(address string) ! {
	return e.client.send_json_rpc[[]string, string]('btc.ImportAddress', [
		address,
	], explorer.default_timeout)!
}

pub fn (mut e BtcClient) import_address_rescan(args ImportAddressRescan) ! {
	return e.client.send_json_rpc[[]ImportAddressRescan, string]('btc.ImportAddressRescan', [
		args,
	], explorer.default_timeout)!
}

pub fn (mut e BtcClient) import_priv_key(wif string) ! {
	return e.client.send_json_rpc[[]string, string]('btc.ImportPrivKey', [
		wif,
	], explorer.default_timeout)!
}

pub fn (mut e BtcClient) import_priv_key_label(args ImportPrivKeyLabel) ! {
	return e.client.send_json_rpc[[]string, string]('btc.ImportPrivKeyLabel', [
		args,
	], explorer.default_timeout)!
}

pub fn (mut e BtcClient) import_priv_key_rescan(args ImportPrivKeyRescan) ! {
	return e.client.send_json_rpc[[]string, string]('btc.ImportPrivKeyRescan', [
		args,
	], explorer.default_timeout)!
}