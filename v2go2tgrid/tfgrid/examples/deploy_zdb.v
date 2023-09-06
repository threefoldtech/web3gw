module main

import log
import threefoldtech.tfgrid

fn main() {
	mut logger := log.Logger(&log.Log{
		level: .debug
	})

	mnemonics := 'route visual hundred rabbit wet crunch ice castle milk model inherit outside'
	chain_network := tfgrid.ChainNetwork.dev // User your desired network
	mut deployer := tfgrid.new_deployer(mnemonics, chain_network)!

	zdb := tfgrid.Zdb{
		size: u64(2) * 1024 * 1024
		mode: 'user'
		password: 'pass'
	}

	wl := zdb.to_workload(name: 'mywlname')

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

	node_contract_id := deployer.deploy(33, mut deployment, '', 0)!
	logger.info('node contract created with id ${node_contract_id}')
}
