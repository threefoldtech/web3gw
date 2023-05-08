module btclightning

pub struct ChainServiceConfig {
pub:
	data_dir          string
	db_path           string
	chain_net         string
	connect_peers     []string
	add_peers         []string
	filter_cache_size u64
	block_cache_size  u64
	persist_to_disk   bool
	broadcast_timeout i64
}

pub struct BlockStamp {
pub:
	height    int
	hash      []byte
	timestamp i64
}

pub struct BlockHeader {
pub:
	version     int
	prev_block  []byte
	merkle_root []byte
	timestamp   i64
	bits        u32
	nonce       u32
}

pub struct BanPeerInfo {
pub:
	address string
	reason  u8
}

pub struct NetTotals {
pub:
	received u64
	sent     u64
}

pub struct UpdatePeerHeightsRequest {
pub:
	latest_block_hash   []byte
	latest_height       int
	source_peer_address string
}

pub struct RescanOptions {
pub:
	end_block       BlockStamp
	start_block     BlockStamp
	start_time      i64
	tx_idx          u32
	watch_inputs    []InputWithScript
	query_options   QueryOptions
	watch_addresses []string
}

pub struct InputWithScript {
pub:
	outpoint  OutPoint
	pk_script string
}

pub struct OutPoint {
pub:
	hash  []byte
	index u32
}

pub struct QueryOptions {
pub:
	invalid_tx_threshold     &f32
	max_batch_size           &i64
	num_retries              &u8
	optimistic_batch         &bool
	optimistic_reverse_batch &bool
	peer_connect_timeout     &i64
	reject_timeout           &i64
	timeout                  &i64
}

pub struct RescanUpdateOptions {
pub:
	add_addresses              []string
	watch_inputs               []InputWithScript
	disable_disconnected_ntfns &bool
	rewind                     &u32
}
