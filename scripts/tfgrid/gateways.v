module tfgrid

const (
	default_timeout = 500000
)

// Deploys a fully qualified domain on gateway (for example site.com) and returns gateway model
// with some extra data related to the created fqdn.
pub fn (mut t TFGridClient) gateways_deploy_fqdn(model GatewayFQDN) !GatewayFQDNResult {
	return t.client.send_json_rpc[[]GatewayFQDN, GatewayFQDNResult]('tfgrid.GatewayFQDNDeploy',
		[model], tfgrid.default_timeout)!
}

// Gets the fqdn info using the name given when created. It returns an object containing the
// fully qualified domain on gateway information.
pub fn (mut t TFGridClient) gateways_get_fqdn(model_name string) !GatewayFQDNResult {
	return t.client.send_json_rpc[[]string, GatewayFQDNResult]('tfgrid.GatewayFQDNGet',
		[model_name], tfgrid.default_timeout)!
}

// Deletes the fully qualified domain on gateway given the name used to create it. An error
// is returned if the attempt failed.
pub fn (mut t TFGridClient) gateways_delete_fqdn(model_name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.GatewayFQDNDelete', [
		model_name,
	], tfgrid.default_timeout)!
}

// Deploys a gateway name given the configuration and returns an object containing the
// gateway name configuration.
pub fn (mut t TFGridClient) gateways_deploy_name(model GatewayName) !GatewayNameResult {
	return t.client.send_json_rpc[[]GatewayName, GatewayNameResult]('tfgrid.GatewayNameDeploy',
		[model], tfgrid.default_timeout)!
}

// Gets the gateway name object given the name used when deploying.
pub fn (mut t TFGridClient) gateways_get_name(model_name string) !GatewayNameResult {
	return t.client.send_json_rpc[[]string, GatewayNameResult]('tfgrid.GatewayNameGet',
		[model_name], tfgrid.default_timeout)!
}

// Deletes the gateway name given the name used when deploying. This call will return an error
// if it fails to do so.
pub fn (mut t TFGridClient) gateways_delete_name(model_name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.GatewayNameDelete', [
		model_name,
	], tfgrid.default_timeout)!
}
