module zos

pub struct Zmachine {
pub mut:
	flist            string // if full url means custom flist meant for containers, if just name should be an official vm
	network          ZmachineNetwork
	size             u8
	compute_capacity ComputeCapacity
	mounts           []Mount
	entrypoint       string // how to invoke that in a vm?
	env              map[string]string // environment for the zmachine
    corex 			 bool
    gpu				 []string
}

pub struct ZmachineNetwork {
pub mut:
	public_ip  string
	interfaces []ZNetworkInterface
	planetary  bool
}

pub struct ZNetworkInterface {
pub mut:
	network string
	ip      string
}

pub fn (mut n ZmachineNetwork) challenge() string {
	mut out := ''
	out += n.public_ip
	out += n.planetary.str()

	for iface in n.interfaces {
		out += iface.network
		out += iface.ip
	}
	return out
}

pub struct Mount {
pub mut:
	name       string
	mountpoint string
}

pub fn (mut m Mount) challenge() string {
	mut out := ''
	out += m.name
	out += m.mountpoint
	return out
}

pub fn (mut m Zmachine) challenge() string {
	mut out := ''

	out += m.flist
	out += m.network.challenge()
	out += '${m.size}'
	out += m.compute_capacity.challenge()

	for mut mnt in m.mounts {
		out += mnt.challenge()
	}
	out += m.entrypoint
	for k, v in m.env {
		out += k
		out += '='
		out += v
	}
	return out
}

// response of the deployment
pub struct ZmachineResult {
pub mut:
	// name unique per deployment, re-used in request & response
	id string
	ip string
	ygg_ip string
	console_url string
}
