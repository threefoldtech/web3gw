module main

import threefoldtech.threebot.tfgrid {TFGridClient, VM, RemoveVMWithGWArgs}
import log

fn run_vm_ops(mut t TFGridClient, mut logger log.Logger) ! {
	network_name := 'hamadavm'

	defer {
		t.delete_vm(network_name) or { logger.error('failed to vm network: ${err}') }
	}

	deploy_res := t.deploy_vm(VM{
		name: 'myfirstvm'
		network: network_name
		capacity: 'small'
		ssh_key: 'hamada ssh'
		gateway: true
		add_wireguard_access: true
	})!
	logger.info('${deploy_res}')

	add_res := t.deploy_vm(VM{
		name: 'mysecondvm'
		network: network_name
		capacity: 'small'
		ssh_key: 'hamada ssh2'
		times: 2
		gateway: true
		add_wireguard_access: true
	})!
	logger.info('${add_res}')

	remove_res := t.remove_vm(RemoveVMWithGWArgs{
		network: network_name
		vm_name: 'myfirstvm'
	})!
	logger.info('${remove_res}')

	get_res := t.get_vm(network_name)!
	logger.info('${get_res}')
}

fn main() {
	mut logger := log.Log{
		level: .info
	}

	mut tfgrid_client, _ := tfgrid.cli(mut logger) or {
		logger.error('failed to initialize tfgrid client: ${err}')
		exit(1)
	}

	run_vm_ops(mut tfgrid_client, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
