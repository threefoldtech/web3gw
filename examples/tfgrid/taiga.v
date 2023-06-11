module main

import threefoldtech.threebot.tfgrid {TFGridClient, Taiga}
import log

fn run_taiga_ops(mut t TFGridClient, mut logger log.Logger) ! {
	model_name := 'hamadataiga'
	deploy_res := t.deploy_taiga(Taiga{
		name: model_name
		farm_id: 1
		capacity: 'small'
		disk_size: 20
		ssh_key: 'key'
		admin_username: 'user1'
		admin_password: 'pass'
		admin_email: 'email@gmail.com'
	})!
	logger.info('${deploy_res}')

	defer {
		t.delete_taiga(model_name) or { logger.error('failed to delete taiga: ${err}') }
	}

	get_res := t.get_taiga(model_name)!
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

	run_taiga_ops(mut tfgrid_client, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
