module main

import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.explorer

pub struct Peertube {
pub:
	name          string
	farm_id       u64
	cpu           u32
	memory        u32 // in mega bytes
	rootfs_size   u32 // in mega bytes
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

fn deploy_peertube(mut client tfgrid.TFGridClient, mut explorer_client explorer.ExplorerClient, peertube Peertube) !PeertubeResult {
	gateway_nodes := explorer_client.nodes(explorer.NodesRequestParams{
		filters: explorer.NodeFilter{
			status: 'up'
			dedicated: false
			farm_ids: [peertube.farm_id]
			domain: true
		}
		pagination: explorer.Limit{
			size: 1
		}
	})!

	if gateway_nodes.nodes.len == 0 {
		return error('failed to find an eligible node for gateway')
	}

	gateway_node_id := gateway_nodes.nodes[0].node_id
	domain := gateway_nodes.nodes[0].public_config.domain

	machine := client.machines_deploy(tfgrid.MachinesModel{
		name: generate_peertube_machine_name(peertube.name)
		network: tfgrid.Network{
			add_wireguard_access: false
		}
		machines: [
			tfgrid.Machine{
				name: 'peertube_vm'
				farm_id: u32(peertube.farm_id)
				cpu: peertube.cpu
				memory: peertube.memory
				rootfs_size: peertube.rootfs_size
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
		client.machines_delete(generate_peertube_machine_name(peertube.name))!
		return error('failed to deploy peertube instance: ${err}')
	}

	gateway := client.gateways_deploy_name(tfgrid.GatewayName{
		name: peertube.name
		backends: ['http://${machine.machines[0].ygg_ip}:9000']
		node_id: u32(gateway_node_id)
	}) or {
		// if either deployment failed, delete all created contracts
		client.machines_delete(generate_peertube_machine_name(peertube.name))!
		client.gateways_delete_name(peertube.name)!
		return error('failed to deploy peertube instance: ${err}')
	}

	return PeertubeResult{
		name: peertube.name
		machine_ygg_ip: machine.machines[0].ygg_ip
		gateway_name: gateway.fqdn
	}
}

fn delete_peertube(mut client tfgrid.TFGridClient, peertube_name string) ! {
	client.gateways_delete_name(peertube_name)!
	client.machines_delete(generate_peertube_machine_name(peertube_name))!
}

fn get_peertube(mut client tfgrid.TFGridClient, peertube_name string) !PeertubeResult {
	machine := client.machines_get(generate_peertube_machine_name(peertube_name))!
	gateway := client.gateways_get_name(peertube_name)!

	return PeertubeResult{
		name: peertube_name
		machine_ygg_ip: machine.machines[0].ygg_ip
		gateway_name: gateway.fqdn
	}
}

fn generate_peertube_machine_name(peertube_name string) string {
	return '${peertube_name}_machine'
}