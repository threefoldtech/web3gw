module tfchain

import freeflowuniverse.crystallib.actionsparser { Action }
import threefoldtech.threebot.tfchain { CreateTwin }

const (
	relay_map = {
		'dev': 'wss://relay.dev.grid.tf',
		'test': 'wss://relay.test.grid.tf',
		'main': 'wss://relay.grid.tf',
		'qa': 'wss://relay.qa.grid.tf',
	}
)

fn (mut t TFChainHandler) twins(action Action) ! {
	match action.name {
		'get' {
			id := action.params.get_u32_default('id', 0)!
			pubkey := action.params.get_default('pubkey', '')!

			if id != 0 {
				twin := t.tfchain.get_twin(id)!
				t.logger.info('twin: ${twin}')

			} else if pubkey != '' {
				twin := t.tfchain.get_twin_by_pubkey(pubkey)!
				t.logger.info('twin: ${twin}')

			} else {
				return error('id or pubkey is required')
			}

		} 'create' {
			network := action.params.get_default('network', 'dev')!
			pubkey := action.params.get('pubkey')!

			relay := relay_map[network]!
			
			twin := t.tfchain.create_twin(CreateTwin{
				relay: relay
				pk: pubkey.bytes()
			})!

			t.logger.info('twin: ${twin}')

		} else {
			return error('core action ${action.name} is invalid')
		}
	}
}
