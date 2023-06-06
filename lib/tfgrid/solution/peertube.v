module solution

import threefoldtech.threebot.explorer
import threefoldtech.threebot.tfgrid { GatewayName, Machine, MachinesModel, Network }

const peertube_cap = {
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

pub struct Peertube {
pub:
	name          string
	farm_id       u64
	capacity Capacity
	ssh_key       string
	db_username   string
	db_password   string
	admin_email   string
	smtp_hostname string
	smtp_username string
	smtp_password string
}

pub struct PeertubeResult {
pub:
	name           string
	machine_ygg_ip string
	gateway_name   string
}

pub fn (mut s SolutionHandler) deploy_peertube(peertube Peertube) !PeertubeResult {
	mut filter := explorer.NodeFilter{
		status: 'up'
		dedicated: false
		domain: true
	}

	if peertube.farm_id != 0 {
		filter = explorer.NodeFilter{
			farm_ids: [peertube.farm_id]
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

	machine := s.tfclient.machines_deploy(MachinesModel{
		name: generate_peertube_machine_name(peertube.name)
		network: Network{
			add_wireguard_access: false
		}
		machines: [
			Machine{
				name: 'peertube_vm'
				farm_id: u32(peertube.farm_id)
				cpu: peertube_cap[peertube.capacity].cpu
				memory: peertube_cap[peertube.capacity].memory
				rootfs_size: peertube_cap[peertube.capacity].size
				flist: 'https://hub.grid.tf/tf-official-apps/peertube-v3.1.1.flist'
				env_vars: {
					'SSH_KEY':                     peertube.ssh_key
					'PEERTUBE_DB_SUFFIX':          '_prod'
					'PEERTUBE_DB_USERNAME':        peertube.db_username
					'PEERTUBE_DB_PASSWORD':        peertube.db_password
					'PEERTUBE_ADMIN_EMAIL':        peertube.admin_email
					'PEERTUBE_WEBSERVER_HOSTNAME': '${peertube.name}.${domain}'
					'PEERTUBE_WEBSERVER_PORT':     '443'
					'PEERTUBE_SMTP_HOSTNAME':      peertube.smtp_hostname
					'PEERTUBE_SMTP_USERNAME':      peertube.smtp_username
					'PEERTUBE_SMTP_PASSWORD':      peertube.smtp_password
					'PEERTUBE_BIND_ADDRESS':       '::'
				}
				planetary: true
			},
		]
	}) or {
		s.tfclient.machines_delete(generate_peertube_machine_name(peertube.name))!
		return error('failed to deploy peertube instance: ${err}')
	}

	gateway := s.tfclient.gateways_deploy_name(GatewayName{
		name: peertube.name
		backends: ['http://${machine.machines[0].ygg_ip}:9000']
		node_id: u32(gateway_node_id)
	}) or {
		// if either deployment failed, delete all created contracts
		s.tfclient.machines_delete(generate_peertube_machine_name(peertube.name))!
		s.tfclient.gateways_delete_name(peertube.name)!
		return error('failed to deploy peertube instance: ${err}')
	}

	return PeertubeResult{
		name: peertube.name
		machine_ygg_ip: machine.machines[0].ygg_ip
		gateway_name: gateway.fqdn
	}
}

pub fn (mut s SolutionHandler) delete_peertube(peertube_name string) ! {
	s.tfclient.gateways_delete_name(peertube_name)!
	s.tfclient.machines_delete(generate_peertube_machine_name(peertube_name))!
}

pub fn (mut s SolutionHandler) get_peertube(peertube_name string) !PeertubeResult {
	machine := s.tfclient.machines_get(generate_peertube_machine_name(peertube_name))!
	gateway := s.tfclient.gateways_get_name(peertube_name)!

	return PeertubeResult{
		name: peertube_name
		machine_ygg_ip: machine.machines[0].ygg_ip
		gateway_name: gateway.fqdn
	}
}

fn generate_peertube_machine_name(peertube_name string) string {
	return '${peertube_name}_peertube_machine'
}
