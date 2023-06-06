module tfgrid

import freeflowuniverse.crystallib.actionsparser {Action}
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { VM }
import log

fn (mut t TFGridHandler) vm(action Action) ! {
	mut logger := log.Logger(&log.Log{
		level: .info
	})

	match action.name {
		'create' {
			network := action.params.get('network')!
			farm_id := action.params.get_int_default('farm_id', 0)!
			capacity := action.params.get_default('capacity', 'meduim')!
			times := action.params.get_int_default('times', 1)!
			disk_size := action.params.get_int_default('disk_size', 10)!
			gateway := action.params.get_default_false('gateway')
			wg := action.params.get_default_false('add_wireguard_access')
			public_ip := action.params.get_default_false('add_public_ips')

			ssh_key_name := action.params.get_default('sshkey', 'default')!
			ssh_key := t.ssh_keys[ssh_key_name] or{
				return error('ssh key ${ssh_key_name} does not exist')
			}

			deploy_res := t.solution_handler.create_vm(VM{
				network: network
				capacity: capacity
				ssh_key: ssh_key
				gateway: gateway
				add_wireguard_access: wg
				add_public_ips: public_ip
			})!

			logger.info('${deploy_res}')
		}
		'get' {
			network := action.params.get('network')!

			get_res := t.solution_handler.get_vm(network)!

			logger.info('${get_res}')
		}
		'remove' {
			network := action.params.get('network')!
			machine := action.params.get('machine')!

			remove_res := t.solution_handler.remove_vm(network, machine)!
			logger.info('${remove_res}')
		}
		'delete' {
			network := action.params.get('network')!

			t.solution_handler.delete_vm(network) or { return error('failed to delete vm network: ${err}') }
		}
		else {
			return error('operation ${action.name} is not supported on vms')
		}
	}
}
