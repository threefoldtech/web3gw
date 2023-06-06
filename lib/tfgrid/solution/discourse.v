module solution

import threefoldtech.threebot.explorer
import threefoldtech.threebot.tfgrid { AddMachine, Disk, GatewayName, GatewayNameResult, Machine, MachineResult, MachinesModel, MachinesResult, Network, RemoveMachine }
import rand
import time

const discourse_cap = {
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

pub struct Discourse {
pub:
	name            string
	farm_id         u64
	capacity Capacity
	disk_size       u32 // in giga bytes
	ssh_key         string
	developer_email string
	smtp_username   string
	smtp_password   string
	smtp_address    string = 'smtp.gmail.com'
	smtp_enable_tls bool   = true
	smtp_port       u32    = 587

	threebot_private_key string
	flask_secret_key     string
}

pub struct DiscourseResult {
pub:
	name           string
	machine_ygg_ip string
	gateway_name   string
}

pub fn (mut s SolutionHandler) deploy_discourse(discourse Discourse) !DiscourseResult {
	mut filter := explorer.NodeFilter{
		status: 'up'
		dedicated: false
		domain: true
	}

	if discourse.farm_id != 0 {
		filter = explorer.NodeFilter{
			farm_ids: [discourse.farm_id]
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
	smtp_enable_tls := if discourse.smtp_enable_tls { 'true' } else { 'false' }

	mut gw := GatewayName{
		name: rand.string(8)
		backends: []string{}
	}

	mut disks := []Disk{}
	if discourse.disk_size > 0{
		disks << Disk{
			size: discourse.disk_size
			mountpoint: '/var/lib/docker'
		}
	}

	machine := s.tfclient.machines_deploy(MachinesModel{
		name: generate_discourse_machine_name(discourse.name)
		network: Network{
			add_wireguard_access: false
		}
		machines: [
			Machine{
				name: 'discourse_vm'
				farm_id: u32(discourse.farm_id)
				cpu: discourse_cap[discourse.capacity].cpu
				memory: discourse_cap[discourse.capacity].memory
				rootfs_size: discourse_cap[discourse.capacity].size
				flist: 'https://hub.grid.tf/tf-official-apps/forum-docker-v3.1.2.flist'
				disks: disks
				env_vars: {
					'SSH_KEY':                         discourse.ssh_key
					'DISCOURSE_HOSTNAME':              domain
					'DISCOURSE_DEVELOPER_EMAILS':      discourse.developer_email
					'DISCOURSE_SMTP_ADDRESS':          discourse.smtp_address
					'DISCOURSE_SMTP_PORT':             '${discourse.smtp_port}'
					'DISCOURSE_SMTP_ENABLE_START_TLS': smtp_enable_tls
					'DISCOURSE_SMTP_USER_NAME':        discourse.smtp_username
					'DISCOURSE_SMTP_PASSWORD':         discourse.smtp_password
					'THREEBOT_PRIVATE_KEY':            discourse.threebot_private_key
					'FLASK_SECRET_KEY':                discourse.flask_secret_key
					'GW_PROJECT_NAME': gw.name
				}
				planetary: true
			},
		]
	}) or {
		s.tfclient.machines_delete(generate_discourse_machine_name(discourse.name))!
		return error('failed to deploy discourse instance: ${err}')
	}

	gateway := s.tfclient.gateways_deploy_name(GatewayName{
		name: discourse.name
		backends: ['http://${machine.machines[0].ygg_ip}:88']
		node_id: u32(gateway_node_id)
	}) or {
		// if either deployment failed, delete all created contracts
		s.tfclient.machines_delete(generate_discourse_machine_name(discourse.name))!
		s.tfclient.gateways_delete_name(discourse.name)!
		return error('failed to deploy discourse instance: ${err}')
	}

	return DiscourseResult{
		name: discourse.name
		machine_ygg_ip: machine.machines[0].ygg_ip
		gateway_name: gateway.fqdn
	}
}

pub fn (mut s SolutionHandler) delete_discourse(discourse_name string) ! {
	s.tfclient.machines_delete(generate_discourse_machine_name(discourse_name))!
	s.tfclient.gateways_delete_name(discourse_name)!
	// time.sleep(10 * time.second)
}

pub fn (mut s SolutionHandler) get_discourse(discourse_name string) !DiscourseResult {
	machine := s.tfclient.machines_get(generate_discourse_machine_name(discourse_name))!
	gateway := s.tfclient.gateways_get_name(discourse_name)!

	return DiscourseResult{
		name: discourse_name
		machine_ygg_ip: machine.machines[0].ygg_ip
		gateway_name: gateway.fqdn
	}
}

fn generate_discourse_machine_name(discourse_name string) string {
	return '${discourse_name}_discourse_machine'
}