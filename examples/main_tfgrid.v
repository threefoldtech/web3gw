module main

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

import threefoldtech.threebot.tfgrid

import flag
import log
import os
import time

const (
	default_server_address = 'ws://127.0.0.1:8080'
)

fn test_machines_ops(mut client tfgrid.TFGridClient, mut logger log.Logger) ! {
	project_name := 'testMachinesOps'

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
			'SSH_KEY': 'ssh-rsa ...'
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

	res := client.machines_deploy(machines_model)!
	logger.info('${res}')

	// get
	time.sleep(20 * time.second)
	res_2 := client.machines_get(project_name)!
	logger.info('${res_2}')

	// delete
	client.machines_delete(project_name)!
}

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
	time.sleep(20 * time.second)
	res_2 := client.k8s_get(project_name)!
	logger.info('${res_2}')

	// delete
	client.k8s_delete(project_name)!
}

fn test_zdb_ops(mut client tfgrid.TFGridClient, mut logger log.Logger) ! {
	project_name := 'testZdbOps'

	// deploy
	zdb_model := tfgrid.ZDB{
		name: project_name
		node_id: 33
		password: 'strongPass'
		size: 10
	}

	res := client.zdb_deploy(zdb_model)!
	logger.info('${res}')

	// get
	time.sleep(10 * time.second)
	res_2 := client.zdb_get(project_name)!
	logger.info('${res_2}')

	// delete
	client.zdb_delete(project_name)!
}

fn test_name_gw_ops(mut client tfgrid.TFGridClient, mut logger log.Logger) ! {
	project_name := 'testGWNameOps'

	// deploy

	gw_model := tfgrid.GatewayName{
		name: project_name
		backends: ['http://1.1.1.1:9000']
	}

	res := client.gateways_deploy_name(gw_model)!
	logger.info('${res}')

	// get
	time.sleep(10 * time.second)
	res_2 := client.gateways_get_name(project_name)!
	logger.info('${res_2}')

	// delete
	client.gateways_delete_name(project_name)!
}

fn test_fqdn_gw_ops(mut client tfgrid.TFGridClient, mut logger log.Logger) ! {
	project_name := 'testGWFQDNOps'

	// deploy
	mut backends := []string{}
	backends << 'http://1.1.1.1:9000'
	gw_model := tfgrid.GatewayFQDN{
		name: project_name
		node_id: 11
		backends: backends
		fqdn: 'gw.test.io'
	}

	res := client.gateways_deploy_fqdn(gw_model)!
	logger.info('${res}')

	// get
	time.sleep(10 * time.second)
	res_2 := client.gateways_get_fqdn(project_name)!
	logger.info('${res_2}')

	// delete
	client.gateways_delete_fqdn(project_name)!
}

fn test_capacity_filter(mut client tfgrid.TFGridClient, mut logger log.Logger) ! {
	filters := tfgrid.FilterOptions{
		farm_id: 1
		mru: 1024 * 4
	}

	res := client.filter_nodes(filters)!
	logger.info('${res}')
}

fn test_zos_node_calls(mut client tfgrid.TFGridClient, mut logger log.Logger) ! {
	mut request := tfgrid.ZOSNodeRequest{
		node_id: 11
	}

	statistics := client.zos_node_statistics(request)!
	logger.info('node statistics: ${statistics}')

	wg_ports := client.zos_network_list_wg_ports(request)!
	logger.info('wg ports: ${wg_ports}')

	network_interfaces := client.zos_network_interfaces(request)!
	logger.info('network interfaces: ${network_interfaces}')

	public_config := client.zos_network_public_config(request)!
	logger.info('public config: ${public_config}')

	dmi := client.zos_system_dmi(request)!
	logger.info('dmi: ${dmi}')

	hypervisor := client.zos_system_hypervisor(request)!
	logger.info('hypervisor: ${hypervisor}')

	version := client.zos_system_version(request)!
	logger.info('version: ${version}')

	// deploy deployment

	deploy_deployment := tfgrid.Deployment{
		version: 0
		twin_id: 49
		contract_id: 1623847
		metadata: 'hamada_meta'
		description: 'hamada_desc'
		expiration: 1234
		signature_requirement: tfgrid.SignatureRequirement{
			weight_required: 1
			requests: [
				tfgrid.SignatureRequest{
					twin_id: 49
					required: true
					weight: 1
				},
			]
		}
		workloads: [
			tfgrid.Workload{
				version: 0
				name: 'wl12'
				workload_type: tfgrid.zdb_workload_type
				data: tfgrid.ZDBWorkload{
					password: ''
					mode: 'seq'
					size: 1
					public: false
				}
				metadata: 'hamada_meta'
				description: 'hamada_res'
			},
		]
	}

	request = tfgrid.ZOSNodeRequest{
		node_id: 11
		data: deploy_deployment
	}
	client.zos_deployment_deploy(request)!

	// update deployment

	update_deployment := tfgrid.Deployment{
		version: 2
		contract_id: 23559
		twin_id: 49
		metadata: 'hamada_meta'
		description: 'hamada_desc'
		expiration: 1234
		signature_requirement: tfgrid.SignatureRequirement{
			weight_required: 1
			requests: [
				tfgrid.SignatureRequest{
					twin_id: 49
					weight: 1
				},
			]
		}
		workloads: [
			tfgrid.Workload{
				version: 2
				name: 'wl1234'
				workload_type: tfgrid.zdb_workload_type
				data: tfgrid.ZDBWorkload{
					password: ''
					mode: 'seq'
					size: 1
					public: false
				}
				metadata: 'hamada_meta'
				description: 'hamada_res'
			},
		]
	}
	request = tfgrid.ZOSNodeRequest{
		node_id: 28
		data: update_deployment
	}
	client.zos_deployment_update(request)!

	// get deployment changes

	request = tfgrid.ZOSNodeRequest{
		node_id: 28
		data: u64(23559)
	}
	deployment_changes := client.zos_deployment_changes(request)!
	logger.info('deployment changes: ${deployment_changes}')

	// // get deployment

	request = tfgrid.ZOSNodeRequest{
		node_id: 28
		data: u64(23559)
	}
	deployment_get := client.zos_deployment_get(request)!
	logger.info('got deployment: ${deployment_get.workloads[0]}')

	// // delete deployment

	request = tfgrid.ZOSNodeRequest{
		node_id: 11
		data: u64(23559)
	}
	client.zos_deployment_delete(request)!
}

fn execute_rpcs(mut client RpcWsClient, mut logger log.Logger, mnemonic string) ! {
	mut tfgrid_client := tfgrid.new(mut client)

	// ADD YOUR CALLS HERE
	tfgrid_client.load(tfgrid.Credentials{
		mnemonic: mnemonic // FILL IN YOUR MNEMONIC HERE
		network: 'dev'
	})!

	// test_machines_ops(mut tfgrid_client, mut logger) or {
	// 	logger.error("Failed executing machines ops: $err")
	// 	exit(1)
	// }

	// test_k8s_ops(mut tfgrid_client, mut logger) or {
	// 	logger.error("Failed executing k8s ops: $err")
	// 	exit(1)
	// }

	// test_zdb_ops(mut tfgrid_client, mut logger) or {
	// 	logger.error("Failed executing zdb ops: $err")
	// 	exit(1)
	// }

	// test_name_gw_ops(mut tfgrid_client, mut logger) or {
	// 	logger.error("Failed executing name gw ops: $err")
	// 	exit(1)
	// }

	// test_fqdn_gw_ops(mut tfgrid_client, mut logger) or {
	// 	logger.error("Failed executing fqdn gw ops: $err")
	// 	exit(1)
	// }

	// test_capacity_filter(mut tfgrid_client, mut logger) or {
	// 	logger.error("Failed executing capacity filter: $err")
	// 	exit(1)
	// }

	test_zos_node_calls(mut tfgrid_client, mut logger) or {
		logger.error('Failed executing zos node calls: ${err}')
		exit(1)
	}
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
	
	
	execute_rpcs(mut myclient, mut logger, mnemonic) or {
		logger.error("Failed executing calls: $err")
		exit(1)
	}
}