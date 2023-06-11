module tfgrid

// Deploys a vms with gateways
pub fn (mut t TFGridClient) deploy_vm(vm VM) !VMResult {
	return t.client.send_json_rpc[[]VM, VMResult]('tfgrid.DeployVMWithGW', [vm], default_timeout)!
}

// Removes a vm from a network
pub fn (mut t TFGridClient) remove_vm(args RemoveVMWithGWArgs) !VMResult {
	return t.client.send_json_rpc[[]RemoveVMWithGWArgs, VMResult]('tfgrid.RemoveVMWithGW',
		[args], default_timeout)!
}

// Gets a deployed vms with gateways
pub fn (mut t TFGridClient) get_vm(vm_name string) !VMResult {
	return t.client.send_json_rpc[[]string, VMResult]('tfgrid.GetVMWithGW', [
		vm_name,
	], default_timeout)!
}

// Deletes a deployed vms with gateways.
pub fn (mut t TFGridClient) delete_vm(vm_name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.DeleteVMWithGW', [
		vm_name,
	], default_timeout)!
}
