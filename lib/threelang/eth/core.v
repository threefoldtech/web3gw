module eth

import freeflowuniverse.crystallib.actionsparser { Action }
import threefoldtech.threebot.eth { Load }

fn (mut h EthHandler) core(action Action) ! {
	match action.name {
		'load' {
			url := action.params.get('url')!
			secret := action.params.get('secret')!

			h.client.load(Load{
				url: url
				secret: secret
			})!
		}
		else {
			return error('core action ${action.name} is invalid')
		}
	}
}
