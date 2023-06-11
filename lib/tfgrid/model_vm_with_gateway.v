module tfgrid

pub struct VM {
pub mut:
	name                 string // this is the vm's name, if multiple vms are to be deployed, and index is appended to the vm's name
	network              string
	farm_id              u32
	capacity             string
	times                u32 = 1
	disk_size            u32
	ssh_key              string
	gateway              bool
	add_wireguard_access bool
	add_public_ips       bool
}

struct VMResult {
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

pub struct RemoveVMWithGWArgs {
pub:
	network string
	vm_name string
}
