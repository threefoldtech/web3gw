module tfgrid

[params]
pub struct VM {
pub mut:
	name                 string // This is the vm's name. If multiple vms are to be deployed, index is appended to the vm's name. If not provided, a random name is generated.
	network              string // Identifier for the network that these VMs will be a part of
	farm_id              u32
	capacity             string // capacity for the vms, could be 'small', 'medium', 'large', or 'extra-large'
	times                u32 = 1 // indicates how many vms to be deployed
	disk_size            u32    // disk size to mount on vms in GB
	ssh_key              string // this is the public key that will allow you to ssh into the VM at a later stage
	gateway              bool   // true to deploy a gateway for each vm. vms should listen for traffic coming from the gateway at port 9000
	add_wireguard_access bool   // true to add wireguard access to the network
	add_public_ips       bool   // true to add public ips for the vms
	public_ipv6          bool   // true to add public ipv6 for the vms
}

[params]
pub struct RemoveVMWithGWArgs {
pub:
	network string
	vm_name string
}

// Deploys a vm with the posibility to add a gateway. if the there is already a network with the same name, the the vms are added to this network
pub fn (mut t TFGridClient) deploy_vm(vm VM) !VMResult {
	return t.client.send_json_rpc[[]VM, VMResult]('tfgrid.DeployVM', [vm], default_timeout)!
}

// Removes a vm from a network
pub fn (mut t TFGridClient) remove_vm(args RemoveVMWithGWArgs) !VMResult {
	return t.client.send_json_rpc[[]RemoveVMWithGWArgs, VMResult]('tfgrid.RemoveVM', [
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