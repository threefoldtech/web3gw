module web3gw

import threefoldtech.threebot.tfchain { TfChainClient }
import threefoldtech.threebot.stellar { StellarClient }
import threefoldtech.threebot.eth { EthClient }
import threefoldtech.threebot.btc { BtcClient }
import log { Logger }
import freeflowuniverse.crystallib.actionsparser { Action }
import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

import threefoldtech.threebot.threelang.clients { Clients }

[heap]
pub struct Web3GWHandler {
pub mut:
	logger     Logger
	clients    Clients
	handlers  map[string]fn(&Action)!
}

pub fn new(mut rpc RpcWsClient, logger &Logger, mut wg_clients Clients) Web3GWHandler {
	mut h := Web3GWHandler{
		logger: logger
		clients: wg_clients
	}
	h.handlers = {
		"keys": h.keys,
		"money": h.money,
	}
	return h
}

pub fn (mut h Web3GWHandler) handle(action &Action) ! {
	if action.actor in h.handlers {
		handler := h.handlers[action.actor]
		handler(action)!
	} else {
		h.logger.error("unknown actor: ${action.actor}")
	}
}
