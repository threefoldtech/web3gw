module btc

pub struct EstimateSmartFeeResult {
	feerate f64
	errors  []string
	blocks  i64
}

pub struct GetAddressInfoResult {
	EmbeddedAddressInfo
	ismine      bool
	iswatchonly bool
	timestamp   int
	hdkeypath   string
	hdseedid    string
	embedded    EmbeddedAddressInfo
}

pub struct EmbeddedAddressInfo {
	address             string
	script_pub_key      string   [json: 'scriptPubKey']
	solvable            bool
	desc                string
	isscript            bool
	ischange            bool
	iswitness           bool
	witness_version     int
	witness_program     string
	script              byte
	hex                 string
	pubkeys             []string
	sigsrequired        int
	pubkey              string
	iscompressed        bool
	hdmasterfingerprint string
	labels              []string
}

pub struct GetBlockStatsResult {
	avgfee              i64
	avgfeerate          i64
	avgtxsize           i64
	feerate_percentiles []i64
	blockhash           string
	height              i64
	ins                 i64
	maxfee              i64
	maxfeerate          i64
	maxtxsize           i64
	medianfee           i64
	mediantime          i64
	mediantxsize        i64
	minfee              i64
	minfeerate          i64
	mintxsize           i64
	outs                i64
	swtotal_size        i64
	swtotal_weight      i64
	swtxs               i64
	subsidy             i64
	time                i64
	total_out           i64
	total_size          i64
	total_weight        i64
	txs                 i64
	utxo_increase       i64
	utxo_size_inc       i64
}

// GetBlockVerboseTxResult models the data from the getblock command when the
pub struct GetBlockVerboseTxResult {
	hash              string
	confirmations     i64
	strippedsize      int
	size              int
	weight            int
	height            i64
	version           int
	version_hex       string        [json: 'versionHex']
	merkleroot        string
	tx                []TxRawResult
	rawtx             []TxRawResult
	time              i64
	nonce             u32
	bits              string
	difficulty        f64
	previousblockhash string
	nextblockhash     string
}

// TxRawResult models the data from the getrawtransaction command.
pub struct TxRawResult {
	hex           string
	txid          string
	hash          string
	size          int
	vsize         int
	weight        int
	version       u32
	locktime      u32
	vin           []Vin
	vout          []Vout
	blockhash     string
	confirmations u64
	time          i64
	blocktime     i64
}

pub struct Vin {
	coinbase    string
	txid        string
	vout        u32
	script_sig  &ScriptSig [json: 'scriptSig']
	sequence    u32
	txinwitness []string
}

pub struct ScriptSig {
	asm_ string [json: 'asm']
	hex  string
}

pub struct Vout {
	value          f64
	n              u32
	script_pub_key ScriptPubKeyResult [json: 'scriptPubKey']
}

pub struct ScriptPubKeyResult {
	asm_      string   [json: 'asm']
	hex       string
	req_sigs  int      [json: 'reqSigs']
	type_     string   [json: 'type']
	addresses []string
}

// GetChainTxStatsResult models the data from the getchaintxstats command.
pub struct GetChainTxStatsResult {
	time                      i64
	txcount                   i64
	window_final_block_hash   string
	window_final_block_height int
	window_block_count        int
	window_tx_count           int
	window_interval           int
	txrate                    f64
}

// GetMiningInfoResult models the data from the getmininginfo command.
pub struct GetMiningInfoResult {
	blocks             i64
	currentblocksize   u64
	currentblockweight u64
	currentblocktx     u64
	difficulty         f64
	errors             string
	generate           bool
	genproclimit       int
	hashespersec       f64
	networkhashps      f64
	pooledtx           u64
	testnet            bool
}

// GetNodeAddressesResult models the data returned from the getnodeaddresses command.
pub struct GetNodeAddressesResult {
	// Timestamp in seconds since epoch (Jan 1 1970 GMT) keeping track of when the node was last seen
	time     i64
	services u64
	address  string
	port     u16
}

// GetPeerInfoResult models the data returned from the getpeerinfo command.
pub struct GetPeerInfoResult {
	id             int
	addr           string
	addrlocal      string
	services       string
	relaytxes      bool
	lastsend       i64
	lastrecv       i64
	bytessent      u64
	bytesrecv      u64
	conntime       i64
	timeoffset     i64
	pingtime       f64
	pingwait       f64
	version        u32
	subver         string
	inbound        bool
	startingheight int
	currentheight  int
	banscore       int
	feefilter      i64
	syncnode       bool
}

pub struct Transaction {
	msg_tx          MsgTx  // Underlying MsgTx
	tx_hash         []byte // Cached transaction hash
	tx_hash_witness []byte // Cached transaction witness hash
	tx_has_witness  bool   // If the transaction has witness data
	tx_index        int    // Position within a block or TxIndexUnknown
}

pub struct MsgTx {
	version   int
	tx_in     []TxIn
	tx_out    []TxOut
	lock_time u32
}

pub struct TxIn {
	previous_out_point OutPoint
	signature_script   []byte
	witness            [][]byte
	sequence           u32
}

pub struct OutPoint {
	hash  []byte
	index u32
}

pub struct TxOut {
	value     i64
	pk_script []byte
}

pub struct LoadWalletResult {
	name string 
	warning string
}

pub struct GetWalletInfoResult {
	walletname            string
	walletversion         int
	txcount  int 
	keypoololdest int
	keypoolsize int
	keypoolsize_hd_internal ?int
	unlocked_until ?int
	paytxfee f64
	hdseedid string
	private_keys_enabled bool
	avoid_reuse bool
}

// CreateWalletResult models the result of the createwallet command.
pub struct CreateWalletResult {
	name    string
	warning string
}

pub struct GetBlockChainInfo {
	chain string
	blocks int
	headers int
	bestblockhash string
	difficulty f64
	mediantime i64
	verificationprogress f64
	initialblockdownload bool
	pruned bool
	pruneheight int
	chainwork string
	size_on_disk i64
}

pub struct ListReceivedByLabelResult {
	account string
	amount  f64
	confirmations u64
}

pub struct ListReceivedByAddressResult {
	account string
	address string
	amount f64
	confirmations u64
	txids []string
	involveswatchonly bool [json: 'involvesWatchonly']
}

pub struct ListSinceBlockResult {
	transactions []ListTransactionsResult
	lastblock string
}

pub struct ListTransactionsResult{
	abandoned bool
	account string
	address string
	amount f64
	bip125_replaceable string [json: 'bip125-replaceable']
	blockhash string
	blockheight int
	blockindex i64
	blocktime i64
	category string
	confirmations i64
	fee f64
	generated bool
	involveswatchonly bool
	label string
	time i64
	timereceived i64
	trusted bool
	txid string
	vout u32
	walletconflicts []string
	otheraccount string
}