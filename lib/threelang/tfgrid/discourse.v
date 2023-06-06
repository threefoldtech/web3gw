module tfgrid

import freeflowuniverse.crystallib.actionsparser { Action }
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { Discourse }
import rand

fn (mut t TFGridHandler) discourse(action Action) ! {
	match action.name {
		'create' {
			name := action.params.get_default('name', rand.string(10).to_lower())!
			farm_id := action.params.get_int_default('farm_id', 0)!
			capacity_str := action.params.get_default('capacity', 'meduim')!
			capacity := solution.get_capacity(capacity_str)!
			ssh_key_name := action.params.get_default('sshkey', 'default')!
			ssh_key := t.get_ssh_key(ssh_key_name)!
			developer_email := action.params.get('developer_email')!
			smtp_address := action.params.get('smtp_address')!
			smtp_port := action.params.get_int('smtp_port')!
			smtp_username := action.params.get('smtp_username')!
			smtp_password := action.params.get('smtp_password')!
			smtp_tls := action.params.get_default_false('smtp_tls')

			deploy_res := t.solution_handler.deploy_discourse(Discourse{
				name: name
				farm_id: u64(farm_id)
				capacity: capacity
				ssh_key: ssh_key
				developer_email: developer_email
				smtp_address: smtp_address
				smtp_port: u32(smtp_port)
				smtp_username: smtp_username
				smtp_password: smtp_password
				smtp_enable_tls: smtp_tls
			})!

			t.logger.info('${deploy_res}')
		}
		'get' {
			name := action.params.get('name')!

			get_res := t.solution_handler.get_discourse(name)!

			t.logger.info('${get_res}')
		}
		
		'delete' {
			name := action.params.get('name')!

			t.solution_handler.delete_discourse(name) or {
				return error('failed to delete discourse instance: ${err}')
			}
		}
		else {
			return error('operation ${action.name} is not supported on discourse')
		}
	}
}
