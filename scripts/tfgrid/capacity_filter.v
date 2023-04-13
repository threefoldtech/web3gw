module tfgrid

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

[params]
pub struct FilterOptions {
	farm_id          u64
	public_config    bool
	public_ips_count u32
	dedicated        bool
	mru              u64
	hru              u64
	sru              u64
}

pub struct FilterResult {
	filter_options  FilterOptions
	available_nodes []u32
}

pub fn filter_nodes(mut client RpcWsClient, filters FilterOptions) !FilterResult {
	return client.send_json_rpc[[]FilterOptions, FilterResult]('tfgrid.FilterNodes', [
		filters,
	], default_timeout)
}
