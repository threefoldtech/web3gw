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

pub fn (mut b BtcClient) import_pub_key(pub_key string) ! {
	b.client.send_json_rpc[[]string, string]('btc.ImportPubKey', [pub_key], default_timeout)!
}

pub fn (mut b BtcClient) import_pub_key_rescan(args ImportPubKeyRescan) ! {
	b.client.send_json_rpc[[]ImportAddressRescan, string]('btc.ImportPubKeyRescan', [
		args,
	], default_timeout)!
}

pub fn (mut b BtcClient) invalidate_block(hash string) ! {
	b.client.send_json_rpc[[]string, string]('btc.InvalidateBlock', [hash], default_timeout)!
}

pub fn (mut b BtcClient) rename_account(args RenameAccount) ! {
	b.client.send_json_rpc[[]RenameAccount, string]('btc.RenameAccount', [args], default_timeout)!
}

pub fn (mut b BtcClient) submit_block(block SubmitBlock) ! {
	b.client.send_json_rpc[[]SubmitBlock, string]('btc.SubmitBlock', [block], default_timeout)!
}

pub fn (mut b BtcClient) send_to_address(args SendToAddress) ! {
	b.client.send_json_rpc[[]SendToAddress, string]('btc.SendToAddress', [args], default_timeout)!
}

pub fn (mut b BtcClient) send_to_address_comment(args SendToAddress) ! {
	b.client.send_json_rpc[[]SendToAddress, string]('btc.SendToAddressComment', [args],
		default_timeout)!
}
