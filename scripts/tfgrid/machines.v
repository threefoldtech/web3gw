module tfgrid


// machiens_deploy Deploy machines workload
// - model: the machines model
// returns machines model with the computed fileds from the grid
pub fn (mut client TFGridClient) machines_deploy(model MachinesModel) !MachinesResult {
	return client.send_json_rpc[[]MachinesModel, MachinesResult]('tfgrid.MachinesDeploy',
		[model], default_timeout)!
}

// machines_get Get machines deployment info using deployment name
// - model_name: the machines model name
// returns machines model info
pub fn (mut client TFGridClient) machines_get(model_name string) !MachinesResult {
	return client.send_json_rpc[[]string, MachinesResult]('tfgrid.MachinesGet', [
		model_name,
	], default_timeout)!
}

// machines_delete Delete a deployed machines using project name
// - model_name: the machines model name
pub fn (mut client TFGridClient) machines_delete(model_name string) ! {
	_ := client.send_json_rpc[[]string, string]('tfgrid.MachinesDelete', [model_name],
		default_timeout)!
}

// NOTE: not implemented
// // Add new machine to a machines deployment
// pub fn machines_add_machine(mut client RpcWsClient, params AddMachine) !MachinesResult {
// 	return client.send_json_rpc[[]AddMachine, MachinesResult]('tfgrid.MachinesAdd', [params], default_timeout)!
// }

// // // Delete machine from a machines deployment
// pub fn machines_delete_machine(mut client RpcWsClient, params RemoveMachine) !MachinesResult {
// 	return client.send_json_rpc[[]RemoveMachine, MachinesResult]('tfgrid.MachinesRemove', [params], default_timeout)!
// }
