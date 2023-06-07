module tfgrid

import freeflowuniverse.crystallib.actionsparser { Action }
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { Presearch }
import rand

fn (mut t TFGridHandler) presearch(action Action) ! {
	match action.name {
		'create' {
			name := action.params.get_default('name', rand.string(10).to_lower())!
			farm_id := action.params.get_int_default('farm_id', 0)!
			ssh_key_name := action.params.get_default('sshkey', 'default')!
			ssh_key := t.get_ssh_key(ssh_key_name)!
			disk_size := action.params.get_storagecapacity_in_bytes('disk_size')! / u32(1024*1024*1024)
			public_ipv4 := action.params.get_default_false('public_ip')
			
			deploy_res := t.solution_handler.deploy_presearch(Presearch{
				name: name
				farm_id: u64(farm_id)
				ssh_key: ssh_key
				disk_size: u32(disk_size)
				public_ipv4: public_ipv4
			})!

			t.logger.info('${deploy_res}')
		}
		'get' {
			name := action.params.get('name')!

			get_res := t.solution_handler.get_presearch(name)!

			t.logger.info('${get_res}')
		}
		
		'delete' {
			name := action.params.get('name')!

			t.solution_handler.delete_presearch(name) or {
				return error('failed to delete presearch instance: ${err}')
			}
		}
		else {
			return error('operation ${action.name} is not supported on presearch')
		}
	}
}
