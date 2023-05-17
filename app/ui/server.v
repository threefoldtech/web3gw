module main

import os
import vweb
import time
import chat
// import threefoldtech.chat
import freeflowuniverse.crystallib.pathlib

struct App {
	vweb.Context
	vweb.Controller
}

pub fn (mut app App) before_request() {
	referer := app.get_header('Referer')
	hx_request := app.get_header('Hx-Request') == 'true'
	if !(hx_request) && !app.req.url.contains('static') && !app.req.url.contains('/playground'){
		app.index()
	}
}

struct NavItem {
	label string
	url string
}

pub fn (mut app App) index() vweb.Result {

	route := if app.req.url == '/' {
		'/dashboard'
	} else {
		app.req.url
	}
	sidebar := [
		NavItem{
			label: 'Dashboard'
			url: '/dashboard'
		},
		NavItem{
			label: 'Domains'
			url: '/domains'
		},
		NavItem{
			label: 'Manual'
			url: 'https://threefoldtech.github.io/web3_proxy/index.html'
		},
		NavItem{
			label: 'Manual2'
			url: '/manual2'
		},
		NavItem{
			label: 'Playground'
			url: '/playground2'
		}
	]

	server_ip := '127.0.0.1:8000'
	last_deployment := time.now().format()
	server_status := if app.server_is_running() {
		'Connected'
	} else {
		'Disconnected'
	}
	return $vweb.html()
}

fn (mut app App) server_is_running() bool {
	return true
}

pub fn main() {

	mut app := &App{
        controllers: [
			vweb.controller('/chat', &chat.Chat{}),
        ]
	}
	// app.mount_static_folder_at('../build/static', '/static')
	// playground_
	// os.execute('ln -s ')
	// mount_playground()!
	app.mount_static_folder_at('static', '/static')
	vweb.run(app, 8081)
}

pub fn mount_playground() ! {
	os.chdir(os.dir(@FILE))!
	playground_path := pathlib.get('../playground').path
	static_path := pathlib.get('static').path
	os.execute('ln -s $playground_path/css/* $static_path/css/')
	os.execute('ln -s $static_path/js/* static/js/')
	os.execute('cp ../playground/index.html templates')
}