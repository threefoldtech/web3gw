module gridprocessor

struct Credentials {
	mnemonic string
	network  string
}

fn get_credentials(grid_op GridOp, param_map map[string]string, args_set map[string]bool) !Credentials {
	match grid_op {
		.login {}
		else {
			return error('invalid login operation. operation should be empty')
		}
	}

	mnemonic := param_map['mnemonic'] or { return error('mnemonic phrase is missing') }
	network := param_map['network'] or { 'main' }

	return Credentials{
		mnemonic: mnemonic
		network: network
	}
}
