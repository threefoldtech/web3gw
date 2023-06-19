module eth

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h EthHandler) fungible(action Action) ! {
	match action.name {
		'get_balance' {
			contract_address := action.params.get('contract_address')!
			target := action.params.get('target')!

			res := h.client.get_fungible_balance(contract_address: contract_address, target: target)!

			h.logger.info('${res}')
		}
		'get_owner' {
			contract_address := action.params.get('contract_address')!
			token_id := action.params.get_int('token_id')!

			res := h.client.owner_of_fungible(
				contract_address: contract_address
				token_id: i64(token_id)
			)!

			h.logger.info('${res}')
		}
		'safe_transfer' {
			contract_address := action.params.get('contract_address')!
			from := action.params.get('from')!
			to := action.params.get('to')!
			token_id := action.params.get_int('token_id')!

			res := h.client.safe_transfer_fungible(
				contract_address: contract_address
				from: from
				to: to
				token_id: i64(token_id)
			)!

			h.logger.info('${res}')
		}
		'transfer' {
			contract_address := action.params.get('contract_address')!
			from := action.params.get('from')!
			to := action.params.get('to')!
			token_id := action.params.get_int('token_id')!

			res := h.client.transfer_fungible(
				contract_address: contract_address
				from: from
				to: to
				token_id: i64(token_id)
			)!

			h.logger.info('${res}')
		}
		'set_approval' {
			contract_address := action.params.get('contract_address')!
			from := action.params.get('from')!
			to := action.params.get('to')!
			amount := action.params.get_int('amount')!

			res := h.client.set_fungible_approval(
				contract_address: contract_address
				from: from
				to: to
				amount: i64(amount)
			)!

			h.logger.info('${res}')
		}
		'set_approval_for_all' {
			contract_address := action.params.get('contract_address')!
			from := action.params.get('from')!
			to := action.params.get('to')!
			approved := action.params.get_default_false('approved')

			res := h.client.set_fungible_approval_for_all(
				contract_address: contract_address
				from: from
				to: to
				approved: approved
			)!

			h.logger.info('${res}')
		}
		'get_approval' {
			contract_address := action.params.get('contract_address')!
			owner := action.params.get('owner')!
			operator := action.params.get('operator')!

			res := h.client.get_approval_for_fungible(
				contract_address: contract_address
				owner: owner
				operator: operator
			)!

			h.logger.info('${res}')
		}
		'get_approval_for_all' {
			contract_address := action.params.get('contract_address')!
			owner := action.params.get('owner')!
			operator := action.params.get('operator')!

			res := h.client.get_approval_for_all_fungible(
				contract_address: contract_address
				owner: owner
				operator: operator
			)!

			h.logger.info('${res}')
		}
		else {
			return error('fungible action ${action.name} is invalid')
		}
	}
}
