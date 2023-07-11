module eth

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h EthHandler) block(action Action) ! {
	match action.name {
		'height' {
			res := h.client.height()!

			h.logger.info('${res}')
		}
		else {
			return error('block action ${action.name} is invalid')
		}
	}
}
