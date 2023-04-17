module explorer

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

const (
	default_timeout = 500000
)

pub fn load(mut client RpcWsClient, network string) ! {
	_ := client.send_json_rpc[[]string, string]('explorer.Load', [network], default_timeout)!
}

pub fn ping(mut client RpcWsClient) !string {
	return client.send_json_rpc[[]string, string]('explorer.Ping', []string{}, default_timeout)!
}

pub fn nodes(mut client RpcWsClient, filters NodeFilter, pagination) ![]Node {
	return client.send_json_rpc[[]NodeFilter, []Node]('explorer.Nodes', [filters], default_timeout)!
}

pub fn farms(mut client RpcWsClient) ![]Farm {
	return client.send_json_rpc[[]string, []Farm]('explorer.Farm', []string{}, default_timeout)!
}

pub fn contracts(mut client RpcWsClient) ![]Contract {
	return client.send_json_rpc[[]string, []Contract]('explorer.Contracts', []string{}, default_timeout)!
}

pub fn twins(mut client RpcWsClient) ![]Twin {
	return client.send_json_rpc[[]string, []Twin]('explorer.Twins', []string{}, default_timeout)!
}

pub fn node(mut client RpcWsClient) ![]NodeWithNestedCapacity {
	return client.send_json_rpc[[]string, []NodeWithNestedCapacity]('explorer.Node', []string{}, default_timeout)!
}

pub fn node_status(mut client RpcWsClient) !NodeStatus {
	return client.send_json_rpc[[]string, NodeStatus]('explorer.NodeStatus', []string{}, default_timeout)!
}

pub fn counters(mut client RpcWsClient) !Counters {
	return client.send_json_rpc[[]string, Counters]('explorer.Counters', []string{}, default_timeout)!
}