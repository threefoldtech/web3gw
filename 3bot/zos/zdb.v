module zos

type ZdbMode = string

pub struct ZdbModes {
pub:
	seq  string = 'seq'
	user string = 'user'
}

pub const zdb_modes = ZdbModes{}

type DeviceType = string

pub struct DeviceTypes {
pub:
	hdd string = 'hdd'
	ssd string = 'ssd'
}

pub const device_types = DeviceTypes{}

pub struct Zdb {
pub mut:
	// size in bytes
	size      u64
	mode      ZdbMode
	password  string
	public    bool
}

pub fn (mut z Zdb) challenge() string {
	mut out := ''
	out += '${z.size}'
	out += '${z.mode}'
	out += z.password
	out += '${z.public}'

	return out
}

pub struct ZdbResult {
pub mut:
	namespace string
	ips       []string
	port      u32
}
