module main

import os
import vweb
import time
// import threefoldtech.chat
import freeflowuniverse.crystallib.pathlib

pub fn (mut app App) manual() vweb.Result {
	iframe_source := "https://threefoldtech.github.io/web3_proxy/index.html"
	return $vweb.html()
}

[GET; '/manual/:path']
pub fn (mut app App) manual_path(path string) vweb.Result {
	println(path)
	manual_url := "https://threefoldtech.github.io/web3_proxy"
	iframe_source := if path == '' {
		'$manual_url/index.html' 
	} else {
		'$manual_url/$path'
	}
	return app.html($tmpl('templates/manual2.html'))
}
