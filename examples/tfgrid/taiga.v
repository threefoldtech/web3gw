module main

import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.explorer
import log

fn test_taiga_ops(mut client tfgrid.TFGridClient, mut exp explorer.ExplorerClient, mut logger log.Logger) ! {
	model_name := 'hamadataiga'
	deploy_res := client.deploy_taiga(mut exp, tfgrid.Taiga{
		name: model_name
		farm_id: 1
		cpu: 2
		memory: 4096
		rootfs_size: 10240
		disk_size: 20
		ssh_key: 'key'
		admin_username: 'user1'
		admin_password: 'pass'
		admin_email: 'email@gmail.com'
	})!
	logger.info('${deploy_res}')

	defer {
		client.delete_taiga(model_name) or { logger.error('failed to delete taiga: ${err}') }
	}

	get_res := client.get_taiga(model_name)!
	logger.info('${get_res}')
}

fn main() {
	mut logger := log.Log{
		level: .info
	}

	mut tfgrid_client, mut exp := tfgrid.cli(mut logger) or {
		logger.error('failed to initialize tfgrid client: ${err}')
		exit(1)
	}

	test_taiga_ops(mut tfgrid_client, mut exp, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
