module solution

import threefoldtech.threebot.explorer
import threefoldtech.threebot.tfgrid { Disk, GatewayName, Machine, MachinesModel, Network }

const taiga_cap = {
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

pub struct Taiga {
pub:
	name           string
	farm_id        u64
	capacity Capacity
	disk_size      u32 // in giga bytes
	ssh_key        string
	admin_username string
	admin_password string
	admin_email    string
}

pub struct TaigaResult {
pub:
	name           string
	machine_ygg_ip string
	gateway_name   string
}

pub fn (mut s SolutionHandler) deploy_taiga(taiga Taiga) !TaigaResult {
	mut filter := explorer.NodeFilter{
		status: 'up'
		dedicated: false
		domain: true
	}

	if taiga.farm_id != 0 {
		filter = explorer.NodeFilter{
			farm_ids: [taiga.farm_id]
		}
	}

	gateway_nodes := s.explorer.nodes(explorer.NodesRequestParams{
		filters: filter
		pagination: explorer.Limit{
			size: 1
		}
	})!

	if gateway_nodes.nodes.len == 0 {
		return error('failed to find an eligible node for gateway')
	}

	gateway_node_id := gateway_nodes.nodes[0].node_id
	domain := gateway_nodes.nodes[0].public_config.domain

	mut disks := []Disk{}
	if taiga.disk_size > 0{
		disks << Disk{
			size: taiga.disk_size
			mountpoint: '/mnt/disk1'
		}
	}

	machine := s.tfclient.machines_deploy(MachinesModel{
		name: generate_taiga_machine_name(taiga.name)
		network: Network{
			add_wireguard_access: true
		}
		machines: [
			Machine{
				name: 'taiga_vm'
				farm_id: u32(taiga.farm_id)
				cpu: taiga_cap[taiga.capacity].cpu
				memory: taiga_cap[taiga.capacity].memory
				rootfs_size: taiga_cap[taiga.capacity].size
				flist: 'https://hub.grid.tf/tf-official-apps/grid3_taiga_docker-latest.flist'
				env_vars: {
					'DOMAIN_NAME':         '${taiga.name}.${domain}'
					'ADMIN_USERNAME':      taiga.admin_username
					'ADMIN_PASSWORD':      taiga.admin_password
					'ADMIN_EMAIL':         taiga.admin_email
					'SSH_KEY':             taiga.ssh_key
					'DEFAULT_FROM_EMAIL':  ''
					'EMAIL_USE_TLS':       ''
					'EMAIL_USE_SSL':       ''
					'EMAIL_HOST':          ''
					'EMAIL_PORT':          ''
					'EMAIL_HOST_USER':     ''
					'EMAIL_HOST_PASSWORD': ''
				}
				disks: disks
				planetary: true
			},
		]
	}) or {
		s.tfclient.machines_delete(generate_taiga_machine_name(taiga.name))!
		return error('failed to deploy taiga instance: ${err}')
	}

	gateway := s.tfclient.gateways_deploy_name(GatewayName{
		name: taiga.name
		backends: ['http://${machine.machines[0].ygg_ip}:9000']
		node_id: u32(gateway_node_id)
	}) or {
		s.tfclient.machines_delete(generate_taiga_machine_name(taiga.name))!
		s.tfclient.gateways_delete_name(taiga.name)!
		return error('failed to deploy taiga instance: ${err}')
	}

	return TaigaResult{
		name: taiga.name
		machine_ygg_ip: machine.machines[0].ygg_ip
		gateway_name: gateway.fqdn
	}
}

pub fn (mut s SolutionHandler) delete_taiga(taiga_name string) ! {
	s.tfclient.gateways_delete_name(taiga_name)!
	s.tfclient.machines_delete(generate_taiga_machine_name(taiga_name))!
}

pub fn (mut s SolutionHandler) get_taiga(taiga_name string) !TaigaResult {
	machine := s.tfclient.machines_get(generate_taiga_machine_name(taiga_name))!
	gateway := s.tfclient.gateways_get_name(taiga_name)!

	return TaigaResult{
		name: taiga_name
		machine_ygg_ip: machine.machines[0].ygg_ip
		gateway_name: gateway.fqdn
	}
}

fn generate_taiga_machine_name(taiga_name string) string {
	return '${taiga_name}_taiga_machine'
}
