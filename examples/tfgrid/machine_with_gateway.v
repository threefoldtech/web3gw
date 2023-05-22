module main

import threefoldtech.threebot.tfgrid

pub struct MachineWithGateway {
pub:
	machine tfgrid.Machine
	gateway bool
}

pub struct MachineWithGatewayResult {
pub mut:
	machine tfgrid.MachineResult
	gateway string
}

pub struct MachinesWithGateways {
pub:
	name                 string
	add_wireguard_access bool
	machines             []MachineWithGateway
}

pub struct MachinesWithGatewaysResult {
pub mut:
	name             string
	wireguard_config string
	machines         []MachineWithGatewayResult
}

fn deploy_machines_with_gateways(mut client tfgrid.TFGridClient, machines_with_gateways MachinesWithGateways) !MachinesWithGatewaysResult {
	mut machines_list := []tfgrid.Machine{}
	for machine_with_gateway in machines_with_gateways.machines {
		machines_list << machine_with_gateway.machine
	}

	machines_model := client.machines_deploy(tfgrid.MachinesModel{
		name: machines_with_gateways.name
		network: tfgrid.Network{
			add_wireguard_access: machines_with_gateways.add_wireguard_access
		}
		machines: machines_list
	}) or {
		client.machines_delete(machines_with_gateways.name)!
		return error('failed to deploy machines: ${err}')
	}

	// maps machine name to its ip
	mut machines_map := map[string]tfgrid.MachineResult{}
	for machine in machines_model.machines {
		machines_map[machine.name] = machine
	}

	mut machines_result := MachinesWithGatewaysResult{
		name: machines_with_gateways.name
		wireguard_config: machines_model.network.wireguard_config
	}

	for machine_with_gateway in machines_with_gateways.machines {
		mut machine_res := MachineWithGatewayResult{
			machine: machines_map[machine_with_gateway.machine.name]
		}

		if machine_with_gateway.gateway {
			ip := get_machine_ip(machines_map[machine_with_gateway.machine.name]) or {
				delete_machines_with_gateways(mut client, machines_with_gateways.name)!
				return error('failed to get machine ${machine_with_gateway.machine.name} ip: ${err}')
			}

			gw := client.gateways_deploy_name(tfgrid.GatewayName{
				name: generate_gateway_name(machine_with_gateway.machine.name)
				backends: ['http://${ip}:9000']
			}) or {
				delete_machines_with_gateways(mut client, machines_with_gateways.name)!
				return error('failed to deploy gateways: ${err}')
			}

			machine_res.gateway = gw.fqdn
		}

		machines_result.machines << machine_res
	}

	return machines_result
}

fn delete_machines_with_gateways(mut client tfgrid.TFGridClient, machines_with_gateways_name string) ! {
	// the gateways must be deleted first
	machines_model := client.machines_get(machines_with_gateways_name)!

	for machine in machines_model.machines {
		client.gateways_delete_name(generate_gateway_name(machine.name))!
	}

	client.machines_delete(machines_with_gateways_name)!
}

fn get_machine_ip(machine tfgrid.MachineResult) !string {
	if machine.computed_ip4 != '' {
		return machine.computed_ip4
	}

	if machine.ygg_ip != '' {
		return machine.ygg_ip
	}

	return error('machine ${machine.name} neither has a public ipv4, nor a ygg ip')
}

fn get_machines_with_gateways(mut client tfgrid.TFGridClient, machines_with_gateways_name string) !MachinesWithGatewaysResult {
	machines_model := client.machines_get(machines_with_gateways_name)!

	mut result := MachinesWithGatewaysResult{
		name: machines_model.name
		wireguard_config: machines_model.network.wireguard_config
	}

	for machine in machines_model.machines {
		gw := client.gateways_get_name(generate_gateway_name(machine.name)) or {
			// TODO: error should idicate whether the grid client did not find gateway with this name, or some other error
			continue
		}

		result.machines << MachineWithGatewayResult{
			machine: machine
			gateway: gw.fqdn
		}
	}

	return result
}

fn generate_gateway_name(machine_name string) string {
	return '${machine_name}gw'
}
