module tfchain

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

const (
	default_timeout = 500000
)

[params]
struct Transfer{
pub:
	amount u64
	destination string
}

[params]
struct CreateTwin {
pub:
	relay string
	pk    []byte
}

[params]
struct AcceptTermsAndConditions {
	link string
	hash string
}

[params]
struct GetContractWithHash {
	node_id u32
	hash    []byte
}

[params]
struct CreateNodeContract {
	node_id              u32
	body                 string
	hash                 string
	public_ips           u32
	solution_provider_id &u64
}

[params]
struct CreateRentContract {
	node_id	u32
	solution_provider_id &u64
}

[params]
struct ServiceContractCreate {
	service  []byte
	consumer []byte
}

[params]
struct ServiceContractBill {
	contract_id     u64
	variable_amount u64
	metadata        string
}

[params]
struct SetServiceContractFees {
	contract_id  u64
	base_fee     u64
	variable_fee u64
}

[params]
struct ServiceContractSetMetadata {
	contract_id u64
	metadata    string
}

struct PublicIPInput {
	ip      string
	gateway string
}

[params]
struct CreateFarm {
	name       string
	public_ips []PublicIPInput
}

pub fn load(mut client RpcWsClient, network string, passphrase string) ! {
	_ := client.send_json_rpc[[]string, string]('tfchain.Load', [network, passphrase], tfchain.default_timeout)!
}

pub fn transer(mut client RpcWsClient, args Transfer) ! {
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


/*
func (c *Client) GetFarmByName(ctx context.Context, name string) u32, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.GetFarmByName(name)
}

func (c *Client) CreatetFarm(ctx context.Context, args CreateFarm) error {

func (c *Client) GetContract(ctx context.Context, contract_id u64) (*substrate.Contract, error) {

func (c *Client) GetContractIDByNameRegistration(ctx context.Context, name string) (u64, error) {

func (c *Client) GetContractWithHash(ctx context.Context, args GetContractWithHash) (u64, error) {

func (c *Client) CreateNameContract(ctx context.Context, name string) (u64, error) {

func (c *Client) CreateNodeContract(ctx context.Context, args CreateNodeContract) (u64, error) {

func (c *Client) CreateRentContract(ctx context.Context, args CreateRentContract) (u64, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.CreateRentContract(*state.identity, args.node_id, args.solution_provider_id)
}

func (c *Client) ServiceContractCreate(ctx context.Context, args ServiceContractCreate) (u64, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return 0, pkg.ErrClientNotConnected{}
	}

	return state.client.ServiceContractCreate(*state.identity, args.service, args.consumer)
}

func (c *Client) ServiceContractApprove(ctx context.Context, contract_id u64) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ServiceContractApprove(*state.identity, contract_id)
}

func (c *Client) ServiceContractBill(ctx context.Context, args ServiceContractBill) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ServiceContractBill(*state.identity, args.contract_id, args.variable_amount, args.metadata)
}

func (c *Client) ServiceContractCancel(ctx context.Context, contract_id u64) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ServiceContractCancel(*state.identity, contract_id)
}

func (c *Client) ServiceContractReject(ctx context.Context, contract_id u64) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ServiceContractReject(*state.identity, contract_id)
}

func (c *Client) ServiceContractSetFees(ctx context.Context, args SetServiceContractFees) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ServiceContractSetFees(*state.identity, args.contract_id, args.base_fee, args.variable_fee)
}

func (c *Client) ServiceContractSetMetadata(ctx context.Context, args ServiceContractSetMetadata) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.ServiceContractSetMetadata(*state.identity, args.contract_id, args.metadata)
}

func (c *Client) CancelContract(ctx context.Context, contract_id u64) error {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return pkg.ErrClientNotConnected{}
	}

	return state.client.CancelContract(*state.identity, contract_id)
}

func (c *Client) GetZosVersion(ctx context.Context) (string, error) {
	state, ok := c.state.Get(state.IDFromContext(ctx))
	if !ok || state.client == nil {
		return "", pkg.ErrClientNotConnected{}
	}

	return state.client.GetZosVersion()
}
*/