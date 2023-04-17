module tfgrid

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

const (
	default_timeout = 500000
)

struct GatewayClient {
	RpcWsClient
}

// Deploy a fully qualified domain on gateway ex: site.com
pub fn (mut client GatewayClient) gateways_deploy_fqdn(model GatewayFQDN) !GatewayFQDNResult {
	return client.send_json_rpc[[]GatewayFQDN, GatewayFQDNResult]('tfgrid.GatewayFQDNDeploy',
		[model], tfgrid.default_timeout)!
}

// Get fqdn info using deployment name.
pub fn (mut client GatewayClient) gateways_get_fqdn(model_name string) !GatewayFQDNResult {
	return client.send_json_rpc[[]string, GatewayFQDNResult]('tfgrid.GatewayFQDNGet',
		[model_name], tfgrid.default_timeout)!
}

// Delete fqdn using deployment name
pub fn (mut client GatewayClient) gateways_delete_fqdn(model_name string) ! {
	_ := client.send_json_rpc[[]string, string]('tfgrid.GatewayFQDNDelete', [
		model_name,
	], tfgrid.default_timeout)!
}

// Deploy name domain on gateway ex: name.gateway.com
pub fn (mut client GatewayClient) gateways_deploy_name(model GatewayName) !GatewayNameResult {
	return client.send_json_rpc[[]GatewayName, GatewayNameResult]('tfgrid.GatewayNameDeploy',
		[model], tfgrid.default_timeout)!
}

// Get fqdn info using deployment name.
pub fn (mut client GatewayClient) gateways_get_name(model_name string) !GatewayNameResult {
	return client.send_json_rpc[[]string, GatewayNameResult]('tfgrid.GatewayNameGet',
		[model_name], tfgrid.default_timeout)!
}

// Delete gateway name using the name of the gateway name
pub fn (mut client GatewayClient) gateways_delete_name(model_name string) ! {
	_ := client.send_json_rpc[[]string, string]('tfgrid.GatewayNameDelete', [
		model_name,
	], tfgrid.default_timeout)!
}
