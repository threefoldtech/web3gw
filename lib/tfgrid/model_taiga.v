module tfgrid

pub struct Taiga {
pub:
	name           string
	farm_id        u64
	capacity       string
	disk_size      u32 // in giga bytes
	ssh_key        string
	admin_username string
	admin_password string
	admin_email    string
}

pub struct TaigaResult {
pub:
	name           string
	machine_ygg_ip string
	gateway_name   string
}
