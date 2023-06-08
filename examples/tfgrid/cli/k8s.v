module main

import cli
import os
import rand
import log
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution
import threefoldtech.threebot.explorer
import freeflowuniverse.crystallib.rpcwebsocket

fn get_k8s_commands() []cli.Command {
	return [
		cli.Command{
			name: 'create'
			description: 'Create a kubernetes cluster on the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'cluster'
					description: 'Cluster name'
					required: false
				},
				cli.Flag{
					flag: .string
					name: 'capacity'
					description: 'Capacity package (small, medium, large, extra-large)'
					required: false
				},
				cli.Flag{
					flag: .int
					name: 'replica'
					description: 'Number of workers on the cluster'
					required: false
				},
				cli.Flag{
					flag: .bool
					name: 'wg'
					description: 'Either or not to add a wireguard access to the nodes'
					required: false
				},
				cli.Flag{
					flag: .string
					name: 'ssh_key'
					description: 'Public ssh key to add to the machine'
					required: false
				},
				cli.Flag{
					flag: .bool
					name: 'public_ip'
					description: 'Either or not to add a public ip to the master'
					required: false
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_k8s_create(cmd)!
				return
			}
		},
		cli.Command{
			name: 'delete'
			description: 'Delete a kubernetes cluster on the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'cluster'
					description: 'Cluster name'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_k8s_delete(cmd)!
				return
			}
		},
		cli.Command{
			name: 'add'
			description: 'Add a worker to deployed cluster'
			flags: [
				cli.Flag{
					flag: .string
					name: 'cluster'
					description: 'Cluster name to add worker to'
					required: false
				},
				cli.Flag{
					flag: .string
					name: 'capacity'
					description: 'Capacity package (small, medium, large, extra-large)'
					required: false
				},
				cli.Flag{
					flag: .bool
					name: 'wg'
					description: 'Either or not to add a wireguard access to the machine'
					required: false
				},
				cli.Flag{
					flag: .string
					name: 'ssh_key'
					description: 'Public ssh key to add to the machine'
					required: false
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_k8s_add(cmd)!
				return
			}
		},
		cli.Command{
			name: 'remove'
			description: 'Remove a single worker from a deployed cluster'
			flags: [
				cli.Flag{
					flag: .string
					name: 'cluster'
					description: 'Cluster name'
					required: true
				},
				cli.Flag{
					flag: .string
					name: 'worker'
					description: 'Worker name to remove'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_k8s_remove(cmd)!
				return
			}
		},
		cli.Command{
			name: 'get'
			description: 'Get a deployed cluster from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'cluster'
					description: 'Cluster name'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_k8s_get(cmd)!
				return
			}
		},
	]
}

fn execute_k8s_create(cmd cli.Command) ! {
	// get flags
	mut cluster := cmd.flags.get_string('cluster')!
	mut capacity := cmd.flags.get_string('capacity')!
	mut replica := cmd.flags.get_int('replica')!
	mut wg := cmd.flags.get_bool('wg')!
	mut public_ip := cmd.flags.get_bool('public_ip')!
	mut ssh_key := cmd.flags.get_string('ssh_key')!

	// assign default values
	if cluster == '' {
		cluster = rand.string(8)
	}
	if capacity == '' {
		capacity = 'small'
	}
	if replica == 0 {
		replica = 1
	}
	if ssh_key == '' {
		default_ssh_key := os.execute('cat ~/.ssh/id_rsa.pub')
		ssh_key = default_ssh_key.output
	}

	// execute
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	res := solution_handler.create_k8s(solution.K8s{
		name: cluster
		capacity: capacity
		replica: replica
		wg: wg
		public_ip: public_ip
		ssh_key: ssh_key
	})!

	logger.info('${res}')

	return
}

fn execute_k8s_get(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	mut cluster := cmd.flags.get_string('cluster')!

	if cluster == '' {
		logger.info('Please provide a cluster name')
	}

	res := solution_handler.get_k8s(cluster)!
	logger.info('${res}')
}

fn execute_k8s_remove(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	mut cluster := cmd.flags.get_string('cluster')!
	mut worker := cmd.flags.get_string('worker')!

	if cluster == '' {
		logger.info('Please provide a cluster name')
	}

	if worker == '' {
		logger.info('Please provide a worker name')
	}

	res := solution_handler.remove_k8s_worker(cluster, worker)!
	logger.info('${res}')
}

fn execute_k8s_delete(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	mut cluster := cmd.flags.get_string('cluster')!

	if cluster == '' {
		logger.info('Please provide a cluster name')
	}

	solution_handler.delete_vm(cluster)!
	logger.info('deleted cluster ${cluster}')
}

fn execute_k8s_add(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut cluster := cmd.flags.get_string('cluster')!
	mut capacity := cmd.flags.get_string('capacity')!
	mut wg := cmd.flags.get_bool('wg')!
	mut ssh_key := cmd.flags.get_string('ssh_key')!

	// assign default values
	if cluster == '' {
		logger.error('Please provide a cluster name')
	}
	if capacity == '' {
		capacity = 'small'
	}
	if ssh_key == '' {
		default_ssh_key := os.execute('cat ~/.ssh/id_rsa.pub')
		ssh_key = default_ssh_key.output
	}

	// execute

	res := solution_handler.add_k8s_worker(solution.K8s{
		name: cluster
		capacity: capacity
		wg: wg
		ssh_key: ssh_key
	})!

	logger.info('${res}')
}
