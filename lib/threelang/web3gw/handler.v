module web3gw

import threefoldtech.threebot.tfchain { TfChainClient }
import threefoldtech.threebot.stellar { StellarClient }
import threefoldtech.threebot.eth { EthClient }
import threefoldtech.threebot.btc { BtcClient }

import log { Logger }

import freeflowuniverse.crystallib.actionsparser { Action }
import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

pub struct Web3GWHandler {
pub mut:
	tfc_client  TfChainClient
	btc_client  BtcClient
	eth_client  EthClient
	str_client  StellarClient  
	logger		Logger
}

pub fn new(mut rpc RpcWsClient, logger Logger) Web3GWHandler {
	return Web3GWHandler {
		tfc_client: tfchain.new(mut rpc)
		btc_client: btc.new(mut rpc)
		eth_client: eth.new(mut rpc)
		str_client: stellar.new(mut rpc)
		logger: logger
	}
}

pub fn (mut h Web3GWHandler) handle(action Action) ! {
	match action.actor {
		'client' { h.handle_client(action)! }
		'money' { h.handle_money(action)! }
		else { return error("unknown actor") }
	}
}