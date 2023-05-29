module main

import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.explorer
import log

fn test_discourse_ops(mut client tfgrid.TFGridClient, mut exp explorer.ExplorerClient, mut logger log.Logger) ! {
	model_name := 'hamadadiscourse'
	deploy_res := client.deploy_discourse(mut exp, tfgrid.Discourse{
		name: model_name
		cpu: 4
		memory: 4096
		rootfs_size: 10240
		disk_size: 10
		ssh_key: 'hamada ssh key'
		developer_email: 'em@mail.com'
		smtp_username: 'user1'
		smtp_password: 'pass1'
		threebot_private_key: 'asdfqwer12asdfqwer12asdfqwer12asdfqwer12'
		flask_secret_key: 'asdfqwer12asdfqwer12asdfqwer12asdfqwer12'
	})!
	logger.info('${deploy_res}')

	defer {
		client.delete_discourse(model_name) or {
			logger.error('failed to delete discourse: ${err}')
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

	test_discourse_ops(mut tfgrid_client, mut exp, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
