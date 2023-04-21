module main

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

import explorer

import flag
import log
import os

const (
	default_server_address = 'http://127.0.0.1:8080'
)

fn test_nodes(mut client explorer.ExplorerClient, mut logger log.Logger) ! {
	node_filters := explorer.NodeFilter{
		node_id: 11
	}
	req_limit := explorer.Limit{
		size: 10
		ret_count: true
	}
	params := explorer.NodesRequestParams{
		filters: node_filters
		pagination: req_limit
	}

	nodes := client.nodes(params)!
	logger.info("nodes: ${nodes}")
}

fn test_farms(mut client explorer.ExplorerClient, mut logger log.Logger) ! {
	farm_filters := explorer.FarmFilter{
		farm_id: 1
	}
	req_limit := explorer.Limit{}
	params := explorer.FarmsRequestParams{
		filters: farm_filters
		pagination: req_limit
	}

	farms := client.farms(params)!
	logger.info("farms: ${farms}")
}

fn test_contracts(mut client explorer.ExplorerClient, mut logger log.Logger) ! {
	contract_filters := explorer.ContractFilter{
		contract_id: 4523
	}
	req_limit := explorer.Limit{
		size: 10
		ret_count: true
	}
	params := explorer.ContractsRequestParams{
		filters: contract_filters
		pagination: req_limit
	}

	contracts := client.contracts(params)!
	logger.info("contracts: ${contracts}")
}


fn test_twins(mut client explorer.ExplorerClient, mut logger log.Logger) ! {
	twin_filters := explorer.TwinFilter{
		twin_id: 29
	}
	req_limit := explorer.Limit{}

	params := explorer.TwinsRequestParams{
		filters: twin_filters
		pagination: req_limit
	}

	twins := client.twins(params)!
	logger.info("twins: ${twins}")
}

fn test_node(mut client explorer.ExplorerClient, mut logger log.Logger) ! {
	node := client.node(11)!
	logger.info("node: ${node}")
}

fn test_node_status(mut client explorer.ExplorerClient, mut logger log.Logger) ! {
	status := client.node_status(11)!
	logger.info("nodestatus: ${status}")
}

fn test_counters(mut client explorer.ExplorerClient, mut logger log.Logger) ! {
	filters := explorer.StatsFilter{
		status: 'up'
	}
	counters := client.counters(filters)!
	logger.info("counters: ${counters}")
}

fn execute_rpcs(mut client RpcWsClient, mut logger log.Logger) ! {
	mut explorer_client := explorer.new(mut client)
	explorer_client.load('dev')!

	test_nodes(mut explorer_client, mut logger)!
	// test_farms(mut explorer_client, mut logger)!
	// test_contracts(mut explorer_client, mut logger)!
	// test_twins(mut explorer_client, mut logger)!
	// test_node(mut explorer_client, mut logger)!
	// test_node_status(mut explorer_client, mut logger)!
	// test_counters(mut explorer_client, mut logger)!
}

fn main() {
	mut fp := flag.new_flag_parser(os.args)
	fp.application('Welcome to the web3_proxy client. The web3_proxy client allows you to execute all remote procedure calls that the web3_proxy server can handle.')
	fp.limit_free_args(0, 0)!
	fp.description('')
	fp.skip_executable()
	address := fp.string('address', `a`, '${default_server_address}', 'The address of the web3_proxy server to connect to.')
	debug_log := fp.bool('debug', 0, false, 'By setting this flag the client will print debug logs too.')
	_ := fp.finalize() or {
		eprintln(err)
		println(fp.usage())
		exit(1)
	}

	mut logger := log.Logger(&log.Log{
		level: if debug_log { .debug } else { .info }
	})

	mut myclient := rpcwebsocket.new_rpcwsclient(address, &logger) or {
		logger.error('Failed creating rpc websocket client: ${err}')
		exit(1)
	}

	_ := spawn myclient.run()
	
	
	execute_rpcs(mut myclient, mut logger) or {
		logger.error("Failed executing calls: $err")
		exit(1)
	}
}
