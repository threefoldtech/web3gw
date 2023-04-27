module btc

// configurations to load bitcoin client
pub struct Config {
	host string
	user string
	pass string
}

// args to import bitcoin address
pub struct ImportAddressRescan {
	address string
	account string
	rescan  bool
}

pub struct ImportPrivKeyLabel {
	wif   string
	label string
}

pub struct ImportPrivKeyRescan {
	wif    string
	label  string
	rescan bool
}

pub struct ImportPubKeyRescan {
	pub_key string
	rescan  bool
}

pub struct RenameAccount {
	old_account string
	new_account string
}

// pub struct SubmitBlock {
// 	block   &btcutil.Block
// 	options &btcjson.SubmitBlockOptions
// }

// send amount of token to address, with/without comment
pub struct SendToAddress {
	address    string
	amount     i64
	comment    string // is intended to be used for the purpose of the transaction
	comment_to string // is intended to be used for who the transaction is being sent to.
}

pub struct EstimateSmartFee {
	conf_target i64
	mode        string
}

// EstimateSmartFeeResult models the data returned buy the chain server
pub struct EstimateSmartFeeResult {
	feerate f64
	errors  []string
	blocks  i64
}

pub struct GenerateToAddress {
	num_blocks i64
	address    string
	max_tries  i64
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

struct EmbeddedAddressInfo {
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

// type Block struct {
// 	msg_block                &MsgBlock  // Underlying MsgBlock
// 	serialized_block          []byte          // Serialized bytes for the block
// 	serialized_block_no_witness []byte          // Serialized bytes for block w/o witness data
// 	block_hash                &Hash // Cached block hash
// 	block_height              int           // Height in the main block chain
// 	transactions             []&Tx           // Transactions
// 	txnsGenerated            bool            // ALL wrapped transactions generated
// }

// type MsgBlock struct {
// 	Header       BlockHeader
// 	Transactions []&MsgTx
// }

// type BlockHeader struct {
// 	Version int
// 	PrevBlock Hash
// 	MerkleRoot Hash
// 	Timestamp Time
// 	Bits u32
// 	Nonce u32
// }

type Hash = []byte

// type Time struct {
// 	wall u64
// 	ext  i64
// 	loc &Location
// }

// type Location struct {
// 	name string
// 	zone []zone
// 	tx   []zoneTrans
// 	extend string
// 	cacheStart i64
// 	cacheEnd   i64
// 	cacheZone  &zone
// }

// type zone struct {
// 	name   string
// 	offset int
// 	isDST  bool
// }

// type zoneTrans struct {
// 	when         i64
// 	index        u8
// 	isstd, isutc bool
// }

// type SubmitBlock struct {
// 	Version  int
// 	TxIn     []&TxIn
// 	TxOut    []&TxOut
// 	LockTime u32
// }

// type TxIn struct {
// 	PreviousOutPoint OutPoint
// 	SignatureScript  []byte
// 	Witness          TxWitness
// 	Sequence         u32
// }

// type OutPoint struct {
// 	Hash  Hash
// 	Index u32
// }

// type TxWitness [][]byte

// type TxOut struct {
// 	Value    i64
// 	PkScript []byte
// }

// type Tx struct {
// 	msgTx         &MsgTx
// 	txHash        &Hash
// 	txHashWitness &Hash
// 	txHasWitness  &bool
// 	txIndex       int
// }
