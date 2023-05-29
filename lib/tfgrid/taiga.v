module tfgrid

import threefoldtech.threebot.explorer

pub struct Taiga {
pub:
	name           string
	farm_id        u64
	cpu            u32
	memory         u32 // in mega bytes
	rootfs_size    u32 // in mega bytes
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

pub fn (mut client TFGridClient) deploy_taiga(mut explorer_client explorer.ExplorerClient, taiga Taiga) !TaigaResult {
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

	gateway_nodes := explorer_client.nodes(explorer.NodesRequestParams{
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

	machine := client.machines_deploy(MachinesModel{
		name: generate_taiga_machine_name(taiga.name)
		network: Network{
			add_wireguard_access: true
		}
		machines: [
			Machine{
				name: 'taiga_vm'
				farm_id: u32(taiga.farm_id)
				cpu: taiga.cpu
				memory: taiga.memory
				rootfs_size: taiga.rootfs_size
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
				disks: [
					Disk{
						size: taiga.disk_size
						mountpoint: '/mnt/disk1'
					},
				]
				planetary: true
			},
		]
	}) or {
		client.machines_delete(generate_taiga_machine_name(taiga.name))!
		return error('failed to deploy taiga instance: ${err}')
	}

	gateway := client.gateways_deploy_name(GatewayName{
		name: taiga.name
		backends: ['http://${machine.machines[0].ygg_ip}:9000']
		node_id: u32(gateway_node_id)
	}) or {
		client.machines_delete(generate_taiga_machine_name(taiga.name))!
		client.gateways_delete_name(taiga.name)!
		return error('failed to deploy taiga instance: ${err}')
	}

	return TaigaResult{
		name: taiga.name
		machine_ygg_ip: machine.machines[0].ygg_ip
		gateway_name: gateway.fqdn
	}
}

pub fn (mut client TFGridClient) delete_taiga(taiga_name string) ! {
	client.gateways_delete_name(taiga_name)!
	client.machines_delete(generate_taiga_machine_name(taiga_name))!
}

pub fn (mut client TFGridClient) get_taiga(taiga_name string) !TaigaResult {
	machine := client.machines_get(generate_taiga_machine_name(taiga_name))!
	gateway := client.gateways_get_name(taiga_name)!

	return TaigaResult{
		name: taiga_name
		machine_ygg_ip: machine.machines[0].ygg_ip
		gateway_name: gateway.fqdn
	}
}

fn generate_taiga_machine_name(taiga_name string) string {
	return '${taiga_name}_taiga_machine'
}
