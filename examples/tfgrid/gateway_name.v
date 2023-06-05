module main

import threefoldtech.threebot.tfgrid
import log

fn run_name_gw_ops(mut client tfgrid.TFGridClient, mut logger log.Logger) ! {
	gw_name := 'qowienfoiqw'

	res := client.gateways_deploy_name(tfgrid.GatewayName{
		name: gw_name
		backends: ['http://1.1.1.1:9000']
		node_id: 2
	})!
	logger.info('${res}')

	defer {
		client.gateways_delete_name(gw_name) or {
			logger.error('failed to delete gateway name: ${err}')
		}
	}

	res_2 := client.gateways_get_name(gw_name)!
	logger.info('${res_2}')
}

fn main() {
	mut logger := log.Log{
		level: .info
	}

	mut tfgrid_client, _ := tfgrid.cli(mut logger) or {
		logger.error('failed to initialize tfgrid client: ${err}')
		exit(1)
	}

	run_name_gw_ops(mut tfgrid_client, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
