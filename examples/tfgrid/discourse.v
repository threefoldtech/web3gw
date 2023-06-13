module main

import threefoldtech.threebot.tfgrid {TFGridClient}
import log

fn run_discourse_ops(mut t TFGridClient, mut logger log.Logger) ! {
	model_name := 'hamadadiscourse'
	deploy_res := t.deploy_discourse(tfgrid.Discourse{
		name: model_name
		capacity: 'medium'
		disk_size: 10
		ssh_key: 'hamada ssh key'
		developer_email: 'em@mail.com'
		smtp_username: 'user1'
		smtp_password: 'pass1'
	})!
	logger.info('${deploy_res}')

	defer {
		t.delete_discourse(model_name) or { logger.error('failed to delete discourse: ${err}') }
	}

	get_res := t.get_discourse(model_name)!
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

	run_discourse_ops(mut tfgrid_client, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
