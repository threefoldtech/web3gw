module main

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import threefoldtech.threebot.tfchain

import flag
import log
import os

const (
	default_server_address = 'ws://127.0.0.1:8080'
)

fn execute_rpcs(mut client RpcWsClient, mut logger log.Logger, network tfchain.Network, mnemonic string) ! {
	mut tfchain_client := tfchain.new(mut client)
	tfchain_client.load(network: network, mnemonic: mnemonic)!
	
	address := tfchain_client.address()!
	
	mut balance := tfchain_client.balance(address)!
	logger.info('Balance: ${balance}\n')
}

fn main() {
	mut fp := flag.new_flag_parser(os.args)
	fp.application('Welcome to the web3_proxy client. The web3_proxy client allows you to execute all remote procedure calls that the web3_proxy server can handle.')
	fp.limit_free_args(0, 0)!
	fp.description('')
	fp.skip_executable()
	debug_log := fp.bool('debug', 0, false, 'By setting this flag the client will print debug logs too.')
	mut logger := log.Logger(&log.Log{
		level: if debug_log { .debug } else { .info }
	})
	mnemonic := fp.string('mnemonic', `m`, '', 'The mnemonic to use to load tfchain.')
	network_ := fp.string('network', `n`, '', 'The network to connect to on tfchain.')
	address := fp.string('address', `a`, '${default_server_address}', 'The address of the web3_proxy server to connect to.')
	_ := fp.finalize() or {
		eprintln(err)
		println(fp.usage())
		exit(1)
	}
	mut failed_parsing := false
	if mnemonic == "" {
		logger.error("Argument mnemonic is required!")
		failed_parsing = true
	}
	network := match network_ {
		"mainnet" { tfchain.Network.mainnet }
		"testnet" { tfchain.Network.testnet }
		"qanet" { tfchain.Network.qanet }
		"devnet" { tfchain.Network.devnet }
		else { 
			logger.error("Invalid network, should be one of mainnet, testnet, qanet or devnet")
			failed_parsing = true
			tfchain.Network.devnet
		}
	}
	if failed_parsing {
		println(fp.usage())
		exit(1)
	}

	mut myclient := rpcwebsocket.new_rpcwsclient(address, &logger) or {
		logger.error('Failed creating rpc websocket client: ${err}')
		exit(1)
	}

	_ := spawn myclient.run()
	
	execute_rpcs(mut myclient, mut logger, network, mnemonic) or {
		logger.error("Failed executing calls: $err")
		exit(1)
	}
}