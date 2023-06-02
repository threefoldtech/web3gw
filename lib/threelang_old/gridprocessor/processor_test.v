module gridprocessor

import freeflowuniverse.crystallib.params { Param, Params }

fn test_add_action_valid() {
	mut p := GridProcessor{}

	p.add_action('k8s', 'create', Params{
		params: [
			Param{
				key: 'name'
				value: 'cluster1'
			},
			Param{
				key: 'workers'
				value: '4'
			},
			Param{
				key: 'worker_size'
				value: 'large'
			},
		]
	})!
}
