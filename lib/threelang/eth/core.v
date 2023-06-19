module eth

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h EthHandler) core(action Action) ! {
	match action.name {
		'load' {
			url := action.params.get('url')!
			secret := action.params.get('secret')!

			h.client.load(url: url, secret: secret)!
		}
		else {
			return error('core action ${action.name} is invalid')
		}
	}
}
