module main
import flag
import os
import log
import threefoldtech.threebot.metrics
import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import threefoldtech.threebot.explorer

const(
	default_server_address = 'ws://127.0.0.1:8080'
)
fn main(){

	mut fp := flag.new_flag_parser(os.args)
	fp.application('Welcome to the metrics sal.')
	fp.limit_free_args(0, 0)!
	fp.description('')
	fp.skip_executable()
	network := fp.string('network', `t`, '', 'The network to connect to. Should be development, testing or production.')
	farm_id := fp.int('farm_id', `f`, 0, 'the farm id.')
	node_id := fp.int('node_id', `n`, 0, 'The node id')
	address := fp.string('address', `a`, '${default_server_address}', 'The address of the web3_proxy server to connect to.')
	debug_log := fp.bool('debug', 0, false, 'By setting this flag the client will print debug logs too.')
	_ := fp.finalize() or {
		eprintln(err)
		println(fp.usage())
		exit(1)
	}

	args := metrics.MetricsURLArgs{
		network: network, 
		farm_id: u32(farm_id),
		node_id: u32(node_id)
	}
	mut logger := log.Logger(&log.Log{
		level: if debug_log { .debug } else { .info }
	})

	mut myclient := rpcwebsocket.new_rpcwsclient(address, &logger) or {
		logger.error('Failed creating rpc websocket client: ${err}')
		exit(1)
	}

	_ := spawn myclient.run()
	
	mut explorer_client := explorer.new(mut myclient)
	explorer_client.load(network)!

	url := metrics.get_metrics_url(args, mut explorer_client, mut logger) or {
		println("failed to construct metrics url for this node: ${err}")
		return
	}
	println(url)
	
}
