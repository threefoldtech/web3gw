module main

import threefoldtech.threebot.tfgrid
import log

fn test_machines_with_gateways_ops(mut client tfgrid.TFGridClient, mut logger log.Logger) ! {
	model_name := 'hamada_machines_with_gateways'
	deploy_res := client.deploy_machines_with_gateways(tfgrid.MachinesWithGateways{
		name: model_name
		add_wireguard_access: false
		machines: [
			tfgrid.MachineWithGateway{
				gateway: true
				machine: tfgrid.Machine{
					name: 'vm1'
					memory: 4096
					rootfs_size: 4096
				}
			},
		]
	})!
	logger.info('${deploy_res}')

	defer {
		client.delete_machines_with_gateways(model_name) or {
			logger.error('failed to delete machines_with_gateways: ${err}')
		}
	}

	get_res := client.get_machines_with_gateways(model_name)!
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

	test_machines_with_gateways_ops(mut tfgrid_client, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
