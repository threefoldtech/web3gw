module solution

import threefoldtech.threebot.tfgrid { Disk, Machine, MachinesModel, Network }

const presearch_cap = {
	Capacity.small:       CapacityPackage{
		cpu: 1
		memory: 2048
		size: 4096
	}
	Capacity.medium:      CapacityPackage{
		cpu: 2
		memory: 4096
		size: 8192
	}
	Capacity.large:       CapacityPackage{
		cpu: 4
		memory: 8192
		size: 16384
	}
	Capacity.extra_large: CapacityPackage{
		cpu: 8
		memory: 16384
		size: 32768
	}
}

pub struct Presearch {
pub:
	name        string
	farm_id     u64
	capacity Capacity
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

pub fn (mut s SolutionHandler) deploy_presearch(presearch Presearch) !PresearchResult {
	mut disks := []Disk{}
	if presearch.disk_size > 0{
		disks << Disk{
			size: presearch.disk_size
			mountpoint: '/var/lib/docker'
		}
	}

	machine := s.tfclient.machines_deploy(MachinesModel{
		name: presearch.name
		network: Network{
			add_wireguard_access: false
		}
		machines: [
			Machine{
				name: 'presearch_vm'
				farm_id: u32(presearch.farm_id)
				cpu: presearch_cap[presearch.capacity].cpu
				memory: presearch_cap[presearch.capacity].memory
				rootfs_size: presearch_cap[presearch.capacity].size
				flist: 'https://hub.grid.tf/tf-official-apps/presearch-v2.2.flist'
				env_vars: {
					'SSH_KEY':                     presearch.ssh_key
					'PRESEARCH_REGISTRATION_CODE': ''
				}
				disks: disks
				planetary: true
				public_ip: presearch.public_ipv4
			},
		]
	}) or {
		s.tfclient.machines_delete(presearch.name)!
		return error('failed to deploy presearch instance: ${err}')
	}

	return PresearchResult{
		name: presearch.name
		machine_ygg_ip: machine.machines[0].ygg_ip
		machine_ipv4: machine.machines[0].computed_ip4
	}
}

pub fn (mut s SolutionHandler) delete_presearch(presearch_name string) ! {
	s.tfclient.machines_delete(presearch_name)!
}

pub fn (mut s SolutionHandler) get_presearch(presearch_name string) !PresearchResult {
	machine := s.tfclient.machines_get(presearch_name)!

	return PresearchResult{
		name: presearch_name
		machine_ygg_ip: machine.machines[0].ygg_ip
		machine_ipv4: machine.machines[0].computed_ip4
	}
}
