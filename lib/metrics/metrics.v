module metrics
// import net.http
// import json
import log
import threefoldtech.threebot.explorer


[params]
pub struct MetricsURLArgs {
pub:
	org_id   u32 = 2
	network  string
	farm_id  u32
	node_id  u32
}

pub struct Node{
	twin_id u64 [json: twinId]
}

pub struct Twin {
	account_id string [json: accountId]
}
const (
	envs = {
		'dev': 'development',
		'qa': 'qa'
		'test': 'testing',
		'main': 'production'
	
	}
)

pub fn get_metrics_url(args MetricsURLArgs, mut explorer_cl explorer.ExplorerClient, mut logger log.Logger) !string{

	env := envs[args.network] or {panic('env not found')}

	node := explorer_cl.node(args.node_id) or {
		logger.error('failed to get node: ${err}')
		exit(1)
	}

	twin_filters := explorer.TwinFilter{
		twin_id: node.twin_id
	}
	params := explorer.TwinsRequestParams{
		filters: twin_filters
	}
	res := explorer_cl.twins(params)!
	if res.twins.len < 1{
		panic("twin object node found")
	}
	
	return "https://metrics.grid.tf/d/rYdddlPWkfqwf/zos-host-metrics?orgId=${args.org_id}&refresh=30s&var-network=${env}&var-farm=${args.farm_id}&var-node=${res.twins[0].account_id}&var-diskdevices=%5Ba-z%5D%2B%7Cnvme%5B0-9%5D%2Bn%5B0-9%5D%2B%7Cmmcblk%5B0-9%5D%2B" 
}