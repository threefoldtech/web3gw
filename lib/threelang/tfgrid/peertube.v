module tfgrid

import freeflowuniverse.crystallib.actionsparser { Action }
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { Peertube }
import rand

fn (mut t TFGridHandler) peertube(action Action) ! {
	match action.name {
		'create' {
			name := action.params.get_default('name', rand.string(10).to_lower())!
			farm_id := action.params.get_int_default('farm_id', 0)!
			capacity_str := action.params.get_default('capacity', 'meduim')!
			capacity := solution.get_capacity(capacity_str)!
			ssh_key_name := action.params.get_default('sshkey', 'default')!
			ssh_key := t.get_ssh_key(ssh_key_name)!
			admin_email := action.params.get('admin_email')!
			db_username := action.params.get('db_username')!
			db_password := action.params.get('db_password')!
			smtp_hostname := action.params.get('smtp_hostname')!
			smtp_username := action.params.get('smtp_username')!
			smtp_password := action.params.get('smtp_password')!
			
			deploy_res := t.solution_handler.deploy_peertube(Peertube{
				name: name
				farm_id: u64(farm_id)
				capacity: capacity
				ssh_key: ssh_key
				admin_email: admin_email
				db_username: db_username
				db_password: db_password
				smtp_hostname: smtp_hostname
				smtp_username: smtp_username
				smtp_password: smtp_password
			})!

			t.logger.info('${deploy_res}')
		}
		'get' {
			name := action.params.get('name')!

			get_res := t.solution_handler.get_peertube(name)!

			t.logger.info('${get_res}')
		}
		
		'delete' {
			name := action.params.get('name')!

			t.solution_handler.delete_peertube(name) or {
				return error('failed to delete peertube instance: ${err}')
			}
		}
		else {
			return error('operation ${action.name} is not supported on peertube')
		}
	}
}