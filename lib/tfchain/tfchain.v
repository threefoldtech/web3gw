module tfchain

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

const (
	default_timeout = 500000
)

pub enum Network as u8 {
	mainnet
	testnet
	qanet
	devnet
}

[params]
pub struct Load {
pub:
	network Network
	mnemonic string
}

[params]
pub struct Transfer {
pub:
	amount      u64
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
	node_id              u32
	solution_provider_id ?u64
}

[params]
pub struct ServiceContractCreate {
	service  string
	consumer string
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

[noinit; openrpc: exclude]
pub struct TfChainClient {
mut:
	client &RpcWsClient
}

[openrpc: exclude]
pub fn new(mut client RpcWsClient) TfChainClient {
	return TfChainClient{
		client: &client
	}
}

// Load your mnemonic with this call. Choose the network while doing so. The network should be one of:
// mainnet, testnet, qanet, devnet 
pub fn (mut t TfChainClient) load(args Load) ! {
	_ := t.client.send_json_rpc[[]Load, string]('tfchain.Load', [args],
		tfchain.default_timeout)!
}

// Transfer some amount to some destination. The destionation should be a SS58 address.
pub fn (mut t TfChainClient) transfer(args Transfer) ! {
	_ := t.client.send_json_rpc[[]Transfer, string]('tfchain.Transfer', [args], tfchain.default_timeout)!
}

// Ask for the balance of an entity using this call. The address should be a SS58 address.
pub fn (mut t TfChainClient) balance(address string) !i64 {
	return t.client.send_json_rpc[[]string, i64]('tfchain.Balance', [address], tfchain.default_timeout)!
}

// Get the current height of the chain. 
pub fn (mut t TfChainClient) height() !u64 {
	return t.client.send_json_rpc[[]string, u64]('tfchain.Height', []string{}, tfchain.default_timeout)!
}

// Get a twin by id. 
pub fn (mut t TfChainClient) get_twin(id u32) !Twin {
	return t.client.send_json_rpc[[]u32, Twin]('tfchain.GetTwin', [id], tfchain.default_timeout)!
}

// Get the twin id that is bound to a SS58 address. 
pub fn (mut t TfChainClient) get_twin_by_pubkey(address string) !u32 {
	return t.client.send_json_rpc[[]string, u32]('tfchain.GetTwinByPubKey', [address],
		tfchain.default_timeout)!
}

// Create a twin. Provide the relay and your public key in this call. The result of this call contains
// your twin id. 
pub fn (mut t TfChainClient) create_twin(args CreateTwin) !u32 {
	return t.client.send_json_rpc[[]CreateTwin, u32]('tfchain.CreateTwin', [args], tfchain.default_timeout)!
}

// Accepts terms and conditions. Provide the document link and document hash while executing this call.
pub fn (mut t TfChainClient) accept_terms_and_conditions(args AcceptTermsAndConditions) ! {
	_ := t.client.send_json_rpc[[]AcceptTermsAndConditions, string]('tfchain.AcceptTermsAndConditions',
		[args], tfchain.default_timeout)!
}

// Get a node by id.
pub fn (mut t TfChainClient) get_node(id u32) !Node {
	return t.client.send_json_rpc[[]u32, Node]('tfchain.GetNode', [id], tfchain.default_timeout)!
}

// Get the nodes that belong to the farm with id. Returns a list of node ids. 
pub fn (mut t TfChainClient) get_nodes(farm_id u32) ![]u32 {
	return t.client.send_json_rpc[[]u32, []u32]('tfchain.GetNodes', [farm_id], tfchain.default_timeout)!
}

// Get farm by id. 
pub fn (mut t TfChainClient) get_farm(id u32) !Farm {
	return t.client.send_json_rpc[[]u32, Farm]('tfchain.GetFarm', [id], tfchain.default_timeout)!
}

// Get farm by name. 
pub fn (mut t TfChainClient) get_farm_by_name(name string) !u32 {
	return t.client.send_json_rpc[[]string, u32]('tfchain.GetFarmByName', [name], tfchain.default_timeout)!
}

// Create a farm. Provide a name an a list of public ips if there are any.
pub fn (mut t TfChainClient) create_farm(args CreateFarm) ! {
	_ := t.client.send_json_rpc[[]CreateFarm, string]('tfchain.CreateFarm', [args], tfchain.default_timeout)!
}

// Get a contract by id. Contract can be any of the following types: node contract, name contract or rent 
// contract.
pub fn (mut t TfChainClient) get_contract(contract_id u64) !Contract {
	return t.client.send_json_rpc[[]u64, Contract]('tfchain.GetContract', [
		contract_id,
	], tfchain.default_timeout)!
}

// Get contract by name registration. Returns the id of the contract.
pub fn (mut t TfChainClient) get_contract_id_by_name_registration(name string) !u64 {
	return t.client.send_json_rpc[[]string, u64]('tfchain.GetContractIDByNameRegistration',
		[name], tfchain.default_timeout)!
}

// Get contract by hash. Returns the contract id.
pub fn (mut t TfChainClient) get_contract_with_hash(args GetContractWithHash) !u64 {
	return t.client.send_json_rpc[[]GetContractWithHash, u64]('tfchain.GetContractWithHash',
		[args], tfchain.default_timeout)!
}

// Get contracts belonging to a node. Returns a list of contract ids.
pub fn (mut t TfChainClient) get_node_contracts(node_id u32) ![]u64 {
	return t.client.send_json_rpc[[]u32, []u64]('tfchain.GetNodeContracts', [node_id],
		tfchain.default_timeout)!
}

// Create a name contract. Provide the dns name via this call. Returns the id of the contract it 
// creates. 
pub fn (mut t TfChainClient) create_name_contract(name string) !u64 {
	return t.client.send_json_rpc[[]string, u64]('tfchain.CreateNameContract', [name],
		tfchain.default_timeout)!
}

// Create a node contract. Provide the node id, body, hash, how many public ips you want to use and 
// optionally the solution provider id.
pub fn (mut t TfChainClient) create_node_contract(args CreateNodeContract) !u64 {
	return t.client.send_json_rpc[[]CreateNodeContract, u64]('tfchain.CreateNodeContract',
		[args], tfchain.default_timeout)!
}

// Create a rent contract. Provide the node id of the node you want to rent and optionally the solution
// provider id. 
pub fn (mut t TfChainClient) create_rent_contract(args CreateRentContract) !u64 {
	return t.client.send_json_rpc[[]CreateRentContract, u64]('tfchain.CreateRentContract',
		[args], tfchain.default_timeout)!
}

// Create a service contract. Provide the SS58 addresses of the service and the consumer.
pub fn (mut t TfChainClient) service_contract_create(args ServiceContractCreate) !u64 {
	return t.client.send_json_rpc[[]ServiceContractCreate, u64]('tfchain.ServiceContractCreate',
		[args], tfchain.default_timeout)!
}

// Approve a service contract. Provide the contract id to do so. Approving the contract is only allowed 
// if the agreement is ready and only by the consumer or the service.
pub fn (mut t TfChainClient) service_contract_approve(contract_id u64) ! {
	_ := t.client.send_json_rpc[[]u64, string]('tfchain.ServiceContractApprove', [
		contract_id,
	], tfchain.default_timeout)!
}

// Bill a service contract. Provide the contract id, variable_amount and some metadata. The contract
// can only be billed if both service and consumer have approved the contract. 
pub fn (mut t TfChainClient) service_contract_bill(args ServiceContractBill) ! {
	_ := t.client.send_json_rpc[[]ServiceContractBill, string]('tfchain.ServiceContractBill',
		[args], tfchain.default_timeout)!
}

// Cancel a service contract. Provide the contract id to do so.
pub fn (mut t TfChainClient) service_contract_cancel(contract_id u64) ! {
	_ := t.client.send_json_rpc[[]u64, string]('tfchain.ServiceContractCancel', [
		contract_id,
	], tfchain.default_timeout)!
}

// Reject a service contract. Provide the contract id to do so. Only the service or the consumer is
// allowed to do so and only if it reached a state of agreement. The contract will be deleted when
// at the first call of either the service or the consumer. 
pub fn (mut t TfChainClient) service_contract_reject(contract_id u64) ! {
	_ := t.client.send_json_rpc[[]u64, string]('tfchain.ServiceContractReject', [
		contract_id,
	], tfchain.default_timeout)!
}

// Set the service contract fees. Provide the contract id, the base fee and the variable fee. Only
// the service is allowed to set the fees and only if the contract is not approved by both parties
// yet. The state is set to agreemen if the metadata has been set.
pub fn (mut t TfChainClient) service_contract_set_fees(args SetServiceContractFees) ! {
	_ := t.client.send_json_rpc[[]SetServiceContractFees, string]('tfchain.ServiceContractSetFees',
		[args], tfchain.default_timeout)!
}

// Set the service contract metadata. Provide contract id and the metadata. Only service or consumer
// can set the metadata. The metadata can be modified by calling this method again but only if the
// contract is not yet approved by both parties. The base fee cannot be 0. 
pub fn (mut t TfChainClient) service_contract_set_metadata(args ServiceContractSetMetadata) ! {
	_ := t.client.send_json_rpc[[]ServiceContractSetMetadata, string]('tfchain.ServiceContractSetMetadata',
		[args], tfchain.default_timeout)!
}

// Cancel a contract. Provide the contract id to do so. Only the creator of the contract can 
// cancle the contract. If the contract is a rent contract it can only be canceled if it has no
// more active workloads. 
pub fn (mut t TfChainClient) cancel_contract(contract_id u64) ! {
	_ := t.client.send_json_rpc[[]u64, string]('tfchain.CancelContract', [contract_id],
		tfchain.default_timeout)!
}

// Cancel a list contracts. Provide a list of contract ids to do so. They can only be canceled
// if they were created by the entity calling this function. Rent contracts can not have active 
// workloads in order to be canceled. 
pub fn (mut t TfChainClient) batch_cancel_contract(contract_ids []u64) ! {
	_ := t.client.send_json_rpc[[][]u64, string]('tfchain.BatchCancelContract', [
		contract_ids,
	], tfchain.default_timeout)!
}

// Get the current ZOS version
pub fn (mut t TfChainClient) get_zos_version() !string {
	return t.client.send_json_rpc[[]string, string]('tfchain.GetZosVersion', []string{},
		tfchain.default_timeout)!
}
