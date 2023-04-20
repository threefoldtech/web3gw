module explorer

import freeflowuniverse.crystallib.rpcwebsocket

const (
	default_timeout = 500000
)

pub fn (mut client Explorer) load(network string) ! {
	_ := client.send_json_rpc[[]string, string]('explorer.Load', [network], explorer.default_timeout)!
}

pub fn (mut client Explorer) ping() !string {
	return client.send_json_rpc[[]string, string]('explorer.Ping', []string{}, explorer.default_timeout)!
}

pub fn (mut client Explorer) nodes(params NodesRequestParams) ![]Node {
	return client.send_json_rpc[[]NodesRequestParams, []Node]('explorer.Nodes', [
		params,
	], explorer.default_timeout)!
}

pub fn (mut client Explorer) farms(params FarmsRequestParams) ![]Farm {
	return client.send_json_rpc[[]FarmsRequestParams, []Farm]('explorer.Farms', [
		params,
	], explorer.default_timeout)!
}

pub fn (mut client Explorer) contracts(params ContractsRequestParams) ![]Contract {
	return client.send_json_rpc[[]ContractsRequestParams, []Contract]('explorer.Contracts',
		[params], explorer.default_timeout)!
}

pub fn (mut client Explorer) twins(params TwinsRequestParams) ![]Twin {
	return client.send_json_rpc[[]TwinsRequestParams, []Twin]('explorer.Twins', [
		params,
	], explorer.default_timeout)!
}

pub fn (mut client Explorer) node(node_id u32) ![]NodeWithNestedCapacity {
	return client.send_json_rpc[[]u32, []NodeWithNestedCapacity]('explorer.Node', [
		node_id,
	], explorer.default_timeout)!
}

pub fn (mut client Explorer) node_status(node_id u32) !NodeStatus {
	return client.send_json_rpc[[]u32, NodeStatus]('explorer.NodeStatus', [node_id], explorer.default_timeout)!
}

pub fn (mut client Explorer) counters(filters StatsFilter) !Counters {
	return client.send_json_rpc[[]StatsFilter, Counters]('explorer.Counters', [
		filters,
	], explorer.default_timeout)!
}
