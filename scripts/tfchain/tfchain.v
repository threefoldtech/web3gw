module tfchain

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

const (
	default_timeout = 500000
)

[params]
pub struct Transfer{
pub:
	amount u64
	destination string
}

[params]
pub struct CreateTwin {
pub:
	relay string
	pk    []byte
}

[params]
pub struct AcceptTermsAndConditions {
	link string
	hash string
}

[params]
pub struct GetContractWithHash {
	node_id u32
	hash    []byte
}

[params]
pub struct CreateNodeContract {
	node_id              u32
	body                 string
	hash                 string
	public_ips           u32
	solution_provider_id ?u64
}

[params]
pub struct CreateRentContract {
	node_id	u32
	solution_provider_id ?u64
}

[params]
pub struct ServiceContractCreate {
	service  []byte
	consumer []byte
}

[params]
pub struct ServiceContractBill {
	contract_id     u64
	variable_amount u64
	metadata        string
}

[params]
pub struct SetServiceContractFees {
	contract_id  u64
	base_fee     u64
	variable_fee u64
}

[params]
pub struct ServiceContractSetMetadata {
	contract_id u64
	metadata    string
}

pub struct PublicIPInput {
	ip      string
	gateway string
}

[params]
pub struct CreateFarm {
	name       string
	public_ips []PublicIPInput
}

[noinit]
pub struct TfChainClient {
mut:
	client &RpcWsClient
}

pub fn new(mut client RpcWsClient) TfChainClient {
	return TfChainClient{
		client: &client
	}
}

pub fn (mut t TfChainClient) load(network string, passphrase string) ! {
	_ := t.client.send_json_rpc[[]string, string]('tfchain.Load', [network, passphrase], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) transfer(args Transfer) ! {
	_ := t.client.send_json_rpc[[]Transfer, string]('tfchain.Transfer', [args],
		tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) balance(address string) !i64 {
	return t.client.send_json_rpc[[]string, i64]('tfchain.Balance', [address], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) height() !u64 {
	return t.client.send_json_rpc[[]string, u64]('tfchain.Height', []string{}, tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) get_twin(id u32) !Twin {
	return t.client.send_json_rpc[[]u32, Twin]('tfchain.GetTwin', [id], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) get_twin_by_pubkey(address string) !u32 {
	return t.client.send_json_rpc[[]string, u32]('tfchain.GetTwinByPubKey', [address], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) create_twin(args CreateTwin) !u32 {
	return t.client.send_json_rpc[[]CreateTwin, u32]('tfchain.CreateTwin', [args], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) accept_terms_and_conditions(args AcceptTermsAndConditions) ! {
	_ := t.client.send_json_rpc[[]AcceptTermsAndConditions, string]('tfchain.AcceptTermsAndConditions', [args], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) get_node(id u32) !Node {
	return t.client.send_json_rpc[[]u32, Node]('tfchain.GetNode', [id], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) create_node(node Node) !u32 {
	return t.client.send_json_rpc[[]Node, u32]('tfchain.CreateNode', [node], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) get_nodes(farm_id u32) ![]u32 {
	return t.client.send_json_rpc[[]u32, []u32]('tfchain.GetNodes', [farm_id], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) get_farm(id u32) !Farm {
	return t.client.send_json_rpc[[]u32, Farm]('tfchain.GetFarm', [id], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) get_farm_by_name(name string) !u32 {
	return t.client.send_json_rpc[[]string, u32]('tfchain.GetFarmByName', [name], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) create_farm(args CreateFarm) ! {
	_ := t.client.send_json_rpc[[]CreateFarm, string]('tfchain.CreateFarm', [args], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) get_contract(contract_id u64) !Contract {
	return t.client.send_json_rpc[[]u64, Contract]('tfchain.GetContract', [contract_id], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) get_contract_id_by_name_registration(name string) !u64 {
	return t.client.send_json_rpc[[]string, u64]('tfchain.GetContractIDByNameRegistration', [name], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) get_contract_with_hash(args GetContractWithHash) !u64 {
	return t.client.send_json_rpc[[]GetContractWithHash, u64]('tfchain.GetContractWithHash', [args], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) get_node_contracts(node_id u32) ![]u64 {
	return t.client.send_json_rpc[[]u32, []u64]('tfchain.GetNodeContracts', [node_id], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) create_name_contract(name string) !u64 {
	return t.client.send_json_rpc[[]string, u64]('tfchain.CreateNameContract', [name], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) create_node_contract(args CreateNodeContract) !u64 {
	return t.client.send_json_rpc[[]CreateNodeContract, u64]('tfchain.CreateNodeContract', [args], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) create_rent_contract(args CreateRentContract) !u64 {
	return t.client.send_json_rpc[[]CreateRentContract, u64]('tfchain.CreateRentContract', [args], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) service_contract_create(args ServiceContractCreate) !u64 {
	return t.client.send_json_rpc[[]ServiceContractCreate, u64]('tfchain.ServiceContractCreate', [args], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) service_contract_approve(contract_id u64) !u64 {
	return t.client.send_json_rpc[[]u64, u64]('tfchain.ServiceContractApprove', [contract_id], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) service_contract_bill(args ServiceContractBill) ! {
	_ := t.client.send_json_rpc[[]ServiceContractBill, string]('tfchain.ServiceContractBill', [args], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) service_contract_cancel(contract_id u64) ! {
	_ := t.client.send_json_rpc[[]u64, string]('tfchain.ServiceContractCancel', [contract_id], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) service_contract_reject(contract_id u64) ! {
	_ := t.client.send_json_rpc[[]u64, string]('tfchain.ServiceContractReject', [contract_id], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) service_contract_set_fees(args SetServiceContractFees) ! {
	_ := t.client.send_json_rpc[[]SetServiceContractFees, string]('tfchain.ServiceContractSetFees', [args], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) service_contract_set_metadata(args ServiceContractSetMetadata) ! {
	_ := t.client.send_json_rpc[[]ServiceContractSetMetadata, string]('tfchain.ServiceContractSetMetadata', [args], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) cancel_contract(contract_id u64) ! {
	_ := t.client.send_json_rpc[[]u64, string]('tfchain.CancelContract', [contract_id], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) batch_cancel_contract(contract_ids []u64) ! {
	_ := t.client.send_json_rpc[[][]u64, string]('tfchain.BatchCancelContract', [contract_ids], tfchain.default_timeout)!
}

pub fn (mut t TfChainClient) get_zos_version() !string {
	return t.client.send_json_rpc[[]string, string]('tfchain.GetZosVersion', []string{}, tfchain.default_timeout)!
}
