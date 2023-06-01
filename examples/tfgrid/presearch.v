module main

import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { Presearch, SolutionHandler }
import log

fn test_presearch_ops(mut s SolutionHandler, mut logger log.Logger) ! {
	model_name := 'hamadapresearch'
	deploy_res := s.deploy_presearch(Presearch{
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
		s.delete_presearch(model_name) or { logger.error('failed to delete presearch: ${err}') }
	}

	get_res := s.get_presearch(model_name)!
	logger.info('${get_res}')
}

fn main() {
	mut logger := log.Log{
		level: .info
	}

	mut tfgrid_client, exp := tfgrid.cli(mut logger) or {
		logger.error('failed to initialize tfgrid client: ${err}')
		exit(1)
	}

	mut s := SolutionHandler{
		tfclient: &tfgrid_client
		explorer: &exp
	}

	test_presearch_ops(mut s, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
