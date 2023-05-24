module main
import flag
import os
import threefoldtech.threebot.metrics

fn main(){

	mut fp := flag.new_flag_parser(os.args)
	fp.application('Welcome to the metrics sal.')
	fp.limit_free_args(0, 0)!
	fp.description('')
	fp.skip_executable()
	network := fp.string('network', `t`, '', 'The network to connect to. Should be development, testing or production.')
	farm_id := fp.int('farm_id', `f`, 0, 'the farm id.')
	node_id := fp.int('node_id', `n`, 0, 'The node id')

	_ := fp.finalize() or {
		eprintln(err)
		println(fp.usage())
		exit(1)
	}

	args := metrics.MetricsURLArgs{
		network: network, 
		farm_id:farm_id, 
		node_id: node_id
	}

	url := metrics.get_metrics_url(args) or {
		println("failed to construct metrics url for this node")
		return
	}
	println(url)
	
}
