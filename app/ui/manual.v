module main

import os
import vweb
import time

struct Manual {
	vweb.Context
mut:
	path string
}

pub fn (mut manual Manual) before_request() {
	url := manual.req.url.all_after_first('/')
	path := url.trim_string_left('manual')
	println(path)
	manual.path = path
	manual.index()
}

pub fn (mut manual Manual) index() vweb.Result {
	manual_url := 'https://threefoldtech.github.io/web3_proxy'
	iframe_source := if manual.path == '' {
		'${manual_url}/index.html'
	} else {
		'${manual_url}/${manual.path}'
	}
	return manual.html($tmpl('templates/manual.html'))
}
