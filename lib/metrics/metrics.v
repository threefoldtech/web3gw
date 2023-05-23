module metrics

[params]
pub struct MetricsURLArgs {
pub:
	org_id   int = 2
	network string
	farm_id    int
	node_key string
}


pub fn get_metrics_url(args MetricsURLArgs) string{
	return "https://metrics.grid.tf/d/rYdddlPWkfqwf/zos-host-metrics?orgId=${args.org_id}&refresh=30s&var-network=${args.network}&var-farm=${args.farm_id}&var-node=${args.node_key}&var-diskdevices=%5Ba-z%5D%2B%7Cnvme%5B0-9%5D%2Bn%5B0-9%5D%2B%7Cmmcblk%5B0-9%5D%2B" 
}