module tfgrid

[params]
pub struct VM {
pub mut:
	name                 string // This is the vm's name. If multiple vms are to be deployed, index is appended to the vm's name. If not provided, a random name is generated.
	network              string // Identifier for the network that the vm will be a part of
	farm_id              u32    // farm id to deploy on, if 0, a random eligible node on a random farm will be selected
	capacity             string // capacity for the vm, could be 'small', 'medium', 'large', or 'extra-large'
	times                u32 = 1 // indicates how many vms will be deployed with the configuration defined by this object
	disk_size            u32    // disk size to mount on the vms in GB
	ssh_key              string // this is the public key that will allow you to ssh into the VM at a later stage
	gateway              bool   // if true, a gateway will deployed for the vm. the vm should listen for traffic coming from the gateway at port 9000
	add_wireguard_access bool   // if true, a wireguard access point will be added to the network
	add_public_ipv4      bool   // if true, a public ipv4 will be added the vm
	add_public_ipv6      bool   // if true, a public ipv6 will be added the vm
}

[params]
pub struct RemoveVM {
pub:
	network string
	vm_name string
}

// Deploys a vm with the posibility to add a gateway. if the there is already a network with the same name, the the vms are added to this network
pub fn (mut t TFGridClient) deploy_vm(vm VM) !VMResult {
	return t.client.send_json_rpc[[]VM, VMResult]('tfgrid.DeployVM', [vm], default_timeout)!
}

// Removes a vm from a network
pub fn (mut t TFGridClient) remove_vm(args RemoveVM) !VMResult {
	return t.client.send_json_rpc[[]RemoveVM, VMResult]('tfgrid.RemoveVM', [
		args,
	], default_timeout)!
}

// Gets a deployed network of vms
pub fn (mut t TFGridClient) get_vm(network string) !VMResult {
	return t.client.send_json_rpc[[]string, VMResult]('tfgrid.GetVM', [
		network,
	], default_timeout)!
}

// Deletes a deployed network of vms
pub fn (mut t TFGridClient) delete_vm(network string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.DeleteVM', [
		network,
	], default_timeout)!
}
