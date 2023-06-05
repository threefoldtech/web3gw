module main

import os
import threefoldtech.threebot.threelang { RunnerArgs }
import log
import flag

const (
	default_server_address = 'ws://127.0.0.1:8080'
)

fn main() {
	mut fp := flag.new_flag_parser(os.args)
	fp.application('Welcome to the web3_proxy client. The web3_proxy client allows you to execute all remote procedure calls that the web3_proxy server can handle.')
	fp.limit_free_args(0, 0)!
	fp.description('')
	fp.skip_executable()

	config_file_path := fp.string('file', `f`, '', 'Relative path to the Threelang config file.')
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

	_ := threelang.new(RunnerArgs{ path: config_file_path, address: address }, debug_log) or {
		eprintln(err)
		exit(1)
	}
}

