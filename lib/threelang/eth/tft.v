module eth

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h EthHandler) tft(action Action) ! {
	match action.name {
		'balance' {
			res := h.client.tft_balance()!

			h.logger.info('${res}')
		}
		else {
			return error('tft action ${action.name} is invalid')
		}
	}
}
