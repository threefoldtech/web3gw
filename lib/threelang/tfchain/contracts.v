module tfchain

import freeflowuniverse.crystallib.actionsparser { Action }
import threefoldtech.threebot.tfchain { CreateRentContract, CreateNodeContract, GetContractWithHash }

fn (mut t TFChainHandler) contracts(action Action) ! {
	match action.name {
		'get' {
			id := action.params.get_u64_default('id', 0)!
			if id != 0 {
				contract := t.tfchain.get_contract(id)!
				t.logger.info('contract ${contract}')
			} else {
				hash := action.params.get_default('hash', '')!
				node_id := action.params.get_u32('node_id')!
				if hash != '' {
					contract := t.tfchain.get_contract_with_hash(GetContractWithHash{
						node_id: node_id,
						hash: hash.bytes(),
					})!
					t.logger.info('contract ${contract}')
				} else {
					contracts := t.tfchain.get_node_contracts(node_id)!
					t.logger.info('contracts ${contracts}')
				}
			}
		}
		'create' {
			type_ := action.params.get('type')!
			
			match type_ {
				'name' {
					name := action.params.get('name')!
					
					res := t.tfchain.create_name_contract(name)!
					t.logger.info('created name contract ${res}')
				}
				'rent' {
					node_id := action.params.get_u32('node_id')!
					solution_provider_id := action.params.get_u64_default('solution_provider_id', 0)!

					res := t.tfchain.create_rent_contract(CreateRentContract{
						node_id: node_id,
						solution_provider_id: solution_provider_id,
					})!
					t.logger.info('created rent contract ${res}')
				}
				'node' {
					node_id := action.params.get_u32('node_id')!
					body := action.params.get_default('body', '')!
					hash := action.params.get('hash')!
					public_ip := action.params.get_u32_default('public_ip', 0)!
					solution_provider_id := action.params.get_u64_default('solution_provider_id', 0)!

					res := t.tfchain.create_node_contract(CreateNodeContract{
						node_id: node_id,
						body: body,
						hash: hash,
						public_ips: public_ip,
						solution_provider_id: solution_provider_id,
					})!
					t.logger.info('created node contract ${res}')
				}
				else {
					return error('invalid contract type ${type_}')
				}
			} 
		}
		'cancel' {
			contract_id := action.params.get_u64_default('contract_id', 0)!
			contract_ids_str := action.params.get_default('contract_ids', '')!

			if contract_id != 0 {
				t.tfchain.cancel_contract(contract_id)!
			} 
			
			if contract_ids_str != '' {
				contract_ids_map := contract_ids_str.split(',')
				mut contract_ids := []u64{}
				for id in contract_ids_map {
					cid := id.u64()
					contract_ids << cid
				}
				t.tfchain.batch_cancel_contract(contract_ids)!
			} 
			
			if contract_id == 0 && contract_ids_str == '' {
				return error('you need to provide one of the paramerters contract_id or contract_ids')
			}
		}
		else {
			return error('core action ${action.name} is invalid')
		}
	}
}
