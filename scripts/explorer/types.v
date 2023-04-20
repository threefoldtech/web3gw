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

pub struct NodesResult  {
	nodes      []Node 
	total_count int  
}

pub struct FarmsResult  {
	farms      []Farm 
	total_count int  
}

pub struct TwinsResult  {
	twins      []Twin 
	total_count int  
}

pub struct ContractsResult  {
	contracts      []Contract 
	total_count int  
}

pub struct Node {
	id                 string       [json: 'id']
	node_id            int          [json: 'nodeId']
	farm_id            int          [json: 'farmId']
	twin_id            int          [json: 'twinId']
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

pub struct StatsFilter {
	status string
}

pub struct NodeWithNestedCapacity {
	id                string
	node_id           int    [json: 'nodeId']
	farm_id           int    [json: 'farmId']
	twin_id           int    [json: 'twinId']
	country           string
	grid_version      int
	city              string
	uptime            i64
	created           i64
	farming_policy_id int
	updated_at        i64

	capacity      CapacityResult
	location      Location
	public_config PublicConfig

	status             string
	certification_type string
	dedicated          bool
	rent_contract_id   u32
	rented_by_twin_id  u32
	serial_number      string
}

pub struct CapacityResult {
	total Capacity
	used  Capacity
}

pub struct PublicConfig {
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
	ret_count bool
	randomize bool
}

pub struct NodeFilter {
	status             string
	free_mru           u64    [json: 'freeMru']
	free_hru           u64    [json: 'freeHru']
	free_sru           u64    [json: 'freeSru']
	total_mru          u64    [json: 'totalMru']
	total_hru          u64    [json: 'totalHru']
	total_sru          u64    [json: 'totalSru']
	total_cru          u64    [json: 'totalCru']
	country            string
	country_contains   string [json: 'countryContains']
	city               string
	city_contains      string [json: 'cityContains']
	farm_name          string [json: 'farmName']
	farm_name_contains string
	farm_ids           []u64  [json: 'farmIds']
	free_ips           u64    [json: 'freeIps']
	ipv4               bool
	ipv6               bool
	domain             bool
	dedicated          bool
	rentable           bool
	rented             bool
	rented_by          u64    [json: 'rentedBy']
	available_for      u64    [json: 'availableFor']
	node_id            u64    [json: 'nodeId']
	twin_id            u64
}

pub struct FarmFilter {
	free_ips           u64    [json: 'freeIps']
	total_ips          u64    [json: 'totalIps']
	stellar_address    string [json: 'stellarAddress']
	pricing_policy_id  u64
	farm_id            u64    [json: 'farmId']
	twin_id            u64    [json: 'twinId']
	name               string
	name_contains      string [json: 'nameContains']
	certification_type string [json: 'certificationType']
	dedicated          bool
}

pub struct TwinFilter {
	twin_id    u64    [json: 'TwinID']
	account_id string [json: 'AccountID']
	relay      string
	public_key string [json: 'PublicKey']
}

pub struct Twin {
	twin_id    u64    [json: 'twinId']
	account_id string [json: 'accountId']
	relay      string [json: 'relay']
	public_key string [json: 'publicKey']
}

pub struct ContractFilter {
	contract_id          u64    [json: 'ContractID']
	twin_id              u64    [json: 'TwinID']
	node_id              u64    [json: 'NodeID']
	type_                string [json: 'Type']
	state                string [json: 'State']
	name                 string [json: 'Name']
	number_of_public_ips u64    [json: 'NumberOfPublicIps']
	deployment_data      string [json: 'DeploymentData']
	deployment_hash      string [json: 'DeploymentHash']
}

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
	pricing_policy_id  int
	certification_type string
	stellar_address    string     [json: 'stellarAddress']
	dedicated          bool
	public_ips         []PublicIP
}

pub struct PublicIP {
	id          string
	ip          string
	farm_id     string
	contract_id int
	gateway     string
}
