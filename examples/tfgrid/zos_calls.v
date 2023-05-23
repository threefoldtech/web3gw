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

	test_zos_node_calls(mut tfgrid_client, mut logger) or {
		logger.error('Failed executing zos node calls: ${err}')
		exit(1)
	}
}
