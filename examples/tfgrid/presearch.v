module main

import threefoldtech.threebot.tfgrid
import log

fn test_presearch_ops(mut client tfgrid.TFGridClient, mut logger log.Logger) ! {
	model_name := 'hamadapresearch'
	deploy_res := client.deploy_presearch(tfgrid.Presearch{
		name: model_name
		farm_id: 1
		cpu: 2
		memory: 4096
		rootfs_size: 10240
		disk_size: 10
		ssh_key: 'key'
		public_ipv4: false
	})!
	logger.info('${deploy_res}')

	defer {
		client.delete_presearch(model_name) or {
			logger.error('failed to delete presearch: ${err}')
		}
	}

	get_res := client.get_presearch(model_name)!
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

	test_presearch_ops(mut tfgrid_client, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
