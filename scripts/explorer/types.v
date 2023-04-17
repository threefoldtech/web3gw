module explorer

pub struct Node {
	id                string       
	node_id            int    [json:"nodeId"]      
	farm_id            int    [json:"farmId"]      
	twin_id           int     [json:"twinId"]     
	country           string       
	grid_version       int          
	city              string       
	uptime            i64        
	created           i64        
	farming_policy_id   int          
	updated_at         i64        
	total_resources    Capacity     
	used_resources     Capacity     
	location          Location     
	public_config      PublicConfig 
	status            string       
	certification_type string       
	dedicated         bool         
	rent_contract_id    u32         
	rented_by_twin_id    u32         
	serial_number      string       
}

pub struct NodeWithNestedCapacity  {
	id                string       
	node_id            int    [json:"nodeId"]      
	farm_id            int    [json:"farmId"]      
	twin_id           int     [json:"twinId"]     
	country           string       
	grid_version       int          
	city              string       
	uptime            i64        
	created           i64        
	farming_policy_id   int          
	updated_at         i64        
	
	capacity          CapacityResult 
	location          Location       
	public_config      PublicConfig  

	status            string       
	certification_type string       
	dedicated         bool         
	rent_contract_id    u32         
	rented_by_twin_id    u32         
	serial_number      string
}

pub struct CapacityResult  {
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
	ret_count  bool
	randomize bool
}

pub struct NodeFilter  {
	status           string
	free_mru          u64
	free_hru         u64
	free_sru         u64
	total_mru         u64
	total_hru        u64
	total_sru        u64
	total_cru         u64
	country          string
	country_contains  string
	city             string
	city_contains     string
	farm_name         string
	farm_name_contains string
	farm_ids          []u64
	free_ips          u64
	ipv4             bool
	ipv6             bool
	domain           bool
	dedicated        bool
	rentable         bool
	rented           bool
	rented_by         u64
	available_for     u64
	node_id           u64
	twin_id           u64
}

pub struct FarmFilter  {
	free_ips           u64
	total_ips          u64
	stellar_address    string
	pricing_policy_id   u64
	farm_id           u64
	twin_id            u64
	name              string
	name_contains      string
	certification_type string
	dedicated         bool
}

pub struct TwinFilter  {
	twin_id    u64
	account_id string
	relay     string
	public_key string
}

pub struct Twin  {
	twin_id    u64
	account_id string
	relay     string
	public_key string
}

pub struct ContractFilter  {
	contract_id       u64
	twin_id           u64
	node_id           u64
	type_             string [json:"type"]
	state            string
	name             string
	number_of_public_ips u64
	deployment_data   string
	deployment_hash   string
}

pub struct Counters  {
	nodes             i64            
	farms             i64            
	countries         i64            
	total_cru          i64            
	total_sru         i64            
	total_mru         i64            
	total_hru         i64            
	public_ips         i64            
	access_nodes       i64            
	gateways          i64            
	twins             i64            
	contracts         i64            
	nodes_distribution map[string]i64 
}

pub struct NodeStatus  {
	status string
}

pub struct Contract  {
	contract_id u32             
	twin_id    u32             
	state      string           
	created_at  u32             
	type_       string            [json:"type"]
	details    string    
	billing    []ContractBilling
}

pub struct ContractBilling  {
	amount_billed     u64
	discount_received string
	timestamp        u64
}

pub struct Farm  {
	name              string     
	farm_id            int        [json:"farmId"]
	twin_id            int        [json:"twinId"]
	pricing_policy_id   int    
	certification_type string     
	stellar_address    string     
	dedicated         bool       
	public_ips         []PublicIP 
}

pub struct PublicIP  {
	id         string 
	ip         string 
	farm_id     string 
	contract_id int    
	gateway    string 
}