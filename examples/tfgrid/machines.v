module main

import threefoldtech.threebot.tfgrid
import log

fn test_machines_ops(mut client tfgrid.TFGridClient, mut logger log.Logger) ! {
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

	mut tfgrid_client, _ := tfgrid.cli(mut logger) or {
		logger.error('failed to initialize tfgrid client: ${err}')
		exit(1)
	}

	test_machines_ops(mut tfgrid_client, mut logger) or {
		logger.error('${err}')
		exit(1)
	}
}
