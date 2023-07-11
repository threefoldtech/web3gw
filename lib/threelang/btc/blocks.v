module btc

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h BTCHandler) blocks(action Action) ! {
	match action.name {
		'generate' {
			num_blocks := action.params.get_int('num_blocks')!

			res := h.client.generate_blocks(u32(num_blocks))!

			h.logger.info('${res}')
		}
		'generate_to_address' {
			num_blocks := action.params.get_int_default('num_blocks', 1)!
			address := action.params.get('address')!
			max_tries := action.params.get_int_default('max_tries', 1)!

			res := h.client.generate_blocks_to_address(
				num_blocks: u32(num_blocks)
				address: address
				max_tries: max_tries
			)!

			h.logger.info('${res}')
		}
		else {
			return error('blocks action ${action.name} is invalid')
		}
	}
}
