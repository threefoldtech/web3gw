module web3gw

import freeflowuniverse.crystallib.actionsparser { Action }
import threefoldtech.threebot.tfchain
import threefoldtech.threebot.stellar
import threefoldtech.threebot.eth
import threefoldtech.threebot.btc
import strconv

const (
	default_currencies = {
		'bitcoin':  'btc'
		'ethereum': 'eth'
		'stellar':  'xlm'
		'tfchain':  'tft'
	}
	operations = {
		'send':    send
		'swap':    swap
		'balance': balance
	}
)

pub fn (mut h Web3GWHandler) money(action Action) ! {
	if action.name in web3gw.operations {
		op := web3gw.operations[action.name]
		op(mut h, action)!
	} else {
		return error('Unknown operation: ${action.name}')
	}
}

fn send(mut h Web3GWHandler, action Action) ! {
	channel := action.params.get('channel')!
	bridge_to := action.params.get_default('channel_to', '')!
	to := action.params.get('to')!
	amount := action.params.get('amount')!

	if bridge_to != '' {
		if channel == 'ethereum' && bridge_to == 'stellar' {
			res := h.clients.eth_client.bridge_to_stellar(
				amount: amount
				destination: to
			)!
			h.logger.info(res)
		} else if channel == 'stellar' && bridge_to == 'ethereum' {
			res := h.clients.str_client.bridge_to_eth(
				amount: amount
				destination: to
			)!
			h.logger.info(res)
		} else if channel == 'stellar' && bridge_to == 'tfchain' {
			mut twin_id := strconv.atoi(to) or { 0 }
			if twin_id == 0 {
				// make call for tfchain to get tht twin_id from address
				res := h.clients.tfc_client.get_twin_by_pubkey(to)!
				twin_id = int(res)
			}
			res := h.clients.str_client.bridge_to_tfchain(
				amount: amount
				twin_id: u32(twin_id)
			)!
			h.logger.info(res)
		} else if channel == 'tfchain' && bridge_to == 'stellar' {
			h.clients.tfc_client.swap_to_stellar(
				amount: amount.u64()
				target_stellar_address: to
			)!
		} else {
			return error('unsupported bridge')
		}
	} else {
		match channel {
			'bitcoin' {
				res := h.clients.btc_client.send_to_address(
					address: to
					amount: amount.i64()
				)!
				h.logger.info(res)
			}
			'stellar' {
				res := h.clients.str_client.transfer(
					destination: to
					amount: amount
				)!
				h.logger.info(res)
			}
			'ethereum' {
				res := h.clients.eth_client.transfer(
					destination: to
					amount: amount
				)!
				h.logger.info(res)
			}
			'tfchain' {
				h.clients.tfc_client.transfer(
					destination: to
					amount: amount.u64()
				)!
				h.logger.info('transfered')
			}
			else {
				return error('Unknown channel: ${channel}')
			}
		}
	}
}

fn swap(mut h Web3GWHandler, action Action) ! {
	from := action.params.get('from')!
	to := action.params.get('to')!
	amount := action.params.get('amount')!

	if from == 'eth' && to == 'tft' {
		res := h.clients.eth_client.swap_eth_for_tft(amount)!
		h.logger.info(res)
	} else if from == 'tft' && to == 'eth' {
		res := h.clients.eth_client.swap_tft_for_eth(amount)!
		h.logger.info(res)
	} else if from == 'tft' && to == 'xlm' {
		res := h.clients.str_client.swap(
			amount: amount
			source_asset: from
			destination_asset: to
		)!
		h.logger.info(res)
	} else if from == 'xlm' && to == 'tft' {
		res := h.clients.str_client.swap(
			amount: amount
			source_asset: from
			destination_asset: to
		)!
		h.logger.info(res)
	} else {
		return error('unsupported swap')
	}
}

fn balance(mut h Web3GWHandler, action Action) ! {
	channel := action.params.get('channel')!
	mut currency := action.params.get_default('currency', '')!

	if currency == '' {
		currency = web3gw.default_currencies[channel]!
	}

	if channel == 'bitcoin' {
		account := action.params.get('account')!
		res := h.clients.btc_client.get_balance(account)!
		h.logger.info('balance on ${channel} is ${res}')
	} else if channel == 'ethereum' && currency == 'eth' {
		address := h.clients.eth_client.address()!
		res := h.clients.eth_client.balance(address)!
		h.logger.info('balance on ${channel} is ${res}')
	} else if channel == 'ethereum' && currency == 'tft' {
		res := h.clients.eth_client.tft_balance()!
		h.logger.info('balance on ${channel} is ${res}')
	} else if channel == 'stellar' {
		address := h.clients.str_client.address()!
		res := h.clients.str_client.balance(address)!
		h.logger.info('balance on ${channel} is ${res}')
	} else if channel == 'tfchain' {
		address := h.clients.tfc_client.address()!
		res := h.clients.tfc_client.balance(address)!
		h.logger.info('balance on ${channel} is ${res}')
	} else {
		return error('unsupported balance')
	}
}
