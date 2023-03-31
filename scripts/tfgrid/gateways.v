module tfgrid

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

const (
	default_timeout = 500000
)

// Deploy a fully qualified domain on gateway ex: site.com
pub fn gateways_deploy_fqdn(mut client RpcWsClient, params GatewayFQDN) !GatewayFQDNResult {
	return client.send_json_rpc[[]GatewayFQDN, GatewayFQDNResult]('tfgrid.GatewayFQDNDeploy', [params], default_timeout)!
}

// Get fqdn info using deployment name.
pub fn gateways_get_fqdn(mut client RpcWsClient, params string) !GatewayFQDNResult {
	return client.send_json_rpc[[]string, GatewayFQDNResult]('tfgrid.GatewayFQDN.Get', [params], default_timeout)!
}

// Deploy name domain on gateway ex: name.gateway.com
pub fn gateways_deploy_name(mut client RpcWsClient, params GatewayName) !GatewayNameResult {
	return client.send_json_rpc[[]GatewayName, GatewayNameResult]('tfgrid.GatewayNameDeploy', [params], default_timeout)!
}

pub fn gateways_get_params(mut client RpcWsClient, params string) !GatewayNameResult {
	return client.send_json_rpc[[]string, GatewayNameResult]('tfgrid.GatewayNameGet', [params], default_timeout)!
}

// Delete fqdn using deployment name
pub fn gateways_delete_fqdn(mut client RpcWsClient, params string) ! {
	_ := client.send_json_rpc[[]string, string]('tfgrid.GatewayFQDNDelete', [params], default_timeout)!
}

// Delete gateway name using the name of the gateway name
pub fn gateways_delete_name(mut client RpcWsClient, params string) ! {
	_ := client.send_json_rpc[[]string, string]('tfgrid.GatewayNameDelete', [params], default_timeout)!
}
