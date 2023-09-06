module tfgrid

import json

pub struct ZLogs {
	zmachine string //TODO: format of string
	output   string //TODO: format of string
}

pub fn (z ZLogs) challenge() string {
	mut output := ''
	output += z.zmachine 
	output += z.output  

	return output
}

pub fn (z ZLogs) to_workload(args WorkloadArgs) Workload {
	return Workload{
		version: args.version or { 0 }
		name: args.name
		type_: workload_types.zlogs
		data: json.encode(z)
		metadata: args.metadata or { '' }
		description: args.description or { '' }
		result: args.result or { WorkloadResult{} }
	}
}
