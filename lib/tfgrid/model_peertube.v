module tfgrid

pub struct Peertube {
pub:
	name          string
	farm_id       u64
	capacity      string
	ssh_key       string
	db_username   string
	db_password   string
	admin_email   string
	smtp_hostname string
	smtp_username string
	smtp_password string
}

pub struct PeertubeResult {
pub:
	name           string
	machine_ygg_ip string
	fqdn           string
}
