module zos

// wg network reservation (znet)

pub struct Znet {
pub mut:
	// unique nr for each network chosen, this identified private networks as connected to a container or vm or ...
	// corresponds to the 2nd number of a class B ipv4 address
	// is a class C of a chosen class B
	// form: e.g. 192.168.16.0/24
	// needs to be a private subnet
	subnet   string
	ip_range string
	// wireguard private key, curve25519
	// TODO: is this in libsodium
	wireguard_private_key string
	//>1024?
	wireguard_listen_port u16
	peers                 []Peer
}

pub fn (mut n Znet) challenge() string {
	mut out := ''
	out += n.ip_range
	out += n.subnet
	out += n.wireguard_private_key
	out += n.wireguard_listen_port.str()
	for mut p in n.peers {
		out += p.challenge()
	}

	return out
}

// is a remote wireguard client which can connect to this node
pub struct Peer {
pub mut:
	// is another class C in same class B as above
	subnet string
	// wireguard public key, curve25519
	wireguard_public_key string
	allowed_ips          []string
	// ipv4 or ipv6
	// can be empty, one of the 2 need to be filled in though
	endpoint string
}

// TODO: need API endpoint on ZOS to find open ports
// TODO: reservation for 1 h, after will be released again

pub fn (mut p Peer) challenge() string {
	mut out := ''
	out += p.wireguard_public_key
	out += p.endpoint
	out += p.subnet

	for ip in p.allowed_ips {
		out += ip
	}
	return out
}
