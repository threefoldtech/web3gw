module tfgrid

import json

pub fn (mut client TFGridClient) zos_deployment_deploy(request ZOSNodeRequest) ! {
	client.send_json_rpc[[]ZOSNodeRequest, string]('tfgrid.ZOSDeploymentDeploy', [
		request,
	], default_timeout)!
}

pub fn (mut client TFGridClient) zos_system_version(request ZOSNodeRequest) !SystemVersion {
	return client.send_json_rpc[[]ZOSNodeRequest, SystemVersion]('tfgrid.ZOSSystemVersion',
		[request], default_timeout)!
}

pub fn (mut client TFGridClient) zos_system_hypervisor(request ZOSNodeRequest) !string {
	return client.send_json_rpc[[]ZOSNodeRequest, string]('tfgrid.ZOSSystemHypervisor',
		[request], default_timeout)!
}

pub fn (mut client TFGridClient) zos_system_dmi(request ZOSNodeRequest) !DMI {
	return client.send_json_rpc[[]ZOSNodeRequest, DMI]('tfgrid.ZOSSystemDMI', [
		request,
	], default_timeout)!
}

pub fn (mut client TFGridClient) zos_network_public_config(request ZOSNodeRequest) !PublicConfig {
	return client.send_json_rpc[[]ZOSNodeRequest, PublicConfig]('tfgrid.ZOSNetworkPublicConfigGet',
		[request], default_timeout)!
}

pub fn (mut client TFGridClient) zos_network_interfaces(request ZOSNodeRequest) !map[string][]string {
	return client.send_json_rpc[[]ZOSNodeRequest, map[string][]string]('tfgrid.ZOSNetworkInterfaces',
		[request], default_timeout)!
}

pub fn (mut client TFGridClient) zos_network_list_wg_ports(request ZOSNodeRequest) ![]u16 {
	return client.send_json_rpc[[]ZOSNodeRequest, []u16]('tfgrid.ZOSNetworkListWGPorts',
		[request], default_timeout)!
}

pub fn (mut client TFGridClient) zos_node_statistics(request ZOSNodeRequest) !Statistics {
	return client.send_json_rpc[[]ZOSNodeRequest, Statistics]('tfgrid.ZOSStatisticsGet',
		[request], default_timeout)!
}

pub fn (mut client TFGridClient) zos_deployment_changes(request ZOSNodeRequest) ![]Workload {
	wls := client.send_json_rpc[[]ZOSNodeRequest, []WorkloadRaw]('tfgrid.ZOSDeploymentChanges',
		[request], default_timeout)!
	return decode_workloads(wls)!
}

pub fn (mut client TFGridClient) zos_deployment_update(request ZOSNodeRequest) ! {
	client.send_json_rpc[[]ZOSNodeRequest, string]('tfgrid.ZOSDeploymentUpdate', [
		request,
	], default_timeout)!
}

// this is disabled in zos, user should instead cancel their contract
pub fn (mut client TFGridClient) zos_deployment_delete(request ZOSNodeRequest) ! {
	client.send_json_rpc[[]ZOSNodeRequest, string]('tfgrid.ZOSDeploymentDelete', [
		request,
	], default_timeout)!
}

pub fn (mut client TFGridClient) zos_deployment_get(request ZOSNodeRequest) !Deployment {
	deployment := client.send_json_rpc[[]ZOSNodeRequest, DeploymentRaw]('tfgrid.ZOSDeploymentGet',
		[
		request,
	], default_timeout)!
	return decode_deployment(deployment)!
}

fn decode_deployment(dl DeploymentRaw) !Deployment {
	wls := decode_workloads(dl.workloads)!
	return Deployment{
		version: dl.version
		twin_id: dl.twin_id
		contract_id: dl.contract_id
		metadata: dl.metadata
		description: dl.description
		expiration: dl.expiration
		signature_requirement: dl.signature_requirement
		workloads: wls
	}
}

fn decode_workloads(raw_workloads []WorkloadRaw) ![]Workload {
	mut wls := []Workload{}
	for wl in raw_workloads {
		wls << Workload{
			version: wl.version
			name: wl.name
			workload_type: wl.workload_type
			data: decode_workload_data(wl.workload_type, wl.data)!
			metadata: wl.metadata
			description: wl.description
			result: decode_result(wl.workload_type, wl.result)!
		}
	}

	return wls
}

fn decode_workload_data(workload_type string, data string) !WorkloadData {
	return match workload_type {
		zdb_workload_type { json.decode(ZDBWorkload, data)! }
		zmachine_workload_type { json.decode(ZMachine, data)! }
		zmount_workload_type { json.decode(ZMount, data)! }
		network_workload_type { json.decode(NetworkWorkload, data)! }
		zlogs_workload_type { json.decode(Zlogs, data)! }
		public_ip_worklaod_type { json.decode(PublicIP, data)! }
		gateway_name_proxy_workload_type { json.decode(GatewayNameProxyWorkload, data)! }
		gateway_fqdn_proxy_workload_type { json.decode(GatewayFQDNProxyWorkload, data)! }
		else { error('workload type ${workload_type} is not supported') }
	}
}

fn decode_result(workload_type string, result ResultRaw) !Result {
	mut result_data := ResultData{}
	result_data = match workload_type {
		zdb_workload_type { ResultData(json.decode(ZDBResultData, result.data)!) }
		zmachine_workload_type { json.decode(ZMachineResult, result.data)! }
		zmount_workload_type { '' }
		network_workload_type { '' }
		zlogs_workload_type { '' }
		public_ip_worklaod_type { json.decode(PublicIPResult, result.data)! }
		gateway_name_proxy_workload_type { json.decode(GatewayNameProxyResult, result.data)! }
		gateway_fqdn_proxy_workload_type { '' }
		else { return error('workload type ${workload_type} is not supported') }
	}

	return Result{
		created: result.created
		state: result.state
		message: result.message
		data: result_data
	}
}
