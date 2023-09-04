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
	host string // the address of the node including the port (for example: 185.69.167.219:8332)
	user string // the user name that you can use to connect to the btc node
	pass string // the password needed to connect to the btc node
	wallet string // optional, the wallet name you want to use for future calls (can be changed with the LoadWallet method)
}

// args to import bitcoin address
[params]
pub struct ImportAddress {
	address string // the Bitcoin address (or hex-encoded script)
	label string // an optional label
	rescan bool = true // whether or not to scan the chain and mempool for wallet transactions
	p2sh bool // add the p2sh version of the script as well
}

[params]
pub struct ImportPrivKey {
	wif    string
	label  string // an optional label (default: current label if address exists, otherwise "")
	rescan bool = true // scan the chain and mempool for wallet transactions
}

[params]
pub struct ImportPubKey {
	pub_key string
	label string // an optional label
	rescan  bool // scan the chain and mempool for wallet transactions
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
	conf_target i64 = 1 // confirmation target in blocks (value between 1 and 1008)
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
pub struct GetNewAddress {
	label string
	address_type string
}

[params]
pub struct CreateWallet {
	name                 string // name of the wallet to create
	disable_private_keys bool // disable the possibility of private keys (only watchonlys are possible in this mode)
	create_blank_wallet  bool // create a blank wallet (has no keys or HD seed)
	passphrase           string // encrypt the wallet with this passphrase
	avoid_reuse          bool // keep track of coin reuse, and treat dirty and clean coins differently with privacy considerations in mind
}

[params]
pub struct Move {
	from_account      string
	to_account        string
	amount            i64
	min_confirmations int
	comment           string
}

[params]
pub struct SetLabel {
	label string // the label to assign to the address
	address string // the bitcoin address to be associated with a label
}

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

// Imports the passed public address. When rescan is true,
// the block history is scanned for transactions addressed to provided address.
pub fn (mut c BtcClient) import_address(args ImportAddress) ! {
	_ := c.client.send_json_rpc[[]ImportAddress, string]('btc.ImportAddress', [
		args,
	], btc.default_timeout)!
}

// Imports the passed private key which must be the wallet import
// format (WIF). It sets the account label to the one provided. When rescan is true,
// the block history is scanned for transactions addressed to provided privKey.
// The WIF string must be a base58-encoded string.
pub fn (mut c BtcClient) import_priv_key(args ImportPrivKey) ! {
	_ := c.client.send_json_rpc[[]ImportPrivKey, string]('btc.ImportPrivKey', [
		args,
	], btc.default_timeout)!
}

// Imports the passed public key. When rescan is true, the block history is scanned for transactions addressed to provided pubkey.
pub fn (mut c BtcClient) import_pub_key_rescan(args ImportPubKey) ! {
	_ := c.client.send_json_rpc[[]ImportPubKey, string]('btc.ImportPubKey', [
		args,
	], btc.default_timeout)!
}

// List the accounts of a wallet
pub fn (mut c BtcClient) list_labels() !map[string]i64 {
	return c.client.send_json_rpc[[]string, map[string]i64]('btc.ListLabels', []string{}, btc.default_timeout)!
}

// Allows you to rename an account.
pub fn (mut c BtcClient) rename_account(args RenameAccount) ! {
	_ := c.client.send_json_rpc[[]RenameAccount, string]('btc.RenameAccount', [args], btc.default_timeout)!
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

// Associates the given label to the given address
pub fn (mut c BtcClient) set_label(args SetLabel) ! {
	_ := c.client.send_json_rpc[[]SetLabel, string]('btc.SetLabel', [args], btc.default_timeout)!
}

// Returns information about the given bitcoin address.
pub fn (mut c BtcClient) get_address_info(address string) !GetAddressInfoResult {
	return c.client.send_json_rpc[[]string, GetAddressInfoResult]('btc.GetAddressInfo',
		[address], btc.default_timeout)!
}

// Returns the list of addresses associated with the provided label. The returned list will be the string encoded versions of the addresses.
pub fn (mut c BtcClient) get_addresses_by_label(label string) ![]string {
	return c.client.send_json_rpc[[]string, []string]('btc.GetAddressesByLabel', [
		label,
	], btc.default_timeout)!
}

// Returns the available balance using the default number of minimum confirmations.
pub fn (mut c BtcClient) get_balance() !i64 {
	return c.client.send_json_rpc[[]string, i64]('btc.GetBalance', []string{}, btc.default_timeout)!
}

// Returns the height of the most-work fully-validated chain.
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

// Returns information on the state of the blockchain. 
pub fn (mut c BtcClient) get_blockchain_info() !GetBlockChainInfo {
	return c.client.send_json_rpc[[]string, GetBlockChainInfo]('btc.GetBlockStats',
		[]string{}, btc.default_timeout)!
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

// Returns the number of connections to other nodes.
pub fn (mut c BtcClient) get_connection_count() !i64 {
	return c.client.send_json_rpc[[]string, i64]('btc.GetConnectionCount', []string{}, btc.default_timeout)!
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

// Returns a new address. The returned string will be the encoded address based on the address_type provided. If 
// address_type is left empty the default address type will be used from the chain's parameters.
pub fn (mut c BtcClient) get_new_address(args GetNewAddress) !string {
	return c.client.send_json_rpc[[]GetNewAddress, string]('btc.GetNewAddress', [args], btc.default_timeout)!
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

// Returns the total amount received by the specified label
pub fn (mut c BtcClient) get_received_by_label(label string) !Transaction {
	return c.client.send_json_rpc[[]string, Transaction]('btc.GetReceivedByLabel', [label],
		btc.default_timeout)!
}

// Load a specific wallet. This is needed if you want to switch wallets or if you didn't pass any wallet name at Load time.  
pub fn (mut c BtcClient) load_wallet(wallet_name string) !LoadWalletResult {
	return c.client.send_json_rpc[[]string, LoadWalletResult]('btc.LoadWallet', [wallet_name],
		btc.default_timeout)!
}

// Return wallet information
pub fn (mut c BtcClient) get_wallet_info() !GetWalletInfoResult {
	return c.client.send_json_rpc[[]string, GetWalletInfoResult]('btc.GetWalletInfo', []string{},
		btc.default_timeout)!
}

// Lists the received transactions by label
pub fn (mut c BtcClient) list_received_by_label() ![]ListReceivedByLabelResult {
	return c.client.send_json_rpc[[]string, []ListReceivedByLabelResult]('btc.ListReceivedByLabel', []string{},
		btc.default_timeout)!
}

// Lists the received transactions by address
pub fn (mut c BtcClient) list_received_by_address() ![]ListReceivedByAddressResult {
	return c.client.send_json_rpc[[]string, []ListReceivedByAddressResult]('btc.ListReceivedByAddress', []string{},
		btc.default_timeout)!
}

// List all transactions since block with hash.
pub fn (mut c BtcClient) list_since_block(hash string) !ListSinceBlockResult {
	return c.client.send_json_rpc[[]string, ListSinceBlockResult]('btc.ListSinceBlock', [hash],
		btc.default_timeout)!
}

// List all transactions for a specific label.
pub fn (mut c BtcClient) list_transactions(label string) ![]ListTransactionsResult {
	return c.client.send_json_rpc[[]string, []ListTransactionsResult]('btc.ListTransactions', [label],
		btc.default_timeout)!
}

// Creates a new wallet taking into account the provided arguments. 
pub fn (mut c BtcClient) create_wallet(args CreateWallet) !CreateWalletResult {
	return c.client.send_json_rpc[[]CreateWallet, CreateWalletResult]('btc.CreateWallet',
		[args], btc.default_timeout)!
}

// Sets the transaction fee per kilobyte paid by transactions created by this wallet.
pub fn (mut c BtcClient) set_tx_fee(fee i64) ! {
	_ := c.client.send_json_rpc[[]i64, string]('btc.SetTxFee',
		[fee], btc.default_timeout)!
}

// Deprecated
// Moves specified amount from one account in your wallet to another. Only funds with the default number of minimum confirmations will be used.
// A comment can also be added to the transaction.
pub fn (mut c BtcClient) move(args Move) !bool {
	return c.client.send_json_rpc[[]Move, bool]('btc.Move', [args], btc.default_timeout)!
}
