module main

import os
import threefoldtech.web3gw.threelang { RunnerArgs }
import flag

const (
	default_server_address = 'ws://127.0.0.1:8080'
)
 
fn main() {
	mut fp := flag.new_flag_parser(os.args)
	fp.application('Welcome to The Threelang Parser cli. The Threelang parser is a utility software that reads and executes Threelang formatted actions.')
	fp.limit_free_args(0, 0)!
	fp.description('')
	fp.skip_executable()

	config_file_path := fp.string('file', `f`, '', 'The path to the markdown file containing threelang.')
	address := fp.string('address', `a`, '${default_server_address}', 'The address of the web3gw server to connect to.')
	debug_log := fp.bool('debug', 0, false, 'By setting this flag the client will print debug logs too.')

	_ := fp.finalize() or {
		eprintln(err)
		println(fp.usage())
		exit(1)
	}

	_ := threelang.new(RunnerArgs{ path: config_file_path, address: address }, debug_log) or {
		eprintln(err)
		exit(1)
	}
}

