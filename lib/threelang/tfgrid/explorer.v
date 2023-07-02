module tfgrid

import freeflowuniverse.crystallib.actionsparser { Action }
import threefoldtech.threebot.explorer { ContractFilter, ContractsRequestParams, FarmFilter, FarmsRequestParams, Limit, NodeFilter, NodesRequestParams, StatsFilter, TwinFilter, TwinsRequestParams }

pub fn (mut h TFGridHandler) explorer(action Action) ! {
	match action.name {
		'contracts' {
			network := action.params.get_default('network', 'main')!
			h.explorer.load(network)!

			mut filter := ContractFilter{}
			filter.contract_id = action.params.get_u64('contract_id')!
			filter.twin_id = action.params.get_u64('twin_id')!
			filter.node_id = action.params.get_u64('node_id')!
			filter.type_ = action.params.get('type')!
			filter.state = action.params.get('state')!
			filter.name = action.params.get('name')!
			filter.number_of_public_ips = action.params.get_u64('number_of_public_ips')!
			filter.deployment_data = action.params.get('deployment_data')!
			filter.deployment_hash = action.params.get('deployment_hash')!

			page := action.params.get_u64_default('page', 1)!
			size := action.params.get_u64_default('size', 50)!
			randomize := action.params.get_default_false('randomize')
			count := action.params.get_default_false('count')

			req := ContractsRequestParams{
				filters: filter
				pagination: Limit{
					page: page
					size: size
					randomize: randomize
					ret_count: count
				}
			}

			res := h.explorer.contracts(req)!
			h.logger.info('contracts: ${res}')
		}
		'nodes' {
			network := action.params.get_default('network', 'main')!
			h.explorer.load(network)!

			mut filter := NodeFilter{}
			if action.params.exists('status') {
				filter.status = action.params.get('status')!
			}
			if action.params.exists('free_mru') {
				filter.free_mru = action.params.get_storagecapacity_in_bytes('free_mru')!
			}
			if action.params.exists('free_hru') {
				filter.free_hru = action.params.get_storagecapacity_in_bytes('free_hru')!
			}
			if action.params.exists('free_sru') {
				filter.free_sru = action.params.get_storagecapacity_in_bytes('free_sru')!
			}
			if action.params.exists('total_mru') {
				filter.total_mru = action.params.get_storagecapacity_in_bytes('total_mru')!
			}
			if action.params.exists('total_hru') {
				filter.total_hru = action.params.get_storagecapacity_in_bytes('total_hru')!
			}
			if action.params.exists('total_sru') {
				filter.total_sru = action.params.get_storagecapacity_in_bytes('total_sru')!
			}
			if action.params.exists('total_cru') {
				filter.total_cru = action.params.get_u64('total_cru')!
			}
			if action.params.exists('country') {
				filter.country = action.params.get('country')!
			}
			if action.params.exists('country_contains') {
				filter.country_contains = action.params.get('country_contains')!
			}
			if action.params.exists('city') {
				filter.city = action.params.get('city')!
			}
			if action.params.exists('city_contains') {
				filter.city_contains = action.params.get('city_contains')!
			}
			if action.params.exists('farm_name') {
				filter.farm_name = action.params.get('farm_name')!
			}
			if action.params.exists('farm_name_contains') {
				filter.farm_name_contains = action.params.get('farm_name_contains')!
			}
			if action.params.exists('farm_id') {
				filter.farm_ids = action.params.get_list_u64('farm_id')!
			}
			if action.params.exists('free_ips') {
				filter.free_ips = action.params.get_u64('free_ips')!
			}
			if action.params.exists('ipv4') {
				filter.ipv4 = action.params.get_default_false('ipv4')
			}
			if action.params.exists('ipv6') {
				filter.ipv6 = action.params.get_default_false('ipv6')
			}
			if action.params.exists('domain') {
				filter.domain = action.params.get_default_false('domain')
			}
			if action.params.exists('dedicated') {
				filter.dedicated = action.params.get_default_false('dedicated')
			}
			if action.params.exists('rentable') {
				filter.rentable = action.params.get_default_false('rentable')
			}
			if action.params.exists('rented') {
				filter.rented = action.params.get_default_false('rented')
			}
			if action.params.exists('rented_by') {
				filter.rented_by = action.params.get_u64('rented_by')!
			}
			if action.params.exists('available_for') {
				filter.available_for = action.params.get_u64('available_for')!
			}
			if action.params.exists('node_id') {
				filter.node_id = action.params.get_u64('node_id')!
			}
			if action.params.exists('twin_id') {
				filter.twin_id = action.params.get_u64('twin_id')!
			}

			page := action.params.get_u64_default('page', 1)!
			size := action.params.get_u64_default('size', 50)!
			randomize := action.params.get_default_false('randomize')
			count := action.params.get_default_false('count')

			req := NodesRequestParams{
				filters: filter
				pagination: Limit{
					page: page
					size: size
					randomize: randomize
					ret_count: count
				}
			}

			res := h.explorer.nodes(req)!
			h.logger.info('nodes: ${res}')
		}
		'farms' {
			network := action.params.get_default('network', 'main')!
			h.explorer.load(network)!

			mut filter := FarmFilter{}
			if action.params.exists('free_ips') {
				filter.free_ips = action.params.get_u64('free_ips')!
			}
			if action.params.exists('total_ips') {
				filter.total_ips = action.params.get_u64('total_ips')!
			}
			if action.params.exists('stellar_address') {
				filter.stellar_address = action.params.get('stellar_address')!
			}
			if action.params.exists('pricing_policy_id') {
				filter.pricing_policy_id = action.params.get_u64('pricing_policy_id')!
			}
			if action.params.exists('farm_id') {
				filter.farm_id = action.params.get_u64('farm_id')!
			}
			if action.params.exists('twin_id') {
				filter.twin_id = action.params.get_u64('twin_id')!
			}
			if action.params.exists('name') {
				filter.name = action.params.get('name')!
			}
			if action.params.exists('name_contains') {
				filter.name_contains = action.params.get('name_contains')!
			}
			if action.params.exists('certification_type') {
				filter.certification_type = action.params.get('certification_type')!
			}
			if action.params.exists('dedicated') {
				filter.dedicated = action.params.get_default_false('dedicated')
			}

			page := action.params.get_u64_default('page', 1)!
			size := action.params.get_u64_default('size', 50)!
			randomize := action.params.get_default_false('randomize')
			count := action.params.get_default_false('count')

			req := FarmsRequestParams{
				filters: filter
				pagination: Limit{
					page: page
					size: size
					randomize: randomize
					ret_count: count
				}
			}

			res := h.explorer.farms(req)!
			h.logger.info('farms: ${res}')
		}
		'stats' {
			network := action.params.get_default('network', 'main')!
			h.explorer.load(network)!
			
			mut filter := StatsFilter{}
			if action.params.exists('status') {
				filter.status = action.params.get('status')!
			}

			res := h.explorer.counters(filter)!
			h.logger.info('stats: ${res}')
		}
		'twins' {
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
		'node' {
			network := action.params.get_default('network', 'main')!
			h.explorer.load(network)!

			node_id := action.params.get_u32('node_id')!
			res := h.explorer.node(node_id)!
			h.logger.info('node: ${res}')
		}
		'node_status' {
			network := action.params.get_default('network', 'main')!
			h.explorer.load(network)!

			node_id := action.params.get_u32('node_id')!
			res := h.explorer.node_status(node_id)!
			h.logger.info('node status: ${res}')
		}
		else {
			return error('explorer does not support operation: ${action.name}')
		}
	}
}
