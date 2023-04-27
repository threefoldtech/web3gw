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

// loads a bitcoin netwrok client
pub fn (mut c BtcClient) load(params Config) !string {
	return c.client.send_json_rpc[[]Config, string]('btc.Load', [
		params,
	], btc.default_timeout)!
}

// creates a new wallet account.
pub fn (mut c BtcClient) create_new_account(account string) ! {
	c.client.send_json_rpc[[]string, string]('btc.CreateNewAccount', [
		account,
	], btc.default_timeout)!
}

// requests the creation of an encrypted wallet.
pub fn (mut c BtcClient) create_encrypted_wallet(passphrase string) ! {
	c.client.send_json_rpc[[]string, string]('btc.CreateEncryptedWallet', [
		passphrase,
	], btc.default_timeout)!
}

// imports the passed public address.
pub fn (mut c BtcClient) import_address(address string) ! {
	c.client.send_json_rpc[[]string, string]('btc.ImportAddress', [
		address,
	], btc.default_timeout)!
}

// imports the passed public address. When rescan is true,
// the block history is scanned for transactions addressed to provided address.
pub fn (mut c BtcClient) import_address_rescan(args ImportAddressRescan) ! {
	c.client.send_json_rpc[[]ImportAddressRescan, string]('btc.ImportAddressRescan', [
		args,
	], btc.default_timeout)!
}

// imports the passed private key which must be the wallet import format (WIF).
// The WIF string must be a base58-encoded string.
pub fn (mut c BtcClient) import_priv_key(wif string) ! {
	c.client.send_json_rpc[[]string, string]('btc.ImportPrivKey', [
		wif,
	], btc.default_timeout)!
}

// imports the passed private key which must be the wallet import
// format (WIF). It sets the account label to the one provided.
// The WIF string must be a base58-encoded string.
pub fn (mut c BtcClient) import_priv_key_label(args ImportPrivKeyLabel) ! {
	c.client.send_json_rpc[[]ImportPrivKeyLabel, string]('btc.ImportPrivKeyLabel', [
		args,
	], btc.default_timeout)!
}

// imports the passed private key which must be the wallet import
// format (WIF). It sets the account label to the one provided. When rescan is true,
// the block history is scanned for transactions addressed to provided privKey.
// The WIF string must be a base58-encoded string.
pub fn (mut c BtcClient) import_priv_key_rescan(args ImportPrivKeyRescan) ! {
	c.client.send_json_rpc[[]ImportPrivKeyRescan, string]('btc.ImportPrivKeyRescan', [
		args,
	], btc.default_timeout)!
}

// imports the passed public key.
pub fn (mut c BtcClient) import_pub_key(pub_key string) ! {
	c.client.send_json_rpc[[]string, string]('btc.ImportPubKey', [pub_key], btc.default_timeout)!
}

// imports the passed public key. When rescan is true, the block history is scanned for transactions addressed to provided pubkey.
pub fn (mut c BtcClient) import_pub_key_rescan(args ImportPubKeyRescan) ! {
	c.client.send_json_rpc[[]ImportPubKeyRescan, string]('btc.ImportPubKeyRescan', [
		args,
	], btc.default_timeout)!
}

// invalidates a specific block.
pub fn (mut c BtcClient) invalidate_block(hash string) ! {
	c.client.send_json_rpc[[]string, string]('btc.InvalidateBlock', [hash], btc.default_timeout)!
}

// creates a new wallet account.
pub fn (mut c BtcClient) rename_account(args RenameAccount) ! {
	c.client.send_json_rpc[[]RenameAccount, string]('btc.RenameAccount', [args], btc.default_timeout)!
}

// NOTE: a way to define SubmitBlock struct?
// // attempts to submit a new block into the bitcoin network.
// pub fn (mut c BtcClient) submit_block(block SubmitBlock) ! {
// 	return c.client.send_json_rpc[[]SubmitBlock, string]('btc.SubmitBlock', [block], btc.default_timeout)!
// }

// sends the passed amount to the given address.
pub fn (mut c BtcClient) send_to_address(args SendToAddress) ![]byte {
	return c.client.send_json_rpc[[]SendToAddress, []byte]('btc.SendToAddress', [args],
		btc.default_timeout)!
}

// sends the passed amount to the given address and stored the provided comment and comment to in the wallet
pub fn (mut c BtcClient) send_to_address_comment(args SendToAddress) ![]byte {
	return c.client.send_json_rpc[[]SendToAddress, []byte]('btc.SendToAddressComment',
		[args], btc.default_timeout)!
}
