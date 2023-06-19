module tfchain

import freeflowuniverse.crystallib.actionsparser { Action }
import threefoldtech.threebot.tfchain { ServiceContractBill, ServiceContractCreate, SetServiceContractFees, ServiceContractSetMetadata }

fn (mut t TFChainHandler) service_contract(action Action) ! {
	match action.name {
		'create' {
			service := action.params.get('service')!
			consumer := action.params.get('consumer')!

			res := t.tfchain.service_contract_create(ServiceContractCreate{
				service: service
				consumer: consumer
			})!

			t.logger.info('service contract created: ${res}')
		}
		'approve' {
			contract_id := action.params.get_u64('contract_id')!
			t.tfchain.service_contract_approve(contract_id)!
		}
		'reject' {
			contract_id := action.params.get_u64('contract_id')!
			t.tfchain.service_contract_reject(contract_id)!
		}
		'bill' {
			contract_id := action.params.get_u64('contract_id')!
			variable_amount := action.params.get_u64('variable_amount')!
			metadata := action.params.get_default('metadata', '')!

			t.tfchain.service_contract_bill(ServiceContractBill{
				contract_id: contract_id
				variable_amount: variable_amount
				metadata: metadata
			})!
		}
		'set' {
			contract_id := action.params.get_u64('contract_id')!
			variable_fee := action.params.get_u64_default('variable_fee', 0)!
			base_fee := action.params.get_u64_default('base_fee', 0)!
			metadata := action.params.get_default('metadata', '')!

			if variable_fee != 0 || base_fee != 0 {
				t.tfchain.service_contract_set_fees(SetServiceContractFees{
					contract_id: contract_id
					variable_fee: variable_fee
					base_fee: base_fee
				})!
			} 

			if metadata != '' {
				t.tfchain.service_contract_set_metadata(ServiceContractSetMetadata{
					contract_id: contract_id
					metadata: metadata
				})!
			}
		}
		'cancel' {
			contract_id := action.params.get_u64('contract_id')!
			t.tfchain.cancel_contract(contract_id)!
		}
		else {
			return error('core action ${action.name} is invalid')
		}
	}
}
