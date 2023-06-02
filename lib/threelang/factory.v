import freeflowuniverse.crystallib.actionsparser

pub struct Runner {
pub mut:
	path string
	//TODO: hold client

}

[params]
pub struct RunnerArgs {
pub mut:
	name      string
	path string
}



pub fn new(args RunnerArgs) !Runner {
	mut factory := Runner{
			path:args.path
		}
	factory.run()!
	return factory
}

pub fn (mut r Runner) run() ! {
	ap := actionsparser.new(path: m.params.path, defaultbook: 'aaa')!
	r.core_actions(ap)!
	r.vm_actions(ap)!
}


}

pub fn (mut r Runner) tfgrid_client_get(name string) ! {
	//TODO: work with sumtype, look for right name
}
