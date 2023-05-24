module main

import threefoldtech.threebot.tfgrid

pub struct Presearch {
pub:
	name        string
	farm_id     u64
	cpu         u32
	memory      u32 // in mega bytes
	rootfs_size u32 // in mega bytes
	disk_size   u32 // in giga bytes
	ssh_key     string
	public_ipv4 bool
}

pub struct PresearchResult {
pub:
	name           string
	machine_ygg_ip string
	machine_ipv4   string
}

fn deploy_presearch(mut client tfgrid.TFGridClient, presearch Presearch) !PresearchResult {
	machine := client.machines_deploy(tfgrid.MachinesModel{
		name: presearch.name
		network: tfgrid.Network{
			add_wireguard_access: false
		}
		machines: [
			tfgrid.Machine{
				name: 'presearch_vm'
				farm_id: u32(presearch.farm_id)
				cpu: presearch.cpu
				memory: presearch.memory
				rootfs_size: presearch.rootfs_size
				flist: 'https://hub.grid.tf/tf-official-apps/presearch-v2.2.flist'
				env_vars: {
					'SSH_KEY':                     presearch.ssh_key
					'PRESEARCH_REGISTRATION_CODE': ''
				}
				disks: [
					tfgrid.Disk{
						size: presearch.disk_size
						mountpoint: '/var/lib/docker'
					},
				]
				planetary: true
				public_ip: presearch.public_ipv4
			},
		]
	}) or {
		client.machines_delete(presearch.name)!
		return error('failed to deploy presearch instance: ${err}')
	}

	return PresearchResult{
		name: presearch.name
		machine_ygg_ip: machine.machines[0].ygg_ip
		machine_ipv4: machine.machines[0].computed_ip4
	}
}

fn delete_presearch(mut client tfgrid.TFGridClient, presearch_name string) ! {
	client.machines_delete(presearch_name)!
}

fn get_presearch(mut client tfgrid.TFGridClient, presearch_name string) !PresearchResult {
	machine := client.machines_get(presearch_name)!

	return PresearchResult{
		name: presearch_name
		machine_ygg_ip: machine.machines[0].ygg_ip
		machine_ipv4: machine.machines[0].computed_ip4
	}
}
