module threelang

import freeflowuniverse.crystallib.actionsparser
import freeflowuniverse.crystallib.rpcwebsocket
import log
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.explorer
import threefoldtech.threebot.tfgrid.solution { SolutionHandler }

pub struct Runner {
pub mut:
	path     string
	handler  SolutionHandler
	ssh_keys map[string]string
}

[params]
pub struct RunnerArgs {
pub mut:
	name    string
	path    string
	address string
}

pub fn new(args RunnerArgs, debug_log bool) !Runner {
	
	mut logger := log.Logger(&log.Log{
		level: if debug_log { .debug } else { .info }
	})

	mut myclient := rpcwebsocket.new_rpcwsclient(args.address, &logger) or {
		return error('Failed creating rpc websocket client: ${err}')
	}

	_ := spawn myclient.run()

	mut tfgrid_client := tfgrid.new(mut myclient)
	mut exp := explorer.new(mut myclient)

	mut factory := Runner{
		path: args.path
		handler: SolutionHandler{
			tfclient: tfgrid_client
			explorer: exp
		}
	}

	factory.run(args.address, debug_log)!
	return factory
}

pub fn (mut r Runner) run(address string, debug_log bool) ! {
	mut ap := actionsparser.new(path: r.path, defaultbook: 'aaa')!

	r.helper_actions(mut ap)!
	r.core_actions(mut ap)!
	r.vm_actions(mut ap)!
	r.gateway_name_actions(mut ap)!
	r.gateway_fqdn_actions(mut ap)!
	r.zdb_actions(mut ap)!
	r.k8s_actions(mut ap)!
}
