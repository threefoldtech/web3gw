module main

//
// import os
import vweb
import json
import encoding.base64
import net.urllib

pub fn (mut app App) chat() vweb.Result {
	return $vweb.html()
}

struct Prompt {
	prompt string
}

['/chat/message'; POST]
pub fn (mut app App) message() vweb.Result {
	action := urllib.query_unescape(app.req.data.trim_string_left('prompt=')) or { panic(err) }
	response := respond(action)
	return app.text(response)
}

fn respond(action string) string {
	return action
}
