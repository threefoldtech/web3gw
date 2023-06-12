module tfgrid

import freeflowuniverse.crystallib.actionsparser { Action }
import threefoldtech.threebot.tfgrid { Credentials }
import freeflowuniverse.crystallib.rpcwebsocket
import threefoldtech.threebot.explorer

fn (mut t TFGridHandler) core(action Action) ! {
	match action.name {
		'login' {
			mnemonic := action.params.get_default('mnemonic', '')!
			netstring := action.params.get_default('network', 'main')!

			t.tfgrid.load(Credentials{
				mnemonic: mnemonic
				network: netstring
			})!
		}
		else {
			return error('core action ${action.name} is invalid')
		}
	}
}
