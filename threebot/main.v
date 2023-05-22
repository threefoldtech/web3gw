module main

import os

const project_dir = os.dir(os.dir(@FILE)) 

fn main() {
	do() or {panic(err)}
}

fn do() ! {
	dir := os.dir(@FILE)
	os.chdir('$dir/ui')!
	// os.execute('v run run.vsh')

	os.execute('go build $project_dir/server')
	os.execute('$project_dir/server/server')
	println('Running Web3Proxy JSON-RPC WS API Server on port:8081')

	mut server := new(
		port: 30000,
		handler: handler
	)

	spawn (&server).run()
	println('Running 3Bot JSON-RPC WS API Server on port:8000')

	os.execute('v run run.vsh')
	println('Running 3Bot User Interface on port:8081')
	for{}
}