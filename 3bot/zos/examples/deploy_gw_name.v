module main

import log
import threefoldtech.zos

fn main() {
	mut logger := log.Logger(&log.Log{
		level: .debug
	})

	mnemonics := '<YOUR MNEMONICS>'
	chain_network := zos.ChainNetwork.dev // User your desired network
	mut deployer := zos.new_deployer(mnemonics, chain_network)!

	gw := zos.GatewayNameProxy{
		tls_passthrough: false
		backends: ['http://1.1.1.1']
		name: 'hamada_gw'
	}

	wl := gw.to_workload(name: 'hamada_gw')

	name_contract_id := deployer.create_name_contract(wl.name)!
	logger.info('name contract ${wl.name} created with id ${name_contract_id}')

	signature_requirement := zos.SignatureRequirement{
		weight_required: 1
		requests: [
			zos.SignatureRequest{
				twin_id: deployer.twin_id
				weight: 1
			},
		]
	}

	mut deployment := zos.new_deployment(
		twin_id: deployer.twin_id
		workloads: [wl]
		signature_requirement: signature_requirement
	)

	node_contract_id := deployer.deploy(11, mut deployment, '', 0)!
	logger.info('node contract created with id ${node_contract_id}')
}
