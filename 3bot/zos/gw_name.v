module zos

struct GatewayNameProxy {
	tls_passthrough bool
	backends        []string
	network         ?string
	name            string
}

pub fn (g GatewayNameProxy) challenge() string {
	mut output := ''
	output += g.name
	output += '${g.tls_passthrough}'
	for b in g.backends {
		output += b
	}
	output += g.network or { '' }

	return output
}

// GatewayProxyResult results
pub struct GatewayProxyResult {
pub mut:
	fqdn string
}
