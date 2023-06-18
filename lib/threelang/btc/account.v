module btc

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h BTCHandler) account(action Action) ! {
	match action.name {
		'rename' {
			old_account := action.params.get('old_account')!
			new_account := action.params.get('new_account')!

			h.client.rename_account(old_account: old_account, new_account: new_account)!
		}
		'send_to_address' {
			address := action.params.get('address')!
			amount := action.params.get_int('amount')!
			comment := action.params.get_default('comment', '')!
			comment_to := action.params.get_default('comment_to', '')!

			res := h.client.send_to_address(
				address: address
				amount: i64(amount)
				comment: comment
				comment_to: comment_to
			)!

			h.logger.info('${res}')
		}
		else {
			return error('account action ${action.name} is invalid')
		}
	}
}
