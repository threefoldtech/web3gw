module eth

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

import math.unsigned { Uint128, uint128_from_dec_str }

const (
	default_timeout = 500000
)

[params]
pub struct Load {
	url    string
	secret string
}

[params]
pub struct Transfer {
	destination string
	amount      string
}

[params]
pub struct TokenTransfer {
	contract_address string
	destination      string
	amount           string
}

[params]
pub struct TokenTransferFrom {
	contract_address string
	from             string
	destination      string
	amount           string
}

[params]
pub struct ApproveTokenSpending {
	contract_address string
	spender          string
	amount           string
}

[params]
pub struct MultisigOwner {
	contract_address string
	target           string
	threshold        i64
}

[params]
pub struct ApproveHash {
	contract_address string
	hash             string
}

[params]
pub struct InitiateMultisigEthTransfer {
	contract_address string
	destination      string
	amount           string
}

[params]
pub struct InitiateMultisigTokenTransfer {
	contract_address string
	token_address    string
	destination      string
	amount           string
}

[params]
pub struct GetFungibleBalance {
	contract_address string
	target           string
}

[params]
pub struct OwnerOfFungible {
	contract_address string
	token_id         i64
}

[params]
pub struct TransferFungible {
	contract_address string
	from             string
	to               string
	token_id         i64
}

[params]
pub struct SetFungibleApproval {
	contract_address string
	from             string
	to               string
	amount           i64
}

[params]
pub struct SetFungibleApprovalForAll {
	contract_address string
	from             string
	to               string
	approved         bool
}

[params]
pub struct ApprovalForFungible {
	contract_address string
	owner            string
	operator         string
}

[params]
pub struct TftEthTransfer {
	destination string
	amount      string
}

[openrpc: exclude]
[noinit]
pub struct EthClient {
mut:
	client &RpcWsClient
}

[openrpc: exclude]
pub fn new(mut client RpcWsClient) EthClient {
	return EthClient{
		client: &client
	}
}

// loads a new eth client
pub fn (mut e EthClient) load(args Load) ! {
	_ := e.client.send_json_rpc[[]Load, string]('eth.Load', [args], eth.default_timeout)!
}

// transfer eth
pub fn (mut e EthClient) transer(args Transfer) !string {
	return e.client.send_json_rpc[[]Transfer, string]('eth.Transfer', [args], eth.default_timeout)!
}

pub fn (mut e EthClient) balance(address string) !Uint128 {
	balance := e.client.send_json_rpc[[]string, string]('eth.Balance', [address], eth.default_timeout)!
	return uint128_from_dec_str(balance)
}

// height returns the current block height.
pub fn (mut e EthClient) height() !u64 {
	return e.client.send_json_rpc[[]string, u64]('eth.Height', []string{}, eth.default_timeout)!
}

// address returns the current loaded eth address.
pub fn (mut e EthClient) address() !string {
	return e.client.send_json_rpc[[]string, string]('eth.Address', []string{}, eth.default_timeout)!
}

// get_hex_seed returns the hex seed for the current loaded account.
pub fn (mut e EthClient) get_hex_seed() !string {
	return e.client.send_json_rpc[[]string, string]('eth.GetHexSeed', []string{}, eth.default_timeout)!
}

// token_balance returns balance for the given token contract.
pub fn (mut e EthClient) token_balance(contractAddress string) !Uint128 {
	balance := e.client.send_json_rpc[[]string, string]('eth.GetTokenBalance', [contractAddress], eth.default_timeout)!
	return uint128_from_dec_str(balance)
	], eth.default_timeout)!
}

// token_transfer transfers tokens to the given address.
pub fn (mut e EthClient) token_transer(args TokenTransfer) !string {
	return e.client.send_json_rpc[[]TokenTransfer, string]('eth.TransferTokens', [
		args,
	], eth.default_timeout)!
}

// token_transfer_from transfers tokens from the given address.
pub fn (mut e EthClient) token_transer_from(args TokenTransferFrom) !string {
	return e.client.send_json_rpc[[]TokenTransferFrom, string]('eth.TransferFromTokens',
		[args], eth.default_timeout)!
}

// approve_token_spending approves token spending for the given address.
pub fn (mut e EthClient) approve_token_spending(args ApproveTokenSpending) !string {
	return e.client.send_json_rpc[[]ApproveTokenSpending, string]('eth.ApproveTokenSpending',
		[args], eth.default_timeout)!
}

// get_multisig_owners returns the owners of the given multisig contract.
pub fn (mut e EthClient) get_multisig_owners(contractAddress string) ![]string {
	return e.client.send_json_rpc[[]string, []string]('eth.GetMultisigOwners', [
		contractAddress,
	], eth.default_timeout)!
}

pub fn (mut e EthClient) get_multisig_threshold(contractAddress string) !Uint128 {
	threshold := e.client.send_json_rpc[[]string, string]('eth.GetMultisigThreshold', [contractAddress],
		eth.default_timeout)!
	return uint128_from_dec_str(threshold)
}

// add_multisig_owner adds a new owner to the given multisig contract.
pub fn (mut e EthClient) add_multisig_owner(args MultisigOwner) !string {
	return e.client.send_json_rpc[[]MultisigOwner, string]('eth.AddMultisigOwner', [
		args,
	], eth.default_timeout)!
}

// remove_multisig_owner removes an owner from the given multisig contract.
pub fn (mut e EthClient) remove_multisig_owner(args MultisigOwner) !string {
	return e.client.send_json_rpc[[]MultisigOwner, string]('eth.RemoveMultisigOwner',
		[args], eth.default_timeout)!
}

// approve_hash approves a hash for the given multisig contract.
pub fn (mut e EthClient) approve_hash(args ApproveHash) !string {
	return e.client.send_json_rpc[[]ApproveHash, string]('eth.ApproveHash', [args], eth.default_timeout)!
}

// is_approved returns true if the given hash is approved for the given multisig contract.
pub fn (mut e EthClient) is_approved(args ApproveHash) !bool {
	return e.client.send_json_rpc[[]ApproveHash, bool]('eth.IsApproved', [args], eth.default_timeout)!
}

// initiate_multisig_eth_transfer initiates a multisig eth transfer.
pub fn (mut e EthClient) initiate_multisig_eth_transfer(args InitiateMultisigEthTransfer) !string {
	return e.client.send_json_rpc[[]InitiateMultisigEthTransfer, string]('eth.InitiateMultisigEthTransfer',
		[args], eth.default_timeout)!
}

// initiate_multisig_token_transfer initiates a multisig token transfer.
pub fn (mut e EthClient) initiate_multisig_token_transfer(args InitiateMultisigTokenTransfer) !string {
	return e.client.send_json_rpc[[]InitiateMultisigTokenTransfer, string]('eth.InitiateMultisigTokenTransfer',
		[args], eth.default_timeout)!
}

// get_fungible_balance returns the balance of the given fungible token.
pub fn (mut e EthClient) get_fungible_balance(args GetFungibleBalance) !Uint128 {
	balance := e.client.send_json_rpc[[]GetFungibleBalance, string]('eth.GetFungibleBalance', [args],
		[args], eth.default_timeout)!
	return uint128_from_dec_str(balance)
}

// onwer_of_fungible returns the owner of the given fungible token.
pub fn (mut e EthClient) owner_of_fungible(args OwnerOfFungible) !string {
	return e.client.send_json_rpc[[]OwnerOfFungible, string]('eth.OwnerOfFungible', [
		args,
	], eth.default_timeout)!
}

// safe_transfer_fungible safely transfers the given fungible token.
pub fn (mut e EthClient) safe_transfer_fungible(args TransferFungible) !string {
	return e.client.send_json_rpc[[]TransferFungible, string]('eth.SafeTransferFungible',
		[args], eth.default_timeout)!
}

// transfer_fungible transfers the given fungible token.
pub fn (mut e EthClient) transfer_fungible(args TransferFungible) !string {
	return e.client.send_json_rpc[[]TransferFungible, string]('eth.TransferFungible',
		[args], eth.default_timeout)!
}

// set_funigble_approval sets the fungible approval for the given fungible token.
pub fn (mut e EthClient) set_fungible_approval(args SetFungibleApproval) !string {
	return e.client.send_json_rpc[[]SetFungibleApproval, string]('eth.SetFungibleApproval',
		[args], eth.default_timeout)!
}

// set_fungible_approval_for_all sets the fungible approval for all the given fungible token.
pub fn (mut e EthClient) set_fungible_approval_for_all(args SetFungibleApprovalForAll) !string {
	return e.client.send_json_rpc[[]SetFungibleApprovalForAll, string]('eth.SetFungibleApprovalForAll',
		[args], eth.default_timeout)!
}

// get_fungible_approval gets the fungible approval for the given fungible token.
pub fn (mut e EthClient) get_approval_for_fungible(args ApprovalForFungible) !bool {
	return e.client.send_json_rpc[[]ApprovalForFungible, bool]('eth.GetApprovalForFungible',
		[args], eth.default_timeout)!
}

// get_approval_for_all_fungible gets the fungible approval for all the given fungible token.
pub fn (mut e EthClient) get_approval_for_all_fungible(args ApprovalForFungible) !bool {
	return e.client.send_json_rpc[[]ApprovalForFungible, bool]('eth.GetApprovalForAllFungible',
		[args], eth.default_timeout)!
}

// transfer_tft_eth transfers tft from an account on ethereum to another
pub fn (mut e EthClient) transfer_eth_tft(args TftEthTransfer) !string {
	return e.client.send_json_rpc[[]TftEthTransfer, string]('eth.TransferEthTft', [
		args,
	], eth.default_timeout)!
}

// bridge_to_stellar bridges tft on ethereum to tft stellar
pub fn (mut e EthClient) bridge_to_stellar(args TftEthTransfer) !string {
	return e.client.send_json_rpc[[]TftEthTransfer, string]('eth.BridgeToStellar',
		[args], eth.default_timeout)!
}

// get_tft_eth_balance returns the tft balance on ethereum
pub fn (mut e EthClient) tft_balance() !Uint128 {
	balance := e.client.send_json_rpc[[]string, string]('eth.GetEthTftBalance', []string{}, eth.default_timeout)!
	return uint128_from_dec_str(balance)
}

// approve_eth_tft_spending approves the given amount of TFT to be swapped
pub fn (mut e EthClient) approve_eth_tft_spending(amount string) !string {
	return e.client.send_json_rpc[[]string, string]('eth.ApproveEthTftSpending', [
		amount,
	], eth.default_timeout)!
}

// get_eth_tft_allowance returns the amount of TFT approved to be swapped
pub fn (mut e EthClient) get_eth_tft_allowance() !string {
	return e.client.send_json_rpc[[]string, string]('eth.EthTftSpendingAllowance', []string{},
		eth.default_timeout)!
}

// quote_eth_for_tft returns the amount of tft that would be received for the given amount of eth
pub fn (mut e EthClient) quote_eth_for_tft(amount_in string) !string {
	return e.client.send_json_rpc[[]string, string]('eth.QuoteEthForTft', [amount_in],
		eth.default_timeout)!
}

// swap_eth_for_tft swaps eth for tft
pub fn (mut e EthClient) swap_eth_for_tft(amount_in string) !string {
	return e.client.send_json_rpc[[]string, string]('eth.SwapEthForTft', [amount_in],
		eth.default_timeout)!
}

// quote_tft_for_eth returns the amount of eth that would be received for the given amount of tft
pub fn (mut e EthClient) quote_tft_for_eth(amount_in string) !string {
	return e.client.send_json_rpc[[]string, string]('eth.QuoteTftForEth', [amount_in],
		eth.default_timeout)!
}

// swap_tft_for_eth swaps tft for eth
pub fn (mut e EthClient) swap_tft_for_eth(amount_in string) !string {
	return e.client.send_json_rpc[[]string, string]('eth.SwapTftForEth', [amount_in],
		eth.default_timeout)!
}
