module web3gw

import threefoldtech.threebot.tfchain
import threefoldtech.threebot.stellar
import threefoldtech.threebot.eth
import threefoldtech.threebot.btc
import freeflowuniverse.crystallib.actionsparser { Action }

pub fn (mut h Web3GWHandler) handle_keys(action Action) ! {
	match action.name {
		'define' {
			tfc_mnemonic := action.params.get_default('tfc_mnemonic', '')!
			tfc_network := action.params.get_default('tfc_network', 'main')!
			if tfc_mnemonic != '' {
				h.tfc_client.load(tfchain.Load{
					network: tfc_network
					mnemonic: tfc_mnemonic
				})!
			}

			btc_host := action.params.get_default('btc_host', '')!
			btc_user := action.params.get_default('btc_user', '')!
			btc_pass := action.params.get_default('btc_pass', '')!
			if btc_host != '' && btc_user != '' && btc_pass != '' {
				h.btc_client.load(btc.Load{
					host: btc_host
					user: btc_user
					pass: btc_pass
				})!
			}

			eth_url := action.params.get_default('eth_url', '')!
			eth_secret := action.params.get_default('eth_secret', '')!
			if eth_url != '' && eth_secret != '' {
				h.eth_client.load(eth.Load{
					url: eth_url
					secret: eth_secret
				})!
			}

			str_network := action.params.get_default('str_network', 'public')!
			str_secret := action.params.get_default('str_secret', '')!
			if str_secret != '' {
				h.str_client.load(stellar.Load{
					network: str_network
					secret: str_secret
				})!
			}
		}
		else {
			return error('unknown action')
		}
	}
}
