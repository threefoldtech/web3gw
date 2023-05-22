module main

import os
import vweb
import time
// import threefoldtech.chat
import freeflowuniverse.crystallib.pathlib

pub fn (mut app App) playground() vweb.Result {
	playground_source := 'https://threefoldtech.github.io/web3_proxy/playground/?schemaUrl=../openrpc/openrpc.json'
	return app.html($tmpl('templates/playground2.html'))
}

['/playground/:schema']
pub fn (mut app App) playground_schema(schema string) vweb.Result {
	playground_source := 'https://threefoldtech.github.io/web3_proxy/playground/?schemaUrl=../openrpc/${schema}/openrpc.json'
	return app.html($tmpl('templates/playground2.html'))
}
