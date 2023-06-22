module tfgrid

pub struct Machine {
pub mut:
	name        string            [required] // machine name
	node_id     u32    // node id to deploy on, if 0, a random eligible node will be selected
	farm_id     u32    // farm id to deploy on, if 0, a random eligible farm will be selected
	flist       string = 'https://hub.grid.tf/tf-official-apps/base:latest.flist' // flist of the machine
	entrypoint  string = '/sbin/zinit init' // entry point for the machine
	public_ip   bool   // if true, a public ipv4 will be added to the node
	public_ip6  bool   // if true, a public ipv6 will be added to the node
	planetary   bool = true // if true, a yggdrasil ip will be added to the node
	cpu         u32  = 1 // number of vcpu cores
	memory      u64  = 1024 // memory of the machine in MBs
	rootfs_size u64    // root file system size in MBs
	zlogs       []Zlog // zlogs configs
	disks       []Disk // disks configs
	qsfss       []QSFS // qsfss configs
	env_vars    map[string]string // env vars to attach to the machine
	description string // machine description
}

pub struct Disk {
pub:
	size        u32    [required] // disk size in GBs
	mountpoint  string [required] // mountpoint of the disk on the machine
	description string // disk description
}

pub struct QSFS {
pub:
	mountpoint            string   [required] // mountpoint of the qsfs on the machine
	encryption_key        string   [required] // 64 long hex encoded encryption key (e.g. 0000000000000000000000000000000000000000000000000000000000000000).
	cache                 u32      [required] // The size of the fuse mountpoint on the node in MBs (holds qsfs local data before pushing).
	minimal_shards        u32      [required] // The minimum amount of shards which are needed to recover the original data.
	expected_shards       u32      [required] // The amount of shards which are generated when the data is encoded. Essentially, this is the amount of shards which is needed to be able to recover the data, and some disposable shards which could be lost. The amount of disposable shards can be calculated as expected_shards - minimal_shards.
	redundant_groups      u32      [required] // The amount of groups which one should be able to loose while still being able to recover the original data.
	redundant_nodes       u32      [required] // The amount of nodes that can be lost in every group while still being able to recover the original data.
	encryption_algorithm  string = 'AES' // configuration to use for the encryption stage. Currently only AES is supported.
	compression_algorithm string = 'snappy' // configuration to use for the compression stage. Currently only snappy is supported.
	metadata              Metadata [required] // metadata configs
	description           string // qsfs description

	max_zdb_data_dir_size u32     [required] // Maximum size of the data dir in MiB, if this is set and the sum of the file sizes in the data dir gets higher than this value, the least used, already encoded file will be removed.
	groups                []Group [required] // groups configs
}

pub struct Zlog {
pub:
	output string // Url of the remote location receiving logs. URLs should use one of `redis, ws, wss` schema. e.g. wss://example_ip.com:9000
}

pub struct MachinesResult {
pub:
	name        string          // name of the machines model
	metadata    string          // metadata for the model
	description string          // model description
	network     NetworkResult   // network configs
	machines    []MachineResult // machines configs
}

pub struct MachineResult {
pub:
	name        string            // machine name
	node_id     u32               // node id that this machine was deployed on
	farm_id     u32               // farm id that this machine was deployed on
	flist       string            // flist used by this machine
	entrypoint  string            // entry point of the machine
	public_ip   bool              // whether or not a public ipv4 was requested on this machine
	public_ip6  bool              // whether or not a public ipv6 was requested on this machine
	planetary   bool              // whether or not a yggdrasil ip was requested on this machine
	cpu         u32               // number of vCPUs of this machine
	memory      u64               // machine's memory in MBs
	rootfs_size u64               // machine's root file system size
	zlogs       []Zlog            // zlogs configs
	disks       []DiskResult      // disks configs
	qsfss       []QSFSResult      // qsfs configs
	env_vars    map[string]string // env vars attached to this machine
	description string // machine description
	// computed
	computed_ip4 string // public ipv4 attached to this machine, if any
	computed_ip6 string // public ipv6 attached to this machine, if any
	wireguard_ip string // private wireguard ip of this machine
	ygg_ip       string // yggdrasil ip attached to this machine, if any
}

pub struct QSFSResult {
pub:
	mountpoint       string // mountpoint of the qsfs on the machine
	encryption_key   string // 64 long hex encoded encryption key
	cache            u32    // The size of the fuse mountpoint on the node in MBs (holds qsfs local data before pushing).
	minimal_shards   u32    // The minimum amount of shards which are needed to recover the original data.
	expected_shards  u32    // The amount of shards which are generated when the data is encoded. Essentially, this is the amount of shards which is needed to be able to recover the data, and some disposable shards which could be lost. The amount of disposable shards can be calculated as expected_shards - minimal_shards.
	redundant_groups u32    // The amount of groups which one should be able to loose while still being able to recover the original data.
	redundant_nodes  u32    // The amount of nodes that can be lost in every group while still being able to recover the original data.

	encryption_algorithm  string   // configuration to use for the encryption stage.
	compression_algorithm string   // configuration to use for the compression stage.
	metadata              Metadata // metadata configs
	description           string   // qsfs description

	max_zdb_data_dir_size u32     // Maximum size of the data dir in MiB, if this is set and the sum of the file sizes in the data dir gets higher than this value, the least used, already encoded file will be removed.
	groups                []Group // groups configs
	// computed
	name             string // qsfs name
	metrics_endpoint string // metrics endpoint for the qsfs
}

[params]
pub struct DiskResult {
pub:
	size        u32    [required] // disk size in GBs
	mountpoint  string [required] // mount point of the disk on the machine
	description string // disk description
	// computed
	name string [required] // disk name
}

pub struct Metadata {
	type_                string    [json: 'type'] = 'zdb' // configuration for the metadata store to use, currently only ZDB is supported.
	prefix               string    [required] // Data stored on the remote metadata is prefixed with.
	encryption_algorithm string = 'AES' // configuration to use for the encryption stage. Currently only AES is supported.
	encryption_key       string    [required] // 64 long hex encoded encryption key (e.g. 0000000000000000000000000000000000000000000000000000000000000000).
	backends             []Backend // backends configs
}

pub struct Group {
	backends []Backend
}

pub struct Backend {
	address   string [required] // Address of backend ZDB (e.g. [300:a582:c60c:df75:f6da:8a92:d5ed:71ad]:9900 or 60.60.60.60:9900).
	namespace string [required] // ZDB namespace.
	password  string [required] // Namespace password.
}

pub struct Network {
pub:
	ip_range             string = '10.1.0.0/16' // network ip range, must have a subnet mask of 16
	add_wireguard_access bool   // if true, a wireguard access point will be added to the network
}

pub struct NetworkResult {
pub:
	name     string // network name
	ip_range string // network ip range
	// computed
	wireguard_config string // wireguard configuration, if any
}
