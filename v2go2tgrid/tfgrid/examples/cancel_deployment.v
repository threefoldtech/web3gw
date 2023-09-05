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

	contract_id := u64(37459) // replace with contract id that you want to cancel
	deployer.cancel_contract(contract_id)!

	logger.info('contract ${contract_id} is canceled')
}
