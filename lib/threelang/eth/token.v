module eth

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h EthHandler) token(action Action) ! {
	match action.name {
		'balance' {
			contract_address := action.params.get('contract_address')!

			res := h.client.token_balance(contract_address)!

			h.logger.info('${res}')
		}
		'transfer' {
			contract_address := action.params.get('contract_address')!
			destination := action.params.get('destination')!
			amount := action.params.get('amount')!

			res := h.client.token_transfer(
				contract_address: contract_address
				destination: destination
				amount: amount
			)!

			h.logger.info('${res}')
		}
		'transer_from' {
			contract_address := action.params.get('contract_address')!
			from := action.params.get('from')!
			destination := action.params.get('destination')!
			amount := action.params.get('amount')!

			res := h.client.token_transer_from(
				contract_address: contract_address
				from: from
				destination: destination
				amount: amount
			)!

			h.logger.info('${res}')
		}
		'approve_spending' {
			contract_address := action.params.get('contract_address')!
			spender := action.params.get('from')!
			amount := action.params.get('amount')!

			res := h.client.approve_token_spending(
				contract_address: contract_address
				spender: spender
				amount: amount
			)!

			h.logger.info('${res}')
		}
		else {
			return error('token action ${action.name} is invalid')
		}
	}
}
