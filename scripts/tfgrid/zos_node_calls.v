module tfgrid

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

pub fn zos_deployment_deploy(mut client RpcWsClient, request ZOSNodeRequest) ! {
	client.send_json_rpc[[]ZOSNodeRequest, string]('tfgrid.ZOSDeploymentDeploy', [
		request,
	], default_timeout)!
}

pub fn zos_system_version(mut client RpcWsClient, request ZOSNodeRequest) !SystemVersion {
	return client.send_json_rpc[[]ZOSNodeRequest, SystemVersion]('tfgrid.ZOSSystemVersion',
		[request], default_timeout)!
}

pub fn zos_system_hypervisor(mut client RpcWsClient, request ZOSNodeRequest) !string {
	return client.send_json_rpc[[]ZOSNodeRequest, string]('tfgrid.ZOSSystemHypervisor',
		[request], default_timeout)!
}

pub fn zos_system_dmi(mut client RpcWsClient, request ZOSNodeRequest) !DMI {
	return client.send_json_rpc[[]ZOSNodeRequest, DMI]('tfgrid.ZOSSystemDMI', [
		request,
	], default_timeout)!
}

pub fn zos_network_public_config(mut client RpcWsClient, request ZOSNodeRequest) !PublicConfig {
	return client.send_json_rpc[[]ZOSNodeRequest, PublicConfig]('tfgrid.ZOSNetworkPublicConfigGet',
		[request], default_timeout)!
}

pub fn zos_network_interfaces(mut client RpcWsClient, request ZOSNodeRequest) !map[string][]string {
	return client.send_json_rpc[[]ZOSNodeRequest, map[string][]string]('tfgrid.ZOSNetworkInterfaces',
		[request], default_timeout)!
}

pub fn zos_network_list_wg_ports(mut client RpcWsClient, request ZOSNodeRequest) ![]u16 {
	return client.send_json_rpc[[]ZOSNodeRequest, []u16]('tfgrid.ZOSNetworkListWGPorts',
		[request], default_timeout)!
}

pub fn zos_node_statistics(mut client RpcWsClient, request ZOSNodeRequest) !Statistics {
	return client.send_json_rpc[[]ZOSNodeRequest, Statistics]('tfgrid.ZOSStatisticsGet',
		[request], default_timeout)!
}

pub fn zos_deployment_changes(mut client RpcWsClient, request ZOSNodeRequest) ![]Workload {
	return client.send_json_rpc[[]ZOSNodeRequest, []Workload]('tfgrid.ZOSDeploymentChanges',
		[request], default_timeout)!
}

pub fn zos_deployment_update(mut client RpcWsClient, request ZOSNodeRequest) ! {
	client.send_json_rpc[[]ZOSNodeRequest, string]('tfgrid.ZOSDeploymentUpdate', [
		request,
	], default_timeout)!
}

// this is disabled in zos, user should instead cancel their contract
pub fn zos_deployment_delete(mut client RpcWsClient, request ZOSNodeRequest) ! {
	client.send_json_rpc[[]ZOSNodeRequest, string]('tfgrid.ZOSDeploymentDelete', [
		request,
	], default_timeout)!
}

pub fn zos_deployment_get(mut client RpcWsClient, request ZOSNodeRequest) !Deployment {
	return client.send_json_rpc[[]ZOSNodeRequest, Deployment]('tfgrid.ZOSDeploymentGet',
		[request], default_timeout)!
}
