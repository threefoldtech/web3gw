// ssd mounts under zmachine

module zos

// ONLY possible on SSD
pub struct Zmount {
pub mut:
	size i64 // bytes
}

pub fn (mut mount Zmount) challenge() string {
	return '${mount.size}'
}

pub struct ZmountResult {
pub mut:
	volume_id string
}
