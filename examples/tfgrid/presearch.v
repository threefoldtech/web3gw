module main

import threefoldtech.threebot.tfgrid {TFGridClient}
import log

fn run_presearch_ops(mut t TFGridClient, mut logger log.Logger) ! {
	model_name := 'hamadapresearch'
	deploy_res := t.deploy_presearch(tfgrid.Presearch{
		name: model_name
		farm_id: 1
		ssh_key: 'key'
		public_ipv4: false
	})!
	logger.info('${deploy_res}')

	defer {
		t.delete_presearch(model_name) or { logger.error('failed to delete presearch: ${err}') }
	}

	get_res := t.get_presearch(model_name)!
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

	run_presearch_ops(mut tfgrid_client, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
