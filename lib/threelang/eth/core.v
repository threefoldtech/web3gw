module eth

import freeflowuniverse.crystallib.actionsparser { Action }
import threefoldtech.threebot.eth { Load, Transfer }

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
		'transfer' {
			destination := action.params.get('destination')!
			amount := action.params.get('amount')!

			res := h.client.transfer(Transfer{
				destination: destination
				amount: amount
			})!

			h.logger.info('${res}')
		}
		'balance' {
			address := action.params.get('address')!

			res := h.client.balance(address)!

			h.logger.info('${res}')
		}
		'height' {
			res := h.client.height()!

			h.logger.info('${res}')
		}
		'address' {
			res := h.client.address()!

			h.logger.info('${res}')
		}
		'get_hex_seed' {
			res := h.client.get_hex_seed()!

			h.logger.info('${res}')
		}
		'token_balance' {
			contract_address := action.params.get('contract_address')!

			res := h.client.token_balance(contract_address)!

			h.logger.info('${res}')
		}
		'token_transer' {
			contract_address := action.params.get('contract_address')!
			destination := action.params.get('destination')!
			amount := action.params.get('amount')!

			res := h.client.token_transer(
				contract_address: contract_address
				destination: destination
				amount: amount
			)!

			h.logger.info('${res}')
		}
		'token_transer_from' {
			contract_address := action.params.get('contract_address')!
			from := action.params.get('from')!
			destination := action.params.get('destination')!
			amount := action.params.get('amount')!

			res := h.client.token_transer_from(
				contract_address: contract_address
				from: from
				destination: destination
				amount: amount
			)!

			h.logger.info('${res}')
		}
		'approve_token_spending' {
			contract_address := action.params.get('contract_address')!
			spender := action.params.get('from')!
			amount := action.params.get('amount')!

			res := h.client.approve_token_spending(
				contract_address: contract_address
				spender: spender
				amount: amount
			)!

			h.logger.info('${res}')
		}
		'get_multisig_owners' {
			contract_address := action.params.get('contract_address')!

			res := h.client.get_multisig_owners(contract_address)!

			h.logger.info('${res}')
		}
		'get_multisig_threshold' {
			contract_address := action.params.get('contract_address')!

			res := h.client.get_multisig_threshold(contract_address)!

			h.logger.info('${res}')
		}
		'add_multisig_owner' {
			contract_address := action.params.get('contract_address')!
			target := action.params.get('target')!
			threshold := action.params.get_int('threshold')!

			res := h.client.add_multisig_owner(
				contract_address: contract_address
				target: target
				threshold: i64(threshold)
			)!

			h.logger.info('${res}')
		}
		'remove_multisig_owner' {
			contract_address := action.params.get('contract_address')!
			target := action.params.get('target')!
			threshold := action.params.get_int('threshold')!

			res := h.client.remove_multisig_owner(
				contract_address: contract_address
				target: target
				threshold: i64(threshold)
			)!

			h.logger.info('${res}')
		}
		'approve_hash' {
			contract_address := action.params.get('contract_address')!
			hash := action.params.get('hash')!

			res := h.client.approve_hash(contract_address: contract_address, hash: hash)!

			h.logger.info('${res}')
		}
		'is_approved' {
			contract_address := action.params.get('contract_address')!
			hash := action.params.get('hash')!

			res := h.client.is_approved(contract_address: contract_address, hash: hash)!

			h.logger.info('${res}')
		}
		'initiate_multisig_eth_transfer' {
			contract_address := action.params.get('contract_address')!
			destination := action.params.get('destination')!
			amount := action.params.get('amount')!

			res := h.client.initiate_multisig_eth_transfer(
				contract_address: contract_address
				destination: destination
				amount: amount
			)!

			h.logger.info('${res}')
		}
		'initiate_multisig_token_transfer' {
			contract_address := action.params.get('contract_address')!
			token_address := action.params.get('token_address')!
			destination := action.params.get('destination')!
			amount := action.params.get('amount')!

			res := h.client.initiate_multisig_token_transfer(
				contract_address: contract_address
				token_address: token_address
				destination: destination
				amount: amount
			)!

			h.logger.info('${res}')
		}
		'get_fungible_balance' {
			contract_address := action.params.get('contract_address')!
			target := action.params.get('target')!

			res := h.client.get_fungible_balance(contract_address: contract_address, target: target)!

			h.logger.info('${res}')
		}
		'owner_of_fungible' {
			contract_address := action.params.get('contract_address')!
			token_id := action.params.get_int('token_id')!

			res := h.client.owner_of_fungible(
				contract_address: contract_address
				token_id: i64(token_id)
			)!

			h.logger.info('${res}')
		}
		'safe_transfer_fungible' {
			contract_address := action.params.get('contract_address')!
			from := action.params.get('from')!
			to := action.params.get('to')!
			token_id := action.params.get_int('token_id')!

			res := h.client.safe_transfer_fungible(
				contract_address: contract_address
				from: from
				to: to
				token_id: i64(token_id)
			)!

			h.logger.info('${res}')
		}
		'transfer_fungible' {
			contract_address := action.params.get('contract_address')!
			from := action.params.get('from')!
			to := action.params.get('to')!
			token_id := action.params.get_int('token_id')!

			res := h.client.transfer_fungible(
				contract_address: contract_address
				from: from
				to: to
				token_id: i64(token_id)
			)!

			h.logger.info('${res}')
		}
		'set_fungible_approval' {
			contract_address := action.params.get('contract_address')!
			from := action.params.get('from')!
			to := action.params.get('to')!
			amount := action.params.get_int('amount')!

			res := h.client.set_fungible_approval(
				contract_address: contract_address
				from: from
				to: to
				amount: i64(amount)
			)!

			h.logger.info('${res}')
		}
		'set_fungible_approval_for_all' {
			contract_address := action.params.get('contract_address')!
			from := action.params.get('from')!
			to := action.params.get('to')!
			approved := action.params.get_default_false('approved')

			res := h.client.set_fungible_approval_for_all(
				contract_address: contract_address
				from: from
				to: to
				approved: approved
			)!

			h.logger.info('${res}')
		}
		'get_approval_for_fungible' {
			contract_address := action.params.get('contract_address')!
			owner := action.params.get('owner')!
			operator := action.params.get('operator')!

			res := h.client.get_approval_for_fungible(
				contract_address: contract_address
				owner: owner
				operator: operator
			)!

			h.logger.info('${res}')
		}
		'get_approval_for_all_fungible' {
			contract_address := action.params.get('contract_address')!
			owner := action.params.get('owner')!
			operator := action.params.get('operator')!

			res := h.client.get_approval_for_all_fungible(
				contract_address: contract_address
				owner: owner
				operator: operator
			)!

			h.logger.info('${res}')
		}
		'transfer_eth_tft' {
			destination := action.params.get('destination')!
			amount := action.params.get('amount')!

			res := h.client.transfer_eth_tft(destination: destination, amount: amount)!

			h.logger.info('${res}')
		}
		'bridge_to_stellar' {
			destination := action.params.get('destination')!
			amount := action.params.get('amount')!

			res := h.client.bridge_to_stellar(destination: destination, amount: amount)!

			h.logger.info('${res}')
		}
		'tft_balance' {
			res := h.client.tft_balance()!

			h.logger.info('${res}')
		}
		'approve_eth_tft_spending' {
			amount := action.params.get('amount')!

			res := h.client.approve_eth_tft_spending(amount)!

			h.logger.info('${res}')
		}
		'get_eth_tft_allowance' {
			res := h.client.get_eth_tft_allowance()!

			h.logger.info('${res}')
		}
		'quote_eth_for_tft' {
			amount_in := action.params.get('amount_in')!

			res := h.client.quote_eth_for_tft(amount_in)!

			h.logger.info('${res}')
		}
		'swap_eth_for_tft' {
			amount_in := action.params.get('amount_in')!

			res := h.client.swap_eth_for_tft(amount_in)!

			h.logger.info('${res}')
		}
		'quote_tft_for_eth' {
			amount_in := action.params.get('amount_in')!

			res := h.client.quote_tft_for_eth(amount_in)!

			h.logger.info('${res}')
		}
		'swap_tft_for_eth' {
			amount_in := action.params.get('amount_in')!

			res := h.client.swap_tft_for_eth(amount_in)!

			h.logger.info('${res}')
		}
		else {
			return error('core action ${action.name} is invalid')
		}
	}
}
