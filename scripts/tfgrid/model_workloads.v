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

pub struct ZMount {
pub:
	size u64 [json: 'size']
}

pub struct ZMachine {
pub:
	flist            string            [json: 'flist']
	network          MachineNetwork    [json: 'network']
	size             u64               [json: 'size']
	compute_capacity MachineCapacity   [json: 'compute_capacity']
	mounts           []MachineMount    [json: 'mounts']
	entrypoint       string            [json: 'entrypoint']
	env              map[string]string [json: 'env']
	corex            bool              [json: 'corex']
}

pub struct ZMachineResult {
pub:
	id          string [json: 'id']
	ip          string [json: 'ip']
	ygg_ip      string [json: 'ygg_ip']
	console_url string [json: 'console_url']
}

struct MachineNetwork {
	public_ip  string             [json: 'public_ip']
	planetary  bool               [json: 'planetary']
	interfaces []MachineInterface [json: 'interfaces']
}

struct MachineInterface {
	network string [json: 'network']
	ip      string [json: 'ip']
}

struct MachineCapacity {
	cpu    u8  [json: 'cpu']
	memory u64 [json: 'memory']
}

struct MachineMount {
	name       string [json: 'name']
	mountpoint string [json: 'mountpoint']
}

pub struct Zlogs {
pub:
	zmachine string [json: 'zmachine']
	output   string [json: 'output']
}

pub struct ZDBWorkload {
pub:
	size     u64    [json: 'size']
	mode     string [json: 'mode']
	password string [json: 'password']
	public   bool   [json: 'public']
}


// TODO: change json representation snake case after resolving #1952 in zos
pub struct ZDBResultData {
pub:
	namespace string   [json: 'Namespace']
	ips       []string [json: 'IPs']
	port      u32      [json: 'Port']
}

pub struct NetworkWorkload {
pub:
	network_ip_range IPNet  [json: 'ip_range']
	subnet           IPNet  [json: 'subnet']
	wg_private_key   string [json: 'wireguard_private_key']
	wg_listen_port   u16    [json: 'wireguard_listen_port']
	peers            []Peer [json: 'peers']
}

struct IPNet {
	ip   string [json: 'ip']
	mask string [json: 'mask']
}

struct Peer {
	subnet        IPNet   [json: 'subnet']
	wg_public_key string  [json: 'wireguard_public_key']
	allowed_ips   []IPNet [json: 'allowed_ips']
	endpoint      string  [json: 'endpoint']
}

pub struct PublicIP {
pub:
	v4 bool [json: 'v4']
	v6 bool [json: 'v6']
}

pub struct PublicIPResult {
pub:
	ip      IPNet  [json: 'ip']
	ipv6    IPNet  [json: 'ipv6']
	gateway string [json: 'gateway']
}

pub struct GatewayNameProxyWorkload {
pub:
	tls_passthrough bool     [json: 'tls_passthrough']
	backends        []string [json: 'backends']
	network         string   [json: 'network'; omitempty]
}

pub struct GatewayNameProxyResult {
pub:
	fqdn string [json: 'fqdn']
}

pub struct GatewayFQDNProxyWorkload {
pub:
	tls_passthrough bool     [json: 'tls_passthrough']
	backends        []string [json: 'backends']
	fqdn            string   [json: 'fqdn']
}
