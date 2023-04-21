module main

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

import eth
import explorer
import stellar
import tfchain
import tfgrid
import nostr
import explorer
//ADD NEW CLIENTS HERE

import flag
import log
import os
import time

const (
	default_server_address = 'http://127.0.0.1:8080'
)


fn execute_rpcs(mut client RpcWsClient, mut logger log.Logger, mnemonic string) ! {
	// ADD YOUR CALLS HERE
}

fn execute_nostr_rpcs(mut client RpcWsClient, mut logger log.Logger) ! {
	key := nostr.generate_keypair(mut client)!
	println(key)

	nostr.load(mut client, key)!

	nostr.connect_to_relay(mut client, "ws://localhost:8081")!
	nostr.subscribe_to_relays(mut client)!

	nostr.publish_to_relays(mut client, [""], "hello world 1!")!
	nostr.publish_to_relays(mut client, [""], "hello world 2!")!

	time.sleep(5 * time.second)

	events := nostr.get_events(mut client)!
	println("events")
	println(events)

	// Close subscriptions
	subscription_ids := nostr.get_subscription_ids(mut client)!
	println(subscription_ids)
	for id in subscription_ids {
		nostr.close_subscription(mut client, id)!
	}
}


fn main() {
	mut fp := flag.new_flag_parser(os.args)
	fp.application('Welcome to the web3_proxy client. The web3_proxy client allows you to execute all remote procedure calls that the web3_proxy server can handle.')
	fp.limit_free_args(0, 0)!
	fp.description('')
	fp.skip_executable()
	mnemonic := fp.string('mnemonic', `m`, '', 'The mnemonic to be used to call any function')
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
	_ := spawn myclient.run() //QUESTION: why is that in thread?
	execute_nostr_rpcs(mut myclient, mut logger) or {
		logger.error("Failed executing calls: $err")
		exit(1)
	}
}
