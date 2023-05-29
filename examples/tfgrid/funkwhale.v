module main

import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.explorer
import log

fn test_funkwhale_ops(mut client tfgrid.TFGridClient, mut exp explorer.ExplorerClient, mut logger log.Logger) ! {
	model_name := 'hamadafunkwhale'
	deploy_res := client.deploy_funkwhale(mut exp, tfgrid.Funkwhale{
		name: model_name
		farm_id: 1
		cpu: 2
		memory: 4096
		rootfs_size: 10240
		admin_email: 'admin@gmail.com'
		admin_username: 'user1'
		admin_password: 'pass1'
	})!
	logger.info('${deploy_res}')

	defer {
		client.delete_funkwhale(model_name) or {
			logger.error('failed to delete funkwhale: ${err}')
		}
	}

	get_res := client.get_discourse(model_name)!
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

	test_funkwhale_ops(mut tfgrid_client, mut exp, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
