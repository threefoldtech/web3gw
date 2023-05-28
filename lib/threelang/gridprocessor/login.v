module gridprocessor

import freeflowuniverse.crystallib.params
import strconv

struct Credentials {
	mnemonic string
	network  string
}

fn (mut g GridProcessor) login(grid_op GridOp, param_map map[string]string, args_set map[string]bool) ! {
	match grid_op {
		.login {}
		else {
			return error('invalid login operation. operation should be empty')
		}
	}

	mnemonic := param_map['mnemonic'] or { return error('mnemonic phrase is missing') }
	network := param_map['network'] or { 'main' }
	g.credentials = Credentials{
		mnemonic: mnemonic
		network: network
	}
}
