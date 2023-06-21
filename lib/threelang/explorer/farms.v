module explorer

import freeflowuniverse.crystallib.actionparser { Action }
import threefoldtech.threebot.tfgrid { Limit, FarmsRequestParams, FarmFilter }

fn (mut h ExplorerHandler) farms(action Action) ! {
	match action.name {
		'filter' {
			free_ips := action.params.get_int_default('free_ips', 0)!
			total_ips := action.params.get_int_default('total_ips', 0)!
			stellar_address := action.params.get_default('stellar_address', '')!
			pricing_policy_id := action.params.get_default('pricing_policy_id', '')!
			farm_id := action.params.get_u32_default('farm_id', 0)!
			twin_id := action.params.get_u32_default('twin_id', 0)!
			name := action.params.get_default('name', '')!
			name_contains := action.params.get_default('name_contains', '')!
			certification_type := action.params.get_default('certification_type', '')!
			dedicated := action.params.get_default('dedicated', false)!
			
			page := action.params.get_default('page', 1)!
			size := action.params.get_default('size', 50)!
			randomize := action.params.get_default('randomize', false)!
			count := action.params.get_default('count', false)!

			req := FarmsRequestParams{
				filters: FarmFilter{
					free_ips: free_ips,
					total_ips: total_ips,
					stellar_address: stellar_address,
					pricing_policy_id: pricing_policy_id,
					farm_id: farm_id,
					twin_id: twin_id,
					name: name,
					name_contains: name_contains,
					certification_type: certification_type,
					dedicated: dedicated,
				},
				pagination: Limit{
					page: page,
					size: size,
					randomize: randomize,
					ret_count: count,
				}
			}

			res := h.explorer.farms(req)!
			h.logger.info('farms: ${res}')
		}
		else {
			return error('unknown action: ${action.name}')
		}
	}
}