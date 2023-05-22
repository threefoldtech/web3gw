module main

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

import threefoldtech.threebot.tfchain

import flag
import log
import os

const (
	default_server_address = 'ws://127.0.0.1:8080'
)


fn execute_rpcs(mut client RpcWsClient, mut logger log.Logger, network string) ! {
	mut tfchain_client := tfchain.new(mut client)
	mnemonic := tfchain_client.create_account(network)!
	logger.info("Account created with mnemonic:\n------------\n${mnemonic}\n------------\n KEEP THIS SAFE!")


	my_address := tfchain_client.address()!

	balance := tfchain_client.balance(my_address)!
	logger.info("Your balance is ${balance}")
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
	network := fp.string('network', `n`, '', 'The network to choose on tfchain')
	address := fp.string('address', `a`, '${default_server_address}', 'The address of the web3_proxy server to connect to.')
	_ := fp.finalize() or {
		eprintln(err)
		println(fp.usage())
		exit(1)
	}
	
	if network == "" {
		logger.error("Argument network is required!")
		exit(1)
	}

	mut myclient := rpcwebsocket.new_rpcwsclient(address, &logger) or {
		logger.error('Failed creating rpc websocket client: ${err}')
		println(fp.usage())
		exit(1)
	}
	_ := spawn myclient.run()
	execute_rpcs(mut myclient, mut logger, network) or {
		logger.error("Failed executing calls: $err")
		exit(1)
	}
}
