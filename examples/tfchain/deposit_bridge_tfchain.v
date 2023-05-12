module main

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

import threefoldtech.threebot.stellar

import flag
import log
import os

const (
	default_server_address = 'ws://127.0.0.1:8080'
)


fn execute_rpcs(mut client RpcWsClient, mut logger log.Logger, network string, secret string, amount string, destination string, twin_id u32) ! {
	mut stellar_client := stellar.new(mut client)

	stellar_client.load(secret: secret, network: network)!

	balance := stellar_client.balance("")!
	logger.info("My balance is: ${balance}")

	// Amount in stroops (1 TFT = 10^7 stroops)
	// Destination is the tfchain address
	stellar_client.bridge_to_tfchain(amount: amount, destination: destination, twin_id: twin_id)!
}


fn main() {
	// This is some code that allows us to quickly create a commmand line tool with arguments. 
	mut fp := flag.new_flag_parser(os.args)
	fp.application('Welcome to the web3_proxy client. The web3_proxy client allows you to execute all remote procedure calls that the web3_proxy server can handle.')
	fp.limit_free_args(0, 0)!
	fp.description('')
	fp.skip_executable()
	debug_log := fp.bool('debug', 0, false, 'By setting this flag the client will print debug logs too.')
	mut logger := log.Logger(&log.Log{
		level: if debug_log { .debug } else { .info }
	})
	network := fp.string_opt('network', `n`, 'The network to choose on stellar') or {
		logger.error("Argument network is required!")
		exit(1)
	}
	secret := fp.string_opt('secret', `s`, 'Your secret on stellar') or {
		logger.error("Argument secret is required!")
		exit(1)
	}
	amount := fp.string_opt('amount', `a`, 'The amount to bridge') or {
		logger.error("Argument amount is required!")
		exit(1)
	}
	destination := fp.string_opt('destination', `d`, 'The destination address on tfchain') or {
		logger.error("Argument destination is required!")
		exit(1)
	}
	twin_id := fp.int_opt('twinid', `t`, 'Your twin id in tfchain') or {
		logger.error("Argument twinid is required!")
		exit(1)
	}	
	if twin_id <= 0 {
		logger.error("Invalid twinid, it must be greater than 0!")	
		exit(1)
	}
	address := fp.string('address', `a`, '${default_server_address}', 'The address of the web3_proxy server to connect to.')
	_ := fp.finalize() or {
		eprintln(err)
		println(fp.usage())
		exit(1)
	}

	mut myclient := rpcwebsocket.new_rpcwsclient(address, &logger) or {
		logger.error('Failed creating rpc websocket client: ${err}')
		exit(1)
	}
	_ := spawn myclient.run()
	execute_rpcs(mut myclient, mut logger, network, secret, amount, destination, u32(twin_id)) or {
		logger.error("Failed executing calls: $err")
		exit(1)
	}
}
