module jsonrpc_server

import time
import term
import log
import net.websocket {Server}
import freeflowuniverse.crystallib.jsonrpc

pub struct JSONRPCServer {
	Server
	logger log.Logger
	handle_request fn (jsonrpc.JsonRpcRequest)
}

struct Config {
	port int
	handle_request fn (jsonrpc.JsonRpcRequest)
}

pub fn new(config Config) JSONRPCServer {
	return JSONRPCServer {
		Server: websocket.new_server(.ip6, config.port, '')
		handle_request: config.handle_request
	}
}

pub fn (mut server JSONRPCServer[T]) run()! {
	defer {
		unsafe {
			server.free()
		}
	}
	// Make that in execution test time give time to execute at least one time
	server.ping_interval = 100
	server.on_connect(handle_connection)!
	server.on_message_ref(fn[T](mut ws websocket.Client, msg &websocket.Message, mut server &JSONRPCServer[T]) ! {
		server.handle_message[T](mut ws, msg) or {panic(err)}
	}, &server)
	server.on_close(handle_close)
	server.listen() or {handle_error}
	slog('s.listen finished')
}

// fn main() {
// 	spawn start_server()
// 	// time.sleep(100 * time.millisecond)
// 	// start_client()!
// 	for{}
// }

fn slog(message string) {
	eprintln(term.colorize(term.bright_yellow, message))
}

fn clog(message string) {
	eprintln(term.colorize(term.cyan, message))
}

fn wlog(message string) {
	eprintln(term.colorize(term.bright_blue, message))
}

// handle connection logs connection and returns whether the connection should be accepted or not
fn handle_connection(mut s websocket.ServerClient) !bool {
	$if debug {
		wlog('New connection by client key: ${s.client_key}')
	}
	// TODO: handle accpet/deny
	return true
}

fn (mut server JSONRPCServer[T]) handle_message(mut ws websocket.Client, msg &websocket.Message) ! {
	slog('s.on_message msg.opcode: ${msg.opcode} | msg.payload: ${msg.payload}')
	req := jsonrpc.new_jsonrpcrequest[string]('method', 'hello')
	server.handle_request[string](req)
	ws.write(msg.payload, msg.opcode) or {
		eprintln('ws.write err: ${err}')
		return err
	}
}

fn handle_close(mut ws websocket.Client, code int, reason string) ! {
	slog('s.on_close code: ${code}, reason: ${reason}')
	println('client ($ws.id) closed connection')
}

fn handle_error(err IError) ! {
	slog('s.listen err: ${err}')
	return err
}

// // start_server starts the websocket server, it receives messages
// // and send it back to the client that sent it
// fn start_server() ! {
// 	mut s := websocket.new_server(.ip6, 30000, '')
// 	defer {
// 		unsafe {
// 			s.free()
// 		}
// 	}
// 	// Make that in execution test time give time to execute at least one time
// 	s.ping_interval = 100
// 	s.on_connect(handle_connection)!
// 	s.on_message(handle_message)
// 	s.on_close(handle_close)
// 	s.listen() or {handle_error}
// 	slog('s.listen finished')
// }

// // // start_client starts the websocket client, it writes a message to
// // // the server and prints all the messages received
// // fn start_client() ! {
// // 	mut ws := websocket.new_client('ws://localhost:30000')!
// // 	defer {
// // 		unsafe {
// // 			ws.free()
// // 		}
// // 	}
// // 	// mut ws := websocket.new_client('wss://echo.websocket.org:443')?
// // 	// use on_open_ref if you want to send any reference object
// // 	ws.on_open(fn (mut ws websocket.Client) ! {
// // 		clog('ws.on_open')
// // 	})
// // 	// use on_error_ref if you want to send any reference object
// // 	ws.on_error(fn (mut ws websocket.Client, err string) ! {
// // 		clog('ws.on_error error: ${err}')
// // 	})
// // 	// use on_close_ref if you want to send any reference object
// // 	ws.on_close(fn (mut ws websocket.Client, code int, reason string) ! {
// // 		clog('ws.on_close')
// // 	})
// // 	// use on_message_ref if you want to send any reference object
// // 	ws.on_message(fn (mut ws websocket.Client, msg &websocket.Message) ! {
// // 		if msg.payload.len > 0 {
// // 			message := msg.payload.bytestr()
// // 			clog('ws.on_message client got type: ${msg.opcode} payload: `${message}`')
// // 		}
// // 	})
// // 	// you can add any pointer reference to use in callback
// // 	// t := TestRef{count: 10}
// // 	// ws.on_message_ref(fn (mut ws websocket.Client, msg &websocket.Message, r &SomeRef) ? {
// // 	// // eprintln('type: $msg.opcode payload:\n$msg.payload ref: $r')
// // 	// }, &r)
// // 	ws.connect() or {
// // 		clog('ws.connect err: ${err}')
// // 		return err
// // 	}
// // 	clog('ws.connect succeeded')
// // 	spawn write_echo(mut ws) // or { println('error on write_echo $err') }
// // 	ws.listen() or {
// // 		clog('ws.listen err: ${err}')
// // 		return err
// // 	}
// // 	clog('ws.listen finished')
// // }

// // fn write_echo(mut ws websocket.Client) ! {
// // 	wlog('write_echo, start')
// // 	message := 'echo this'
// // 	for i := 0; i <= 5; i++ {
// // 		// Server will send pings every 30 seconds
// // 		wlog('write_echo, writing message: `${message}` ...')
// // 		ws.write_string(message) or {
// // 			wlog('write_echo, ws.write_string err: ${err}')
// // 			return err
// // 		}
// // 		time.sleep(100 * time.millisecond)
// // 	}
// // 	ws.close(1000, 'normal') or {
// // 		wlog('write_echo, close err: ${err}')
// // 		return err
// // 	}
// // 	wlog('write_echo, done')
// // }
