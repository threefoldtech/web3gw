module main

import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { Peertube, SolutionHandler }
import log

fn test_peertube_ops(mut s SolutionHandler, mut logger log.Logger) ! {
	model_name := 'hamadapeertube'
	deploy_res := s.deploy_peertube(Peertube{
		name: model_name
		cpu: 2
		memory: 4096
		rootfs_size: 10240
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
		s.delete_peertube(model_name) or { logger.error('failed to delete peertube: ${err}') }
	}

	get_res := s.get_peertube(model_name)!
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

	mut s := SolutionHandler{
		tfclient: &tfgrid_client
		explorer: &exp
	}

	test_peertube_ops(mut s, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
