module tfchain

import freeflowuniverse.crystallib.actionsparser { Action }
import threefoldtech.threebot.tfchain { CreateFarm, PublicIPInput }

fn (mut t TFChainHandler) farms(action Action) ! {
	match action.name {
		'get' {
			id := action.params.get_u32_default('id', 0)!
			name := action.params.get_default('name', '')!

			if id != 0 {
				farm := t.tfchain.get_farm(id)!
				t.logger.info('farm ${farm}')

			} 
			if name != '' {
				farm := t.tfchain.get_farm_by_name(name)!
				t.logger.info('farm ${farm}')

			} 
			if id == 0 && name == '' {
				return error('id or name should be provided')
			}
		}
		'create' {
			name := action.params.get('name')!
			public_ips := action.params.get_default('public_ips', '')!
			gateways := action.params.get_default('gateways', '')!

			ips_arr := public_ips.split(',')
			gateways_arr := gateways.split(',')

			if ips_arr.len != gateways_arr.len {
				return error('public_ips and gateways should have the same length')
			}

			mut ips := []PublicIPInput{}
			for i in 0 .. ips_arr.len {
				ips << PublicIPInput{ip: ips_arr[i], gateway: gateways_arr[i]}
			}

			t.tfchain.create_farm(CreateFarm{name: name, public_ips: ips})!

		}
		else {
			return error('core action ${action.name} is invalid')
		}
	}
}
