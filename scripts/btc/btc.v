module btc

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

const (
	default_timeout = 500000
)

[noinit]
pub struct BtcClient {
mut:
	client &RpcWsClient
}

pub fn new(mut client RpcWsClient) BtcClient {
	return BtcClient{
		client: &client
	}
}

// Needed calls
// Load(ctx context.Context, conState jsonrpc.State, args Load)
// CreateNewAccount(ctx context.Context, conState jsonrpc.State, account string) 
// CreateEncryptedWallet(ctx context.Context, conState jsonrpc.State, passphrase string) 
// ImportAddress(ctx context.Context, conState jsonrpc.State, address string) 
// ImportAddressRescan(ctx context.Context, conState jsonrpc.State, args ImportAddressRescan) 
// ImportPrivKey(ctx context.Context, conState jsonrpc.State, wif string) 
// ImportPrivKeyLabel(ctx context.Context, conState jsonrpc.State, args ImportPrivKeyLabel)
// ImportPrivKeyRescan(ctx context.Context, conState jsonrpc.State, args ImportPrivKeyRescan) 

// ImportPubKey(ctx context.Context, conState jsonrpc.State, pubKey string) 
// ImportPubKeyRescan(ctx context.Context, conState jsonrpc.State, args ImportPubKeyRescan) 
// InvalidateBlock(ctx context.Context, conState jsonrpc.State, hash string) 
// RenameAccount(ctx context.Context, conState jsonrpc.State, args RenameAccount)
// SubmitBlock(ctx context.Context, conState jsonrpc.State, args SubmitBlock)
// SendToAddress(ctx context.Context, conState jsonrpc.State, args SendToAddress)
// SendToAddressComment(ctx context.Context, conState jsonrpc.State, args SendToAddress) 


pub fn (mut e BtcClient) load(params Load) !string {
	return e.client.send_json_rpc[[]Load, string]('btc.Load', [
		params,
	], explorer.default_timeout)!
}