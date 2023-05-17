module chat


import os
import vweb
import time
// import threefoldtech.chat
import freeflowuniverse.crystallib.pathlib

interface IChatBot {
	respond(string) string
}

pub struct Chat {
    vweb.Context
mut:
	bot IChatBot
}

pub fn (mut chat Chat) hi() vweb.Result {
	return chat.html('hello')
}

pub fn (mut chat Chat) yo() vweb.Result {
	return $vweb.html()
}

[POST]
pub fn (mut chat Chat) message() vweb.Result {
	message := 'hey'
	resp := chat.bot.respond(message)
	return $vweb.html()
}