module tfgrid

import freeflowuniverse.crystallib.actionsparser { Action }
import threefoldtech.web3gw.explorer { Limit, TwinFilter, TwinsRequestParams }

pub fn (mut h TFGridHandler) twins(action Action) ! {
	match action.name {
		'get' {
			network := action.params.get_default('network', 'main')!
			h.explorer.load(network)!

			mut filter := TwinFilter{}
			if action.params.exists('twin_id') {
				filter.twin_id = action.params.get_u64('twin_id')!
			}
			if action.params.exists('account_id') {
				filter.account_id = action.params.get('account_id')!
			}
			if action.params.exists('relay') {
				filter.relay = action.params.get('relay')!
			}
			if action.params.exists('public_key') {
				filter.public_key = action.params.get('public_key')!
			}

			page := action.params.get_u64_default('page', 1)!
			size := action.params.get_u64_default('size', 50)!
			randomize := action.params.get_default_false('randomize')
			count := action.params.get_default_false('count')

			req := TwinsRequestParams{
				filters: filter
				pagination: Limit{
					page: page
					size: size
					randomize: randomize
					ret_count: count
				}
			}

			res := h.explorer.twins(req)!
			h.logger.info('twins: ${res}')
		}
		else {
			return error('explorer does not support operation: ${action.name}')
		}
	}
}
