module tfgrid

import freeflowuniverse.crystallib.params
import threefoldtech.threebot.tfgrid {TFGridClient}
import strconv

struct GWNameCreateParams {
	name            string
	node_id         u32
	tls_passthrough bool
	backend         string
}

fn (gw GWNameCreateParams) execute(mut client TFGridClient) ! {
	client.gateways_deploy_name(tfgrid.GatewayName{
		name: gw.name
		backends: [gw.backend]
		node_id: gw.node_id
		tls_passthrough: gw.tls_passthrough
	})!
}

fn (mut g GridProcessor) build_gateway_name(op GridOp, param_map map[string]string) ! {
	match op {
		.create {
			g.gateway_name_create(param_map)!
		}
		.read {
			g.gateway_name_read(param_map)!
		}
		.delete {
			g.gateway_name_delete(param_map)!
		}
		else {
			return error('gateway names do not support updates')
		}
	}
}

fn (mut g GridProcessor) gateway_name_create(param_map map[string]string) ! {
	name := param_map['name'] or { return error('gateway name is missing') }
	node_id := strconv.parse_uint(param_map['node_id'], 10, 32)!
	tls_passthrough := match param_map['tls_passthrough'] {
		'true' {
			true
		}
		'false', '' {
			false
		}
		else {
			return error('invalid value for tls_passthrough')
		}
	}
	backend := param_map['backend'] or { return error('gateway backend is missing') }

	gw := GWNameCreateParams{
		name: name
		node_id: u32(node_id)
		tls_passthrough: tls_passthrough
		backend: backend
	}

	g.projects[gw.name] = gw
}

fn (mut g GridProcessor) gateway_name_read(param_map map[string]string) ! {
}

fn (mut g GridProcessor) gateway_name_delete(param_map map[string]string) ! {
}
