module threelang

import freeflowuniverse.crystallib.markdowndocs { Action, NewDocArgs, new }
import freeflowuniverse.crystallib.params { Params }
import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import threefoldtech.threebot.threelang.gridprocessor
import threefoldtech.threebot.threelang.tfchain { ChainProcessor }
import log

pub struct ThreeLangParser {
mut:
	grid_processor  Processor
	chain_processor Processor
	// other modules parsers
}

// Processor is an interface for all tfgrid modules to implement
// each module should be able to store an action through the add_action method
// then execute the action through the execute method
interface Processor {
mut:
	add_action(namespace string, operation string, params Params) !
	execute(mut rpc_client RpcWsClient) !
}

const (
	tfgrid_module          = 'tfgrid'
	tfchain_module         = 'tfchain'

	default_server_address = 'ws://127.0.0.1:8080'
)

// parse takes an md file path as input, preprocesses it, returns a ThreeLangParser instance
pub fn parse(args NewDocArgs) !ThreeLangParser {
	mut t := ThreeLangParser{
		grid_processor: gridprocessor.new()
		chain_processor: ChainProcessor{}
	}

	doc := new(args)!

	for item in doc.items {
		match item {
			Action {
				t.delegate(item)!
			}
			else {
				return error('invalid item. document should only contain actions')
			}
		}
	}

	return t
}

// delegate decides which parser should this action belong to, and delegates it to this parser
fn (mut t ThreeLangParser) delegate(action Action) ! {
	// validate action name
	mut action_name := parse_action_name(action)!

	mod := action_name[0]
	ns := action_name[1]
	op := if action_name.len == 3 { action_name[2] } else { '' }

	match mod {
		threelang.tfgrid_module {
			t.grid_processor.add_action(ns, op, action.params)!
		}
		threelang.tfchain_module {
			t.chain_processor.add_action(ns, op, action.params)!
		}
		else {
			return error('invalid module name ${mod}')
		}
	}
}

// execute performs all actions specified inside the md file
pub fn (mut t ThreeLangParser) execute() ! {
	// create a websocket connection, pass it to all processors

	mut logger := log.Logger(&log.Log{
		level: .debug
	})

	mut rpc_client := rpcwebsocket.new_rpcwsclient(threelang.default_server_address, &logger) or {
		return error('failed to create rpc websocket client: ${err}')
	}

	_ := spawn rpc_client.run()

	t.grid_processor.execute(mut rpc_client)!
	t.chain_processor.execute(mut rpc_client)!
}

// parse_action_name validates action name, splits it into 3 parts (module, namespace, operation)
fn parse_action_name(action Action) ![]string {
	mut action_array := action.name.split('.')
	if action_array.len != 2 && action_array.len != 3 {
		return error('invalid action name ${action.name}')
	}

	return action_array
}
