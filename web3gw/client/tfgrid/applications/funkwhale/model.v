module funkwhale

pub struct FunkwhaleResult {
pub:
	name   string // identifier for the instance
	ygg_ip string // instance ygg ip
	ipv6   string // instance ipv6, if any
	fqdn   string // fully qualified domain name pointing to the instance
}
