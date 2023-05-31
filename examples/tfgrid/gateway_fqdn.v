module main

import threefoldtech.threebot.tfgrid
import log

fn test_fqdn_gw_ops(mut client tfgrid.TFGridClient, mut logger log.Logger) ! {
	model_name := '3omarName'

	res := client.gateways_deploy_fqdn(tfgrid.GatewayFQDN{
		name: model_name
		node_id: 2
		backends: ['http://1.1.1.1:9000']
		fqdn: 'gw.test.io'
	})!
	logger.info('${res}')

	defer {
		client.gateways_delete_fqdn(model_name) or {
			logger.error('failed to delete gateway fqdn: ${err}')
		}
	}

	res_2 := client.gateways_get_fqdn(model_name)!
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

	test_fqdn_gw_ops(mut tfgrid_client, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
