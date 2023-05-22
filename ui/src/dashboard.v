module main

import vweb
import time

pub fn (mut app App) dashboard() vweb.Result {
	server_ip := '127.0.0.1:8000'
	last_deployment := time.now().format()
	server_status := if app.server_is_running() {
		'Connected'
	} else {
		'Disconnected'
	}
	return $vweb.html()
}
