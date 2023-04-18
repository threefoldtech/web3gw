module tfgrid

const (
	default_timeout = 500000
)

// gateways_deploy_fqdn Deploys a fully qualified domain on gateway ex: site.com
// - model: the gateway model
// returns gateway model with the computed fileds form the grid.
pub fn (mut client TFGridClient) gateways_deploy_fqdn(model GatewayFQDN) !GatewayFQDNResult {
	return client.send_json_rpc[[]GatewayFQDN, GatewayFQDNResult]('tfgrid.GatewayFQDNDeploy',
		[model], tfgrid.default_timeout)!
}

// gateways_get_fqdn Get fqdn info using deployment name.
// - model_name: the name of gateway model
// returns the full gateway model
pub fn (mut client TFGridClient) gateways_get_fqdn(model_name string) !GatewayFQDNResult {
	return client.send_json_rpc[[]string, GatewayFQDNResult]('tfgrid.GatewayFQDNGet',
		[model_name], tfgrid.default_timeout)!
}

// gateways_delete_fqdn Delete fqdn using deployment name
// - model_name: the name of gateway model
pub fn (mut client TFGridClient) gateways_delete_fqdn(model_name string) ! {
	_ := client.send_json_rpc[[]string, string]('tfgrid.GatewayFQDNDelete', [
		model_name,
	], tfgrid.default_timeout)!
}

// gateways_deploy_name Deploys a fully qualified domain on gateway
// - model: the gateway model
// returns gateway model with the computed fileds form the grid.
pub fn (mut client TFGridClient) gateways_deploy_name(model GatewayName) !GatewayNameResult {
	return client.send_json_rpc[[]GatewayName, GatewayNameResult]('tfgrid.GatewayNameDeploy',
		[model], tfgrid.default_timeout)!
}

// gateways_get_name Get fqdn info using deployment name.
// - model_name: the name of gateway model
// returns the full gateway model
pub fn (mut client TFGridClient) gateways_get_name(model_name string) !GatewayNameResult {
	return client.send_json_rpc[[]string, GatewayNameResult]('tfgrid.GatewayNameGet',
		[model_name], tfgrid.default_timeout)!
}

// gateways_delete_fqdn Delete fqdn using deployment name
// - model_name: the name of gateway model
pub fn (mut client TFGridClient) gateways_delete_name(model_name string) ! {
	_ := client.send_json_rpc[[]string, string]('tfgrid.GatewayNameDelete', [
		model_name,
	], tfgrid.default_timeout)!
}
