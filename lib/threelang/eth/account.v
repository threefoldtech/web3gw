module eth

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h EthHandler) account(action Action) ! {
	match action.name {
		'address' {
			res := h.client.address()!

			h.logger.info('${res}')
		}
		'hex_seed' {
			res := h.client.get_hex_seed()!

			h.logger.info('${res}')
		}
		else {
			return error('account action ${action.name} is invalid')
		}
	}
}
