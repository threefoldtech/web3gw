module gridprocessor

import freeflowuniverse.crystallib.params { Params }
import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import threefoldtech.threebot.tfgrid { TFGridClient }

type Builder = fn (grid_op GridOp, param_map map[string]string, args_set map[string]bool) !(string, Process)

// GridProcessor should handle processing all tfgrid related actions
pub struct GridProcessor {
mut:
	credentials     ?Credentials
	projects        map[string][]Process
	process_builder map[int]Builder
}

// Process is an interface for all tfgrid operations
pub interface Process {
mut:
	execute(mut client TFGridClient) !
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

pub fn new() GridProcessor {
	mut g := GridProcessor{
		projects: map[string][]Process{}
		process_builder: map[int]Builder{}
		credentials: none
	}

	g.process_builder[int(GridNS.gateway_name)] = build_gateway_name_process
	// record other solutions

	return g
}

// add_action validates the provided namespace, operation, and action_params, then adds the extracted information to the processor
fn (mut g GridProcessor) add_action(ns string, op string, action_params Params) ! {
	grid_ns := get_grid_ns(ns)!
	grid_op := get_grid_op(op)!
	param_map := get_param_map(action_params)
	args_set := get_argument_set(action_params)

	match grid_ns {
		.login {
			g.credentials = get_credentials(grid_op, param_map, args_set)!
		}
		.construct {
			// for customizability
		}
		// other namespaces
		else {
			process_builder := g.process_builder[int(grid_ns)]

			project, process := process_builder(grid_op, param_map, args_set)!

			g.projects[project] << process
		}
	}
}

fn (mut g GridProcessor) execute(mut rpc_client RpcWsClient) ! {
	mut tfgrid_client := tfgrid.new(mut rpc_client)

	cred := g.credentials or { return error('Unauthorized. You must add a login action') }

	tfgrid_client.load(tfgrid.Credentials{
		mnemonic: cred.mnemonic
		network: cred.network
	})!

	for _, mut processes in g.projects {
		for mut p in processes {
			// TODO: all processes should try to run before returning an error
			p.execute(mut tfgrid_client)!
		}
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
		'' {
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
	for p in action_params.args {
		mp[p] = true
	}

	return mp
}
