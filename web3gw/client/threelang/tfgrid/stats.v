module tfgrid

import freeflowuniverse.crystallib.baobab.actions { Action }
import threefoldtech.web3gw.explorer { StatsFilter }

pub fn (mut h TFGridHandler) stats(action Action) ! {
	match action.name {
		'get' {
			network := action.params.get_default('network', 'main')!
			h.explorer.load(network)!

			mut filter := StatsFilter{}
			if action.params.exists('status') {
				filter.status = action.params.get('status')!
			}

			res := h.explorer.counters(filter)!
			h.logger.info('stats: ${res}')
		}
		else {
			return error('explorer does not support operation: ${action.name}')
		}
	}
}
