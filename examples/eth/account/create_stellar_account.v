module main

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import threefoldtech.threebot.eth
import threefoldtech.threebot.stellar

import flag
import log
import os

const (
	default_server_address = 'ws://127.0.0.1:8080'
	mainnet_ethereum_node_url = 'ws://185.69.167.224:8546'
	goerli_node_url        = 'ws://45.156.243.137:8546'
)

fn execute_rpcs(mut client RpcWsClient, mut logger log.Logger, network string, secret string) ! {
	eth_url := match network {
		"public" {
			mainnet_ethereum_node_url
		}
		"testnet" {
			goerli_node_url
		}
		else {
			return error('Invalid network ${network}')
		}
	}
	mut eth_client := eth.new(mut client)
	mut stellar_client := stellar.new(mut client)
	eth_client.load(url: eth_url, secret: secret)!

	stellar_secret := eth_client.create_and_activate_stellar_account(network)!
	logger.info("Secret: ${stellar_secret} (keep it safe!!!)")

	stellar_client.load(network: network, secret: stellar_secret)!

	address := stellar_client.address()!
	logger.info('Address: ${address}')

	balance := stellar_client.balance(address)!
	logger.info('Balance: ${balance}')
}

fn main() {
	mut fp := flag.new_flag_parser(os.args)
	fp.application('Welcome to the web3_proxy client. The web3_proxy client allows you to execute all remote procedure calls that the web3_proxy server can handle.')
	fp.limit_free_args(0, 0)!
	fp.description('')
	fp.skip_executable()
	network := fp.string('network', `n`, 'public', '')
	secret := fp.string('secret', `s`, '', 'The secret to use for eth.')
	address := fp.string('address', `a`, '${default_server_address}', 'The address of the web3_proxy server to connect to.')
	debug_log := fp.bool('debug', 0, false, 'By setting this flag the client will print debug logs too.')
	_ := fp.finalize() or {
		eprintln(err)
		println(fp.usage())
		exit(1)
	}

	mut logger := log.Logger(&log.Log{
		level: if debug_log { .debug } else { .info }
	})

	mut myclient := rpcwebsocket.new_rpcwsclient(address, &logger) or {
		logger.error('Failed creating rpc websocket client: ${err}')
		exit(1)
	}

	_ := spawn myclient.run()

	execute_rpcs(mut myclient, mut logger, network, secret) or {
		logger.error('Failed executing calls: ${err}')
		exit(1)
	}
}
