module tfchain

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut t TFChainHandler) account(action Action) ! {
	match action.name {
		'create' {
			network := action.params.get_default('network', 'dev')!

			new_acc := t.tfchain.create_account(network)!

			t.logger.info('created account ${new_acc}')
		}
		'address' {
			address := t.tfchain.address()!
			t.logger.info('address is ${address}')
		}
		else {
			return error('core action ${action.name} is invalid')
		}
	}
}
