module eth

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h EthHandler) balance(action Action) ! {
	match action.name {
		'get' {
			address := action.params.get('address')!

			res := h.client.balance(address)!

			h.logger.info('${res}')
		}
		else {
			return error('balance action ${action.name} is invalid')
		}
	}
}
