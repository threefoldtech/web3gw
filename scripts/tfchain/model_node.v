module tfchain

pub struct Node {
pub:
	version           u32
	node_id           u32
	farm_id           u32
	twin_id           u32
	resources         Resources
	location		  Location
	public_config     ?PublicConfig
	created           u64
	farming_policy    u32
	interfaces        []Interface
	certification     string  // DIY or Certified
	secure_boot       bool
	vertualized       bool
	board_serial      ?string
	connection_price  u32
}


pub struct Resources{
pub mut:
	hru u64
	sru u64
	cru u64
	mru u64
}

pub struct Location{
pub:
	city string
	country string
	latitude string
	longitude string
}

pub struct PublicConfig{
pub:
	ip4 IP
	ip6 ?IP
	domain ?string
}

pub struct IP{
	ip string 
	gw string
}

pub struct Interface{
	name string
	mac string
	ips []string
}
