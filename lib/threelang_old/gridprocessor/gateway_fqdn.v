module gridprocessor

import threefoldtech.threebot.tfgrid { GatewayFQDN }
import threefoldtech.threebot.tfgrid.solution { SolutionHandler }
import strconv
import rand
import encoding.utf8

struct GWFQDNCreateParams {
	name            string
	node_id         u32
	tls_passthrough bool
	backend         string
	fqdn            string
}

fn (gw_create GWFQDNCreateParams) execute(mut s SolutionHandler) !string {
	ret := s.tfclient.gateways_deploy_fqdn(GatewayFQDN{
		name: gw_create.name
		backends: [gw_create.backend]
		node_id: gw_create.node_id
		fqdn: gw_create.fqdn
		tls_passthrough: gw_create.tls_passthrough
	})!

	return ret.str()
}

struct GWFQDNGetParams {
	name string
}

fn (gw_get GWFQDNGetParams) execute(mut s SolutionHandler) !string {
	ret := s.tfclient.gateways_get_fqdn(gw_get.name)!
	return ret.str()
}

struct GWFQDNDeleteParams {
	name string
}

fn (gw_delete GWFQDNDeleteParams) execute(mut s SolutionHandler) !string {
	s.tfclient.gateways_delete_fqdn(gw_delete.name)!
	return 'gateway fqdn ${gw_delete.name} is deleted'
}

fn build_gateway_fqdn_process(op GridOp, param_map map[string]string, args_set map[string]bool) !(string, Process) {
	match op {
		.create {
			return gateway_fqdn_create(param_map, args_set)!
		}
		.get {
			return gateway_fqdn_read(param_map, args_set)!
		}
		.delete {
			return gateway_fqdn_delete(param_map, args_set)!
		}
		else {
			return error('gateway FQDN does not support updates')
		}
	}
}

fn gateway_fqdn_create(param_map map[string]string, args_set map[string]bool) !(string, Process) {
	name := param_map['name'] or { utf8.to_lower(rand.string(10)) }
	node_id_str := param_map['node_id'] or { return error('gateway ${name} node id is missing') }
	node_id := strconv.parse_uint(node_id_str, 10, 32)!
	tls_passthrough := args_set['tls_passthrough']
	backend := param_map['backend'] or { return error('gateway backend is missing') }
	fqdn := param_map['fqdn'] or { return error('gateway fqdn is missing') }

	gw := GWFQDNCreateParams{
		name: name
		node_id: u32(node_id)
		tls_passthrough: tls_passthrough
		backend: backend
		fqdn: fqdn
	}

	return name, gw
}

fn gateway_fqdn_read(param_map map[string]string, args_set map[string]bool) !(string, Process) {
	name := param_map['name'] or { return error('gateway name is missing') }

	gw := GWFQDNGetParams{
		name: name
	}

	return name, gw
}

fn gateway_fqdn_delete(param_map map[string]string, args_set map[string]bool) !(string, Process) {
	name := param_map['name'] or { return error('gateway name is missing') }

	gw := GWFQDNDeleteParams{
		name: name
	}

	return name, gw
}
