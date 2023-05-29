module main

import threefoldtech.threebot.tfgrid
import log

fn test_zdb_ops(mut client tfgrid.TFGridClient, mut logger log.Logger) ! {
	model_name := 'testZdbOps'

	res := client.zdb_deploy(tfgrid.ZDB{
		name: model_name
		node_id: 83
		password: 'strongPass'
		size: 10
	})!
	logger.info('${res}')

	defer {
		client.zdb_delete(model_name) or { logger.error('failed to delete zdb: ${err}') }
	}

	res_2 := client.zdb_get(model_name)!
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

	test_zdb_ops(mut tfgrid_client, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
