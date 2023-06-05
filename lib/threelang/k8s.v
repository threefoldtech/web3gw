module threelang

import freeflowuniverse.crystallib.actionsparser
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { K8s }
import log

fn (mut r Runner) k8s_actions(mut actions actionsparser.ActionsParser) ! {
	mut actions2 := actions.filtersort(actor: 'kubernetes', book: 'tfgrid')!
	for action in actions2 {
		match action.name {
			'create' {
				name := action.params.get('name')!
				farm_id := action.params.get_int_default('farm_id', 0)!
				capacity := action.params.get_default('capacity', 'small')!
				replica := action.params.get_int_default('replica', 1)!
				wg := action.params.get_default('add_wireguard_access', 'no')!
				public_ip := action.params.get_default('add_public_ips', 'no')!

				ssh_key_name := action.params.get_default('sshkey', 'default')!
				ssh_key := r.ssh_keys[ssh_key_name]

				deploy_res := r.handler.create_k8s(K8s{
					name: name
					farm_id: farm_id
					capacity: capacity
					replica: replica
					wg: if wg == 'yes' { true } else { false }
					public_ip: if public_ip == 'yes' { true } else { false }
					ssh_key: ssh_key
				})!

				println('${deploy_res}')
			}
			'get' {
				name := action.params.get('name')!

				get_res := r.handler.get_k8s(name)!

				println('${get_res}')
			}
			'add' {
				name := action.params.get('name')!
				farm_id := action.params.get_int_default('farm_id', 0)!
				capacity := action.params.get_default('capacity', 'small')!
				wg := action.params.get_default('add_wireguard_access', 'no')!

				ssh_key_name := action.params.get_default('sshkey', 'default')!
				ssh_key := r.ssh_keys[ssh_key_name]

				add_res := r.handler.add_k8s_worker(K8s{
					name: name
					farm_id: farm_id
					capacity: capacity
					wg: if wg == 'yes' { true } else { false }
					ssh_key: ssh_key
				})!

				println('${add_res}')
			}
			'remove' {
				name := action.params.get('name')!
				worker_name := action.params.get('worker_name')!

				remove_res := r.handler.remove_k8s_worker(name, worker_name)!
				println('${remove_res}')
			}
			'delete' {
				name := action.params.get('name')!

				r.handler.delete_k8s(name) or { println('failed to delete k8s cluster: ${err}') }
			}
			else {
				return error('operation ${action.name} is not supported on k8s')
			}
		}
	}
}
