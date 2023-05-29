module tfgrid

import freeflowuniverse.crystallib.rpcwebsocket
import threefoldtech.threebot.explorer
import flag
import log
import os

const (
	default_server_address = 'ws://127.0.0.1:8080'
)

pub fn cli(mut logger log.Logger) !(TFGridClient, explorer.ExplorerClient) {
	mut fp := flag.new_flag_parser(os.args)
	fp.application('Welcome to the web3_proxy client. The web3_proxy client allows you to execute all remote procedure calls that the web3_proxy server can handle.')
	fp.limit_free_args(0, 0)!
	fp.description('')
	fp.skip_executable()
	mnemonic := fp.string('mnemonic', `m`, '', 'The mnemonic to be used to call any function')
	network := fp.string('network', `n`, 'dev', 'TF network to use')
	address := fp.string('address', `a`, '${tfgrid.default_server_address}', 'The address of the web3_proxy server to connect to.')
	debug_log := fp.bool('debug', 0, false, 'By setting this flag the client will print debug logs too.')
	_ := fp.finalize() or {
		println(fp.usage())
		return error('${err}')
	}

	logger.set_level(if debug_log { .debug } else { .info })

	mut myclient := rpcwebsocket.new_rpcwsclient(address, &logger) or {
		return error('Failed creating rpc websocket client: ${err}')
	}

	_ := spawn myclient.run()

	mut tfgrid_client := new(mut myclient)
	mut exp := explorer.new(mut myclient)

	tfgrid_client.load(Credentials{
		mnemonic: mnemonic
		network: network
	})!

	exp.load(network)!

	return tfgrid_client, exp
}
