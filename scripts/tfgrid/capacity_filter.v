module tfgrid

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

// the available nodes filters
[params]
pub struct FilterOptions {
	farm_id          u64 // will try to use farmerbot if found on this farm.
	public_config    bool // useful with filtering gateway nodes.
	public_ips_count u32 // free ips for the node.
	dedicated        bool // if the node is dedicated or not.
	mru              u64 // free memory on node in MB.
	hru              u64 // free hdd storage on node in GB.
	sru              u64 // free ssd storage on node in GB.
}

// the output result for filter_nodes
pub struct FilterResult {
	filter_options  FilterOptions // the passed request filters.
	available_nodes []u32 // list of available nodes ids.
}

struct Filter {
	RpcWsClient
}

// filter_nodes is filters the grid nodes with farmerbot or gridproxy
// - filters: instance of FilterOptions
// returns list of available nodes ids and the filters
pub fn (mut client Filter) filter_nodes(filters FilterOptions) !FilterResult {
	return client.send_json_rpc[[]FilterOptions, FilterResult]('tfgrid.FilterNodes', [
		filters,
	], default_timeout)
}
