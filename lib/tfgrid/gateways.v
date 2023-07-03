module tfgrid

[params]
pub struct GatewayFQDN {
pub:
	name            string   [required] // identifier for the gateway, must be unique
	node_id         u32      [required] // node to deploy the gateway workload on
	tls_passthrough bool // True to enable TLS encryption
	backends        []string [required] // The backend that the gateway will point to
	fqdn            string   [required] // The fully qualified domain name that points to this gateway
}

[params]
pub struct GatewayName {
pub mut:
	name            string   [json: 'name'; required] // identifier for the gateway, must be unique
	node_id         u32      [json: 'node_id'] // node to deploy the gateway workload on, if 0, a random elibile node will be selected
	tls_passthrough bool     [json: 'tls_passthrough'] // True to enable TLS encryption
	backends        []string [json: 'backends'; required] // The backend that the gateway wwill point to
}

// Deploys a fully qualified domain on gateway (for example site.com) and returns gateway model
// with some extra data related to the created fqdn.
pub fn (mut t TFGridClient) gateways_deploy_fqdn(model GatewayFQDN) !GatewayFQDNResult {
	return t.client.send_json_rpc[[]GatewayFQDN, GatewayFQDNResult]('tfgrid.GatewayFQDNDeploy',
		[model], t.timeout)!
}

// Gets the fqdn info using the name given when created. It returns an object containing the
// fully qualified domain on gateway information.
pub fn (mut t TFGridClient) gateways_get_fqdn(model_name string) !GatewayFQDNResult {
	return t.client.send_json_rpc[[]string, GatewayFQDNResult]('tfgrid.GatewayFQDNGet',
		[model_name], t.timeout)!
}

// Deletes the fully qualified domain on gateway given the name used to create it. An error
// is returned if the attempt failed.
pub fn (mut t TFGridClient) gateways_delete_fqdn(model_name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.GatewayFQDNDelete', [
		model_name,
	], t.timeout)!
}

// Deploys a gateway name given the configuration and returns an object containing the
// gateway name configuration.
pub fn (mut t TFGridClient) gateways_deploy_name(model GatewayName) !GatewayNameResult {
	return t.client.send_json_rpc[[]GatewayName, GatewayNameResult]('tfgrid.GatewayNameDeploy',
		[model], t.timeout)!
}

// Gets the gateway name object given the name used when deploying.
pub fn (mut t TFGridClient) gateways_get_name(model_name string) !GatewayNameResult {
	return t.client.send_json_rpc[[]string, GatewayNameResult]('tfgrid.GatewayNameGet',
		[model_name], t.timeout)!
}

// Deletes the gateway name given the name used when deploying. This call will return an error
// if it fails to do so.
pub fn (mut t TFGridClient) gateways_delete_name(model_name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.GatewayNameDelete', [
		model_name,
	], t.timeout)!
}
