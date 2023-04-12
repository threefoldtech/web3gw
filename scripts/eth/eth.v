module eth

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

const (
	default_timeout = 500000
)

// CORE

pub fn load(mut client RpcWsClient, url string, secret string) ! {
	_ := client.send_json_rpc[[]string, string]('eth.Load', [url, secret], eth.default_timeout)!
}

pub fn transer(mut client RpcWsClient, destination string, amount i64) !string {
	_ := client.send_json_rpc[[]string, string]('eth.Transfer', [destination, amount],
		eth.default_timeout)!
}

pub fn balance(mut client RpcWsClient, address string) !i64 {
	return client.send_json_rpc[[]string, i64]('eth.Balance', [address], eth.default_timeout)!
}

pub fn height(mut client RpcWsClient) !u64 {
	return client.send_json_rpc[[]string, u64]('eth.Height', []string{}, eth.default_timeout)!
}

// ERC20

pub fn token_balance(mut client RpcWsClient, contractAddress string, address string) !i64 {
	return client.send_json_rpc[[]string, i64]('eth.GetTokenBalance', [contractAddress, address], eth.default_timeout)!
}

pub fn token_transer(mut client RpcWsClient, contractAddress string, destination string, amount i64) !string {
	_ := client.send_json_rpc[[]string, string]('eth.TransferTokens', [contractAddress, destination, amount],
		eth.default_timeout)!
}

pub fn token_transer_from(mut client RpcWsClient, contractAddress string, from string, destination string, amount i64) !string {
	_ := client.send_json_rpc[[]string, string]('eth.TransferFromTokens', [contractAddress, from, destination, amount],
		eth.default_timeout)!
}

pub fn approve_token_spending(mut client RpcWsClient, contractAddress string, destination string, amount i64) !string {
	_ := client.send_json_rpc[[]string, string]('eth.ApproveTokenSpending', [contractAddress, destination, amount],
		eth.default_timeout)!
}

// Multisig

pub fn get_multisig_owners(mut client RpcWsClient, contractAddress string) ![]string {
	_ := client.send_json_rpc[[]string, string]('eth.GetMultisigOwners', [contractAddress],
		eth.default_timeout)!
}

pub fn get_multisig_threshold(mut client RpcWsClient, contractAddress string) !i64 {
	_ := client.send_json_rpc[[]string, string]('eth.GetMultisigThreshold', [contractAddress],
		eth.default_timeout)!
}

pub fn add_multisig_owner(mut client RpcWsClient, contractAddress string, target string, threshold i64) !string {
	_ := client.send_json_rpc[[]string, string]('eth.AddMultisigOwner', [contractAddress, target, threshold],
		eth.default_timeout)!
}

pub fn remove_multisig_owner(mut client RpcWsClient, contractAddress string, target string, threshold i64) !string {
	_ := client.send_json_rpc[[]string, string]('eth.RemoveMultisigOwner', [contractAddress, target, threshold],
		eth.default_timeout)!
}

pub fn approve_hash(mut client RpcWsClient, contractAddress string, hash string) !string {
	_ := client.send_json_rpc[[]string, string]('eth.ApproveHash', [contractAddress, target, threshold],
		eth.default_timeout)!
}

pub fn is_approved(mut client RpcWsClient, contractAddress string, hash string) !bool {
	_ := client.send_json_rpc[[]string, string]('eth.IsApproved', [contractAddress, target, threshold],
		eth.default_timeout)!
}

pub fn initiate_multisig_token_transfer(mut client RpcWsClient, contractAddress string, destination string, amount i64) !string {
	_ := client.send_json_rpc[[]string, string]('eth.InitiateMultisigEthTransfer', [contractAddress, destination, amount],
		eth.default_timeout)!
}

pub fn initiate_multisig_token_transfer(mut client RpcWsClient, contractAddress string, tokenAddress string, destination string, amount i64) !string {
	_ := client.send_json_rpc[[]string, string]('eth.InitiateMultisigTokenTransfer', [contractAddress, tokenAddress, destination, amount],
		eth.default_timeout)!
}

// Fungibles

// GetFungibleBalance (ERC721)
pub fn get_fungible_balance(mut client RpcWsClient, contractAddress string, address string) !i64 {
	_ := client.send_json_rpc[[]string, string]('eth.GetFungibleBalance', [contractAddress, address],
		eth.default_timeout)!
}

// OwnerOfFungible (ERC721)
pub fn owner_of_fungible(mut client RpcWsClient, contractAddress string, tokenId i64) !string {
	_ := client.send_json_rpc[[]string, string]('eth.OwnerOfFungible', [contractAddress, tokenId],
		eth.default_timeout)!
}

// SafeTransferFungible (ERC721)
pub fn safe_transfer_fungible(mut client RpcWsClient, contractAddress string, from string, to string, tokenId i64) !string {
	_ := client.send_json_rpc[[]string, string]('eth.SafeTransferFungible', [contractAddress, from, to, tokenId],
		eth.default_timeout)!
}

// TransferFungible (ERC721)
pub fn transfer_fungible(mut client RpcWsClient, contractAddress string, from string, to string, tokenId i64) !string {
	_ := client.send_json_rpc[[]string, string]('eth.TransferFungible', [contractAddress, from, to, tokenId],
		eth.default_timeout)!
}

// SetFungibleApproval (ERC721)
pub fn set_fungible_approval(mut client RpcWsClient, contractAddress string, from string, to string, approved bool) !string {
	_ := client.send_json_rpc[[]string, string]('eth.SetFungibleApproval', [contractAddress, from, to, approved],
		eth.default_timeout)!
}

// SetFungibleApprovalForAll (ERC721)
pub fn set_fungible_approval_for_all(mut client RpcWsClient, contractAddress string, from string, to string, approved bool) !string {
	_ := client.send_json_rpc[[]string, string]('eth.SetFungibleApprovalForAll', [contractAddress, from, to, approved],
		eth.default_timeout)!
}

// GetApprovalForFungible (ERC721)
pub fn get_fungible_approval(mut client RpcWsClient, contractAddress string, from string, to string) !bool {
	_ := client.send_json_rpc[[]string, string]('eth.GetApprovalForFungible', [contractAddress, from, to],
		eth.default_timeout)!
}

// GetApprovalForAllFungible (ERC721)
pub fn get_fungible_approval_for_all(mut client RpcWsClient, contractAddress string, from string, to string) !bool {
	_ := client.send_json_rpc[[]string, string]('eth.GetApprovalForAllFungible', [contractAddress, from, to],
		eth.default_timeout)!
}