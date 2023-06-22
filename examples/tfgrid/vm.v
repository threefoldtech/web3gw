module main

import threefoldtech.threebot.tfgrid { RemoveVM, TFGridClient, VM, VMResult }
import log { Logger }
import flag { FlagParser }
import os
import freeflowuniverse.crystallib.rpcwebsocket
import rand

const (
	default_server_address = 'ws://127.0.0.1:8080'
)

fn deploy_vm(mut fp FlagParser, mut t TFGridClient) !VMResult {
	fp.usage_example('deploy [options]')

	name := fp.string('vm_name', `m`, rand.string(6), 'VM name')
	network := fp.string_opt('vm_network', `v`, 'Name of the VM network')!
	farm_id := fp.int('farm_id', `f`, 0, 'Farm ID to deploy on')
	capacity := fp.string('capacity', `c`, 'medium', 'Capacity of the instance')
	times := fp.int('times', `t`, 1, 'The number of vms to deploy')
	disk_size := fp.int('disk_size', `d`, 0, 'Size of disk the will be mounted on each vm')
	gateway := fp.bool('gateway', `g`, false, 'True to add a gateway for each vm')
	wg := fp.bool('wg', `w`, false, 'True to add a wireguard access point to the network')
	add_public_ipv4 := fp.bool('add_public_ipv4', `4`, false, 'True to add a public ipv4 to each vm')
	add_public_ipv6 := fp.bool('add_public_ipv6', `6`, false, 'True to add a public ipv6 to each vm')
	ssh_key := fp.string('ssh_key', `s`, '', 'Public SSH Key to access the instance')
	_ := fp.finalize()!

	vm := VM{
		name: name
		network: network
		farm_id: u32(farm_id)
		capacity: capacity
		ssh_key: ssh_key
		times: u32(times)
		disk_size: u32(disk_size)
		gateway: gateway
		add_wireguard_access: wg
		add_public_ipv4: add_public_ipv4
		add_public_ipv6: add_public_ipv6
	}

	return t.deploy_vm(vm)!
}

fn get_vm(mut fp FlagParser, mut t TFGridClient) !VMResult {
	fp.usage_example('get [options]')

	network := fp.string_opt('vm_network', `v`, 'Name of the VM network')!
	_ := fp.finalize()!

	return t.get_vm(network)!
}

fn delete_vm(mut fp FlagParser, mut t TFGridClient) ! {
	fp.usage_example('delete [options]')

	network := fp.string_opt('vm_network', `v`, 'Name of the VM network')!
	_ := fp.finalize()!

	return t.delete_vm(network)
}

fn remove_vm(mut fp FlagParser, mut t TFGridClient) !VMResult {
	fp.usage_example('remove [options]')

	network := fp.string_opt('vm_network', `v`, 'Name of the VM network')!
	vm_name := fp.string_opt('vm', `v`, 'Name of the VM to be removed')!
	_ := fp.finalize()!

	return t.remove_vm(RemoveVM{
		network: network
		vm_name: vm_name
	})!
}

fn main() {
	mut fp := flag.new_flag_parser(os.args)
	fp.application('Welcome to the web3_proxy client. The web3_proxy client allows you to execute all remote procedure calls that the web3_proxy server can handle.')
	fp.description('')
	fp.skip_executable()
	fp.allow_unknown_args()

	mnemonic := fp.string_opt('mnemonic', `m`, 'The mnemonic to be used to call any function') or {
		eprintln('${err}')
		exit(1)
	}
	network := fp.string('network', `n`, 'dev', 'TF network to use')
	address := fp.string('address', `a`, '${default_server_address}', 'The address of the web3_proxy server to connect to.')
	debug_log := fp.bool('debug', 0, false, 'By setting this flag the client will print debug logs too.')
	operation := fp.string_opt('operation', `o`, 'Required operation to perform ')!
	remainig_args := fp.finalize() or {
		eprintln('${err}')
		exit(1)
	}

	mut logger := Logger(&log.Log{
		level: if debug_log { .debug } else { .info }
	})

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

	match operation {
		'deploy' {
			mut new_fp := flag.new_flag_parser(remainig_args)
			res := deploy_vm(mut new_fp, mut tfgrid_client) or {
				logger.error('${err}')
				exit(1)
			}
			logger.info('${res}')
		}
		'get' {
			mut new_fp := flag.new_flag_parser(remainig_args)
			res := get_vm(mut new_fp, mut tfgrid_client) or {
				logger.error('${err}')
				exit(1)
			}
			logger.info('${res}')
		}
		'delete' {
			mut new_fp := flag.new_flag_parser(remainig_args)
			delete_vm(mut new_fp, mut tfgrid_client) or {
				logger.error('${err}')
				exit(1)
			}
		}
		'remove' {
			mut new_fp := flag.new_flag_parser(remainig_args)
			res := remove_vm(mut new_fp, mut tfgrid_client) or {
				logger.error('${err}')
				exit(1)
			}
			logger.info('${res}')
		}
		else {
			logger.error('operation ${operation} is invalid')
			exit(1)
		}
	}
}
