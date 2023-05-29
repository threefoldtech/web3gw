module gridprocessor

import rand
import threefoldtech.threebot.tfgrid { TFGridClient }

struct GatewayedVMsCreate {
	model_name string
	farm_id i64
	number_of_machines int
	capacity string
	disk_size int
	ssh_key string
	backend_port int
	public_network bool
	public_ips bool
}

fn (vms GatewayedVMsCreate) execute(mut client TFGridClient) {
	client.gatewayed_vms_create(vms)
}

struct GatewayedVMsRead {
	model_name string
}

fn (vms GatewayedVMsRead) execute(mut client TFGridClient) {
	client.gatewayed_vms_read(vms)
}

struct GatewayedVMsDelete {
	model_name string
}

fn (vms GatewayedVMsDelete) execute(mut client TFGridClient) {
	client.gatewayed_vms_delete(vms)
}

fn build_gatewayed_vms_processes(op GridOp, param_map map[string]string, args_set map[string]bool) !(string, Process) {
	match op {
		.create {
			return gatewayed_vms_create(param_map, args_set)!
		}
		.read {
			return gatewayed_vms_read(param_map, args_set)!
		}
		.delete {
			return gatewayed_vms_delete(param_map, args_set)!
		}
		else {
			return error("Invalid operation")
		}
	}
}

fn gatewayed_vms_create(param_map map[string]string, args_set map[string]bool) !(string, Process) {
	model_name := param_map["model_name"] or { return error("Missing model_name") }
	farm_id := param_map["farm_id"] or { '0' }
	number_of_machines := param_map["number_of_machines"] or { '1' }
	capacity := param_map["capacity"] or { return error("Missing capacity") }
	disk_size := param_map["disk_size"]
	ssh_key := param_map["ssh_key"] or { return error("Missing ssh_key") }
	backend_port := param_map["backend_port"] or { '80' }

	public_network := false
	public_ips := false
	if args_set["public_network"] {
		public_network = true
	}
	if args_set["public_ips"] {
		public_ips = true
	}
	
	model := GatewayedVMsCreate{
		model_name: model_name,
		farm_id: farm_id,
		number_of_machines: number_of_machines,
		capacity: capacity,
		disk_size: disk_size,
		ssh_key: ssh_key,
		backend_port: backend_port,
		public_network: public_network,
		public_ips: public_ips
	}

	return model_name, model
}

fn gatewayed_vms_read(param_map map[string]string, args_set map[string]bool) !(string, Process) {
	model_name := param_map["model_name"] or { return error("Missing model_name") }

	model := GatewayedVMsRead{
		model_name: model_name
	}

	return model_name, model
}

fn gatewayed_vms_delete(param_map map[string]string, args_set map[string]bool) !(string, Process) {
	model_name := param_map["model_name"] or { return error("Missing model_name") }

	model := GatewayedVMsDelete{
		model_name: model_name
	}

	return model_name, model
}
