module explorer

import freeflowuniverse.crystallib.rpcwebsocket

const (
	default_timeout = 500000
)

// load create gridproxy client for the provided network
// - network: tf grid network name [dev, qa, test, main]
pub fn (mut client Explorer) load(network string) ! {
	_ := client.send_json_rpc[[]string, string]('explorer.Load', [network], explorer.default_timeout)!
}

// ping pings the gridproxy server 
pub fn (mut client Explorer) ping() !string {
	return client.send_json_rpc[[]string, string]('explorer.Ping', []string{}, explorer.default_timeout)!
}

// nodes fetches grid nodes based on some filters
// - params: nodes filters and paginations
// returns NodesResult which is a list of the filtered nodes and the total count of filterd nodes on the grid.
pub fn (mut client Explorer) nodes(params NodesRequestParams) !NodesResult {
	return client.send_json_rpc[[]NodesRequestParams, NodesResult]('explorer.Nodes', [
		params,
	], explorer.default_timeout)!
}

// farms fetches grid farms based on some filters
// - params: farms filters and paginations
// returns FarmsResult which is a list of the filtered farms and the total count of filterd farms on the grid.
pub fn (mut client Explorer) farms(params FarmsRequestParams) !FarmsResult {
	return client.send_json_rpc[[]FarmsRequestParams, FarmsResult]('explorer.Farms', [
		params,
	], explorer.default_timeout)!
}

// contracts fetches grid contracts based on some filters
// - params: contracts filters and paginations
// returns ContractsResult which is a list of the filtered contracts and the total count of filterd contracts on the grid.
pub fn (mut client Explorer) contracts(params ContractsRequestParams) !ContractsResult {
	return client.send_json_rpc[[]ContractsRequestParams, ContractsResult]('explorer.Contracts',
		[params], explorer.default_timeout)!
}

// twins fetches grid twins based on some filters
// - params: twins filters and paginations
// returns TwinsResult which is a list of the filtered twins and the total count of filterd twins on the grid.
pub fn (mut client Explorer) twins(params TwinsRequestParams) !TwinsResult {
	return client.send_json_rpc[[]TwinsRequestParams, TwinsResult]('explorer.Twins', [
		params,
	], explorer.default_timeout)!
}

// node fetch specific grid node
// - node_id: the reqested node id
// returns node info
pub fn (mut client Explorer) node(node_id u32) ![]NodeWithNestedCapacity {
	return client.send_json_rpc[[]u32, []NodeWithNestedCapacity]('explorer.Node', [
		node_id,
	], explorer.default_timeout)!
}

// node_status check the status of node
// - node_id: the requested node id
// returns node status [up, down]
pub fn (mut client Explorer) node_status(node_id u32) !NodeStatus {
	return client.send_json_rpc[[]u32, NodeStatus]('explorer.NodeStatus', [node_id], explorer.default_timeout)!
}

// counters fetch the total counts of grid statistics
// - filters: include/exclude up/down nodes from the counting
// returns total counts of grid statistics
pub fn (mut client Explorer) counters(filters StatsFilter) !Counters {
	return client.send_json_rpc[[]StatsFilter, Counters]('explorer.Counters', [
		filters,
	], explorer.default_timeout)!
}
