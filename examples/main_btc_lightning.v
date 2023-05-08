module main

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import threefoldtech.threebot.btclightning

import flag
import log
import os

const (
	default_server_address = 'http://127.0.0.1:8080'
)

fn execute_rpcs(mut client RpcWsClient, mut logger log.Logger) ! {
	mut btc_lightning_client := btclightning.new(mut client)

	chain_service_cfg := btclightning.ChainServiceConfig{

	}
	_ := btc_lightning_client.new_chain_service(chain_service_cfg)!

	best_block := btc_lightning_client.best_block()!
	logger.info("Best block: ${best_block}")

	height := 100
	block_hash := btc_lightning_client.get_block_hash(height)!
	logger.info("Block hash: ${block_hash}")

	hash := "abc"
	block_header := btc_lightning_client.get_block_header(hash)!
	logger.info("Block header: ${block_header}")


	block_height := btc_lightning_client.get_block_height(hash)!
	logger.info("Block height: ${block_height}")

	ban_peer_info := btclightning.BanPeerInfo{

	}
	_ := btc_lightning_client.ban_peer(ban_peer_info)!
	
	peer_address := "abc"
	is_banned := btc_lightning_client.is_banned(peer_address)!
	logger.info("Is banned: ${is_banned}")
	
	is_server_peer_persistent := true
	_ := btc_lightning_client.add_peer(is_server_peer_persistent)!
	
	bytes_sent := 100
	_ := btc_lightning_client.add_bytes_sent(bytes_sent)!
	
	bytes_received := 100
	_ := btc_lightning_client.add_bytes_received(bytes_received)!
	
	net_totals := btc_lightning_client.net_totals()!
	logger.info("Net totals: ${net_totals}")

	update_peer_heights_request := btclightning.UpdatePeerHeightsRequest{

	}
	_ := btc_lightning_client.update_peer_heights(update_peer_heights_request)!
	
	_ := btc_lightning_client.start_chain_service()!
	
	_ := btc_lightning_client.stop_chain_service()!
	
	is_current := btc_lightning_client.is_current()!
	logger.info("Is current: ${is_current}")

	rescan_options := btclightning.RescanOptions{

	}
	_ := btc_lightning_client.new_rescan(rescan_options)!
	
	_ := btc_lightning_client.start_rescan()!
	
	rescan_update_options := btclightning.RescanUpdateOptions{

	}
	_ := btc_lightning_client.udpate_rescan(rescan_update_options)!
	
	_ := btc_lightning_client.shutdown_rescan()!
}

fn main() {
	mut fp := flag.new_flag_parser(os.args)
	fp.application('Welcome to the web3_proxy client. The web3_proxy client allows you to execute all remote procedure calls that the web3_proxy server can handle.')
	fp.limit_free_args(0, 0)!
	fp.description('')
	fp.skip_executable()
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

	execute_rpcs(mut myclient, mut logger) or {
		logger.error('Failed executing calls: ${err}')
		exit(1)
	}
}
