module tfgrid

// the needed configurations for twin login
pub struct Credentials {
	mnemonic string // secret mnemonics 
	network  string // grid network [dev, qa, test, main]
}

// load logins with twin credentials. must be executed on each session
// - credentials: user mnemonics and grid network
pub fn (mut client TFGridClient) load(credentials Credentials) ! {
	_ := client.send_json_rpc[[]string, string]('tfgrid.Load', [credentials.mnemonic, credentials.network],
		default_timeout)!
}
