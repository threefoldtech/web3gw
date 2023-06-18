module stellar

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h StellarHandler) core(action Action) ! {
	match action.name {
		'login' {
			secret := action.params.get('secret')!
			network := action.params.get_default('network', 'public')!

			h.client.load(secret: secret, network: network)!
		}
		else {
			return error('core action ${action.name} is invalid')
		}
	}
}
