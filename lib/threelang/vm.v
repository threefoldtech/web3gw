module threelang

import freeflowuniverse.crystallib.actionsparser
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { VM }
import log

fn (mut r Runner) vm_actions(mut actions actionsparser.ActionsParser) ! {
	mut actions2 := actions.filtersort(actor: 'machines', book: 'tfgrid')!
	for action in actions2 {
		match action.name {
			'create' {
				network := action.params.get('network')!
				farm_id := action.params.get_int_default('farm_id', 0)!
				capacity := action.params.get_default('capacity', 'meduim')!
				times := action.params.get_int_default('times', 1)!
				disk_size := action.params.get_int_default('disk_size', 10)!
				gateway := action.params.get_default('gateway', 'no')!
				wg := action.params.get_default('add_wireguard_access', 'no')!
				public_ip := action.params.get_default('add_public_ips', 'no')!

				ssh_key_name := action.params.get_default('sshkey', 'default')!
				ssh_key := r.ssh_keys[ssh_key_name]

				deploy_res := r.handler.create_vm(VM{
					network: network
					capacity: capacity
					ssh_key: ssh_key
					gateway: if gateway == 'yes' { true } else { false }
					add_wireguard_access: if wg == 'yes' { true } else { false }
					add_public_ips: if public_ip == 'yes' { true } else { false }
				})!

				println('${deploy_res}')
			}
			'get' {
				network := action.params.get('network')!

				get_res := r.handler.get_vm(network)!

				println('${get_res}')
			}
			'remove' {
				network := action.params.get('network')!
				machine := action.params.get('machine')!

				remove_res := r.handler.remove_vm(network, machine)!
				println('${remove_res}')
			}
			'delete' {
				network := action.params.get('network')!

				r.handler.delete_vm(network) or { println('failed to delete vm network: ${err}') }
			}
			else {
				return error('operation ${action.name} is not supported on vms')
			}
		}
	}
}
