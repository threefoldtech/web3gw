module threelang

import freeflowuniverse.crystallib.actionsparser
import threefoldtech.threebot.tfgrid {ZDB}
import threefoldtech.threebot.tfgrid.solution { Capacity, SolutionHandler, VM }
import rand
import encoding.utf8

const GB = 1024 * 1024 * 1024

fn (mut r Runner) zdb_actions(mut actions actionsparser.ActionsParser) ! {
	mut zdb_actions := actions.filtersort(actor: 'zdbs', book:'tfgrid')!
	for action in zdb_actions {
		match action.name {
			'create' {
				node_id := action.params.get_int_default('node_id', 0)!
				name := action.params.get_default('name', utf8.to_lower(rand.string(10)))!
				password := action.params.get('password')!
				public := action.params.get_default('public', 'no')!
				size := action.params.get_storagecapacity_in_bytes('size')!
				mode := action.params.get_default('mode', 'user')!

				zdb_deploy := r.handler.tfclient.zdb_deploy(ZDB{
					node_id: u32(node_id)
					name: name
					password: password
					public: if public == 'yes' {true} else {false}
					size: u32(size / GB)
					mode: mode
				})!

			}
			'delete' {
				name := action.params.get('name')!
				r.handler.tfclient.zdb_delete(name)!
			}
			'get' {
				name := action.params.get('name')!
				zdb_get := r.handler.tfclient.zdb_get(name)!
			}
			else {
				return error('action ${action.name} is not supported on zdbs')
			}
		}
	}
}