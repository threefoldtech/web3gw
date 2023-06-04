module main

import freeflowuniverse.crystallib.actionsparser
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { Capacity, SolutionHandler, VM }
import log

// fn test_vm_ops(mut s SolutionHandler, mut logger log.Logger) ! {
// 	network_name := 'hamadavm'

// 	defer {
// 		s.delete_vm(network_name) or { logger.error('failed to vm network: ${err}') }
// 	}

// 	deploy_res := s.create_vm(VM{
// 		network: network_name
// 		capacity: Capacity.small
// 		ssh_key: 'hamada ssh'
// 		gateway: true
// 		add_wireguard_access: true
// 	})!
// 	logger.info('${deploy_res}')

// 	add_res := s.create_vm(VM{
// 		network: network_name
// 		capacity: Capacity.small
// 		ssh_key: 'hamada ssh2'
// 		times: 2
// 		gateway: true
// 		add_wireguard_access: true
// 	})!
// 	logger.info('${add_res}')

// 	remove_res := s.remove_vm(network_name, add_res.vms[0].machine.name)!
// 	logger.info('${remove_res}')

// 	get_res := s.get_vm(network_name)!
// 	logger.info('${get_res}')
// }
// import freeflowuniverse.crystallib.texttools

fn (mut r Runner) vm_actions(mut actions actionsparser.ActionsParser) ! {
	mut actions2 := actions.filtersort(actor: 'machines', book:'tfgrid')!
	for action in actions2 {
		match action.name {
			'create' {
				name := action.params.get('name')!
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
					network: name
					capacity: capacity
					ssh_key: ssh_key
					gateway: false
					add_wireguard_access: true
				})!

				println('${deploy_res}')
			}
			'delete' {
				println('deleting')
			}
			'get' {
				println('getting')
			}
			'add' {
				println('adding')
			}
			'remove' {
				println('removing')
			}
			else {
				println("error")
			}
		}
	}
}