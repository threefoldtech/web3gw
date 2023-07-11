module btc

import threefoldtech.threebot.btc as btc_client { BtcClient }
import freeflowuniverse.crystallib.actionsparser { Action }
import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import log { Logger }

pub struct BTCHandler {
pub mut:
	client BtcClient
	logger Logger
}

pub fn new(mut rpc_client RpcWsClient, logger Logger) BTCHandler {
	mut client := btc_client.new(mut rpc_client)

	return BTCHandler{
		client: client
		logger: logger
	}
}

pub fn (mut h BTCHandler) handle_action(action Action) ! {
	match action.actor {
		'core' {
			// load
			h.core(action)!
		}
		'import'{
			// import_address, import_address_rescan, import_priv_key, 
			// import_priv_key_label, import_priv_key_rescan, import_pub_key
			// import_pub_key_rescan, 
			h.imports(action)!

		}
		'account'{
			// rename_account, send_to_address
			h.account(action)!
		}
		'estimate'{
			// estimate_smart_fee
			h.estimate(action)!
		}
		'blocks'{
			// generate_blocks, generate_blocks_to_address
			h.blocks(action)!
		}
		'get'{
			/*
				get_account, get_account_address, get_address_info, 
				get_addresses_by_account, get_balance, get_block_count, 
				get_block_hash, get_block_stats, get_block_verbose_tx,
				get_chain_tx_stats, get_difficulty, get_mining_info,
				get_new_address, get_node_addresses, get_peer_info,
				get_raw_transaction
			*/
			h.get(action)!
		}
		'wallet'{
			// create_wallet, move
			h.wallet(action)!
		}
		else {
			return error('action actor ${action.actor} is invalid')
		}
	}
}
