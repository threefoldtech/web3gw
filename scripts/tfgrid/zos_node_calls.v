module tfgrid

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
	return client.send_json_rpc[[]ZOSNodeRequest, []Workload]('tfgrid.ZOSDeploymentChanges',
		[request], default_timeout)!
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
	return client.send_json_rpc[[]ZOSNodeRequest, Deployment]('tfgrid.ZOSDeploymentGet',
		[request], default_timeout)!
}
