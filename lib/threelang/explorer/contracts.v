module explorer

import freeflowuniverse.crystallib.actionparser { Action }
import threefoldtech.threebot.tfgrid { Limit, ContractsRequestParams, ContractFilter }

fn (mut h ExplorerHandler) contracts(action Action) ! {
	match action.name {
		'filter' {
			contract_id := action.params.get_u32_default('contract_id', 0)!
			twin_id := action.params.get_u32_default('twin_id', 0)!
			node_id := action.params.get_u32_default('node_id', 0)!
			type_ := action.params.get_default('type', '')!
			state := action.params.get_default('state', '')!
			name := action.params.get_default('name', '')!
			number_of_public_ips := action.params.get_u32_default('number_of_public_ips', 0)!
			deployment_data := action.params.get_default('deployment_data', '')!
			deployment_hash := action.params.get_default('deployment_hash', '')!
			
			page := action.params.get_default('page', 1)!
			size := action.params.get_default('size', 50)!
			randomize := action.params.get_default('randomize', false)!
			count := action.params.get_default('count', false)!

			req := ContractsRequestParams{
				filters: ContractFilter{
					contract_id: contract_id,
					twin_id: twin_id,
					node_id: node_id,
					type_: type_,
					state: state,
					name: name,
					number_of_public_ips: number_of_public_ips,
					deployment_data: deployment_data,
					deployment_hash: deployment_hash,
				},
				pagination: Limit{
					page: page,
					size: size,
					randomize: randomize,
					ret_count: count,
				}
			}

			res := h.explorer.contracts(req)!
			h.logger.info('contracts: ${res}')
		}
		else {
			return error('unknown action: ${action.name}')
		}
	}
}