module btc

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h BTCHandler) estimate(action Action) ! {
	match action.name {
		'smart_fee' {
			conf_target := action.params.get_int('conf_target')!
			mode := action.params.get_default('mode', 'CONSERVATIVE')!

			res := h.client.estimate_smart_fee(conf_target: i64(conf_target), mode: mode)!

			h.logger.info('${res}')
		}
		else {
			return error('estimate action ${action.name} is invalid')
		}
	}
}
