module gridprocessor

import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { SolutionHandler }
import strconv
import rand
import encoding.utf8

struct GWNameCreateParams {
	name            string
	node_id         u32
	tls_passthrough bool
	backend         string
}

fn (gw_create GWNameCreateParams) execute(mut s SolutionHandler) !string {
	ret := s.tfclient.gateways_deploy_name(tfgrid.GatewayName{
		name: gw_create.name
		backends: [gw_create.backend]
		node_id: gw_create.node_id
		tls_passthrough: gw_create.tls_passthrough
	})!

	return ret.str()
}

struct GWNameGetParams {
	name string
}

fn (gw_get GWNameGetParams) execute(mut s SolutionHandler) !string {
	ret := s.tfclient.gateways_get_name(gw_get.name)!
	return ret.str()
}

struct GWNameDeleteParams {
	name string
}

fn (gw_delete GWNameDeleteParams) execute(mut s SolutionHandler) !string {
	s.tfclient.gateways_delete_name(gw_delete.name)!
	return 'gateway name ${gw_delete.name} deleted'
}

fn build_gateway_name_process(op GridOp, param_map map[string]string, args_set map[string]bool) !(string, Process) {
	match op {
		.create {
			return gateway_name_create(param_map, args_set)!
		}
		.get {
			return gateway_name_read(param_map, args_set)!
		}
		.delete {
			return gateway_name_delete(param_map, args_set)!
		}
		else {
			return error('gateway names do not support updates')
		}
	}
}

fn gateway_name_create(param_map map[string]string, args_set map[string]bool) !(string, Process) {
	name := param_map['name'] or { return error('gateway name is missing') }
	node_id_str := param_map['node_id'] or { '0' }
	node_id := strconv.parse_uint(node_id_str, 10, 32)!
	tls_passthrough := args_set['tls_passthrough']
	backend := param_map['backend'] or { return error('gateway backend is missing') }

	gw := GWNameCreateParams{
		name: name
		node_id: u32(node_id)
		tls_passthrough: tls_passthrough
		backend: backend
	}

	return name, gw
}

fn gateway_name_read(param_map map[string]string, args_set map[string]bool) !(string, Process) {
	name := param_map['name'] or { return error('gateway name is missing') }

	gw := GWNameGetParams{
		name: name
	}

	return name, gw
}

fn gateway_name_delete(param_map map[string]string, args_set map[string]bool) !(string, Process) {
	name := param_map['name'] or { return error('gateway name is missing') }

	gw := GWNameDeleteParams{
		name: name
	}

	return name, gw
}
