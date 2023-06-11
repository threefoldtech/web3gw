module main

import threefoldtech.threebot.tfgrid {TFGridClient}
import log

fn run_peertube_ops(mut t TFGridClient, mut logger log.Logger) ! {
	model_name := 'hamadapeertube'
	deploy_res := t.deploy_peertube(tfgrid.Peertube{
		name: model_name
		capacity: 'small'
		ssh_key: 'key'
		db_username: 'dbuser'
		db_password: 'dbpass'
		admin_email: 'email@gmail.com'
		smtp_hostname: 'smtphost'
		smtp_username: 'user'
		smtp_password: 'pass'
	})!
	logger.info('${deploy_res}')

	defer {
		t.delete_peertube(model_name) or { logger.error('failed to delete peertube: ${err}') }
	}

	get_res := t.get_peertube(model_name)!
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


	run_peertube_ops(mut tfgrid_client, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
