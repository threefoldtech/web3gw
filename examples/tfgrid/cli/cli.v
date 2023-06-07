module main

import cli
import os
import rand
import log
import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { SolutionHandler }
import threefoldtech.threebot.explorer
import freeflowuniverse.crystallib.rpcwebsocket

fn print_manual(cmd cli.Command, level string) string {
	levels := {
		'root':      'modules'
		'module':    'operations'
		'operation': 'flags'
	}

	mut msg := ''

	msg += cmd.name + ' ' + cmd.description
	msg += '\n\n'
	msg += 'Usage: ' + cmd.usage
	msg += '\n\n'

	msg += 'Available ' + levels[level] or { 'commands' } + ':'

	for sub in cmd.commands {
		msg += '\n\t- ' + sub.name + '\t: ' + sub.description
	}

	msg += '\n\n'
	msg += 'Available flags:'
	for flag in cmd.flags {
		msg += '\n\t- ' + flag.name + '\t: ' + flag.description
	}

	println(msg)
	return msg
}

fn get_grid_modules_commands() []cli.Command {
	return [
		cli.Command{
			name: 'vms'
			description: 'machines module on the grid'
			usage: 'tfgrid-cli vms [operation] [flags]'
			disable_help: true
			disable_man: true
			execute: fn (cmd cli.Command) ! {
				print_manual(cmd, 'module')
				return
			}
			commands: get_vms_commands()
		},
		cli.Command{
			name: 'k8s'
			description: 'kubernetes module on the grid'
			usage: 'tfgrid-cli k8s [operation] [flags]'
			disable_help: true
			disable_man: true
			execute: fn (cmd cli.Command) ! {
				print_manual(cmd, 'module')
				return
			}
			commands: get_k8s_commands()
		},
		cli.Command{
			name: 'gws_fqdn'
			description: 'FQDN Gateway module on the grid'
			usage: 'tfgrid-cli gws_fqdn [operation] [flags]'
			disable_help: true
			disable_man: true
			execute: fn (cmd cli.Command) ! {
				print_manual(cmd, 'module')
				return
			}
			commands: get_gws_fqdn_commands()
		},
		cli.Command{
			name: 'gws_name'
			description: 'Name Gateway module on the grid'
			usage: 'tfgrid-cli gws_name [operation] [flags]'
			disable_help: true
			disable_man: true
			execute: fn (cmd cli.Command) ! {
				print_manual(cmd, 'module')
				return
			}
			commands: get_gws_name_commands()
		},
		cli.Command{
			name: 'zdb'
			description: 'ZDB module on the grid'
			usage: 'tfgrid-cli zdb [operation] [flags]'
			disable_help: true
			disable_man: true
			execute: fn (cmd cli.Command) ! {
				print_manual(cmd, 'module')
				return
			}
			commands: get_zdb_commands()
		},
	]
}

fn get_vms_commands() []cli.Command {
	return [
		cli.Command{
			name: 'create'
			description: 'Create a machine module on the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'network'
					description: 'Network name to deploy the machine on'
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
					name: 'times'
					description: 'Number of machines to deploy'
					required: false
				},
				cli.Flag{
					flag: .bool
					name: 'gateway'
					description: 'Either or not to deploy a gateway'
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
				execute_vms_create(cmd)!
				return
			}
		},
		cli.Command{
			name: 'delete'
			description: 'Delete a machine module on the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'network'
					description: 'Network name to delete the machines on'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_vms_delete(cmd)!
				return
			}
		},
		cli.Command{
			name: 'add'
			description: 'Add a machine to deployed network'
			flags: [
				cli.Flag{
					flag: .string
					name: 'network'
					description: 'Network name to deploy the machine on'
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
					name: 'times'
					description: 'Number of machines to deploy'
					required: false
				},
				cli.Flag{
					flag: .bool
					name: 'gateway'
					description: 'Either or not to deploy a gateway'
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
				execute_vms_create(cmd)!
				return
			}
		},
		cli.Command{
			name: 'remove'
			description: 'Remove a single machine from a deployed network'
			flags: [
				cli.Flag{
					flag: .string
					name: 'network'
					description: 'Network name where the machine is deployed'
					required: true
				},
				cli.Flag{
					flag: .string
					name: 'machine_name'
					description: 'Machine name to remove'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_vms_remove(cmd)!
				return
			}
		},
		cli.Command{
			name: 'get'
			description: 'Get a deployed machine from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'network'
					description: 'Network name where the machine is deployed'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_vms_get(cmd)!
				return
			}
		},
	]
}

fn get_solution_handler(cmd cli.Command) !(SolutionHandler, log.Logger) {
	mut logger := log.Log{
		level: .info
	}

	mut rpc_client := rpcwebsocket.new_rpcwsclient('ws://127.0.0.1:8080', &logger)!
	_ := spawn rpc_client.run()

	mut grid_network := cmd.flags.get_string('grid')!
	mnemonic := cmd.flags.get_string('mnemonic')!

	if grid_network == '' {
		grid_network = 'dev'
	}

	mut tfgrid_client := tfgrid.new(mut rpc_client)
	mut explorer_client := explorer.new(mut rpc_client)

	tfgrid_client.load(tfgrid.Credentials{
		mnemonic: mnemonic
		network: grid_network
	})!
	explorer_client.load(grid_network)!

	mut sl := SolutionHandler{
		tfclient: &tfgrid_client
		explorer: &explorer_client
	}

	return sl, logger
}

fn main() {
	mut app := cli.Command{
		name: 'tfgrid-cli'
		description: 'is a command line interface for the ThreeFold Grid.'
		usage: 'tfgrid-cli [module] [operation] [flags]'
		disable_help: true
		disable_man: true
		execute: fn (cmd cli.Command) ! {
			print_manual(cmd, 'root')
			return
		}
		commands: get_grid_modules_commands()
		flags: [
			cli.Flag{
				flag: .string
				name: 'grid'
				description: 'Grid Network (dev, qa, test, main)'
				required: false
				global: true
			},
			cli.Flag{
				flag: .string
				name: 'mnemonic'
				description: 'Secret mnemonic to use to access the grid'
				required: true
				global: true
			},
		]
	}
	app.setup()
	app.parse(os.args)
}

fn execute_vms_create(cmd cli.Command) ! {
	// get flags
	mut network := cmd.flags.get_string('network')!
	mut capacity := cmd.flags.get_string('capacity')!
	mut times := cmd.flags.get_int('times')!
	mut gateway := cmd.flags.get_bool('gateway')!
	mut wg := cmd.flags.get_bool('wg')!
	mut ssh_key := cmd.flags.get_string('ssh_key')!

	// assign default values
	if network == '' {
		network = rand.string(8)
	}
	if capacity == '' {
		capacity = 'small'
	}
	if times == 0 {
		times = 1
	}
	if ssh_key == '' {
		default_ssh_key := os.execute('cat ~/.ssh/id_rsa.pub')
		ssh_key = default_ssh_key.output
	}

	// execute
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	res := solution_handler.create_vm(solution.VM{
		network: network
		capacity: capacity
		times: u32(times)
		gateway: gateway
		add_wireguard_access: wg
		ssh_key: ssh_key
	})!
	logger.info('${res}')

	return
}

fn execute_vms_get(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	mut network := cmd.flags.get_string('network')!

	if network == '' {
		logger.info('Please provide a network name')
	}

	res := solution_handler.get_vm(network)!
	logger.info('${res}')
}

fn execute_vms_remove(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	mut network := cmd.flags.get_string('network')!
	mut machine_name := cmd.flags.get_string('machine_name')!

	if network == '' {
		logger.info('Please provide a network name')
	}

	if machine_name == '' {
		logger.info('Please provide a machine name')
	}

	res := solution_handler.remove_vm(network, machine_name)!
	logger.info('${res}')
}

fn execute_vms_delete(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	mut network := cmd.flags.get_string('network')!

	if network == '' {
		logger.info('Please provide a network name')
	}

	solution_handler.delete_vm(network)!
	logger.info('deleted network ${network}')
}

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

fn get_gws_fqdn_commands() []cli.Command {
	return [
		cli.Command{
			name: 'create'
			description: 'Get a deployed cluster from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'FQDN Gateway name'
				},
				cli.Flag{
					flag: .string
					name: 'fqdn'
					description: 'The full domain name'
					required: true
				},
				cli.Flag{
					flag: .string
					name: 'backend'
					description: 'backend of your deployment (ex: "http://1.1.1.1:9000")'
					required: true
				},
				cli.Flag{
					flag: .int
					name: 'node_id'
					description: 'grid node to deploy the fqdn Gateway on'
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_gws_fqdn_get(cmd)!
				return
			}
		},
		cli.Command{
			name: 'delete'
			description: 'delete a deployed fqdn Gateway from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'FQDN Gateway name'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_gws_fqdn_get(cmd)!
				return
			}
		},
		cli.Command{
			name: 'get'
			description: 'Get a deployed fqdn Gateway from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'FQDN Gateway name'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_gws_fqdn_get(cmd)!
				return
			}
		},
	]
}

fn execute_gws_fqdn_create(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!
	mut fqdn := cmd.flags.get_string('fqdn')!
	mut backend := cmd.flags.get_string('backend')!
	mut node_id := cmd.flags.get_int('node_id')!

	// assign default values
	if fqdn == '' {
		logger.error('Please provide a fqdn')
	}
	if backend == '' {
		logger.error('Please provide a backend')
	}
	if name == '' {
		name = rand.string(8)
	}

	res := solution_handler.tfclient.gateways_deploy_fqdn(tfgrid.GatewayFQDN{
		name: name
		node_id: node_id
		backends: [backend]
		fqdn: fqdn
	})!

	logger.info('${res}')
}

fn execute_gws_fqdn_delete(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!

	// assign default values
	if name == '' {
		logger.error('Please provide a name')
	}

	solution_handler.tfclient.gateways_delete_fqdn(name)!
	logger.info('deleted fqdn ${name}')
}

fn execute_gws_fqdn_get(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!

	// assign default values
	if name == '' {
		logger.error('Please provide a name')
	}

	res := solution_handler.tfclient.gateways_get_fqdn(name)!

	logger.info('${res}')
}

fn get_gws_name_commands() []cli.Command {
	return [
		cli.Command{
			name: 'create'
			description: 'Get a deployed cluster from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Name Gateway name'
				},
				cli.Flag{
					flag: .string
					name: 'backend'
					description: 'backend of your deployment (ex: "http://1.1.1.1:9000")'
					required: true
				},
				cli.Flag{
					flag: .int
					name: 'node_id'
					description: 'grid node to deploy the name Gateway on'
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_gws_name_get(cmd)!
				return
			}
		},
		cli.Command{
			name: 'delete'
			description: 'delete a deployed name Gateway from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Name Gateway name'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_gws_name_get(cmd)!
				return
			}
		},
		cli.Command{
			name: 'get'
			description: 'Get a deployed name Gateway from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Name Gateway name'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_gws_name_get(cmd)!
				return
			}
		},
	]
}

fn execute_gws_name_create(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!
	mut backend := cmd.flags.get_string('backend')!
	mut node_id := cmd.flags.get_int('node_id')!

	// assign default values
	if backend == '' {
		logger.error('Please provide a backend')
	}
	if name == '' {
		name = rand.string(8)
	}

	res := solution_handler.tfclient.gateways_deploy_name(tfgrid.GatewayName{
		name: name
		node_id: node_id
		backends: [backend]
	})!

	logger.info('${res}')
}

fn execute_gws_name_delete(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!

	// assign default values
	if name == '' {
		logger.error('Please provide a name')
	}

	solution_handler.tfclient.gateways_delete_name(name)!
	logger.info('deleted name ${name}')
}

fn execute_gws_name_get(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!

	// assign default values
	if name == '' {
		logger.error('Please provide a name')
	}

	res := solution_handler.tfclient.gateways_get_name(name)!

	logger.info('${res}')
}

fn get_zdb_commands() []cli.Command {
	return [
		cli.Command{
			name: 'create'
			description: 'deploy a zdb on the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Zdb name'
					required: true
				},
				cli.Flag{
					flag: .int
					name: 'node_id'
					description: 'Node id to deploy the zdb on'
					required: true
				},
				cli.Flag{
					flag: .int
					name: 'size'
					description: 'size of zdb in GB'
					required: true
				},
				cli.Flag{
					flag: .string
					name: 'password'
					description: 'secure password for zdb'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_zdb_create(cmd)!
				return
			}
		},
		cli.Command{
			name: 'get'
			description: 'get a zdb from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Zdb name'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_zdb_get(cmd)!
				return
			}
		},
		cli.Command{
			name: 'delete'
			description: 'delete a zdb from the grid'
			flags: [
				cli.Flag{
					flag: .string
					name: 'name'
					description: 'Zdb name'
					required: true
				},
			]
			execute: fn (cmd cli.Command) ! {
				execute_zdb_delete(cmd)!
				return
			}
		},
	]
}

fn execute_zdb_create(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!
	mut node_id := cmd.flags.get_int('node_id')!
	mut password := cmd.flags.get_string('password')!
	mut size := cmd.flags.get_int('size')!

	// assign default values
	if name == '' {
		name = rand.string(8)
	}
	if password == '' {
		password = rand.string(8)
	}
	if size == 0 {
		size = 10
	}

	res := solution_handler.tfclient.zdb_deploy(tfgrid.ZDB{
		name: name
		node_id: node_id
		password: password
		size: size
	})!

	logger.info('${res}')
}

fn execute_zdb_get(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!

	// assign default values
	if name == '' {
		logger.error('Please provide a name')
	}
	res := solution_handler.tfclient.zdb_get(name)!

	logger.info('${res}')
}

fn execute_zdb_create(cmd cli.Command) ! {
	mut solution_handler, mut logger := get_solution_handler(cmd)!

	// get flags
	mut name := cmd.flags.get_string('name')!

	// assign default values
	if name == '' {
		logger.error('Please provide a name')
	}
	solution_handler.tfclient.zdb_delete(name)!

	logger.info('deleted zdb ${name}')
}
