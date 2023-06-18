module stellar

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h StellarHandler) account(action Action) ! {
	match action.name {
		'address' {
			res := h.client.address()!

			h.logger.info('${res}')
		}
		'transfer' {
			amount := action.params.get('amount')!
			destination := action.params.get('destination')!
			memo := action.params.get_default('memo', '')!

			res := h.client.transfer(amount: amount, destination: destination, memo: memo)!

			h.logger.info('${res}')
		}
		else {
			return error('account action ${action.name} is invalid')
		}
	}
}
