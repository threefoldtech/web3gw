module main

import log
import threefoldtech.tfgrid

fn main() {
	mut logger := log.Log{
		level: .debug
	}

	mnemonics := tfgrid.get_mnemonics() or {
		logger.error(err.str())
		exit(1)
	}
	chain_network := tfgrid.ChainNetwork.dev // User your desired network
	mut deployer := tfgrid.new_deployer(mnemonics, chain_network, mut logger)!

	gw := tfgrid.GatewayFQDNProxy{
		tls_passthrough: false
		backends: ['http://1.1.1.1:9000']
		fqdn: 'domaind.gridtesting.xyz'
	}
	wl := gw.to_workload(name: 'mywlname')
	node_id := u32(11)
	logger.info('trying to get node ${node_id} public configuration')
	deployer.get_node_pub_config(node_id) or {
		logger.error('please select another node: ${err}')
		exit(1)
	}
	logger.info('preparing the deployment..')
	signature_requirement := tfgrid.SignatureRequirement{
		weight_required: 1
		requests: [
			tfgrid.SignatureRequest{
				twin_id: deployer.twin_id
				weight: 1
			},
		]
	}
	mut deployment := tfgrid.new_deployment(
		twin_id: deployer.twin_id
		workloads: [wl]
		signature_requirement: signature_requirement
	)

	node_contract_id := deployer.deploy(node_id, mut deployment, '', 0) or {
		logger.error(err.str())
		exit(1)
	}
	logger.info('node contract created with id ${node_contract_id}')
}
