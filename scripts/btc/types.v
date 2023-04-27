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

// type Hash []byte

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
