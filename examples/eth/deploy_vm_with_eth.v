module main

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import threefoldtech.threebot.eth
import threefoldtech.threebot.stellar
import threefoldtech.threebot.tfchain

import flag
import log
import os

const (
	default_server_address = 'http://127.0.0.1:8080'
	goerli_node_url = 'ws://45.156.243.137:8546'
)

[params]
pub struct Arguments {
	eth_secret 			string
	eth_url 			string

	stellar_secret      string
	stellar_network     string

	tfchain_network 	tfchain.Network
	tfchain_mnemonic 	string
}

fn execute_rpcs(mut client RpcWsClient, mut logger log.Logger, args Arguments) ! {
	mut eth_client := eth.new(mut client)
	mut tfchain_client := tfchain.new(mut client)
	mut stellar_client := stellar.new(mut client)

	eth_client.load(url: args.eth_url, secret: args.eth_secret)!
	tfchain_client.load(network: args.tfchain_network, mnemonic: args.tfchain_mnemonic)!
	stellar_client.load(network: args.stellar_network, secret: args.stellar_secret)!

	address := eth_client.address()!

	mut eth_balance := eth_client.balance(address)!
	logger.info('eth balance before swap: ${eth_balance}')
	
	mut eth_tft_balance := eth_client.tft_balance()!
	logger.info('tft balance before swap: ${eth_tft_balance}')
	
	/*
	amount_in := "0.0001"

	quote := eth_client.quote_eth_for_tft(amount_in)!
	logger.info('will receive: ${quote} tft')

	tx := eth_client.swap_eth_for_tft(amount_in)!
	logger.info('tx: ${tx}')
	*/
	quote := 183943396
	
	/*
	eth_tft_balance = eth_client.tft_balance()!
	logger.info('tft balance after swap: ${eth_tft_balance}')

	eth_balance = eth_client.balance(address)!
	logger.info('eth balance after swap: ${eth_balance}')

	eth_client.withdraw_eth_tft_to_stellar(destination: args.stellar_address, amount: quote)!
	logger.info('withdrawn eth to stellar')

	eth_tft_balance = eth_client.tft_balance()!
	logger.info('tft balance after bridge: ${eth_tft_balance}')
	*/
	stellar_address := stellar_client.address()!

	stellar_balance := stellar_client.balance(stellar_address)!
	logger.info('stellar balance after bridge: ${stellar_balance}')

	tfchain_address := tfchain_client.address()!
	logger.info('Bridge to tfchain address ${tfchain_address}')

	tfchain_twinid := tfchain_client.get_twin_by_pubkey(tfchain_address)!
	logger.info('Twin ID on tfchain is ${tfchain_twinid}')

	/*tfchain_client.bridge_to_tfchain(amount:stellar_balance.str(), twin_id:tfchain_twinid)!
	logger.info('Bridge to tfchain done')*/

	tfchain_balance := tfchain_client.balance(tfchain_address)!
	logger.info('Balance on TF Chain is ${tfchain_balance}')
}

fn main() {
	mut fp := flag.new_flag_parser(os.args)
	fp.application('TODO')
	fp.limit_free_args(0, 0)!
	fp.description('')
	fp.skip_executable()
	
	address := fp.string('address', `a`, '${default_server_address}', 'The address of the web3_proxy server to connect to.')
	debug_log := fp.bool('debug', 0, false, 'By setting this flag the client will print debug logs too.')

	eth_secret := fp.string('eth-secret', 0, '', 'The secret to use for eth.')
	eth_url := fp.string('eth-node', 0, '${goerli_node_url}', 'The url of the ethereum node to connect to.')

	stellar_secret := fp.string('stellar-secret', 0, '', 'The secret of the stellar address to send the TFT to.')
	stellar_network := fp.string('stellar-network', 0, '', 'The stellar network of the provided stellar address.')

	tfchain_mnemonic := fp.string('tfchain-mnemonic', 0, '', 'The mnemonic of your tfchain account.')
	tfchain_network := fp.string('tfchain-network', 0, '', 'The tfchain network to use.')

	//amount := fp.int('amount', `m`, 0, 'The amount of TFT to send to the destination address.')

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


	arguments := Arguments {
		eth_secret: eth_secret
		eth_url: eth_url

		stellar_secret: stellar_secret
		stellar_network: stellar_network

		tfchain_network: tfchain.parse_network(tfchain_network) or {
			logger.error('${err}')
			exit(1)
		}
		tfchain_mnemonic: tfchain_mnemonic
	}

	execute_rpcs(mut myclient, mut logger, arguments) or {
		logger.error("Failed executing calls: $err")
		exit(1)
	}
}