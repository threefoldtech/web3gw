module tfgrid

pub struct VMResult {
pub mut:
	network          string // vm network
	wireguard_config string // vm wireguard configuration, if any
	vms              []GatewayedMachines // vms configs
}

pub struct GatewayedMachines {
pub:
	machine MachineResult     // machine configs
	gateway GatewayNameResult // gateway configs
}
