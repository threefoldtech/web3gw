module metrics
import net.http
import json



[params]
pub struct MetricsURLArgs {
pub:
	org_id   int = 2
	network string
	farm_id    int
	node_id string
}

pub struct Node{
	twin_id int [json: twinId]
}

pub struct Twin {
	account_id string [json: accountId]
}

pub fn get_metrics_url(args MetricsURLArgs) !string{

	urls := {
	'development': 'https://gridproxy.dev.grid.tf'
	'qa': 'https://gridproxy.qa.grid.tf'
	'testing' : 'https://gridproxy.test.grid.tf'
	'production' : 'https://gridproxy.grid.tf'
	}

	url := urls[args.network] or { panic('grid proxy url not found') }

	header_config := http.HeaderConfig{
		key: http.CommonHeader.content_type
		value: 'application/json'
	}
	req := http.Request{
		method: http.Method.get
		header: http.new_header(header_config)
		url: '${url}/nodes/${args.node_id}'
	}
	resp := req.do()!
	node := json.decode(Node, resp.body)!

	twin_req := http.Request{
		method: http.Method.get
		header: http.new_header(header_config)
		url: '${url}/twins?twin_id=${node.twin_id}'
	}
	twin_resp := twin_req.do()!
	twins := json.decode([]Twin, twin_resp.body)!
	if twins.len < 1 {
		panic("twin object node found")
	}

	return "https://metrics.grid.tf/d/rYdddlPWkfqwf/zos-host-metrics?orgId=${args.org_id}&refresh=30s&var-network=${args.network}&var-farm=${args.farm_id}&var-node=${twins[0].account_id}&var-diskdevices=%5Ba-z%5D%2B%7Cnvme%5B0-9%5D%2Bn%5B0-9%5D%2B%7Cmmcblk%5B0-9%5D%2B" 
}