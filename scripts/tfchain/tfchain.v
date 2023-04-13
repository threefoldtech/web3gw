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

pub fn load(mut client RpcWsClient, network string, passphrase string) ! {
	_ := client.send_json_rpc[[]string, string]('tfchain.Load', [network, passphrase], tfchain.default_timeout)!
}

pub fn transfer(mut client RpcWsClient, args Transfer) ! {
	_ := client.send_json_rpc[[]Transfer, string]('tfchain.Transfer', [args],
		tfchain.default_timeout)!
}

pub fn balance(mut client RpcWsClient, address string) !i64 {
	return client.send_json_rpc[[]string, i64]('tfchain.Balance', [address], tfchain.default_timeout)!
}

pub fn height(mut client RpcWsClient) !u64 {
	return client.send_json_rpc[[]string, u64]('tfchain.Height', []string{}, tfchain.default_timeout)!
}

pub fn get_twin(mut client RpcWsClient, id u32) !Twin {
	return client.send_json_rpc[[]u32, Twin]('tfchain.GetTwin', [id], tfchain.default_timeout)!
}

pub fn get_twin_by_pubkey(mut client RpcWsClient, pk []byte) !u32 {
	return client.send_json_rpc[[][]byte, u32]('tfchain.GetTwinByPubKey', [pk], tfchain.default_timeout)!
}

pub fn create_twin(mut client RpcWsClient, args CreateTwin) !u32 {
	return client.send_json_rpc[[]CreateTwin, u32]('tfchain.GetTwinByPubKey', [args], tfchain.default_timeout)!
}

pub fn accept_terms_and_conditions(mut client RpcWsClient, args AcceptTermsAndConditions) ! {
	_ := client.send_json_rpc[[]AcceptTermsAndConditions, string]('tfchain.AcceptTermsAndConditions', [args], tfchain.default_timeout)!
}

pub fn get_node(mut client RpcWsClient, id u32) !Node {
	return client.send_json_rpc[[]u32, Node]('tfchain.GetNode', [id], tfchain.default_timeout)!
}

pub fn create_node(mut client RpcWsClient, node Node) !u32 {
	return client.send_json_rpc[[]Node, u32]('tfchain.CreateNode', [node], tfchain.default_timeout)!
}

pub fn get_nodes(mut client RpcWsClient, farm_id u32) ![]u32 {
	return client.send_json_rpc[[]u32, []u32]('tfchain.GetNodes', [farm_id], tfchain.default_timeout)!
}


pub fn get_farm(mut client RpcWsClient, id u32) !Farm {
	return client.send_json_rpc[[]u32, Farm]('tfchain.GetFarm', [id], tfchain.default_timeout)!
}

pub fn get_farm_by_name(mut client RpcWsClient, name string) !u32 {
	return client.send_json_rpc[[]string, u32]('tfchain.GetFarmByName', [name], tfchain.default_timeout)!
}

pub fn create_farm(mut client RpcWsClient, args CreateFarm) ! {
	_ := client.send_json_rpc[[]CreateFarm, string]('tfchain.CreateFarm', [args], tfchain.default_timeout)!
}


pub fn get_contract(mut client RpcWsClient, contract_id u64) !Contract {
	return client.send_json_rpc[[]u64, Contract]('tfchain.GetContract', [contract_id], tfchain.default_timeout)!
}

pub fn get_contract_id_by_name_registration(mut client RpcWsClient, name string) !u64 {
	return client.send_json_rpc[[]string, u64]('tfchain.GetContractIDByNameRegistration', [name], tfchain.default_timeout)!
}

pub fn get_contract_with_hash(mut client RpcWsClient, args GetContractWithHash) !u64 {
	return client.send_json_rpc[[]GetContractWithHash, u64]('tfchain.GetContractWithHash', [args], tfchain.default_timeout)!
}

pub fn create_name_contract(mut client RpcWsClient, name string) !u64 {
	return client.send_json_rpc[[]string, u64]('tfchain.CreateNameContract', [name], tfchain.default_timeout)!
}

pub fn create_node_contract(mut client RpcWsClient, args CreateNodeContract) !u64 {
	return client.send_json_rpc[[]CreateNodeContract, u64]('tfchain.CreateNodeContract', [args], tfchain.default_timeout)!
}

pub fn create_rent_contract(mut client RpcWsClient, args CreateRentContract) !u64 {
	return client.send_json_rpc[[]CreateRentContract, u64]('tfchain.CreateRentContract', [args], tfchain.default_timeout)!
}

pub fn service_contract_create(mut client RpcWsClient, args ServiceContractCreate) !u64 {
	return client.send_json_rpc[[]ServiceContractCreate, u64]('tfchain.ServiceContractCreate', [args], tfchain.default_timeout)!
}

pub fn service_contract_approve(mut client RpcWsClient, contract_id u64) !u64 {
	return client.send_json_rpc[[]u64, u64]('tfchain.ServiceContractApprove', [contract_id], tfchain.default_timeout)!
}

pub fn service_contract_bill(mut client RpcWsClient, args ServiceContractBill) ! {
	_ := client.send_json_rpc[[]ServiceContractBill, string]('tfchain.ServiceContractBill', [args], tfchain.default_timeout)!
}

pub fn service_contract_cancel(mut client RpcWsClient, contract_id u64) ! {
	_ := client.send_json_rpc[[]u64, string]('tfchain.ServiceContractCancel', [contract_id], tfchain.default_timeout)!
}

pub fn service_contract_reject(mut client RpcWsClient, contract_id u64) ! {
	_ := client.send_json_rpc[[]u64, string]('tfchain.ServiceContractReject', [contract_id], tfchain.default_timeout)!
}

pub fn service_contract_set_fees(mut client RpcWsClient, args SetServiceContractFees) ! {
	_ := client.send_json_rpc[[]SetServiceContractFees, string]('tfchain.ServiceContractSetFees', [args], tfchain.default_timeout)!
}

pub fn service_contract_set_metadata(mut client RpcWsClient, args ServiceContractSetMetadata) ! {
	_ := client.send_json_rpc[[]ServiceContractSetMetadata, string]('tfchain.ServiceContractSetMetadata', [args], tfchain.default_timeout)!
}

pub fn cancel_contract(mut client RpcWsClient, contract_id u64) ! {
	_ := client.send_json_rpc[[]u64, string]('tfchain.CancelContract', [contract_id], tfchain.default_timeout)!
}

pub fn get_zos_version(mut client RpcWsClient) !string {
	return client.send_json_rpc[[]string, string]('tfchain.CancelContract', []string{}, tfchain.default_timeout)!
}
