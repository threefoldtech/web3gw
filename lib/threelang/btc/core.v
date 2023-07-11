module btc

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h BTCHandler) core(action Action) ! {
	match action.name {
		'load' {
			host := action.params.get('host')!
			user := action.params.get('user')!
			pass := action.params.get('pass')!

			h.client.load(host: host, user: user, pass: pass)!
		}
		else {
			return error('core action ${action.name} is invalid')
		}
	}
}
