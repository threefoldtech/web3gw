module btc

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h BTCHandler) wallet(action Action) ! {
	match action.name {
		'create' {
			name := action.params.get('name')!
			disable_private_keys := action.params.get_default_false('disable_private_keys')
			create_blank_wallet := action.params.get_default_false('create_blank_wallet')
			passphrase := action.params.get('passphrase')!
			avoid_reuse := action.params.get_default_false('avoid_reuse')

			res := h.client.create_wallet(
				name: name
				disable_private_keys: disable_private_keys
				create_blank_wallet: create_blank_wallet
				passphrase: passphrase
				avoid_reuse: avoid_reuse
			)!

			h.logger.info('${res}')
		}
		'move' {
			from_account := action.params.get('from_account')!
			to_account := action.params.get('to_account')!
			amount := action.params.get_int('amount')!
			min_confirmations := action.params.get_int('min_confirmations')!
			comment := action.params.get_default('comment', '')!

			res := h.client.move(
				from_account: from_account
				to_account: to_account
				amount: i64(amount)
				min_confirmations: min_confirmations
				comment: comment
			)!

			h.logger.info('${res}')
		}
		else {
			return error('wallet action ${action.name} is invalid')
		}
	}
}
