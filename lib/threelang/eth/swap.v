module eth

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h EthHandler) swap(action Action) ! {
	match action.name {
		'approve_eth_tft_spending' {
			amount := action.params.get('amount')!

			res := h.client.approve_eth_tft_spending(amount)!

			h.logger.info('${res}')
		}
		'get_eth_tft_allowance' {
			res := h.client.get_eth_tft_allowance()!

			h.logger.info('${res}')
		}
		'quote_eth_for_tft' {
			amount_in := action.params.get('amount_in')!

			res := h.client.quote_eth_for_tft(amount_in)!

			h.logger.info('${res}')
		}
		'eth_for_tft' {
			amount_in := action.params.get('amount_in')!

			res := h.client.swap_eth_for_tft(amount_in)!

			h.logger.info('${res}')
		}
		'quote_tft_for_eth' {
			amount_in := action.params.get('amount_in')!

			res := h.client.quote_tft_for_eth(amount_in)!

			h.logger.info('${res}')
		}
		'tft_for_eth' {
			amount_in := action.params.get('amount_in')!

			res := h.client.swap_tft_for_eth(amount_in)!

			h.logger.info('${res}')
		}
		else {
			return error('swap action ${action.name} is invalid')
		}
	}
}
