module main

import freeflowuniverse.crystallib.rpcwebsocket
import threefoldtech.threebot.tfgrid
import flag
import log
import os
import time

const (
	default_server_address = 'ws://127.0.0.1:8080'
)

fn test_k8s_ops(mut client tfgrid.TFGridClient, mut logger log.Logger) ! {
	project_name := 'testK8sOps2'

	// deploy
	master := tfgrid.K8sNode{
		name: 'master'
		node_id: 33
		cpu: 1
		memory: 1024
	}

	mut workers := []tfgrid.K8sNode{}
	workers << tfgrid.K8sNode{
		name: 'w1'
		node_id: 33
		cpu: 1
		memory: 1024
	}

	cluster := tfgrid.K8sCluster{
		name: project_name
		token: 'token6'
		ssh_key: 'SSH-Key'
		master: master
		workers: workers
	}

	res := client.k8s_deploy(cluster)!
	logger.info('${res}')

	// get
	res_2 := client.k8s_get(project_name)!
	logger.info('${res_2}')

	// delete
	client.k8s_delete(project_name)!
}

fn main() {
	mut fp := flag.new_flag_parser(os.args)
	fp.application('Welcome to the web3_proxy client. The web3_proxy client allows you to execute all remote procedure calls that the web3_proxy server can handle.')
	fp.limit_free_args(0, 0)!
	fp.description('')
	fp.skip_executable()
	mnemonic := fp.string('mnemonic', `m`, '', 'The mnemonic to be used to call any function')
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

	mut tfgrid_client := tfgrid.new(mut client)

	// ADD YOUR CALLS HERE
	tfgrid_client.load(tfgrid.Credentials{
		mnemonic: mnemonic
		network: 'dev'
	})!

	test_k8s_ops(mut tfgrid_client, mut logger) or {
		logger.error('Failed executing k8s ops: ${err}')
		exit(1)
	}
}
