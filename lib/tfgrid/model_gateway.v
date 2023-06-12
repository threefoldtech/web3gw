module tfgrid

pub struct GatewayFQDNResult {
pub:
	name            string
	node_id         u32
	tls_passthrough bool
	backends        []string
	fqdn            string
	// computed
	contract_id u32
}

pub struct GatewayNameResult {
pub:
	name            string   [json: 'name']
	node_id         u32      [json: 'node_id']
	tls_passthrough bool     [json: 'tls_passthrough']
	backends        []string [json: 'backends']
	// computed
	fqdn             string [json: 'fqdn'] // the full domain name
	name_contract_id u32    [json: 'name_contract_id']
	contract_id      u32    [json: 'contract_id']
}
