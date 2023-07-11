module tfchain

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut t TFChainHandler) metadata(action Action) ! {
	match action.name {
		'zos_version' {
			version := t.tfchain.get_zos_version()!
			t.logger.info('zos version: ${version}')

		} 'chain_height' {
			height := t.tfchain.height()!
			t.logger.info('chain height: ${height}')

		} else {
			return error('core action ${action.name} is invalid')
		}
	}
}
