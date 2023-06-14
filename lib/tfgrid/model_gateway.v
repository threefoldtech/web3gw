module tfgrid

pub struct GatewayFQDNResult {
pub:
	name            string   // name of the instance
	node_id         u32      // node id that the instance was deployed on
	tls_passthrough bool     // whether or not tls was enables
	backends        []string // backends that this gateway is pointing to
	fqdn            string   // fully qualified domain name pointing to this gatewat
	// computed
	contract_id u32 // contract id for the gateway
}

pub struct GatewayNameResult {
pub:
	name            string   // name of the instance
	node_id         u32      // node id that the instance was deployed on
	tls_passthrough bool     // whether or not tls was enabled
	backends        []string // backends that this gateway is pointing to
	// computed
	fqdn             string // the full domain name for this instance
	name_contract_id u32    // name contract id
	contract_id      u32    // contract id for the gateway
}
