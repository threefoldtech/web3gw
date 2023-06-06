module tfgrid

import freeflowuniverse.crystallib.actionsparser { Action }
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { K8s }

fn (mut t TFGridHandler) k8s(action Action) ! {

	match action.name {
		'create' {
			name := action.params.get('name')!
			farm_id := action.params.get_int_default('farm_id', 0)!
			capacity := action.params.get_default('capacity', 'small')!
			replica := action.params.get_int_default('replica', 1)!
			wg := action.params.get_default_false('add_wireguard_access')
			public_ip := action.params.get_default_false('add_public_ips')

			ssh_key_name := action.params.get_default('sshkey', 'default')!
			ssh_key := t.get_ssh_key(ssh_key_name)!

			deploy_res := t.solution_handler.create_k8s(K8s{
				name: name
				farm_id: farm_id
				capacity: capacity
				replica: replica
				wg: wg
				public_ip: public_ip
				ssh_key: ssh_key
			})!

			t.logger.info('${deploy_res}')
		}
		'get' {
			name := action.params.get('name')!

			get_res := t.solution_handler.get_k8s(name)!

			t.logger.info('${get_res}')
		}
		'add' {
			name := action.params.get('name')!
			farm_id := action.params.get_int_default('farm_id', 0)!
			capacity := action.params.get_default('capacity', 'small')!
			wg := action.params.get_default_false('add_wireguard_access')
			public_ip := action.params.get_default_false('add_public_ips')

			ssh_key_name := action.params.get_default('sshkey', 'default')!
			ssh_key := t.get_ssh_key(ssh_key_name)!

			add_res := t.solution_handler.add_k8s_worker(K8s{
				name: name
				farm_id: farm_id
				capacity: capacity
				wg: wg
				public_ip: public_ip
				ssh_key: ssh_key
			})!

			t.logger.info('${add_res}')
		}
		'remove' {
			name := action.params.get('name')!
			worker_name := action.params.get('worker_name')!

			remove_res := t.solution_handler.remove_k8s_worker(name, worker_name)!
			t.logger.info('${remove_res}')
		}
		'delete' {
			name := action.params.get('name')!

			t.solution_handler.delete_k8s(name) or {
				return error('failed to delete k8s cluster: ${err}')
			}
		}
		else {
			return error('operation ${action.name} is not supported on k8s')
		}
	}
}
