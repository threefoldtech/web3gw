module main

import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { Discourse, SolutionHandler }
import log

fn run_discourse_ops(mut s SolutionHandler, mut logger log.Logger) ! {
	model_name := 'hamadadiscourse'
	deploy_res := s.deploy_discourse(Discourse{
		name: model_name
		capacity: 'medium'
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
		s.delete_discourse(model_name) or { logger.error('failed to delete discourse: ${err}') }
	}

	get_res := s.get_discourse(model_name)!
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

	run_discourse_ops(mut s, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
