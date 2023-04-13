module main

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import stellar
import tfgrid

import flag
import log
import os
import time
import json

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

fn test_name_gw_ops(mut client RpcWsClient, mut logger log.Logger) ! {
	project_name := "testGWNameOps"

	// deploy 
	mut backends := []string{}
	backends << 'http://1.1.1.1:9000'
	gw_model := tfgrid.GatewayName{
		name: project_name
		backends: backends
	}

	res := tfgrid.gateways_deploy_name(mut client, gw_model)!
	logger.info("${res}")

	// get
	time.sleep(10 * time.second)
	res_2 := tfgrid.gateways_get_name(mut client, project_name)!
	logger.info("${res_2}")

	// delete
	tfgrid.gateways_delete_name(mut client, project_name)!
}

fn test_fqdn_gw_ops(mut client RpcWsClient, mut logger log.Logger) ! {
	project_name := "testGWFQDNOps"

	// deploy 
	mut backends := []string{}
	backends << 'http://1.1.1.1:9000'
	gw_model := tfgrid.GatewayFQDN{
		name: project_name
		node_id: 11
		backends: backends
		fqdn: 'gw.test.io'
	}

	res := tfgrid.gateways_deploy_fqdn(mut client, gw_model)!
	logger.info("${res}")

	// get
	time.sleep(10 * time.second)
	res_2 := tfgrid.gateways_get_fqdn(mut client, project_name)!
	logger.info("${res_2}")

	// delete
	tfgrid.gateways_delete_fqdn(mut client, project_name)!
}

fn test_capacity_filter(mut client RpcWsClient, mut logger log.Logger) ! {
	filters := tfgrid.FilterOptions {
		farm_id: 1 
		mru: 1024*4
	}

	res := tfgrid.filter_nodes(mut client, filters)!
	logger.info("${res}")
}

fn test_zos_node_calls(mut client RpcWsClient, mut logger log.Logger) !{
	mut request := tfgrid.ZOSNodeRequest{
		node_id: 11,
	}

	statistics := tfgrid.zos_node_statistics(mut client, request)!
	logger.info('node statistics: ${statistics}')

	wg_ports := tfgrid.zos_network_list_wg_ports(mut client, request)!
	logger.info('wg ports: ${wg_ports}')

	network_interfaces := tfgrid.zos_network_interfaces(mut client, request)!
	logger.info('network interfaces: ${network_interfaces}')

	public_config := tfgrid.zos_network_public_config(mut client, request)!
	logger.info('public config: ${public_config}')

	dmi := tfgrid.zos_system_dmi(mut client, request)!
	logger.info('dmi: ${dmi}')

	hypervisor := tfgrid.zos_system_hypervisor(mut client, request)!
	logger.info('hypervisor: ${hypervisor}')

	version := tfgrid.zos_system_version(mut client, request)!
	logger.info('version: ${version}')

	// deploy deployment

	// deploy_deployment := tfgrid.Deployment{
	// 	version: 0
	// 	twin_id: 1
	// 	metadata: 'hamada_meta'
	// 	description: 'hamada_desc'
	// 	expiration: 1234
	// 	signature_requirement: tfgrid.SignatureRequirement{
	// 		weight_required: 1
	// 		requests: [tfgrid.SignatureRequest{
	// 			twin_id: 49
	// 			required: true
	// 			weight: 1
	// 		}]
	// 	}
	// 	workloads: [tfgrid.Workload{
	// 		version: 0
	// 		name: 'wl1'
	// 		workload_type: 'invalid_type1'
	// 		data: 'hamada_data'
	// 		metadata: 'hamada_meta'
	// 		description: 'hamada_res'
	// 	}]
	// }
	// request = tfgrid.ZOSNodeRequest{
	// 	node_id: 11
	// 	data: json.encode(deploy_deployment)
	// }
	// tfgrid.zos_deployment_deploy(mut client, request)!

	// delete deployment

	// request = tfgrid.ZOSNodeRequest{
	// 	node_id: 11
	// 	data: "12345"
	// }
	// tfgrid.zos_deployment_delete(mut client, request)!

	// update deployment

	// update_deployment := tfgrid.Deployment{
	// 	version: 0
	// 	contract_id: 22226
	// 	twin_id: 49
	// 	metadata: 'hamada_meta'
	// 	description: 'hamada_desc'
	// 	expiration: 1234
	// 	signature_requirement: tfgrid.SignatureRequirement{
	// 		weight_required: 1
	// 		requests: [tfgrid.SignatureRequest{
	// 			twin_id: 1
	// 			required: true
	// 			weight: 1
	// 		}]
	// 	}
	// 	workloads: [tfgrid.Workload{
	// 		version: 0
	// 		name: 'wl1'
	// 		workload_type: 'typ1'
	// 		data: 'hamada_data'
	// 		metadata: 'hamada_meta'
	// 		description: 'hamada_res'
	// 		result: tfgrid.Result{
	// 			created: 123345
	// 			state: 'ok'
	// 			error: 'err1'
	// 			data: 'datadatadata'
	// 		}
	// 	}]
	// }
	// request = tfgrid.ZOSNodeRequest{
	// 	node_id: 11
	// 	data: json.encode(update_deployment)
	// }
	// tfgrid.zos_deployment_update(mut client, request)!


	// get deployment

	// request = tfgrid.ZOSNodeRequest{
	// 	node_id: 33
	// 	data: "123456"
	// }
	// deployment_get := tfgrid.zos_deployment_get(mut client, request)!
	// logger.info('got deployment: ${deployment_get}')

	
	// get deployment changes

	// request = tfgrid.ZOSNodeRequest{
	// 	node_id: 11
	// 	data: "54132"
	// }
	// deployment_changes := tfgrid.zos_deployment_changes(mut client, request)!
	// logger.info('deployment changes: ${deployment_changes}')

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

	// test_name_gw_ops(mut client, mut logger) or {
	// 	logger.error("Failed executing name gw ops: $err")
	// 	exit(1)
	// }

	// test_fqdn_gw_ops(mut client, mut logger) or {
	// 	logger.error("Failed executing fqdn gw ops: $err")
	// 	exit(1)
	// }

	// test_capacity_filter(mut client, mut logger) or {
	// 	logger.error("Failed executing capacity filter: $err")
	// 	exit(1)
	// }

	test_zos_node_calls(mut client, mut logger) or {
		logger.error("Failed executing zos node calls: $err")
		exit(1)
	}
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