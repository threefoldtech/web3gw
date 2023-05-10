module btc

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

const (
	default_timeout = 500000
)

[noinit; openrpc: exclude]
pub struct BtcClient {
mut:
	client &RpcWsClient
}

// configurations to load bitcoin client
[params]
pub struct Load {
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
	comment    string // is intended to be used for the purpose of the transaction, keep empty if you don't wat to provide any comment.
	comment_to string // is intended to be used for who the transaction is being sent to.
}

[params]
pub struct EstimateSmartFee {
	conf_target i64 = 1 // confirmation target in blocks
	mode        string = "CONSERVATIVE" // defines the different fee estimation modes, should be one of UNSET, ECONOMICAL or CONSERVATIVE
}

[params]
pub struct GenerateToAddress {
	num_blocks i64 = 1 // the amount of blocks to generate
	address    string
	max_tries  i64 = 1 // the maximum amount of times to try again when generating fails
}

[params]
pub struct GetChainTxStats {
	amount_of_blocks int // provide statistics for amount_of_blocks blocks, if 0 for all blocks
	block_hash_end   string // provide statistics for amount_of_blocks blocks up until the block with the hash provided in block_hash_end
}

[params]
pub struct CreateWallet {
	name                 string
	disable_private_keys bool
	create_blank_wallet  bool
	passphrase           string
	avoid_reuse          bool
}

[params]
pub struct Move {
	from_account      string
	to_account        string
	amount            i64
	min_confirmations int
	comment           string
}

[openrpc: exclude]
pub fn new(mut client RpcWsClient) BtcClient {
	return BtcClient{
		client: &client
	}
}

// Connects to the bitcoin node. This should be the first call to execute.
pub fn (mut c BtcClient) load(params Load) !string {
	return c.client.send_json_rpc[[]Load, string]('btc.Load', [
		params,
	], btc.default_timeout)!
}

// Imports the passed public address.
pub fn (mut c BtcClient) import_address(address string) ! {
	_ := c.client.send_json_rpc[[]string, string]('btc.ImportAddress', [
		address,
	], btc.default_timeout)!
}

// Imports the passed public address. When rescan is true,
// the block history is scanned for transactions addressed to provided address.
pub fn (mut c BtcClient) import_address_rescan(args ImportAddressRescan) ! {
	_ := c.client.send_json_rpc[[]ImportAddressRescan, string]('btc.ImportAddressRescan', [
		args,
	], btc.default_timeout)!
}

// Imports the passed private key which must be the wallet import format (WIF).
// The WIF string must be a base58-encoded string.
pub fn (mut c BtcClient) import_priv_key(wif string) ! {
	_ := c.client.send_json_rpc[[]string, string]('btc.ImportPrivKey', [
		wif,
	], btc.default_timeout)!
}

// Imports the passed private key which must be the wallet import
// format (WIF). It sets the account label to the one provided.
// The WIF string must be a base58-encoded string.
pub fn (mut c BtcClient) import_priv_key_label(args ImportPrivKeyLabel) ! {
	_ := c.client.send_json_rpc[[]ImportPrivKeyLabel, string]('btc.ImportPrivKeyLabel', [
		args,
	], btc.default_timeout)!
}

// Imports the passed private key which must be the wallet import
// format (WIF). It sets the account label to the one provided. When rescan is true,
// the block history is scanned for transactions addressed to provided privKey.
// The WIF string must be a base58-encoded string.
pub fn (mut c BtcClient) import_priv_key_rescan(args ImportPrivKeyRescan) ! {
	_ := c.client.send_json_rpc[[]ImportPrivKeyRescan, string]('btc.ImportPrivKeyRescan', [
		args,
	], btc.default_timeout)!
}

// Imports the passed public key.
pub fn (mut c BtcClient) import_pub_key(pub_key string) ! {
	_ := c.client.send_json_rpc[[]string, string]('btc.ImportPubKey', [pub_key], btc.default_timeout)!
}

// Imports the passed public key. When rescan is true, the block history is scanned for transactions addressed to provided pubkey.
pub fn (mut c BtcClient) import_pub_key_rescan(args ImportPubKeyRescan) ! {
	_ := c.client.send_json_rpc[[]ImportPubKeyRescan, string]('btc.ImportPubKeyRescan', [
		args,
	], btc.default_timeout)!
}

// Allows you to rename an account.
pub fn (mut c BtcClient) rename_account(args RenameAccount) ! {
	c.client.send_json_rpc[[]RenameAccount, string]('btc.RenameAccount', [args], btc.default_timeout)!
}

// Sends the passed amount to the given address with a comment if provided and returns the hash of the transaction
pub fn (mut c BtcClient) send_to_address(args SendToAddress) !string {
	return c.client.send_json_rpc[[]SendToAddress, string]('btc.SendToAddress', [args],
		btc.default_timeout)!
}

// Provides a more accurate estimated fee given an estimation mode. 
pub fn (mut c BtcClient) estimate_smart_fee(args EstimateSmartFee) !EstimateSmartFeeResult {
	return c.client.send_json_rpc[[]EstimateSmartFee, EstimateSmartFeeResult]('btc.EstimateSmartFee',
		[args], btc.default_timeout)!
}

// Generates the provided amount of blocks and returns their hashes.
pub fn (mut c BtcClient) generate_blocks(num_blocks u32) ![]string {
	return c.client.send_json_rpc[[]u32, []string]('btc.GenerateBlocks', [num_blocks], btc.default_timeout)!
}

// Generates numBlocks blocks to the given address and returns their hashes.
pub fn (mut c BtcClient) generate_blocks_to_address(args GenerateToAddress) ![]string {
	return c.client.send_json_rpc[[]GenerateToAddress, []string]('btc.GenerateBlocksToAddress',
		[args], btc.default_timeout)!
}

// Returns the account associated with the passed address. The address should be the string encoded version of a valid address.
pub fn (mut c BtcClient) get_account(address string) !string {
	return c.client.send_json_rpc[[]string, string]('btc.GetAccount', [address], btc.default_timeout)!
}

// Returns the current Bitcoin address for receiving payments to the specified account.
pub fn (mut c BtcClient) get_account_address(account string) !string {
	return c.client.send_json_rpc[[]string, string]('btc.GetAccountAddress', [account],
		btc.default_timeout)!
}

// Returns information about the given bitcoin address.
pub fn (mut c BtcClient) get_address_info(address string) !GetAddressInfoResult {
	return c.client.send_json_rpc[[]string, GetAddressInfoResult]('btc.GetAddressInfo',
		[address], btc.default_timeout)!
}

// Returns the list of addresses associated with the provided account. The returned list will be the string encoded versions of the addresses.
pub fn (mut c BtcClient) get_addresses_by_account(account string) ![]string {
	return c.client.send_json_rpc[[]string, []string]('btc.GetAddressesByAccount', [
		account,
	], btc.default_timeout)!
}

// Returns the available balance for the specified account using the default number of minimum confirmations. You can provide * as an account to get the balance of all accounts.
pub fn (mut c BtcClient) get_balance(account string) !i64 {
	return c.client.send_json_rpc[[]string, i64]('btc.GetBalance', [account], btc.default_timeout)!
}

// Returns the number of blocks in the longest block chain.
pub fn (mut c BtcClient) get_block_count() !i64 {
	return c.client.send_json_rpc[[]string, i64]('btc.GetBlockCount', []string{}, btc.default_timeout)!
}

// Returns the hash of the block in the best block chain at the given height.
pub fn (mut c BtcClient) get_block_hash(block_height i64) !string {
	return c.client.send_json_rpc[[]i64, string]('btc.GetBlockHash', [block_height], btc.default_timeout)!
}

// Returns block statistics given the hash of that block. 
pub fn (mut c BtcClient) get_block_stats(hash string) !GetBlockStatsResult {
	return c.client.send_json_rpc[[]string, GetBlockStatsResult]('btc.GetBlockStats',
		[hash], btc.default_timeout)!
}

// Returns information about a block and its transactions given the hash of that block.
pub fn (mut c BtcClient) get_block_verbose_tx(hash string) !GetBlockVerboseTxResult {
	return c.client.send_json_rpc[[]string, GetBlockVerboseTxResult]('btc.GetBlockVerboseTx',
		[hash], btc.default_timeout)!
}

// Returns statistics about the total number and rate of transactions in the chain. 
// Providing the arguments will reduce the amount of blocks to calculate the statistics on.
pub fn (mut c BtcClient) get_chain_tx_stats(args GetChainTxStats) !GetChainTxStatsResult {
	return c.client.send_json_rpc[[]GetChainTxStats, GetChainTxStatsResult]('btc.GetChainTxStats',
		[args], btc.default_timeout)!
}

// Returns the proof-of-work difficulty as a multiple of the minimum difficulty.
pub fn (mut c BtcClient) get_difficulty() !f64 {
	return c.client.send_json_rpc[[]string, f64]('btc.GetDifficulty', []string{}, btc.default_timeout)!
}

// Returns mining information.
pub fn (mut c BtcClient) get_mining_info() !GetMiningInfoResult {
	return c.client.send_json_rpc[[]string, GetMiningInfoResult]('btc.GetMiningInfo',
		[]string{}, btc.default_timeout)!
}

// Returns a new address. The returned string will be the encoded address (format will be based on the chain's parameters).
pub fn (mut c BtcClient) get_new_address(account string) !string {
	return c.client.send_json_rpc[[]string, string]('btc.GetNewAddress', [account], btc.default_timeout)!
}

// Returns data about known node addresses.
pub fn (mut c BtcClient) get_node_addresses() ![]GetNodeAddressesResult {
	return c.client.send_json_rpc[[]string, []GetNodeAddressesResult]('btc.GetNodeAddresses',
		[]string{}, btc.default_timeout)!
}

// Returns data about each connected network peer.
pub fn (mut c BtcClient) get_peer_info() ![]GetPeerInfoResult {
	return c.client.send_json_rpc[[]string, []GetPeerInfoResult]('btc.GetPeerInfo', []string{},
		btc.default_timeout)!
}

// Returns a transaction given its hash.
pub fn (mut c BtcClient) get_raw_transaction(tx_hash string) !Transaction {
	return c.client.send_json_rpc[[]string, Transaction]('btc.GetRawTransaction', [tx_hash],
		btc.default_timeout)!
}

// Creates a new wallet account taken into account the provided arguments. 
pub fn (mut c BtcClient) create_wallet(args CreateWallet) !CreateWalletResult {
	return c.client.send_json_rpc[[]CreateWallet, CreateWalletResult]('btc.CreateWallet',
		[args], btc.default_timeout)!
}

// Moves specified amount from one account in your wallet to another. Only funds with the default number of minimum confirmations will be used.
// A comment can also be added to the transaction.
pub fn (mut c BtcClient) move(args Move) !bool {
	return c.client.send_json_rpc[[]Move, bool]('btc.Move', [args], btc.default_timeout)!
}
