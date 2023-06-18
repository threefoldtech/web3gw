module btc

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut h BTCHandler) imports(action Action) ! {
	match action.name {
		'address' {
			address := action.params.get('address')!

			h.client.import_address(address)!
		}
		'address_rescan' {
			address := action.params.get('address')!
			account := action.params.get('account')!
			rescan := action.params.get_default_false('rescan')

			h.client.import_address_rescan(address: address, account: account, rescan: rescan)!
		}
		'priv_key' {
			wif := action.params.get('wif')!

			h.client.import_priv_key(wif)!
		}
		'priv_key_label' {
			wif := action.params.get('wif')!
			label := action.params.get('label')!

			h.client.import_priv_key_label(wif: wif, label: label)!
		}
		'priv_key_rescan' {
			wif := action.params.get('wif')!
			label := action.params.get('label')!
			rescan := action.params.get_default_false('rescan')

			h.client.import_priv_key_rescan(wif: wif, label: label, rescan: rescan)!
		}
		'pub_key' {
			pub_key := action.params.get('pub_key')!

			h.client.import_pub_key(pub_key)!
		}
		'pub_key_rescan' {
			pub_key := action.params.get('pub_key')!
			rescan := action.params.get_default_false('rescan')

			h.client.import_pub_key_rescan(pub_key: pub_key, rescan: rescan)!
		}
		else {
			return error('imports action ${action.name} is invalid')
		}
	}
}
