module main

import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { SolutionHandler, Taiga, Capacity }
import log

fn run_taiga_ops(mut s SolutionHandler, mut logger log.Logger) ! {
	model_name := 'hamadataiga'
	deploy_res := s.deploy_taiga(Taiga{
		name: model_name
		farm_id: 1
		capacity: Capacity.small
		disk_size: 20
		ssh_key: 'key'
		admin_username: 'user1'
		admin_password: 'pass'
		admin_email: 'email@gmail.com'
	})!
	logger.info('${deploy_res}')

	defer {
		s.delete_taiga(model_name) or { logger.error('failed to delete taiga: ${err}') }
	}

	get_res := s.get_taiga(model_name)!
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

	run_taiga_ops(mut s, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
