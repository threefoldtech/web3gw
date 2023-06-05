module main

import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { Capacity, SolutionHandler, VM }
import log

fn run_vm_ops(mut s SolutionHandler, mut logger log.Logger) ! {
	network_name := 'hamadavm'

	defer {
		s.delete_vm(network_name) or { logger.error('failed to vm network: ${err}') }
	}

	deploy_res := s.create_vm(VM{
		network: network_name
		capacity: 'small'
		ssh_key: 'hamada ssh'
		gateway: true
		add_wireguard_access: true
	})!
	logger.info('${deploy_res}')

	add_res := s.create_vm(VM{
		network: network_name
		capacity: 'small'
		ssh_key: 'hamada ssh2'
		times: 2
		gateway: true
		add_wireguard_access: true
	})!
	logger.info('${add_res}')

	remove_res := s.remove_vm(network_name, add_res.vms[0].machine.name)!
	logger.info('${remove_res}')

	get_res := s.get_vm(network_name)!
	logger.info('${get_res}')
}

fn main() {
	mut logger := log.Log{
		level: .info
	}

	mut tfgrid_client, mut exp := tfgrid.cli(mut logger) or {
		logger.error('failed to initialize tfgrid client: ${err}')
		exit(1)
	}

	mut s := SolutionHandler{
		tfclient: &tfgrid_client
		explorer: &exp
	}

	run_vm_ops(mut s, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
