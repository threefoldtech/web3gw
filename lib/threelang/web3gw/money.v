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
			bridge_to := action.params.get_default('bridge_to', '')!
			to := action.params.get('to')!
			amount := action.params.get_f64('amount')!

			if bridge_to != '' {
				if channel == 'ethereum' && bridge_to == 'stellar' {
					res := h.eth_client.bridge_to_stellar(eth.TftEthTransfer{
						amount: amount
						destination: to
					})!
					h.logger.info(res)
				} else if channel == 'stellar' && bridge_to == 'ethereum' {
					res := h.str_client.bridge_to_eth(stellar.BridgeTransfer{
						amount: amount
						destination: to
					})
					h.logger.info(res)
				} else if channel == 'stellar' && bridge_to == 'tfchain' {
					twin_id = strconv.atoi(to) or { 0 }
					if twin_id == 0 {
						// make call for tfchain to get tht twin_id from address
						twin_id = h.tfc_client.get_twin_by_pubkey(to)!
					}
					res := h.str_client.bridge_to_tfchain(stellar.TfchainBridgeTransfer{
						amount: amount
						twin_id: twin_id
					})
					h.logger.info(res)
				} else if channel == 'tfchain' && bridge_to == 'stellar' {
					res := h.tfc_client.swap_to_stellar(stellar.SwapToStellar{
						amount: amount
						target_stellar_address: to
					})
				} else {
					return error('unsupported bridge')
				}
			} else {
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

		}
		'swap' {
			from := action.params.get('from')!
			to := action.params.get('to')!
			amount := action.params.get_f64('amount')!

			if from == 'eth' && to == 'tft' {
				res := h.eth_client.swap_eth_for_tft(amount)!
				h.logger.info(res)
			} else if from == 'tft' && to == 'eth' {
				res := h.eth_client.swap_tft_for_eth(amount)!
				h.logger.info(res)
			// } else if from == 'tft' && to == 'xlm' {
			// 	return error('not supported')
			// } else if from == 'xlm' && to == 'tft' {
			// 	return error('not supported')
			} else {
				return error('unsupported swap')
			}
		}
		'balance' {
			channel := action.params.get('channel')!
			currency := action.params.get_default('currency', '')!

			if channel == 'bitcoin' {
				account := action.params.get('account')!
				res := h.btc_client.get_balance(account)!
				h.logger.info('balance on ${channel} is ${res}')
			} else if channel == 'ethereum' && currency == 'eth' {
				address := h.eth_client.address()!
				res := h.eth_client.balance(address)!
				h.logger.info('balance on ${channel} is ${res}')
			} else if channel == 'ethereum' && currency == 'tft' {
				res := h.eth_client.tft_balance()!
				h.logger.info('balance on ${channel} is ${res}')
			} else if channel == 'stellar' {
				address := h.str_client.address()!
				res := h.str_client.balance(address)!
				h.logger.info('balance on ${channel} is ${res}')
			} else if channel == 'tfchain' {
				address := h.tfc_client.address()!
				res := h.tfc_client.balance(address)!
				h.logger.info('balance on ${channel} is ${res}')
			} else {
				return error('unsupported balance')
			}
		}
		else { error!('Unknown action: ${action.name}') }
	}
}