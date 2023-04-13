module main

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import stellar
import tfgrid

import flag
import log
import os
import time

const (
	default_server_address = "http://127.0.0.1:8080"
)

fn test_machines_ops(mut client RpcWsClient, mut logger log.Logger) ! {
	project_name := "testMachinesOps"

	// deploy 
	mut disks := []tfgrid.Disk{}
	disks << tfgrid.Disk{
		size: 10
		mountpoint: '/mnt/disk1'
	}
	mut machines := []tfgrid.Machine{}
	machines << tfgrid.Machine{
		name: 'vm1'
		node_id: 33
		cpu: 2
		memory: 2048
		rootfs_size: 1024
		env_vars: {
			"SSH_KEY": 'ssh-rsa ...'
		}
		disks: disks
	}
	machines_model := tfgrid.MachinesModel{
		name: project_name
		network: tfgrid.Network{
			add_wireguard_access: true
		}
		machines: machines
		metadata: 'metadata1'
		description: 'description'
	}

	res := tfgrid.machines_deploy(mut client, machines_model)!
	logger.info("${res}")


	// get
	time.sleep(20 * time.second)
	res_2 := tfgrid.machines_get(mut client, project_name)!
	logger.info("${res_2}")

	// delete
	tfgrid.machines_delete(mut client, project_name)!
}

fn test_k8s_ops(mut client RpcWsClient, mut logger log.Logger) ! {
	project_name := "testK8sOps2"

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

	res := tfgrid.k8s_deploy(mut client, cluster)!
	logger.info("${res}")

	// get
	time.sleep(20 * time.second)
	res_2 := tfgrid.k8s_get(mut client, project_name)!
	logger.info("${res_2}")

	// delete
	tfgrid.k8s_delete(mut client, project_name)!
}

fn test_zdb_ops(mut client RpcWsClient, mut logger log.Logger) ! {
	project_name := "testZdbOps"

	// deploy 
	zdb_model := tfgrid.ZDB{
		name: project_name
		node_id: 33
		password: 'strongPass'
		size: 10
	}

	res := tfgrid.zdb_deploy(mut client, zdb_model)!
	logger.info("${res}")

	// get
	time.sleep(10 * time.second)
	res_2 := tfgrid.zdb_get(mut client, project_name)!
	logger.info("${res_2}")

	// delete
	tfgrid.zdb_delete(mut client, project_name)!
}

fn execute_rpcs(mut client RpcWsClient, mut logger log.Logger) ! {
	// ADD YOUR CALLS HERE
	tfgrid.load(mut client, tfgrid.Credentials{
		mnemonic: "" // FILL IN YOUR MNEMONIC HERE 
		network: "dev"
	})!

	// test_machines_ops(mut client, mut logger) or {
	// 	logger.error("Failed executing machines ops: $err")
	// 	exit(1)
	// }

	// test_k8s_ops(mut client, mut logger) or {
	// 	logger.error("Failed executing k8s ops: $err")
	// 	exit(1)
	// }

	// test_zdb_ops(mut client, mut logger) or {
	// 	logger.error("Failed executing zdb ops: $err")
	// 	exit(1)
	// }
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
		logger.error("Failed creating rpc websocket client: $err")
		exit(1)
	}
	_ := spawn myclient.run() //QUESTION: why is that in thread?
	execute_rpcs(mut myclient, mut logger) or {
		logger.error("Failed executing calls: $err")
		exit(1)
	}
}