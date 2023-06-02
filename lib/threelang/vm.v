module threelang

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

fn (mut r Runner) vm_actions(actions actionsparser.ActionsParser) ! {
	mut actions2 := actions.filtersort(actor: 'vm', domain:'tfgrid')!
	for action in actions2 {
		p:=action.params
		if action.name == 'create' {
			//get the relevant args from params
			name := action.params.get('name')!
			// name := action.params.get_default('growth', '1:1')!

			//TODO: call the V client to deploy a VM
		}

	}