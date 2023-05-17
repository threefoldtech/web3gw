module main

import os
import vweb
import time
// import threefoldtech.chat
import freeflowuniverse.crystallib.pathlib

pub fn (mut app App) playground() vweb.Result {
	println('herreuya: ${app.query["schema"]}')
	return $vweb.html()
}

['/playground2/:schema']
pub fn (mut app App) playground2(schema string) vweb.Result {
	playground_source := '/playgroundwindow?schema=$schema'
	return $vweb.html()
}