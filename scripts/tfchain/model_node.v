module tfchain

pub struct Node {
pub:
	version           u32
	id           	  u32
	farm_id           u32
	twin_id           u32
	resources         Resources
	location		  Location
	public_config     OptionPublicConfig
	created           u64
	farming_policy    u32
	interfaces        []Interface
	certification     struct {
		is_diy       bool
		is_certified bool
	}
	secure_boot       bool
	vertualized       bool
	board_serial      OptionBoardSerial
	connection_price  u32
}

pub struct OptionBoardSerial{
pub:
	has_value     bool
	as_value      string
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

pub struct OptionPublicConfig{
	has_value     bool
	as_value      PublicConfig
}

pub struct PublicConfig{
pub:
	ip4 IP
	ip6 struct {
		has_value bool
		as_value  IP
	}
	domain struct {
		has_value bool
		as_value  string
	}
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
