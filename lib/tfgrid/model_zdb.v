module tfgrid

[params]
pub struct ZDB {
pub:
	node_id     u32 // node id to deploy the ZDB on, if 0, a random eligible node will be selected
	name        string [required] // zdb name. must be unique
	password    string [required] // zdb password
	public      bool // Makes it read-only if password is set, writable if no password set.
	size        u32    [required] // size of the zdb in GB
	description string // zdb description
	mode        string = 'user' // Mode of the ZDB, `user` or `seq`. `user` is the default mode where a user can SET their own keys, like any key-value store. All keys are kept in memory. in `seq` mode, keys are sequential and autoincremented.
}

pub struct ZDBResult {
pub:
	node_id     u32    // node id of the ZDB
	name        string // zdb name
	password    string // zdb password
	public      bool   // Makes it read-only if password is set, writable if no password set.
	size        u32    // size of the zdb in GB
	description string // zdb description
	mode        string // Mode of the ZDB, `user` or `seq`. `user` is the default mode where a user can SET their own keys, like any key-value store. All keys are kept in memory. in `seq` mode, keys are sequential and autoincremented.
	// computed
	namespace string   // namespace of the zdb
	port      u32      // port of the zdb
	ips       []string // Computed IPs of the ZDB. Two IPs are returned: a public IPv6, and a YggIP, in this order
}
