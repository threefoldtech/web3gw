module threelang

import freeflowuniverse.crystallib.actionsparser
import freeflowuniverse.crystallib.rpcwebsocket
import log
import threefoldtech.threebot.threelang.tfgrid {TFGridHandler}
import threefoldtech.threebot.threelang.tfchain {TFChainHandler}

const(
	tfgrid_book = 'tfgrid'
	tfchain_book = 'tfchain'
)

pub struct Runner {
pub mut:
	path     string

	tfgrid_handler TFGridHandler
	tfchain_handler TFChainHandler
}

[params]
pub struct RunnerArgs {
pub mut:
	name    string
	path    string
	address string
}

pub fn new(args RunnerArgs, debug_log bool) !Runner {
	mut ap := actionsparser.new(path: args.path, defaultbook: 'aaa')!
	
	mut logger := log.Logger(&log.Log{
		level: if debug_log { .debug } else { .info }
	})

	mut myclient := rpcwebsocket.new_rpcwsclient(args.address, &logger) or {
		return error('Failed creating rpc websocket client: ${err}')
	}
	_ := spawn myclient.run()

	tfgrid_handler := tfgrid.new(mut myclient, logger)
	tfchain_handler := tfchain.new(mut myclient, logger)

	mut runner := Runner{
		path: args.path
		tfgrid_handler: tfgrid_handler
		tfchain_handler: tfchain_handler
	}

	runner.run(mut ap)!
	return runner
}

pub fn (mut r Runner) run(mut action_parser actionsparser.ActionsParser) ! {
	for action in action_parser.actions{
		match action.book{
			'tfgrid'{
				r.tfgrid_handler.handle_action(action)!
			}
			'tfchain'{
				r.tfchain_handler.handle_action(action)!
			}
			else{
				return error('module ${action.book} is invalid')
			}
		}
	}
}
