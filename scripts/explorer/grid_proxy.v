module explorer

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

const (
	default_timeout = 500000
)

pub fn load(mut client RpcWsClient, network string) ! {
	_ := client.send_json_rpc[[]string, string]('explorer.Load', [network], default_timeout)!
}

pub fn nodes(mut client RpcWsClient) ![]Node {
	return client.send_json_rpc[[]string, []Node]('explorer.Nodes', []string{}, default_timeout)!
}