module tfgrid

[params]
pub struct FilterOptions {
	farm_id          u64  // will try to use farmerbot if found on this farm.
	public_config    bool // useful with filtering gateway nodes.
	public_ips_count u32  // free ips for the node.
	dedicated        bool // if the node is dedicated or not.
	mru              u64  // free memory on node in MB.
	hru              u64  // free hdd storage on node in GB.
	sru              u64  // free ssd storage on node in GB.
}

// This call can be used to filter nodes on specific options such as the available memory, cores, etc. It
// returns list of available nodes (their id).
pub fn (mut t TFGridClient) filter_nodes(filters FilterOptions) ![]u32 {
	return t.client.send_json_rpc[[]FilterOptions, []u32]('tfgrid.FilterNodes', [
		filters,
	], t.timeout)
}
