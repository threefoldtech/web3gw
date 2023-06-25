module web3gw

import threefoldtech.threebot.tfchain { TfChainClient }
import threefoldtech.threebot.stellar { StellarClient }
import threefoldtech.threebot.eth { EthClient }
import threefoldtech.threebot.btc { BtcClient }
import log { Logger }
import freeflowuniverse.crystallib.actionsparser { Action }
import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

[heap]
pub struct Web3GWHandler {
pub mut:
	tfc_client TfChainClient
	btc_client BtcClient
	eth_client EthClient
	str_client StellarClient
	logger     Logger
	handlers  map[string]fn(Action)!
}

pub fn new(mut rpc RpcWsClient, logger Logger) Web3GWHandler {
	mut h := Web3GWHandler{
		tfc_client: tfchain.new(mut rpc)
		btc_client: btc.new(mut rpc)
		eth_client: eth.new(mut rpc)
		str_client: stellar.new(mut rpc)
		logger: logger
	}
	h.handlers = {
		"keys": h.handle_keys,
		"money": h.handle_money,
	}
	return h
}

pub fn (mut h Web3GWHandler) handle(action Action) ! {
	if action.actor in h.handlers {
		handler := h.handlers[action.actor]
		handler(action)!
	} else {
		h.logger.error("unknown actor: ${action.actor}")
	}
}
