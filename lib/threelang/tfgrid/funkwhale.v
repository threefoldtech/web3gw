module tfgrid

import freeflowuniverse.crystallib.actionsparser { Action }
import threefoldtech.threebot.tfgrid { Funkwhale }
import rand

fn (mut t TFGridHandler) funkwhale(action Action) ! {
	match action.name {
		'create' {
			name := action.params.get_default('name', rand.string(10).to_lower())!
			farm_id := action.params.get_int_default('farm_id', 0)!
			capacity := action.params.get_default('capacity', 'meduim')!
			ssh_key_name := action.params.get_default('sshkey', 'default')!
			ssh_key := t.get_ssh_key(ssh_key_name)!
			admin_email := action.params.get('admin_email')!
			admin_username := action.params.get('admin_username')!
			admin_password := action.params.get('admin_password')!

			deploy_res := t.tfclient.deploy_funkwhale(Funkwhale{
				name: name
				farm_id: u64(farm_id)
				capacity: capacity
				ssh_key: ssh_key
				admin_email: admin_email
				admin_username: admin_username
				admin_password: admin_password
			})!

			t.logger.info('${deploy_res}')
		}
		'get' {
			name := action.params.get('name')!

			get_res := t.tfclient.get_funkwhale(name)!

			t.logger.info('${get_res}')
		}
		'delete' {
			name := action.params.get('name')!

			t.tfclient.delete_funkwhale(name) or {
				return error('failed to delete funkwhale instance: ${err}')
			}
		}
		else {
			return error('operation ${action.name} is not supported on funkwhale')
		}
	}
}
