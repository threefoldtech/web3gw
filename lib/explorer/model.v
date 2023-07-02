module explorer

pub struct NodesRequestParams {
	filters    NodeFilter
	pagination Limit
}

pub struct FarmsRequestParams {
	filters    FarmFilter
	pagination Limit
}

pub struct TwinsRequestParams {
	filters    TwinFilter
	pagination Limit
}

pub struct ContractsRequestParams {
	filters    ContractFilter
	pagination Limit
}

pub struct NodesResult {
pub:
	nodes       []Node
	total_count int
}

pub struct FarmsResult {
	farms       []Farm
	total_count int
}

pub struct TwinsResult {
pub:
	twins       []Twin
	total_count int
}

pub struct ContractsResult {
	contracts   []Contract
	total_count int
}

// available options for filtering nodes
pub struct NodeFilter {
pub mut:
	status             ?string [json: 'Status']
	free_mru           ?u64    [json: 'FreeMru']
	free_hru           ?u64    [json: 'FreeHru']
	free_sru           ?u64    [json: 'FreeSru']
	total_mru          ?u64    [json: 'TotalMru']
	total_hru          ?u64    [json: 'TotalHru']
	total_sru          ?u64    [json: 'TotalSru']
	total_cru          ?u64    [json: 'TotalCru']
	country            ?string [json: 'Country']
	country_contains   ?string [json: 'CountryContains']
	city               ?string [json: 'City']
	city_contains      ?string [json: 'CityContains']
	farm_name          ?string [json: 'FarmName']
	farm_name_contains ?string [json: 'FarmNameContains']
	farm_ids           ?[]u64  [json: 'FarmIds']
	free_ips           ?u64    [json: 'FreeIps']
	ipv4               ?bool   [json: 'IPv4']
	ipv6               ?bool   [json: 'IPv6']
	domain             ?bool   [json: 'Domain']
	dedicated          ?bool   [json: 'Dedicated']
	rentable           ?bool   [json: 'Rentable']
	rented             ?bool   [json: 'Rented']
	rented_by          ?u64    [json: 'RentedBy']
	available_for      ?u64    [json: 'AvailableFor']
	node_id            ?u64    [json: 'NodeID']
	twin_id            ?u64    [json: 'TwinID']
}

// available options for filtering farms
pub struct FarmFilter {
pub mut:
	free_ips           ?u64    [json: 'FreeIPs']
	total_ips          ?u64    [json: 'TotalIPs']
	stellar_address    ?string [json: 'StellarAddress']
	pricing_policy_id  ?u64    [json: 'PricingPolicyId']
	farm_id            ?u64    [json: 'FarmId']
	twin_id            ?u64    [json: 'TwinId']
	name               ?string [json: 'Name']
	name_contains      ?string [json: 'NameContains']
	certification_type ?string [json: 'CertificationType']
	dedicated          ?bool   [json: 'Dedicated']
}

// available options for filtering twins
pub struct TwinFilter {
pub mut:
	twin_id    ?u64    [json: 'TwinID']
	account_id ?string [json: 'AccountID']
	relay      ?string [json: 'Relay']
	public_key ?string [json: 'PublicKey']
}

// available options for filtering contracts
pub struct ContractFilter {
pub mut:
	contract_id          ?u64    [json: 'ContractID']
	twin_id              ?u64    [json: 'TwinID']
	node_id              ?u64    [json: 'NodeID']
	type_                ?string [json: 'Type']
	state                ?string [json: 'State']
	name                 ?string [json: 'Name']
	number_of_public_ips ?u64    [json: 'NumberOfPublicIps']
	deployment_data      ?string [json: 'DeploymentData']
	deployment_hash      ?string [json: 'DeploymentHash']
}

// filter statistics based on grid node status
pub struct StatsFilter {
pub mut:
	status ?string [json: 'Status']
}

// grid node info
pub struct Node {
pub:
	id                 string       [json: 'id']
	node_id            int          [json: 'nodeId']
	farm_id            int          [json: 'farmId']
	twin_id            u64          [json: 'twinId']
	country            string       [json: 'country']
	grid_version       int          [json: 'gridVersion']
	city               string       [json: 'city']
	uptime             i64          [json: 'uptime']
	created            i64          [json: 'created']
	farming_policy_id  int          [json: 'farmingPolicyId']
	updated_at         i64          [json: 'updatedAt']
	total_resources    Capacity     [json: 'total_resources']
	used_resources     Capacity     [json: 'used_resources']
	location           Location     [json: 'location']
	public_config      PublicConfig [json: 'publicConfig']
	status             string       [json: 'status']
	certification_type string       [json: 'certificationType']
	dedicated          bool         [json: 'dedicated']
	rent_contract_id   u32          [json: 'rentContractId']
	rented_by_twin_id  u32          [json: 'rentedByTwinId']
	serial_number      string       [json: 'serialNumber']
}

// grid node with nested capacity object
pub struct NodeWithNestedCapacity {
pub:
	id                string [json: 'id']
	node_id           int    [json: 'nodeId']
	farm_id           int    [json: 'farmId']
	twin_id           u64    [json: 'twinId']
	country           string [json: 'country']
	grid_version      int    [json: 'gridVersion']
	city              string [json: 'city']
	uptime            i64    [json: 'uptime']
	created           i64    [json: 'created']
	farming_policy_id int    [json: 'farmingPolicyId']
	updated_at        i64    [json: 'updatedAt']

	capacity      CapacityResult [json: 'capacity']
	location      Location       [json: 'location']
	public_config PublicConfig   [json: 'publicConfig']

	status             string [json: 'status']
	certification_type string [json: 'certificationType']
	dedicated          bool   [json: 'dedicated']
	rent_contract_id   u32    [json: 'rentContractId']
	rented_by_twin_id  u32    [json: 'rentedByTwinId']
	serial_number      string [json: 'serialNumber']
}

pub struct CapacityResult {
	total_resources Capacity [json: 'total_resources']
	used_resources  Capacity [json: 'used_resources']
}

pub struct PublicConfig {
pub:
	domain string
	gw4    string
	gw6    string
	ipv4   string
	ipv6   string
}

pub struct Capacity {
	cru u64
	sru u64
	hru u64
	mru u64
}

pub struct Location {
	country   string
	city      string
	longitude f64
	latitude  f64
}

pub struct Limit {
	size      u64
	page      u64
	ret_count bool [json: 'retCount']
	randomize bool
}

// Twin info
pub struct Twin {
pub:
	twin_id    u64    [json: 'twinId']
	account_id string [json: 'accountId']
	relay      string [json: 'relay']
	public_key string [json: 'publicKey']
}

// Counters for grid statistics
pub struct Counters {
	nodes              i64
	farms              i64
	countries          i64
	total_cru          i64            [json: 'totalCru']
	total_sru          i64            [json: 'totalSru']
	total_mru          i64            [json: 'totalMru']
	total_hru          i64            [json: 'totalHru']
	public_ips         i64            [json: 'publicIps']
	access_nodes       i64            [json: 'accessNodes']
	gateways           i64
	twins              i64
	contracts          i64
	nodes_distribution map[string]i64 [json: 'nodesDistribution']
}

pub struct NodeStatus {
	status string
}

pub struct Contract {
	contract_id u32               [json: 'contractId']
	twin_id     u32               [json: 'twinId']
	state       string            [json: 'state']
	created_at  u32               [json: 'created_at']
	type_       string            [json: 'type']
	details     string            [json: 'details']
	billing     []ContractBilling [json: 'billing']
}

pub struct ContractBilling {
	amount_billed     u64
	discount_received string
	timestamp         u64
}

pub struct Farm {
	name               string
	farm_id            int        [json: 'farmId']
	twin_id            int        [json: 'twinId']
	pricing_policy_id  int        [json: 'pricingPolicyId']
	certification_type string     [json: 'certificationType']
	stellar_address    string     [json: 'stellarAddress']
	dedicated          bool       [json: 'dedicated']
	public_ips         []PublicIP [json: 'publicIps']
}

pub struct PublicIP {
	id          string
	ip          string
	farm_id     string
	contract_id int
	gateway     string
}
