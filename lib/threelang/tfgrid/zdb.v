module tfgrid

import freeflowuniverse.crystallib.actionsparser { Action }
import threefoldtech.threebot.tfgrid { ZDB }
import rand

const GB = 1024 * 1024 * 1024

fn (mut t TFGridHandler) zdb(action Action) ! {
	match action.name {
		'create' {
			node_id := action.params.get_int_default('node_id', 0)!
			name := action.params.get_default('name', rand.string(10).to_lower())!
			password := action.params.get('password')!
			public := action.params.get_default_false('public')
			size := action.params.get_storagecapacity_in_bytes('size')!
			mode := action.params.get_default('mode', 'user')!

			zdb_deploy := t.tfgrid.zdb_deploy(ZDB{
				node_id: u32(node_id)
				name: name
				password: password
				public: public
				size: u32(size / GB)
				mode: mode
			})!

			t.logger.info('${zdb_deploy}')
		}
		'delete' {
			name := action.params.get('name')!
			t.tfgrid.zdb_delete(name)!
		}
		'get' {
			name := action.params.get('name')!
			zdb_get := t.tfgrid.zdb_get(name)!

			t.logger.info('${zdb_get}')
		}
		else {
			return error('action ${action.name} is not supported on zdbs')
		}
	}
}
