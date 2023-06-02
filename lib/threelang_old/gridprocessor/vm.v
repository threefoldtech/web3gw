module gridprocessor

import threefoldtech.threebot.tfgrid
import threefoldtech.threebot.tfgrid.solution { Capacity, SolutionHandler, VM, get_capacity }
import strconv
import rand
import encoding.utf8

struct VMCreateParams {
	VM
}

struct VMGetParams {
mut:
	network string
}

struct VMRemoveParams {
mut:
	network      string
	machine_name string
}

struct VMDeleteParams {
mut:
	network string
}

fn (vm_create VMCreateParams) execute(mut s SolutionHandler) !string {
	ret := s.create_vm(vm_create.VM)!
	return ret.str()
}

fn (vm_get VMGetParams) execute(mut s SolutionHandler) !string {
	ret := s.get_vm(vm_get.network)!
	return ret.str()
}

fn (vm_remove VMRemoveParams) execute(mut s SolutionHandler) !string {
	ret := s.remove_vm(vm_remove.network, vm_remove.machine_name)!
	return ret.str()
}

fn (vm_delete VMDeleteParams) execute(mut s SolutionHandler) !string {
	s.delete_vm(vm_delete.network)!
	return 'vm ${vm_delete.network} is deleted'
}

fn build_vm_process(op GridOp, param_map map[string]string, args_set map[string]bool) !(string, Process) {
	match op {
		.create {
			return create_vm(param_map, args_set)
		}
		.get {
			return get_vm(param_map, args_set)
		}
		.remove {
			return remove_vm(param_map, args_set)
		}
		.delete {
			return delete_vm(param_map, args_set)
		}
		else {
			return error('invalid operation for vms')
		}
	}
}

fn create_vm(param_map map[string]string, args_set map[string]bool) !(string, Process) {
	mut vm := VMCreateParams{}
	vm.network = param_map['network'] or { utf8.to_lower(rand.string(10)) }

	farm_id_str := param_map['farm_id'] or { '0' }
	vm.farm_id = u32(strconv.parse_uint(farm_id_str, 10, 32)!)

	capacity_str := param_map['capacity'] or { return error('vm capacity must be specified') }
	vm.capacity = get_capacity(capacity_str)!

	times_str := param_map['times'] or { '1' }
	vm.times = u32(strconv.parse_uint(times_str, 10, 32)!)

	disk_size_str := param_map['disk_size'] or { '0' }
	vm.disk_size = u32(strconv.parse_uint(disk_size_str, 10, 32)!)

	vm.ssh_key = param_map['ssh_key'] or { return error('ssh key must be provided') }

	vm.gateway = args_set['gateway']
	vm.add_wireguard_access = args_set['add_wireguard_access']
	vm.add_public_ips = args_set['add_public_ips']

	return vm.network, vm
}

fn get_vm(param_map map[string]string, args_set map[string]bool) !(string, Process) {
	mut vm_get := VMGetParams{}
	vm_get.network = param_map['network'] or { return error('vm network name is missing') }

	return vm_get.network, vm_get
}

fn delete_vm(param_map map[string]string, args_set map[string]bool) !(string, Process) {
	mut vm_delete := VMDeleteParams{}
	vm_delete.network = param_map['network'] or { return error('vm network name is missing') }

	return vm_delete.network, vm_delete
}

fn remove_vm(param_map map[string]string, args_set map[string]bool) !(string, Process) {
	mut vm_remove := VMRemoveParams{}
	vm_remove.network = param_map['network'] or { return error('vm network name is missing') }

	return vm_remove.network, vm_remove
}
