module tfchain

import freeflowuniverse.crystallib.actionsparser { Action }

fn (mut t TFChainHandler) nodes(action Action) ! {
	match action.name {
		'get' {
			farm_id := action.params.get_u32_default('farm_id', 0)!
			node_id := action.params.get_u32_default('node_id', 0)!

			if farm_id != 0 {
				nodes := t.tfchain.get_nodes(farm_id)!
				t.logger.info('nodes: ${nodes}')
				
			} else if node_id != 0 {
				node := t.tfchain.get_node(node_id)!
				t.logger.info('node: ${node}')

			} else {
				return error('farm_id or node_id is required')
			}

		} else {
			return error('core action ${action.name} is invalid')
		}
	}
}
