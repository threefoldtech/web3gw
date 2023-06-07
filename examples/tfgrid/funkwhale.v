module main

import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { Funkwhale, SolutionHandler, Capacity }
import log

fn run_funkwhale_ops(mut s SolutionHandler, mut logger log.Logger) ! {
	model_name := 'hamadafunkwhale'
	deploy_res := s.deploy_funkwhale(Funkwhale{
		name: model_name
		farm_id: 1
		capacity: Capacity.medium
		admin_email: 'admin@gmail.com'
		admin_username: 'user1'
		admin_password: 'pass1'
	})!
	logger.info('${deploy_res}')

	defer {
		s.delete_funkwhale(model_name) or { logger.error('failed to delete funkwhale: ${err}') }
	}

	get_res := s.get_funkwhale(model_name)!
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

	run_funkwhale_ops(mut s, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
