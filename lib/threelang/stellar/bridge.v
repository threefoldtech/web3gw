module stellar

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h StellarHandler) bridge(action Action) ! {
	match action.name {
		'to_eth' {
			amount := action.params.get('amount')!
			destination := action.params.get('destination')!

			res := h.client.bridge_to_eth(amount: amount, destination: destination)!

			h.logger.info('${res}')
		}
		'to_tfchain' {
			amount := action.params.get('amount')!
			twin_id := action.params.get_int('twin_id')!

			res := h.client.bridge_to_tfchain(amount: amount, twin_id: u32(twin_id))!

			h.logger.info('${res}')
		}
		'await_transaction_on_eth_bridge' {
			memo := action.params.get('memo')!

			h.client.await_transaction_on_eth_bridge(memo)!
		}
		else {
			return error('bridge action ${action.name} is invalid')
		}
	}
}
