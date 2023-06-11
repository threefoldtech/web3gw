module tfgrid

// The credentials required before executing any other tfgrid rpc.
pub struct Credentials {
	mnemonic string // secret mnemonic
	network  string // grid network [dev, qa, test, main]
}

// Loads the mnemonic into the session for a specific network. The call returns an error if the mnemonic or the
// network is invalid.
pub fn (mut t TFGridClient) load(credentials Credentials) ! {
	_ := t.client.send_json_rpc[[]Credentials, string]('tfgrid.Load', [credentials], default_timeout)!
}
