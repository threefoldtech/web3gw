module btc

import freeflowuniverse.crystallib.rpcwebsocket

pub fn (mut b BtcClient) import_pub_key(pub_key string) ! {
	b.client.send_json_rpc[[]string, string]('btc.ImportPubKey', [pub_key], default_timeout)!
}

pub fn (mut b BtcClient) import_pub_key_rescan(args ImportPubKeyRescan) ! {
	b.client.send_json_rpc[[]ImportAddressRescan, string]('btc.ImportPubKeyRescan', [
		args,
	], default_timeout)!
}

pub fn (mut b BtcClient) invalidate_block(hash string) ! {
	b.client.send_json_rpc[[]string, string]('btc.InvalidateBlock', [hash], default_timeout)!
}

pub fn (mut b BtcClient) rename_account(args RenameAccount) ! {
	b.client.send_json_rpc[[]RenameAccount, string]('btc.RenameAccount', [args], default_timeout)!
}

pub fn (mut b BtcClient) submit_block(block SubmitBlock) ! {
	b.client.send_json_rpc[[]SubmitBlock, string]('btc.SubmitBlock', [block], default_timeout)!
}

pub fn (mut b BtcClient) send_to_address(args SendToAddress) ! {
	b.client.send_json_rpc[[]SendToAddress, string]('btc.SendToAddress', [args], default_timeout)!
}

pub fn (mut b BtcClient) send_to_address_comment(args SendToAddress) ! {
	b.client.send_json_rpc[[]SendToAddress, string]('btc.SendToAddressComment', [args],
		default_timeout)!
}
