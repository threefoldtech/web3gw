module tfgrid

import freeflowuniverse.crystallib.baobab.actions { Action }
import rand

fn (mut t TFGridHandler) presearch(action Action) ! {
	mut presearch_client := t.tfgrid.applications().presearch()
	match action.name {
		'create' {
			name := action.params.get_default('name', rand.string(10).to_lower())!
			farm_id := action.params.get_int_default('farm_id', 0)!
			ssh_key_name := action.params.get_default('sshkey', 'default')!
			ssh_key := t.get_ssh_key(ssh_key_name)!
			disk_size := action.params.get_storagecapacity_in_gigabytes('disk_size') or { 0 }
			public_ipv4 := action.params.get_default_false('public_ip')
			registration_code := action.params.get('registration_code')!
			public_restore_key := action.params.get_default('public_restore_key', '')!
			private_restore_key := action.params.get_default('private_restore_key', '')!

			deploy_res := presearch_client.deploy(
				name: name
				farm_id: u64(farm_id)
				ssh_key: ssh_key
				disk_size: u32(disk_size)
				public_ipv4: public_ipv4
				registration_code: registration_code
				public_restore_key: public_restore_key
				private_restore_key: private_restore_key
			)!

			t.logger.info('${deploy_res}')
		}
		'get' {
			name := action.params.get('name')!

			get_res := presearch_client.get(name)!

			t.logger.info('${get_res}')
		}
		'delete' {
			name := action.params.get('name')!

			presearch_client.delete(name) or {
				return error('failed to delete presearch instance: ${err}')
			}
		}
		else {
			return error('operation ${action.name} is not supported on presearch')
		}
	}
}
