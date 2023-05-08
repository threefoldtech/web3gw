package btclightning

import (
	"errors"

	"github.com/lightninglabs/neutrino"
)

var (
	ChainServiceInstanceNotFoundErr = errors.New("no active instance of a chain service")
	RescanInstanceNotFoundErr       = errors.New("no active instance of a rescan")
)

// Config is a struct detailing the configuration of the chain service.
type ChainServiceConfig struct {
	DataDir          string   `json:"data_dir"`
	DatabasePath     string   `json:"db_path"`
	ChainNet         string   `json:"chain_net"`
	ConnectPeers     []string `json:"connect_peers"`
	AddPeers         []string `json:"add_peers"`
	FilterCacheSize  uint64   `json:"filter_cache_size"`
	BlockCacheSize   uint64   `json:"block_cache_size"`
	PersistToDisk    bool     `json:"persist_to_disk"`
	BroadcastTimeout int64    `json:"broadcast_timeout"`
}

type BlockStamp struct {
	// Height is the height of the target block.
	Height int32 `json:"height"`

	// Hash is the hash that uniquely identifies this block.
	Hash []byte `json:"hash"`

	// Timestamp is the timestamp of the block in the chain.
	Timestamp int64 `json:"timestamp"`
}

type BlockHeader struct {
	// Version of the block.  This is not the same as the protocol version.
	Version int32 `json:"version"`

	// Hash of the previous block header in the block chain.
	PrevBlock []byte `json:"prev_block"`

	// Merkle tree reference to hash of all transactions for the block.
	MerkleRoot []byte `json:"merkle_root"`

	// Time the block was created.  This is, unfortunately, encoded as a
	// uint32 on the wire and therefore is limited to 2106.
	Timestamp int64 `json:"timestamp"`

	// Difficulty target for the block.
	Bits uint32 `json:"bits"`

	// Nonce used to generate the block.
	Nonce uint32 `json:"nonce"`
}

type BanPeerInfo struct {
	Address string `json:"address"`
	Reason  uint8  `json:"reason"`
}

type NetTotals struct {
	Received uint64 `json:"received"`
	Sent     uint64 `json:"sent"`
}

type UpdatePeerHeightsRequest struct {
	LatestBlockHash   []byte `json:"latest_block_hash"`
	LatestHeight      int32  `json:"latest_height"`
	SourcePeerAddress string `json:"source_peer_address"`
}

type RescanOptions struct {
	EndBlock       BlockStamp        `json:"end_block"`
	StartBlock     BlockStamp        `json:"start_block"`
	StartTime      int64             `json:"start_time"`
	TxIdx          uint32            `json:"tx_idx"`
	WatchInputs    []InputWithScript `json:"watch_inputs"`
	QueryOptions   QueryOptions      `json:"query_options"`
	WatchAddresses []string          `json:"watch_addresses"`
}

type InputWithScript struct {
	OutPoint OutPoint `json:"outpoint"`
	PkScript string   `json:"pk_script"`
}

type OutPoint struct {
	Hash  []byte `json:"hash"`
	Index uint32 `json:"index"`
}

type rescanContext struct {
	// blocksConnected    chan headerfs.BlockStamp
	// blocksDisconnected chan headerfs.BlockStamp
	rescan  *neutrino.Rescan
	errChan <-chan error
	quit    chan struct{}
}

type QueryOptions struct {
	InvalidTxThreshold     *float32 `json:"invalid_tx_threshold"`
	MaxBatchSize           *int64   `json:"max_batch_size"`
	NumRetries             *uint8   `json:"num_retries"`
	OptimisticBatch        *bool    `json:"optimistic_batch"`
	OptimisticReverseBatch *bool    `json:"optimistic_reverse_batch"`
	PeerConnectTimeout     *int64   `json:"peer_connect_timeout"`
	RejectTimeout          *int64   `json:"reject_timeout"`
	Timeout                *int64   `json:"timeout"`
}

type RescanUpdateOptions struct {
	AddAddresses             []string          `json:"add_addresses"`
	WatchInputs              []InputWithScript `json:"watch_inputs"`
	DisableDisconnectedNtfns *bool             `json:"disable_disconnected_ntfns"`
	Rewind                   *uint32           `json:"rewind"`
}
