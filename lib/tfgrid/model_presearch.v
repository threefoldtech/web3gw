module tfgrid

pub struct Presearch {
pub:
	name        string
	farm_id     u64
	disk_size   u32 // in giga bytes
	ssh_key     string
	public_ipv4 bool
}

pub struct PresearchResult {
pub:
	name           string
	machine_ygg_ip string
	machine_ipv4   string
}
