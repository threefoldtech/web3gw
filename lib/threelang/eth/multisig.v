module eth

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h EthHandler) multisig(action Action) ! {
	match action.name {
		'get_owners' {
			contract_address := action.params.get('contract_address')!

			res := h.client.get_multisig_owners(contract_address)!

			h.logger.info('${res}')
		}
		'get_threshold' {
			contract_address := action.params.get('contract_address')!

			res := h.client.get_multisig_threshold(contract_address)!

			h.logger.info('${res}')
		}
		'add_owner' {
			contract_address := action.params.get('contract_address')!
			target := action.params.get('target')!
			threshold := action.params.get_int('threshold')!

			res := h.client.add_multisig_owner(
				contract_address: contract_address
				target: target
				threshold: i64(threshold)
			)!

			h.logger.info('${res}')
		}
		'remove_owner' {
			contract_address := action.params.get('contract_address')!
			target := action.params.get('target')!
			threshold := action.params.get_int('threshold')!

			res := h.client.remove_multisig_owner(
				contract_address: contract_address
				target: target
				threshold: i64(threshold)
			)!

			h.logger.info('${res}')
		}
		'approve_hash' {
			contract_address := action.params.get('contract_address')!
			hash := action.params.get('hash')!

			res := h.client.approve_hash(contract_address: contract_address, hash: hash)!

			h.logger.info('${res}')
		}
		'is_approved' {
			contract_address := action.params.get('contract_address')!
			hash := action.params.get('hash')!

			res := h.client.is_approved(contract_address: contract_address, hash: hash)!

			h.logger.info('${res}')
		}
		'initiate_eth_transfer' {
			contract_address := action.params.get('contract_address')!
			destination := action.params.get('destination')!
			amount := action.params.get('amount')!

			res := h.client.initiate_multisig_eth_transfer(
				contract_address: contract_address
				destination: destination
				amount: amount
			)!

			h.logger.info('${res}')
		}
		'initiate_token_transfer' {
			contract_address := action.params.get('contract_address')!
			token_address := action.params.get('token_address')!
			destination := action.params.get('destination')!
			amount := action.params.get('amount')!

			res := h.client.initiate_multisig_token_transfer(
				contract_address: contract_address
				token_address: token_address
				destination: destination
				amount: amount
			)!

			h.logger.info('${res}')
		}
		else {
			return error('multisig action ${action.name} is invalid')
		}
	}
}
