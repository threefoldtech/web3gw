module main

import freeflowuniverse.crystallib.rpcwebsocket
import threefoldtech.threebot.tfgrid
import flag
import log
import os
import time

const (
	default_server_address = 'ws://127.0.0.1:8080'
)

fn test_machines_ops(mut client tfgrid.TFGridClient, mut logger log.Logger) ! {
	project_name := 'testMachinesOps'

	// deploy
	mut disks := []tfgrid.Disk{}
	disks << tfgrid.Disk{
		size: 10
		mountpoint: '/mnt/disk1'
	}
	mut machines := []tfgrid.Machine{}
	machines << tfgrid.Machine{
		name: 'vm1'
		node_id: 33
		cpu: 2
		memory: 2048
		rootfs_size: 1024
		env_vars: {
			'SSH_KEY': 'ssh-rsa ...'
		}
		disks: disks
	}
	machines_model := tfgrid.MachinesModel{
		name: project_name
		network: tfgrid.Network{
			add_wireguard_access: true
		}
		machines: machines
		metadata: 'metadata1'
		description: 'description'
	}

	res := client.machines_deploy(machines_model)!
	logger.info('${res}')

	// get
	res_2 := client.machines_get(project_name)!
	logger.info('${res_2}')

	// delete
	client.machines_delete(project_name)!
}

fn main() {
	mut fp := flag.new_flag_parser(os.args)
	fp.application('Welcome to the web3_proxy client. The web3_proxy client allows you to execute all remote procedure calls that the web3_proxy server can handle.')
	fp.limit_free_args(0, 0)!
	fp.description('')
	fp.skip_executable()
	mnemonic := fp.string('mnemonic', `m`, '', 'The mnemonic to be used to call any function')
	address := fp.string('address', `a`, '${default_server_address}', 'The address of the web3_proxy server to connect to.')
	debug_log := fp.bool('debug', 0, false, 'By setting this flag the client will print debug logs too.')
	_ := fp.finalize() or {
		eprintln(err)
		println(fp.usage())
		exit(1)
	}

	mut logger := log.Logger(&log.Log{
		level: if debug_log { .debug } else { .info }
	})

	mut myclient := rpcwebsocket.new_rpcwsclient(address, &logger) or {
		logger.error('Failed creating rpc websocket client: ${err}')
		exit(1)
	}

	_ := spawn myclient.run()

	mut tfgrid_client := tfgrid.new(mut client)

	// ADD YOUR CALLS HERE
	tfgrid_client.load(tfgrid.Credentials{
		mnemonic: mnemonic
		network: 'dev'
	})!

	test_machines_ops(mut tfgrid_client, mut logger) or {
		logger.error('Failed executing machines ops: ${err}')
		exit(1)
	}
}
