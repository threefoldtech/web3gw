module tfgrid

import freeflowuniverse.crystallib.actionsparser {Action}
import threefoldtech.threebot.tfgrid { GatewayName }
import threefoldtech.threebot.tfgrid.solution
import rand
import log

fn (mut t TFGridHandler) gateway_name(action Action) ! {
	mut logger := log.Logger(&log.Log{
		level: .info
	})

	match action.name {
		'create' {
			node_id := action.params.get_int_default('node_id', 0)!
			name := action.params.get_default('name', rand.string(10).to_lower())!
			tls_passthrough := action.params.get_default_false('tls_passthrough')
			backend := action.params.get('backend')!

			gw_deploy := t.solution_handler.tfclient.gateways_deploy_name(GatewayName{
				name: name
				node_id: u32(node_id)
				tls_passthrough: tls_passthrough
				backends: [backend]
			})!
			
			logger.info('${gw_deploy}')
		}
		'delete' {
			name := action.params.get('name')!
			t.solution_handler.tfclient.gateways_delete_name(name)!
		}
		'get' {
			name := action.params.get('name')!
			gw_get := t.solution_handler.tfclient.gateways_get_name(name)!
			
			logger.info('${gw_get}')
		}
		else {
			return error('action ${action.name} is not supported on gateways')
		}
	}
}
