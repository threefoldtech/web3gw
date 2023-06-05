module threelang

import freeflowuniverse.crystallib.actionsparser
import threefoldtech.threebot.tfgrid { GatewayName }
import threefoldtech.threebot.tfgrid.solution
import rand
import encoding.utf8

fn (mut r Runner) gateway_name_actions(mut actions actionsparser.ActionsParser) ! {
	mut gateway_actions := actions.filtersort(actor: 'gateway_name', book: 'tfgrid')!
	for action in gateway_actions {
		match action.name {
			'create' {
				node_id := action.params.get_int_default('node_id', 0)!
				name := action.params.get_default('name', utf8.to_lower(rand.string(10)))!
				tls_passthrough := action.params.get_default('tls_passthrough', 'false')!
				backend := action.params.get('backend')!

				gw_deploy := r.handler.tfclient.gateways_deploy_name(GatewayName{
					name: name
					node_id: u32(node_id)
					tls_passthrough: if tls_passthrough == 'yes' { true } else { false }
					backends: [backend]
				})!
			}
			'delete' {
				name := action.params.get('name')!
				r.handler.tfclient.gateways_delete_name(name)!
			}
			'get' {
				name := action.params.get('name')!
				gw_get := r.handler.tfclient.gateways_get_name(name)!
			}
			else {
				return error('action ${action.name} is not supported on gateways')
			}
		}
	}
}
