module explorer

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

const (
	default_timeout = 500000
)

[noinit; openrpc: exclude]
pub struct ExplorerClient {
mut:
	client &RpcWsClient
}

[openrpc: exclude]
pub fn new(mut client RpcWsClient) ExplorerClient {
	return ExplorerClient{
		client: &client
	}
}

// First call to make to initialize your session. Provide the network you want to use to do so.
pub fn (mut e ExplorerClient) load(network string) ! {
	_ := e.client.send_json_rpc[[]string, string]('explorer.Load', [network], explorer.default_timeout)!
}

// Pings the gridproxy server
pub fn (mut e ExplorerClient) ping() !string {
	return e.client.send_json_rpc[[]string, string]('explorer.Ping', []string{}, explorer.default_timeout)!
}

// Fetches grid nodes based on some filters. A list of nodes is returned and the total amount of nodes too. 
pub fn (mut e ExplorerClient) nodes(params NodesRequestParams) !NodesResult {
	return e.client.send_json_rpc[[]NodesRequestParams, NodesResult]('explorer.Nodes', [
		params,
	], explorer.default_timeout)!
}

// Fetches grid farms based on some filters. A list of farms is returned and the total amount of farms too. 
pub fn (mut e ExplorerClient) farms(params FarmsRequestParams) !FarmsResult {
	return e.client.send_json_rpc[[]FarmsRequestParams, FarmsResult]('explorer.Farms', [
		params,
	], explorer.default_timeout)!
}

// Fetches grid contracts based on some filters. A list of contracts is returned and the total amount of contracts too. 
pub fn (mut e ExplorerClient) contracts(params ContractsRequestParams) !ContractsResult {
	return e.client.send_json_rpc[[]ContractsRequestParams, ContractsResult]('explorer.Contracts',
		[params], explorer.default_timeout)!
}

// Fetches grid twins based on some filters. A list of the twins is returned and the total amount of twins too.
pub fn (mut e ExplorerClient) twins(params TwinsRequestParams) !TwinsResult {
	return e.client.send_json_rpc[[]TwinsRequestParams, TwinsResult]('explorer.Twins', [
		params,
	], explorer.default_timeout)!
}

// Gets the node with the provided id. The result object contains data relevant to the node (resources, farm, etc.)
pub fn (mut e ExplorerClient) node(node_id u32) !NodeWithNestedCapacity {
	return e.client.send_json_rpc[[]u32, NodeWithNestedCapacity]('explorer.Node', [
		node_id,
	], explorer.default_timeout)!
}

// Checks the status of node (if it is up or down).
pub fn (mut e ExplorerClient) node_status(node_id u32) !NodeStatus {
	return e.client.send_json_rpc[[]u32, NodeStatus]('explorer.NodeStatus', [node_id], explorer.default_timeout)!
}

// Counters fetches statistics of the grid (amount nodes, amount of farms, amount of contracts, etc).
pub fn (mut e ExplorerClient) counters(filters StatsFilter) !Counters {
	return e.client.send_json_rpc[[]StatsFilter, Counters]('explorer.Counters', [
		filters,
	], explorer.default_timeout)!
}
