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
	mut logger := log.Log{
		level: .info
	}

	mut tfgrid_client, _ := tfgrid.cli(mut logger) or {
		logger.error('failed to initialize tfgrid client: ${err}')
		exit(1)
	}

	test_zos_node_calls(mut tfgrid_client, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
