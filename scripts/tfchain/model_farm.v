module tfchain

// Farm type
pub struct Farm{
	version              u32
	id                   u32
	name                 string
	twin_id              u32
	pricing_policy_id    u32
	certification_type   string // NotCertified or Gold
	public_ips           []PublicIP
	dedicated_farm       bool
	farming_policies_limit ?FarmingPolicyLimit
}

pub struct PublicIP{
pub:
	ip string
	gateway string
	contract_id u64
}

pub struct FarmingPolicyLimit{
pub:
	farming_policy_id u32
	cu ?u64
	su ?u64
	end ?u64
	node_count ?u32
	node_certification bool
}