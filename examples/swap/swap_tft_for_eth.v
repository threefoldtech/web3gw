module main

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import threefoldtech.threebot.eth

import flag
import log
import os

const (
	default_server_address = 'http://127.0.0.1:8080'
)

fn execute_rpcs(mut client RpcWsClient, mut logger log.Logger, secret string) ! {
	mut eth_client := eth.new(mut client)
	eth_client.load(url:'ws://45.156.243.137:8546', secret: secret)!

	address := eth_client.address()!

	mut eth_balance := eth_client.balance(address)!
	print('eth_balance before swap: ${eth_balance}\n')

	balance := eth_client.tft_balance()!
	print('tft balance before swap: ${balance}\n')

	amount_in := "50"

	// First approve spending by uniswap router!
	print("approving ${amount_in} tft for spending\n")
	t := eth_client.approve_tft_spending(amount_in)!
	print('tx: ${t}\n')

	quote := eth_client.quote_tft_for_eth(amount_in)!
	print('will receive: ${quote} eth\n')

	tx := eth_client.swap_tft_for_eth(amount_in)!
	print('tx: ${tx}\n')

	balance_1 := eth_client.tft_balance()!
	print('tft balance after swap: ${balance_1}\n')

	eth_balance = eth_client.balance(address)!
	print('eth_balance after swap: ${eth_balance}\n')
}

fn main() {
	mut fp := flag.new_flag_parser(os.args)
	fp.application('Welcome to the web3_proxy client. The web3_proxy client allows you to execute all remote procedure calls that the web3_proxy server can handle.')
	fp.limit_free_args(0, 0)!
	fp.description('')
	fp.skip_executable()
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
	
	
	execute_rpcs(mut myclient, mut logger, secret) or {
		logger.error("Failed executing calls: $err")
		exit(1)
	}
}
