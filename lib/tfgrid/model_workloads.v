module tfgrid

pub const (
	zdb_workload_type                = 'zdb'
	zmachine_workload_type           = 'zmachine'
	zmount_workload_type             = 'zmount'
	network_workload_type            = 'network'
	zlogs_workload_type              = 'zlogs'
	public_ip_worklaod_type          = 'ip'
	gateway_name_proxy_workload_type = 'gateway-name-proxy'
	gateway_fqdn_proxy_workload_type = 'gateway-fqdn-proxy'
)

// ZMount defines a mount point
pub struct ZMount {
pub:
	size u64 [json: 'size'] // size of the volume
}

// ZMachine reservation data
pub struct ZMachine {
pub:
	flist            string            [json: 'flist'] // Flist of the zmachine, must be a valid url to an flist.
	network          MachineNetwork    [json: 'network'] // Network configuration for machine network
	size             u64               [json: 'size'] // Size of zmachine disk
	compute_capacity MachineCapacity   [json: 'compute_capacity'] // ComputeCapacity configuration for machine cpu+memory
	mounts           []MachineMount    [json: 'mounts'] // Mounts configure mounts/disks attachments to this machine
	entrypoint       string            [json: 'entrypoint'] // entrypoint of the container, if not set the configured one from the flist is going to be used
	env              map[string]string [json: 'env'] // Env variables available for a container
	corex            bool              [json: 'corex'] // Corex works in container mode which forces replace the entrypoing of the container to use `corex`
}

// ZMachineResult result returned by VM reservation
pub struct ZMachineResult {
pub:
	id          string [json: 'id']
	ip          string [json: 'ip']
	ygg_ip      string [json: 'ygg_ip'] // yggdrasil ip of the machine
	console_url string [json: 'console_url']
}

pub struct MachineNetwork {
	public_ip  string             [json: 'public_ip']
	planetary  bool               [json: 'planetary']
	interfaces []MachineInterface [json: 'interfaces']
}

pub struct MachineInterface {
	network string [json: 'network']
	ip      string [json: 'ip']
}

pub struct MachineCapacity {
	cpu    u8  [json: 'cpu']
	memory u64 [json: 'memory']
}

pub struct MachineMount {
	name       string [json: 'name']
	mountpoint string [json: 'mountpoint']
}

// Zlogs is a workload that enables users to stream logs from a zmachine to some url
pub struct Zlogs {
pub:
	zmachine string [json: 'zmachine'] // zmachine name to stream logs for
	output   string [json: 'output'] // output url
}

// ZDB workload info
pub struct ZDBWorkload {
pub:
	size     u64    [json: 'size'] // size of the zdb in GB
	mode     string [json: 'mode'] // mode of the zdb: "user" or "seq"
	password string [json: 'password'] // password for the zdb
	public   bool   [json: 'public'] // if true, makes the zdb read-only if password is set, writable if no password set.
}

// ZDBResultData contains zdb reservation result
pub struct ZDBResultData {
pub:
	// TODO: change json representation snake case after resolving #1952 in zos
	namespace string   [json: 'Namespace'] // namespace of the zdb
	ips       []string [json: 'IPs'] // Computed IPs of the ZDB. Two IPs are returned: a public IPv6, and a YggIP, in this order
	port      u32      [json: 'Port'] // Port of the ZDB.
}

// NetworkWorkload contains the network workload reservation arguments
pub struct NetworkWorkload {
pub:
	network_ip_range IPNet  [json: 'ip_range'] // IP range of the network, must be an IPv4 /16
	subnet           IPNet  [json: 'subnet'] // IPV4 subnet for this network resource, this must be a valid subnet of the entire network ip range.
	wg_private_key   string [json: 'wireguard_private_key'] // The private wg key of this node (this peer) which is installing this network workload right now.
	wg_listen_port   u16    [json: 'wireguard_listen_port'] // WGListenPort is the wireguard listen port on this node.
	peers            []Peer [json: 'peers'] // Peers is a list of other peers in this network
}

pub struct IPNet {
	ip   string [json: 'ip']
	mask string [json: 'mask']
}

pub struct Peer {
	subnet        IPNet   [json: 'subnet']
	wg_public_key string  [json: 'wireguard_public_key']
	allowed_ips   []IPNet [json: 'allowed_ips']
	endpoint      string  [json: 'endpoint']
}

// PublicIP workload arguments
pub struct PublicIP {
pub:
	v4 bool [json: 'v4'] // V4 use one of the reserved Ipv4 from your contract. The Ipv4 itself costs money + the network traffic
	v6 bool [json: 'v6'] // V6 get an ipv6 for the VM. this is for free but the consumed capacity (network traffic) is not
}

// PublicIPResult result returned by publicIP reservation
pub struct PublicIPResult {
pub:
	ip      IPNet  [json: 'ip'] // IP of the VM. The IP must be part of the subnet available in the network resource defined by the networkID on this node
	ipv6    IPNet  [json: 'ipv6'] // IPv6 of the VM.
	gateway string [json: 'gateway']
}

// GatewayNameProxy workload argurments
pub struct GatewayNameProxyWorkload {
pub:
	tls_passthrough bool     [json: 'tls_passthrough'] // whether to pass tls traffic or not
	backends        []string [json: 'backends'] // Backends are list of backend ips (only one is supported atm)
	network         string   [json: 'network'; omitempty] // Network name to join [optional].
}

// GatewayNameProxy reservation result
pub struct GatewayNameProxyResult {
pub:
	fqdn string [json: 'fqdn'] // fqdn of this gateway
}

// GatewayFQDNProxy workload arguments
pub struct GatewayFQDNProxyWorkload {
pub:
	tls_passthrough bool     [json: 'tls_passthrough'] // whether to pass tls traffic or not
	backends        []string [json: 'backends'] // Backends are list of backend ips (only one is supported atm)
	fqdn            string   [json: 'fqdn'] // // FQDN the fully qualified domain name to use (cannot be present with Name)
}
