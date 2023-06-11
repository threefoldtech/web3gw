module tfgrid

pub struct Discourse {
pub:
	name            string
	farm_id         u64
	capacity        string
	disk_size       u32 // in giga bytes
	ssh_key         string
	developer_email string
	smtp_username   string
	smtp_password   string
	smtp_address    string = 'smtp.gmail.com'
	smtp_enable_tls bool   = true
	smtp_port       u32    = 587

	threebot_private_key string
	flask_secret_key     string
}

pub struct DiscourseResult {
pub:
	name           string
	machine_ygg_ip string
	fqdn           string
}
