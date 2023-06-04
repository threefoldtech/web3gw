module threelang

import freeflowuniverse.crystallib.actionsparser
import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import log

import threefoldtech.threebot.tfgrid.solution { SolutionHandler }

pub struct Runner {
pub mut:
	path string
	handler SolutionHandler
	ssh_keys map[string]string
}

[params]
pub struct RunnerArgs {
pub mut:
	name      string
	path string
	address string
}

pub fn new(args RunnerArgs, debug_log bool) !Runner {
	mut factory := Runner{
			path:args.path
		}


	factory.run(args.address, debug_log)!
	return factory
}

pub fn (mut r Runner) run(address string, debug_log bool) ! {
	mut ap := actionsparser.new(path: r.path, defaultbook: 'aaa')!

	mut logger := log.Logger(&log.Log{
		level: if debug_log { .debug } else { .info }
	})

	mut myclient := rpcwebsocket.new_rpcwsclient(address, &logger) or {
		return error('Failed creating rpc websocket client: ${err}')
	}

	_ := spawn myclient.run()

	r.helper_actions(mut ap)!
	r.core_actions(mut ap, mut myclient)!
	r.vm_actions(mut ap)!
	r.gateway_name_actions(mut ap)!
	r.gateway_fqdn_actions(mut ap)!
	r.zdb_actions(mut ap)!
}

pub fn (mut r Runner) helper_actions(mut ap actionsparser.ActionsParser) ! {
	mut sshkey_action := ap.filtersort(actor: 'sshkeys', book: 'tfgrid')!

	for a in sshkey_action {
		name := a.params.get('name')!
		key := a.params.get('ssh_key')!
		r.ssh_keys[name] = key
	}
}