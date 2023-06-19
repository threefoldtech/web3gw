module tfchain

import freeflowuniverse.crystallib.actionsparser { Action }
import threefoldtech.threebot.tfchain { Load }

fn (mut t TFChainHandler) client(action Action) ! {
	match action.name {
		'load' {
			network := action.params.get_default('network', 'dev')!

			mnemonic := action.params.get('mnemonic')!

			t.tfchain.load(Load{
				network: network
				mnemonic: mnemonic
			})!
		}
		else {
			return error('core action ${action.name} is invalid')
		}
	}
}
