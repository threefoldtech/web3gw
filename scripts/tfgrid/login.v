module tfgrid

pub struct Credentials {
	mnemonic string
	network  string
}

pub fn (mut client TFGridClient) load(credentials Credentials) ! {
	_ := client.send_json_rpc[[]string, string]('tfgrid.Load', [credentials.mnemonic, credentials.network],
		default_timeout)!
}
