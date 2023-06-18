module btc

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h BTCHandler) get(action Action) ! {
	match action.name {
		'account' {
			address := action.params.get('address')!

			res := h.client.get_account(address)!

			h.logger.info('${res}')
		}
		'account_address' {
			account := action.params.get('account')!

			res := h.client.get_account_address(account)!

			h.logger.info('${res}')
		}
		'address_info' {
			address := action.params.get('address')!

			res := h.client.get_address_info(address)!

			h.logger.info('${res}')
		}
		'addresses_by_account' {
			account := action.params.get('account')!

			res := h.client.get_addresses_by_account(account)!

			h.logger.info('${res}')
		}
		'balance' {
			account := action.params.get('account')!

			res := h.client.get_balance(account)!

			h.logger.info('${res}')
		}
		'block_count' {
			res := h.client.get_block_count()!

			h.logger.info('${res}')
		}
		'block_hash' {
			block_height := action.params.get_int('block_height')!

			res := h.client.get_block_hash(i64(block_height))!

			h.logger.info('${res}')
		}
		'block_stats' {
			hash := action.params.get('hash')!

			res := h.client.get_block_stats(hash)!

			h.logger.info('${res}')
		}
		'block_verbose_tx' {
			hash := action.params.get('hash')!

			res := h.client.get_block_verbose_tx(hash)!

			h.logger.info('${res}')
		}
		'chain_tx_stats' {
			amount_of_blocks := action.params.get_int('amount_of_blocks')!
			block_hash_end := action.params.get('block_hash_end')!

			res := h.client.get_chain_tx_stats(
				amount_of_blocks: amount_of_blocks
				block_hash_end: block_hash_end
			)!

			h.logger.info('${res}')
		}
		'difficulty' {
			res := h.client.get_difficulty()!

			h.logger.info('${res}')
		}
		'mining_info' {
			res := h.client.get_mining_info()!

			h.logger.info('${res}')
		}
		'new_address' {
			account := action.params.get('account')!

			res := h.client.get_new_address(account)!

			h.logger.info('${res}')
		}
		'node_addresses' {
			res := h.client.get_node_addresses()!

			h.logger.info('${res}')
		}
		'peer_info' {
			res := h.client.get_peer_info()!

			h.logger.info('${res}')
		}
		'raw_transaction' {
			tx_hash := action.params.get('tx_hash')!

			res := h.client.get_raw_transaction(tx_hash)!

			h.logger.info('${res}')
		}
		else {
			return error('get action ${action.name} is invalid')
		}
	}
}
