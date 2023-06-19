module eth

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h EthHandler) bridge(action Action) ! {
	match action.name {
		'stellar' {
			destination := action.params.get('destination')!
			amount := action.params.get('amount')!

			res := h.client.bridge_to_stellar(destination: destination, amount: amount)!

			h.logger.info('${res}')
		}
		else {
			return error('bridge action ${action.name} is invalid')
		}
	}
}
