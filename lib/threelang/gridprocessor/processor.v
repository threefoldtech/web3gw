module gridprocessor

import freeflowuniverse.crystallib.params { Params }
import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import threefoldtech.threebot.tfgrid {TFGridClient}

type Builder = fn(grid_op GridOp, param_map map[string]string, args_set map[string]bool)!

// GridProcessor should handle processing all tfgrid related actions
[heap]
pub struct GridProcessor {
mut:
	credentials Credentials
	projects    map[string]Process
	namespaces map[int]Builder
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

pub fn new() GridProcessor{
	mut g := GridProcessor{
		projects: map[string]Process{}
		namespaces: map[int]Builder{}
	}
	
	g.namespaces[int(GridNS.gateway_name)] = g.build_gateway_name_process
	g.namespaces[int(GridNS.login)] = g.login
	// record other namespaces

	return g
}

// add_action validates the provided namespace, operation, and action_params, then adds the extracted information to the processor
fn (mut g GridProcessor) add_action(ns string, op string, action_params Params) ! {
	grid_ns := get_grid_ns(ns)!
	grid_op := get_grid_op(op)!
	param_map := get_param_map(action_params)
	args_set := get_argument_set(action_params)

	// builder := g.namespaces[int(grid_ns)]

	// builder(grid_op, param_map, args_set)!
	match grid_ns {
		.gateway_name {
			g.build_gateway_name_process(grid_op, param_map, args_set)!
		}
		.login {
			g.login(grid_op, param_map, args_set)!
		}
		// other namespaces
		else {}
	}
}

fn (mut g GridProcessor) execute(mut rpc_client &RpcWsClient) ! {
	println('cred: ${g.credentials}')
	mut tf_cl := tfgrid.new(mut rpc_client)

	tf_cl.load(tfgrid.Credentials{
		mnemonic: g.credentials.mnemonic
		network: g.credentials.network
	})!

	for _, mut process in g.projects{
		process.execute(mut tf_cl)!
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

fn get_argument_set(action_params Params) map[string]bool {
	mut mp := map[string]bool{}
	for p in action_params.args{
		mp[p] = true
	}

	return mp
}