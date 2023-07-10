module stellar

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h StellarHandler) account(action Action) ! {
	match action.name {
		'address' {
			res := h.client.address()!

			h.logger.info('${res}')
		}
		else {
			return error('account action ${action.name} is invalid')
		}
	}
}
