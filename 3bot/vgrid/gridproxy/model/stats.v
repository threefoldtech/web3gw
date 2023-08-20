module model

pub struct GridStat {
pub:
	nodes              u64
	farms              u64
	countries          u64
	total_cru          u64            [json: totalCru]
	total_sru          ByteUnit       [json: totalSru]
	total_mru          ByteUnit       [json: totalMru]
	total_hru          ByteUnit       [json: totalHru]
	public_ips         u64            [json: publicIps]
	access_nodes       u64            [json: accessNodes]
	gateways           u64
	twins              u64
	contracts          u64
	nodes_distribution map[string]u64 [json: nodesDistribution]
}
