module main

import freeflowuniverse.crystallib.rpcwebsocket
import threefoldtech.web3gw.tfgrid 
import flag
import log
import os

const (
	default_server_address = 'ws://127.0.0.1:8080'
)

fn run_machines_ops(mut client tfgrid.TFGridClient, mut logger log.Logger) ! {
	model_name := 'testMachinesOps'

	res := client.machines_deploy(tfgrid.MachinesModel{
		name: model_name
		network: tfgrid.Network{
			add_wireguard_access: false
		}
		machines: [
			tfgrid.Machine{
				name: 'vm1'
				cpu: 2
				memory: 2048
				rootfs_size: 1024
				env_vars: {
					'SSH_KEY': 'ssh-rsa ...'
				}
				disks: [tfgrid.Disk{
					size: 10
					mountpoint: '/mnt/disk1'
				}]
			},
		]
		metadata: 'metadata1'
		description: 'description'
	})!
	logger.info('${res}')

	defer {
		client.machines_delete(model_name) or {
			logger.error('failed while deleting machines: ${err}')
		}
	}

	add_res := client.machines_add(tfgrid.AddMachine{
		model_name: model_name
		machine: tfgrid.Machine{
			name: 'vm3'
			cpu: 2
			memory: 2048
			rootfs_size: 1024
			env_vars: {
				'SSH_KEY': 'ssh-rsa ...'
			}
			disks: [tfgrid.Disk{
				size: 10
				mountpoint: '/mnt/disk1'
			}]
		}
	})!
	logger.info('${add_res}')

	add_res2 := client.machines_add(tfgrid.AddMachine{
		model_name: model_name
		machine: tfgrid.Machine{
			name: 'vm10'
			cpu: 2
			memory: 2048
			rootfs_size: 1024
			env_vars: {
				'SSH_KEY': 'ssh-rsa ...'
			}
			disks: [tfgrid.Disk{
				size: 10
				mountpoint: '/mnt/disk1'
			}]
		}
	})!
	logger.info('${add_res2}')

	remove_res := client.machines_remove(tfgrid.RemoveMachine{
		model_name: model_name
		machine_name: 'vm3'
	})!
	logger.info('${remove_res}')

	res_3 := client.machines_get(model_name)!
	logger.info('${res_3}')
}

fn main() {
	mut logger := log.Log{
		level: .info
	}

	mut fp := flag.new_flag_parser(os.args)
	fp.application('Welcome to the web3_proxy client. The web3_proxy client allows you to execute all remote procedure calls that the web3_proxy server can handle.')
	fp.limit_free_args(0, 0)!
	fp.description('')
	fp.skip_executable()
	mnemonic := fp.string('mnemonic', `m`, '', 'The mnemonic to be used to call any function')
	network := fp.string('network', `n`, 'dev', 'TF network to use')
	address := fp.string('address', `a`, '${default_server_address}', 'The address of the web3_proxy server to connect to.')
	debug_log := fp.bool('debug', 0, false, 'By setting this flag the client will print debug logs too.')
	_ := fp.finalize() or {
		logger.error('${err}')
		println(fp.usage())
		exit(1)
	}

	logger.set_level(if debug_log { .debug } else { .info })

	mut myclient := rpcwebsocket.new_rpcwsclient(address, &logger) or {
		logger.error('Failed creating rpc websocket client: ${err}')
		exit(1)
	}

	_ := spawn myclient.run()

	mut tfgrid_client := tfgrid.new(mut myclient)

	tfgrid_client.load(tfgrid.Credentials{
		mnemonic: mnemonic
		network: network
	})!

	run_machines_ops(mut tfgrid_client, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
