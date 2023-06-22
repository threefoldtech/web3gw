module web3gw

import freeflowuniverse.crystallib.actionsparser { Action }

pub fn (mut h Web3GWHandler) handle_money(action Action) ! {
	match action.name {
		'send' {
			currencies := {
				'tfchain': 'tft',
				'ethereum': 'eth',
				'bitcoin': 'btc',
				'stellar': 'tft'
			}

			channel := action.params.get('channel')!
			to := action.params.get('to')!
			amount := action.params.get_f64('amount')!

			// is it needed ??
			mut currency := action.params.get_default('currency', '')!
			from := action.params.get('from')!

			if currency == '' {
				currency = currencies[channel]
			}

			match channel {
				'bitcoin' {
					res := h.btc_client.send_to_address(btc.SendToAddress{
						address: to,
						amount: amount,
					})!
					h.logger.info(res)
				}
				'stellar' {
					res := h.str_client.transfer(str.Transfer{
						destination: to,
						amount: amount,
					})!
					h.logger.info(res)
				}
				'ethereum' {
					res := h.eth_client.transfer(eth.Transfer{
						destination: to,
						amount: amount,
					})!
					h.logger.info(res)
				}
				'tfchain' {
					res := h.tft_client.transfer(tft.Transfer{
						destination: to,
						amount: amount,
					})!
					h.logger.info(res)
				}
				else { return error('Unknown channel: ${channel}') }
			}
		}
		'swap' {
			mut currency := action.params.get_default('currency', '')!
			target_currency := action.params.get_default('target_currency', '')!
			amount := action.params.get_f64('amount')!

			// is it needed ??
			channel := action.params.get('channel')!
			to := action.params.get('to')!
			from := action.params.get('from')!			

			if currency == 'eth' && target_currency == 'tft' {
				res := h.eth_client.swap_eth_for_tft(amount)!
				h.logger.info(res)
			} else if currency == 'tft' && target_currency == 'eth' {
				res := h.eth_client.swap_tft_for_eth(amount)!
				h.logger.info(res)
			} else if currency == 'tft' && target_currency == 'xlm' {
				return error('not supported')
			} else if currency == 'xlm' && target_currency == 'tft' {
				return error('not supported')
			} else {
				return error('unsupported swap')
			}
		}
		'bridge' {
			channel := action.params.get('channel')!
			mut target_currency := action.params.get_default('target_currency', '')!
			amount := action.params.get_f64('amount')!

			// is it needed ??
			to := action.params.get('to')!
			from := action.params.get('from')!

			mut twin_id := strconv.atoi(to) or { 0 }
			if twin_id == 0 {
				// make call for tfchain to get tht twin_id from address
				twin_id := h.tfc_client.get_twin_by_pubkey(to)!
			}

			if channel == 'ethereum' && target_channel == 'stellar' {
				res := h.eth_client.bridge_to_stellar(eth.TftEthTransfer{
					amount: amount
					destination: to
				})!
				h.logger.info(res)
			} else if channel == 'stellar' && target_channel == 'ethereum' {
				res := h.str_client.bridge_to_eth(stellar.BridgeTransfer{
					amount: amount
					destination: to
				})
				h.logger.info(res)
			} else if channel == 'stellar' && target_channel == 'tfchain' {
				res := h.str_client.bridge_to_tfchain(stellar.TfchainBridgeTransfer{
					amount: amount
					twin_id: twin_id
				})
				h.logger.info(res)
			} else {
				return error('unsupported bridge')
			}
		}
		else { error!('Unknown action: ${action.name}') }
	}
}