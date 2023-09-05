module zos

import json

// wg network reservation (znet)

pub struct Znet {
pub mut:
	// unique nr for each network chosen, this identified private networks as connected to a container or vm or ...
	// corresponds to the 2nd number of a class B ipv4 address
	// is a class C of a chosen class B
	// form: e.g. 192.168.16.0/24
	// needs to be a private subnet
	subnet   string //TODO: what is format
	ip_range string //TODO: what is format, what is difference with subnet
	// wireguard private key, curve25519
	// TODO: is this in libsodium 
	wireguard_private_key string //TODO: what is format
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
	subnet string //TODO: what is format
	// wireguard public key, curve25519
	wireguard_public_key string   //TODO: what is format
	//TODO: give example ipv4 and ipv6 for allowed_ips
	allowed_ips          []string /is ipv4 or ipv6 address from a wireguard client who connects, with netmark
	// ipv4 or ipv6
	// can be empty, one of the 2 need to be filled in though
	endpoint string //TODO: what is format
}

// TODO: need API endpoint on ZOS to find open ports
// TODO: reservation for 1 h, after will be released again, ??? what does this mean?

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

pub fn (z Znet) to_workload(args WorkloadArgs) Workload {
	return Workload{
		version: args.version or { 0 }
		name: args.name
		type_: workload_types.network
		data: json.encode(z)
		metadata: args.metadata or { '' }
		description: args.description or { '' }
		result: args.result or { WorkloadResult{} }
	}
}
