module tfgrid

import freeflowuniverse.crystallib.params { Params }
import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import threefoldtech.threebot.tfgrid {TFGridClient, new}

// GridProcessor should handle processing all tfgrid related actions
pub struct GridProcessor {
mut:
	credentials Credentials
	projects    map[string]Process
}

struct Project {
	op GridOp
}

pub interface Process {
mut:
	// execute performs the specified operation
	execute(mut client &TFGridClient) !
}

pub enum GridNS {
	k8s
	gateway_name
	gateway_fqdn
	machines
	zdb
	discourse
	taiga
	funkwhale
	presearch
	peertube
	construct
	login
}

pub enum GridOp {
	create
	read
	update
	delete

	login
}

// add_action validates the provided namespace, operation, and action_params, then adds the extracted information to the processor
fn (mut g GridProcessor) add_action(ns string, op string, action_params Params) ! {
	grid_ns := get_grid_ns(ns)!
	grid_op := get_grid_op(op)!
	param_map := get_param_map(action_params)
	match grid_ns {
		.gateway_name {
			g.build_gateway_name(grid_op, param_map)!
		}
		.login {
			g.login(grid_op, param_map)!
		}
		// other namespaces
		else {}
	}
}

fn (mut g GridProcessor) execute(mut rpc_client &RpcWsClient) ! {
	mut tfgrid_client := new(mut rpc_client)

	tfgrid_client.load(tfgrid.Credentials{
		mnemonic: g.credentials.mnemonic
		network: g.credentials.network
	})!
	println('logged in')

	for _, mut process in g.projects{
		process.execute(mut tfgrid_client)!
	}
}

// get_grid_ns validates namespace, returns corresponding enum
fn get_grid_ns(ns string) !GridNS {
	match ns {
		'k8s' {
			return GridNS.k8s
		}
		'gateway_name' {
			return GridNS.gateway_name
		}
		'gateway_fqdn' {
			return GridNS.gateway_fqdn
		}
		'machines' {
			return GridNS.machines
		}
		'zdb' {
			return GridNS.zdb
		}
		'discourse' {
			return GridNS.discourse
		}
		'taiga' {
			return GridNS.taiga
		}
		'funkwhale' {
			return GridNS.funkwhale
		}
		'presearch' {
			return GridNS.presearch
		}
		'peertube' {
			return GridNS.peertube
		}
		'construct' {
			// special namespace for adding customizability
			return GridNS.construct
		}
		'login' {
			return GridNS.login
		}
		else {
			return error('invalid tfgrid namespace ${ns}')
		}
	}
}

// get_grid_op validates operation, returns corresponding enum
fn get_grid_op(op string) !GridOp {
	match op {
		'create' {
			return GridOp.create
		}
		'read' {
			return GridOp.read
		}
		'update' {
			return GridOp.update
		}
		'delete' {
			return GridOp.delete
		}
		''{
			return GridOp.login
		}
		else {
			return error('invalid tfgrid operation ${op}')
		}
	}
}

fn get_param_map(action_params Params) map[string]string {
	mut mp := map[string]string{}
	for p in action_params.params {
		mp[p.key] = p.value
	}

	return mp
}
