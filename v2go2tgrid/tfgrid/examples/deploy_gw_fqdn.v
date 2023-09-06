module main

import log
import threefoldtech.tfgrid

fn main() {
	mut logger := log.Logger(&log.Log{
		level: .debug
	})

	mnemonics := '<YOUR MNEMONICS>'
	chain_network := tfgrid.ChainNetwork.dev // User your desired network
	mut deployer := tfgrid.new_deployer(mnemonics, chain_network)!

	gw := tfgrid.GatewayFQDNProxy{
		tls_passthrough: false
		backends: ['http://1.1.1.1:9000']
		fqdn: 'hamada1.3x0.me'
	}
	wl := gw.to_workload(name: 'mywlname')

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

	node_contract_id := deployer.deploy(15, mut deployment, '', 0)!
	logger.info('node contract created with id ${node_contract_id}')
}
