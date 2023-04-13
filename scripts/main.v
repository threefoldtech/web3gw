module main

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }
import stellar
import tfchain
import tfgrid

import flag
import log
import os

const (
	default_server_address = "http://127.0.0.1:8080"
)

fn mymachines() tfgrid.MachinesModel {
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
			"SSH_KEY": 'ssh-rsa ...'
		}
		disks: disks
	}
	return tfgrid.MachinesModel{
		name: 'project1'
		network: tfgrid.Network{
			add_wireguard_access: true
		}
		machines: machines
		metadata: 'metadata1'
		description: 'description'
	}
}


fn execute_rpcs(mut client RpcWsClient, mut logger log.Logger) ! {
	// ADD YOUR CALLS HERE
	tfgrid.load(mut client, tfgrid.Credentials{
		mnemonic: "" // FILL IN YOUR MNEMONIC HERE 
		network: "dev"
	})!

	machines_model := tfgrid.machines_deploy(mut client, mymachines())!
	logger.info("${machines_model}")

	machines_model_2 := tfgrid.machines_get(mut client, "project1")!
	logger.info("${machines_model_2}")

	tfgrid.machines_delete(mut client, "project1")!
}

fn execute_rpcs_tfchain(mut client RpcWsClient, mut logger log.Logger) ! {
	tfchain.load(mut client, "devnet", "")! // FILL IN YOUR MNEMONIC HERE
	tfchain.transfer(mut client, tfchain.Transfer{amount: 1000, destination: ""})! // FILL IN SOME DESTINATION
	height := tfchain.height(mut client)!
	println("Height is ${height}")

	my_balance := tfchain.balance(mut client, "5Ek9gJ3iQFyr1HB5aTpqThqbGk6urv8Rnh9mLj5PD6GA26MS")! // FILL IN ADDRESS
	println("My balance: ${my_balance}")

	twin_32 := tfchain.get_twin(mut client, 32)! // decoding to json is not yet working but add --debug and you will see the received message from server
	println("Twin with id 32: ${twin_32}")

}


fn main() {
	mut fp := flag.new_flag_parser(os.args)
	fp.application('Welcome to the web3_proxy client. The web3_proxy client allows you to execute all remote procedure calls that the web3_proxy server can handle.')
	fp.limit_free_args(0, 0)!
	fp.description('')
	fp.skip_executable()
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
		logger.error("Failed creating rpc websocket client: $err")
		exit(1)
	}
	_ := spawn myclient.run() //QUESTION: why is that in thread?
	/*
	execute_rpcs(mut myclient, mut logger) or {
		logger.error("Failed executing calls: $err")
		exit(1)
	}
	*/
	execute_rpcs_tfchain(mut myclient, mut logger) or {
		logger.error("Failed executing calls: $err")
		exit(1)
	}
}