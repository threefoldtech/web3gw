module tfchain

import freeflowuniverse.crystallib.actionsparser { Action }
import threefoldtech.threebot.tfchain { Transfer, SwapToStellar }

fn (mut t TFChainHandler) money(action Action) ! {
	match action.name {
		'balance' {
			channel := action.params.get_default('channel', 'tfchain')!
			currency := action.params.get_default('currency', 'tft')!
			mut address := action.params.get_default('address', '')!

			if address == '' {
				address = t.tfchain.address()!
			}

			// currently you can only get tft balance from tfchain

			if channel == 'tfchain' && currency == 'tft' {
				balance := t.tfchain.balance(address)!
				t.logger.info('balance of ${address} is ${balance}')
				return
			} else {
				return error('unsupported channel or currency')
			}
		}
		'send' {
			channel := action.params.get_default('channel', 'tfchain')!
			currency := action.params.get_default('currency', 'tft')!
			amount := action.params.get_u64('amount')!
			to := action.params.get('to')!

			if channel == 'tfchain' && currency == 'tft' {
				t.tfchain.transfer(Transfer{
					amount: amount
					destination: to
				})!
				return
			} else if channel == 'stellar' && currency == 'tft' {
				t.tfchain.swap_to_stellar(SwapToStellar{
					amount: amount
					target_stellar_address: to
				})!
				return
			}

			return error('unsupported channel or currency')

		}
		else {
			return error('core action ${action.name} is invalid')
		}
	}
}
