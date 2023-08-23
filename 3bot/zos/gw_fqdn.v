module zos

struct GatewayFQDNProxy {
	tls_passthrough bool
	backends        []string
	network         ?string
	fqdn            string
}

pub fn (g GatewayFQDNProxy) challenge() string {
	mut output := ''
	output += g.fqdn
	output += '${g.tls_passthrough}'
	for b in g.backends {
		output += b
	}
	output += g.network or { '' }

	return output
}
