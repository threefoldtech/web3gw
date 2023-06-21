module explorer

import freeflowuniverse.crystallib.actionparser { Action }
import threefoldtech.threebot.tfgrid { Limit, TwinsRequestParams, TwinFilter }

fn (mut h ExplorerHandler) twins(action Action) ! {
	match action.name {
		'filter' {
			twin_id := action.params.get_u32_default('twin_id', 0)!
			account_id := action.params.get_default('account_id', '')!
			relay := action.params.get_default('relay', false)!
			public_key := action.params.get_default('public_key', '')!
			
			page := action.params.get_default('page', 1)!
			size := action.params.get_default('size', 50)!
			randomize := action.params.get_default('randomize', false)!
			count := action.params.get_default('count', false)!

			req := TwinsRequestParams{
				filters: TwinFilter{
					twin_id: twin_id,
					account_id: account_id,
					relay: relay,
					public_key: public_key,
				},
				pagination: Limit{
					page: page,
					size: size,
					randomize: randomize,
					ret_count: count,
				}
			}

			res := h.explorer.twins()!
			h.logger.info('twins: ${res}')
		}
		else {
			return error('unknown action: ${action.name}')
		}
	}
}