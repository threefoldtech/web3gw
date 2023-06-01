module solution

import threefoldtech.threebot.tfgrid { AddMachine, Disk, GatewayName, GatewayNameResult, Machine, MachineResult, MachinesModel, MachinesResult, Network, RemoveMachine }
import rand
import encoding.utf8

pub struct VM {
pub mut:
	network              string
	farm_id              u32
	capacity             Capacity
	times                u32 = 1
	disk_size            u32
	ssh_key              string
	gateway              bool
	add_wireguard_access bool
	add_public_ips           bool
}

struct VMResult {
pub mut:
	network          string
	wireguard_config string
	vms              []GatewayedMachines
}

struct GatewayedMachines {
pub:
	machine MachineResult
	gateway ?GatewayNameResult
}

struct CapacityPackage {
	cpu    u32
	memory u32
	size   u32
}

const (
	gateway_project_name_env_var = 'WEB3PROXY_DOMAIN_PROJECT_NAME'

	cap                          = {
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
)

// create create or updates a network of vms
pub fn (mut s SolutionHandler) create_vm(vm VM) !VMResult {
	// create should first check if this is a create or an update operation
	_ := s.tfclient.machines_get(vm.network) or {
		if err.msg().contains('found 0 contracts for model') {
			// this is a new network, do a create op
			return s.create_new_vm(vm)
		}

		return error('${err}')
	}

	return s.add_vm(vm)
}

fn (mut s SolutionHandler) create_new_vm(vm VM) !VMResult {
	mut machines_model := MachinesModel{
		name: vm.network
		network: Network{
			add_wireguard_access: vm.add_wireguard_access
		}
		machines: []Machine{}
	}

	mut gws := map[string]GatewayName{}
	for _ in 0 .. vm.times {
		mut m := Machine{
			name: utf8.to_lower(rand.string(8))
			farm_id: vm.farm_id
			public_ip: vm.add_public_ips
			cpu: solution.cap[vm.capacity].cpu
			memory: solution.cap[vm.capacity].memory
			rootfs_size: solution.cap[vm.capacity].size
			env_vars: {
				'SSH_KEY': vm.ssh_key
			}
		}

		if vm.disk_size != 0 {
			m.disks << Disk{
				size: vm.disk_size
				mountpoint: '/mnt/disk'
			}
		}

		if vm.gateway {
			gw_project_name := utf8.to_lower(rand.string(8))
			m.env_vars[solution.gateway_project_name_env_var] = gw_project_name
			gws[m.name] = GatewayName{
				name: gw_project_name
				backends: []string{}
			}
		}

		machines_model.machines << m
	}

	machines_res := s.tfclient.machines_deploy(machines_model)!
	mut gws_res := map[string]GatewayNameResult{}
	for m in machines_res.machines {
		mut gw := gws[m.name] or { continue }
		gw.backends << 'http://[${m.ygg_ip}]:9000'
		gws_res[m.name] = s.tfclient.gateways_deploy_name(gw)!
	}

	// TODO: build solution result
	return new_vm_result(machines_res, gws_res)
}

// get retrieves the vms solution
pub fn (mut s SolutionHandler) get_vm(network_name string) !VMResult {
	// get the machines model
	model := s.tfclient.machines_get(network_name)!
	mut gws := map[string]GatewayNameResult{}
	// get each gateway
	for machine in model.machines {
		gw_project_name := machine.env_vars['WEB3PROXY_DOMAIN_PROJECT_NAME'] or { continue }
		gw := s.tfclient.gateways_get_name(gw_project_name)!
		gws[machine.name] = gw
	}

	return new_vm_result(model, gws)
}

fn new_vm_result(model MachinesResult, gws map[string]GatewayNameResult) VMResult {
	mut res := VMResult{
		network: model.name
		wireguard_config: model.network.wireguard_config
		vms: []GatewayedMachines{}
	}

	for m in model.machines {
		res.vms << GatewayedMachines{
			machine: m
			gateway: get_gw(gws, m.name)
		}
	}

	return res
}

fn get_gw(gws map[string]GatewayNameResult, name string) ?GatewayNameResult {
	return gws[name] or { return none }
}

// delete deletes the vms solution
pub fn (mut s SolutionHandler) delete_vm(network_name string) ! {
	// get machines model
	model := s.tfclient.machines_get(network_name)!

	// get & delete gateways
	for machine in model.machines {
		gw_project_name := machine.env_vars['WEB3PROXY_DOMAIN_PROJECT_NAME'] or { continue }
		s.tfclient.gateways_delete_name(gw_project_name)!
	}

	// delete machines model
	s.tfclient.machines_delete(model.name)!
}

// remove removes a machine from a network
pub fn (mut s SolutionHandler) remove_vm(network_name string, vm_name string) !VMResult {
	model := s.tfclient.machines_get(network_name)!
	// get each gateway
	for machine in model.machines {
		if machine.name == vm_name {
			gw_project_name := machine.env_vars[solution.gateway_project_name_env_var]
			if gw_project_name != '' {
				s.tfclient.gateways_delete_name(gw_project_name)!
			}

			_ = s.tfclient.machines_remove(RemoveMachine{
				machine_name: vm_name
				model_name: network_name
			})!
		}
	}

	return s.get_vm(network_name)
}

// add adds a vm to a network
fn (mut s SolutionHandler) add_vm(vm VM) !VMResult {
	mut machines_res := MachinesResult{}
	mut gws := map[string]GatewayName{}

	for _ in 0 .. vm.times {
		mut m := Machine{
			name: utf8.to_lower(rand.string(8))
			farm_id: vm.farm_id
			public_ip: vm.add_public_ips
			cpu: solution.cap[vm.capacity].cpu
			memory: solution.cap[vm.capacity].memory
			rootfs_size: solution.cap[vm.capacity].size
			env_vars: {
				'SSH_KEY': vm.ssh_key
			}
		}

		if vm.disk_size != 0 {
			m.disks << Disk{
				size: vm.disk_size
				mountpoint: '/mnt/disk'
			}
		}

		if vm.gateway {
			gw_project_name := utf8.to_lower(rand.string(8))
			m.env_vars[solution.gateway_project_name_env_var] = gw_project_name
			gws[m.name] = GatewayName{
				name: gw_project_name
				backends: []string{}
			}
		}

		machines_res = s.tfclient.machines_add(AddMachine{
			model_name: vm.network
			machine: m
		})!
	}

	for m in machines_res.machines {
		mut gw := gws[m.name] or { continue }
		gw.backends << 'http://[${m.ygg_ip}]:9000'
		s.tfclient.gateways_deploy_name(gw)!
	}

	return s.get_vm(vm.network)
}
