module explorer

import freeflowuniverse.crystallib.actionparser { Action }

fn (mut h ExplorerHandler) nodes(action Action) ! {
	match action.name {
		'get' {
			status := action.params.get_default('status', 'up')!
			node_id := action.params.get_default('node_id', 0)!
			free_mru := action.params.get_storagecapacity_in_bytes('free_mru') or { 0 }
			free_hru := action.params.get_storagecapacity_in_bytes('free_hru') or { 0 }
			free_sru := action.params.get_storagecapacity_in_bytes('free_sru') or { 0 }
			total_mru := action.params.get_storagecapacity_in_bytes('total_mru') or { 0 }
			total_hru := action.params.get_storagecapacity_in_bytes('total_hru') or { 0 }
			total_sru := action.params.get_storagecapacity_in_bytes('total_sru') or { 0 }
			total_cru := action.params.get_default('total_cru', 0)!
			country := action.params.get_default('country', '')!
			country_contains := action.params.get_default('country_contains', '')!
			city := action.params.get_default('city', '')!
			city_contains := action.params.get_default('city_contains', '')!
			farm_name := action.params.get_default('farm_name', '')!
			farm_name_contains := action.params.get_default('farm_name_contains', '')!
			farm_id := action.params.get_default('farm_id', 0)!
			free_ips := action.params.get_default('free_ips', 0)!
			gateway := action.params.get_default('gateway', '')!
			dedicated := action.params.get_default('dedicated', false)!
			rentable := action.params.get_default('rentable', false)!
			rented := action.params.get_default('rented', false)!
			rented_by := action.params.get_default('rented_by', '')!
			available_for := action.params.get_default('available_for', '')!
			twin_id := action.params.get_default('twin_id', 0)!


			page := action.params.get_default('page', 1)!
			size := action.params.get_default('size', 50)!
			randomize := action.params.get_default('randomize', false)!
			count := action.params.get_default('count', false)!

			domain := false
			ipv4 := false
			if gateway {
				domain = true
				ipv4 = true
			}

			req := NodesRequestParams{
				filters: NodeFilter{
					id: node_id,
					status: status,
				},
				pagination: Limit{
					page: page,
					size: size,
					randomize: randomize,
					ret_count: count,
				}
			}

			res := h.explorer.nodes()!
			h.logger.info('nodes: ${res}')
		}
	}
}