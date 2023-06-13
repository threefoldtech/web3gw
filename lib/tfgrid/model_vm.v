module tfgrid

pub struct VMResult {
pub mut:
	network          string
	wireguard_config string
	vms              []GatewayedMachines
}

struct GatewayedMachines {
pub:
	machine MachineResult
	gateway GatewayNameResult
}
