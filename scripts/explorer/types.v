module explorer

pub struct Node {
	id                string       
	node_id            int          
	farm_id            int          
	twin_id           int          
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

struct PublicConfig {
	domain string 
	gw4    string 
	gw6    string 
	ipv4   string 
	ipv6   string 
}

struct Capacity {
	cru u64         
	sru u64 
	hru u64 
	mru u64 
}

struct Location {
	country   string   
	city      string   
	longitude f64
	latitude  f64
}