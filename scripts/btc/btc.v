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

// configurations to load bitcoin client
[params]
pub struct Config {
	host string
	user string
	pass string
}

// args to import bitcoin address
[params]
pub struct ImportAddressRescan {
	address string
	account string
	rescan  bool
}

[params]
pub struct ImportPrivKeyLabel {
	wif   string
	label string
}

[params]
pub struct ImportPrivKeyRescan {
	wif    string
	label  string
	rescan bool
}

[params]
pub struct ImportPubKeyRescan {
	pub_key string
	rescan  bool
}

[params]
pub struct RenameAccount {
	old_account string
	new_account string
}

// send amount of token to address, with/without comment
[params]
pub struct SendToAddress {
	address    string
	amount     i64
	comment    string // is intended to be used for the purpose of the transaction
	comment_to string // is intended to be used for who the transaction is being sent to.
}

[params]
pub struct EstimateSmartFee {
	conf_target i64
	mode        string
}

[params]
pub struct GenerateToAddress {
	num_blocks i64
	address    string
	max_tries  i64
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

// provides an estimated fee  in bitcoins per kilobyte
// num is the number of blocks
pub fn (mut c BtcClient) estimate_fee(num i64) !f64 {
	return c.client.send_json_rpc[[]i64, f64]('btc.EstimateFee', [num], btc.default_timeout)!
}

// requests the server to estimate a fee level based on the given parameters.
pub fn (mut c BtcClient) estimate_smart_fee(args EstimateSmartFee) !EstimateSmartFeeResult {
	return c.client.send_json_rpc[[]EstimateSmartFee, EstimateSmartFeeResult]('btc.EstimateSmartFee',
		[args], btc.default_timeout)!
}

// generates numBlocks blocks and returns their hashes.
pub fn (mut c BtcClient) generate(num u32) ![][]byte {
	return c.client.send_json_rpc[[]u32, [][]byte]('btc.Generate', [num], btc.default_timeout)!
}

// generates numBlocks blocks to the given address and returns their hashes.
pub fn (mut c BtcClient) generate_to_address(args GenerateToAddress) ![][]byte {
	return c.client.send_json_rpc[[]GenerateToAddress, [][]byte]('btc.GenerateToAddress',
		[args], btc.default_timeout)!
}

// returns the account associated with the passed address.
pub fn (mut c BtcClient) get_account(address string) !string {
	return c.client.send_json_rpc[[]string, string]('btc.GetAccount', [address], btc.default_timeout)!
}

// returns the current Bitcoin address for receiving payments to the specified account.
pub fn (mut c BtcClient) get_account_address(account string) !string {
	return c.client.send_json_rpc[[]string, string]('btc.GetAccountAddress', [account],
		btc.default_timeout)!
}

// returns information about the given bitcoin address.
pub fn (mut c BtcClient) get_address_info(address string) !GetAddressInfoResult {
	return c.client.send_json_rpc[[]string, GetAddressInfoResult]('btc.GetAddressInfo',
		[address], btc.default_timeout)!
}

// returns the list of addresses associated with the passed account.
pub fn (mut c BtcClient) get_addresses_by_account(account string) ![]string {
	return c.client.send_json_rpc[[]string, []string]('btc.GetAddressesByAccount', [
		account,
	], btc.default_timeout)!
}

// returns the available balance from the server for the specified account using the default number of minimum confirmations.
// The account may be "*" for all accounts.
pub fn (mut c BtcClient) get_balance(account string) !i64 {
	return c.client.send_json_rpc[[]string, i64]('btc.GetBalance', [account], btc.default_timeout)!
}

// returns the number of blocks in the longest block chain.
pub fn (mut c BtcClient) get_block_count(account string) !i64 {
	return c.client.send_json_rpc[[]string, i64]('btc.GetBlockCount', [account], btc.default_timeout)!
}

// returns the hash of the block in the best block chain at the given height.
pub fn (mut c BtcClient) get_block_hash(block_height i64) ![]byte {
	return c.client.send_json_rpc[[]i64, []byte]('btc.GetBlockHash', [block_height], btc.default_timeout)!
}

// returns block statistics.
// hash argument specifies height or hash of the target block
pub fn (mut c BtcClient) get_block_stats(hash string) !GetBlockStatsResult {
	return c.client.send_json_rpc[[]string, GetBlockStatsResult]('btc.GetBlockStats',
		[hash], btc.default_timeout)!
}

// returns a data structure from the server with information about a block and its transactions given its hash.
pub fn (mut c BtcClient) get_block_verbose_tx(hash string) !GetBlockVerboseTxResult {
	return c.client.send_json_rpc[[]string, GetBlockVerboseTxResult]('btc.GetBlockVerboseTx',
		[hash], btc.default_timeout)!
}
