module solution

import threefoldtech.threebot.explorer
import threefoldtech.threebot.tfgrid { AddMachine, Disk, GatewayName, GatewayNameResult, Machine, MachineResult, MachinesModel, MachinesResult, Network, RemoveMachine }

pub struct Funkwhale {
pub:
	name           string
	farm_id        u64
	cpu            u32
	memory         u32 // in mega bytes
	rootfs_size    u32 // in mega bytes
	admin_email    string
	admin_username string
	admin_password string
}

pub struct FunkwhaleResult {
pub:
	name           string
	machine_ygg_ip string
	gateway_name   string
}

pub fn (mut s SolutionHandler) deploy_funkwhale(mut explorer_client explorer.ExplorerClient, funkwhale Funkwhale) !FunkwhaleResult {
	mut filter := explorer.NodeFilter{
		status: 'up'
		dedicated: false
		domain: true
	}

	if funkwhale.farm_id != 0 {
		filter = explorer.NodeFilter{
			farm_ids: [funkwhale.farm_id]
		}
	}

	gateway_nodes := explorer_s.tfclient.nodes(explorer.NodesRequestParams{
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
		name: generate_funkwhale_machine_name(funkwhale.name)
		network: Network{
			add_wireguard_access: false
		}
		machines: [
			Machine{
				name: 'funkwhale_vm'
				farm_id: u32(funkwhale.farm_id)
				cpu: funkwhale.cpu
				memory: funkwhale.memory
				rootfs_size: funkwhale.rootfs_size
				flist: 'https://hub.grid.tf/tf-official-apps/funkwhale-dec21.flist'
				entrypoint: '/init.sh'
				env_vars: {
					'FUNKWHALE_HOSTNAME':        domain
					'DJANGO_SUPERUSER_EMAIL':    funkwhale.admin_email
					'DJANGO_SUPERUSER_USERNAME': funkwhale.admin_username
					'DJANGO_SUPERUSER_PASSWORD': funkwhale.admin_password
				}
				planetary: true
			},
		]
	}) or {
		s.tfclient.machines_delete(generate_funkwhale_machine_name(funkwhale.name))!
		return error('failed to deploy funkwhale instance: ${err}')
	}

	gateway := s.tfclient.gateways_deploy_name(GatewayName{
		name: funkwhale.name
		backends: ['http://${machine.machines[0].ygg_ip}:80']
		node_id: u32(gateway_node_id)
	}) or {
		// if either deployment failed, delete all created contracts
		s.tfclient.machines_delete(generate_funkwhale_machine_name(funkwhale.name))!
		s.tfclient.gateways_delete_name(funkwhale.name)!
		return error('failed to deploy funkwhale instance: ${err}')
	}

	return FunkwhaleResult{
		name: funkwhale.name
		machine_ygg_ip: machine.machines[0].ygg_ip
		gateway_name: gateway.fqdn
	}
}

pub fn (mut s SolutionHandler) delete_funkwhale(funkwhale_name string) ! {
	s.tfclient.gateways_delete_name(funkwhale_name)!
	s.tfclient.machines_delete(generate_funkwhale_machine_name(funkwhale_name))!
}

pub fn (mut s SolutionHandler) get_funkwhale(funkwhale_name string) !FunkwhaleResult {
	machine := s.tfclient.machines_get(generate_funkwhale_machine_name(funkwhale_name))!
	gateway := s.tfclient.gateways_get_name(funkwhale_name)!

	return FunkwhaleResult{
		name: funkwhale_name
		machine_ygg_ip: machine.machines[0].ygg_ip
		gateway_name: gateway.fqdn
	}
}

fn generate_funkwhale_machine_name(funkwhale_name string) string {
	return '${funkwhale_name}_funkwhale_machine'
}
