module tfgrid

pub struct Funkwhale {
pub:
	name           string
	farm_id        u64
	capacity       string
	ssh_key        string
	admin_email    string
	admin_username string
	admin_password string
}

pub struct FunkwhaleResult {
pub:
	name           string
	machine_ygg_ip string
	fqdn           string
}
