module btclightning

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

// initializes new chain service
pub fn (mut c BtcClient) new_chain_service(cfg ChainServiceConfig) ! {
	_ := c.client.send_json_rpc[[]ChainServiceConfig, string]('btclightning.NewChainService',
		[
		cfg,
	], btc.default_timeout)!
}

// Retrieves the most recent block's height and hash where we
// have both the header and filter header ready.
pub fn (mut c BtcClient) best_block() !BlockStamp {
	return c.client.send_json_rpc[[]string, BlockStamp]('btclightning.BestBlock', [
		[]string{},
	], btc.default_timeout)!
}

// GetBlockHash returns the block hash at the given height.
pub fn (mut c BtcClient) get_block_hash(height i64) !string {
	return c.client.send_json_rpc[[]i64, string]('btclightning.GetBlockHash', [
		height,
	], btc.default_timeout)!
}

// Returns the block header for the given block hash, or an
// error if the hash doesn't exist or is unknown.
pub fn (mut c BtcClient) get_block_header(hash string) !BlockHeader {
	return c.client.send_json_rpc[[]string, BlockHeader]('btclightning.GetBlockHeader',
		[
		hash,
	], btc.default_timeout)!
}

// Gets the height of a block by its hash. An error is returned
// if the given block hash is unknown.
pub fn (mut c BtcClient) get_block_height(hash string) !int {
	return c.client.send_json_rpc[[]string, int]('btclightning.GetBlockHeight', [
		hash,
	], btc.default_timeout)!
}

// BanPeer disconnects and bans a peer due to a specific reason for a duration
// of BanDuration.
pub fn (mut c BtcClient) ban_peer(peer_info BanPeerInfo) ! {
	return c.client.send_json_rpc[[]BanPeerInfo, int]('btclightning.BanPeer', [
		peer_info,
	], btc.default_timeout)!
}

// IsBanned returns true if the peer is banned, and false otherwise.
pub fn (mut c BtcClient) is_banned(address string) !bool {
	return c.client.send_json_rpc[[]string, bool]('btclightning.IsBanned', [
		address,
	], btc.default_timeout)!
}

// AddPeer adds a new peer that has already been connected to the server.
pub fn (mut c BtcClient) add_peer(is_server_peer_persistent bool) ! {
	_ := c.client.send_json_rpc[[]bool, string]('btclightning.AddPeer', [
		is_server_peer_persistent,
	], btc.default_timeout)!
}

// AddBytesSent adds the passed number of bytes to the total bytes sent counter
// for the server.
pub fn (mut c BtcClient) add_bytes_sent(bytes_sent u64) ! {
	_ := c.client.send_json_rpc[[]u64, string]('btclightning.AddBytesSent', [
		bytes_sent,
	], btc.default_timeout)!
}

// AddBytesReceived adds the passed number of bytes to the total bytes received
// counter for the server.
pub fn (mut c BtcClient) add_bytes_received(bytes_received u64) ! {
	_ := c.client.send_json_rpc[[]u64, string]('btclightning.AddBytesReceived', [
		bytes_received,
	], btc.default_timeout)!
}

// NetTotals returns the sum of all bytes received and sent across the network
// for all peers.
pub fn (mut c BtcClient) net_totals() !NetTotals {
	return c.client.send_json_rpc[[]string, NetTotals]('btclightning.NetTotals', [
		[]string{},
	], btc.default_timeout)!
}

// UpdatePeerHeights updates the heights of all peers who have announced the
// latest connected main chain block, or a recognized orphan. These height
// updates allow us to dynamically refresh peer heights, ensuring sync peer
// selection has access to the latest block heights for each peer.
pub fn (mut c BtcClient) update_peer_heights(params UpdatePeerHeightsRequest) ! {
	_ := c.client.send_json_rpc[[]UpdatePeerHeightsRequest, string]('btclightning.UpdatePeerHeights',
		[
		params,
	], btc.default_timeout)!
}

// begins connecting to peers and syncing the blockchain.
pub fn (mut c BtcClient) start_chain_service() ! {
	_ := c.client.send_json_rpc[[]string, string]('btclightning.StartChainService', [
		[]string{},
	], btc.default_timeout)!
}

// gracefully shuts down the server by stopping and disconnecting all
// peers and the main listener.
pub fn (mut c BtcClient) stop_chain_service() ! {
	_ := c.client.send_json_rpc[[]string, string]('btclightning.StopChainService', [
		[]string{},
	], btc.default_timeout)!
}

// IsCurrent lets the caller know whether the chain service's block manager
// thinks its view of the network is current.
pub fn (mut c BtcClient) is_current() !bool {
	return c.client.send_json_rpc[[]string, bool]('btclightning.IsCurrent', [
		[]string{},
	], btc.default_timeout)!
}

// instantiates a new rescan object that runs in another a separate thread and has an
// updatable filter.
pub fn (mut c BtcClient) new_rescan(options RescanOptions) ! {
	_ := c.client.send_json_rpc[[]RescanOptions, string]('btclightning.NewRescan', [
		options,
	], btc.default_timeout)!
}

// kicks off the rescan thread, which will begin to scan the chain
// according to the specified rescan options.
pub fn (mut c BtcClient) start_rescan() ! {
	_ := c.client.send_json_rpc[[]string, string]('btclightning.StartRescan', [
		[]string{},
	], btc.default_timeout)!
}

// sends an update to a long-running rescan/notification thread.
pub fn (mut c BtcClient) udpate_rescan(options RescanUpdateOptions) ! {
	_ := c.client.send_json_rpc[[]RescanUpdateOptions, string]('btclightning.UpdateRescan',
		[
		options,
	], btc.default_timeout)!
}

// waits until all threads associated with the rescan have exited.
pub fn (mut c BtcClient) shutdown_rescan() ! {
	_ := c.client.send_json_rpc[[]string, string]('btclightning.RescanShutdown', [
		[]string{},
	], btc.default_timeout)!
}
