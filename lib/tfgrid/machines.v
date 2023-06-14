module tfgrid

[params]
pub struct MachinesModel {
pub mut:
	name        string    [required] // identifier for the machines model, must be unique
	network     Network   [required] // network cofigs
	machines    []Machine [required] // machines specs
	metadata    string // metadata for the model
	description string // description of the model
}

[params]
pub struct AddMachine {
pub:
	machine    Machine // machine configs
	model_name string  // model name that this machine will be added to
}

[params]
pub struct RemoveMachine {
pub:
	machine_name string // machine name to be removed
	model_name   string // model name that this machine will be removed from
}

// Deploys a machines workload given a machines model. If it does not succeed the call returns an error.
pub fn (mut t TFGridClient) machines_deploy(model MachinesModel) !MachinesResult {
	return t.client.send_json_rpc[[]MachinesModel, MachinesResult]('tfgrid.MachinesDeploy',
		[model], default_timeout)!
}

// Gets the information about a machines deployment using the deployment name
pub fn (mut t TFGridClient) machines_get(model_name string) !MachinesResult {
	return t.client.send_json_rpc[[]string, MachinesResult]('tfgrid.MachinesGet', [
		model_name,
	], default_timeout)!
}

// Deletes a deployed machines given the name used when deploying.
pub fn (mut t TFGridClient) machines_delete(model_name string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfgrid.MachinesDelete', [
		model_name,
	], default_timeout)!
}

// Add new machine to a machines deployment
pub fn (mut t TFGridClient) machines_add(params AddMachine) !MachinesResult {
	return t.client.send_json_rpc[[]AddMachine, MachinesResult]('tfgrid.MachinesAdd',
		[params], default_timeout)!
}

// Remove machine from a machines deployment
pub fn (mut t TFGridClient) machines_remove(params RemoveMachine) !MachinesResult {
	return t.client.send_json_rpc[[]RemoveMachine, MachinesResult]('tfgrid.MachinesRemove',
		[
		params,
	], default_timeout)!
}
