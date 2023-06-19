module eth

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h EthHandler) transfer(action Action) ! {
	match action.name {
		'eth' {
			destination := action.params.get('destination')!
			amount := action.params.get('amount')!

			res := h.client.transfer(Transfer{
				destination: destination
				amount: amount
			})!

			h.logger.info('${res}')
		}
		'tft' {
			destination := action.params.get('destination')!
			amount := action.params.get('amount')!

			res := h.client.transfer_eth_tft(destination: destination, amount: amount)!

			h.logger.info('${res}')
		}
		else {
			return error('transfer action ${action.name} is invalid')
		}
	}
}
