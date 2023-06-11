module main

import threefoldtech.threebot.tfgrid {TFGridClient}
import log

fn run_funkwhale_ops(mut t TFGridClient, mut logger log.Logger) ! {
	model_name := 'hamadafunkwhale'
	deploy_res := t.deploy_funkwhale(tfgrid.Funkwhale{
		name: model_name
		capacity: 'small'
		admin_email: 'admin@gmail.com'
		admin_username: 'user1'
		admin_password: 'pass1'
	})!
	logger.info('${deploy_res}')

	defer {
		t.delete_funkwhale(model_name) or { logger.error('failed to delete funkwhale: ${err}') }
	}

	get_res := t.get_funkwhale(model_name)!
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


	run_funkwhale_ops(mut tfgrid_client, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
