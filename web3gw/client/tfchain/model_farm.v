module tfchain

// Farm type
pub struct Farm {
pub:
	version                u32
	id                     u32
	name                   string
	twin_id                u32
	pricing_policy_id      u32
	certification_type     OptionCertificationType
	public_ips             []PublicIP
	dedicated_farm         bool
	farming_policies_limit OptionFarmingPoliciesLimit
}

pub struct OptionCertificationType {
pub:
	is_not_certified bool
	is_gold          bool
}

pub struct OptionFarmingPoliciesLimit {
pub:
	has_value bool
	as_value  FarmingPoliciesLimit
}

pub struct PublicIP {
pub:
	ip          string
	gateway     string
	contract_id u64
}

pub struct FarmingPoliciesLimit {
pub:
	farming_policy_id  u32
	cu                 OptionU64
	su                 OptionU64
	end                OptionU64
	node_count         OptionU32
	node_certification bool
}

pub struct OptionU32 {
pub:
	has_value bool
	as_value  u32
}

pub struct OptionU64 {
pub:
	has_value bool
	as_value  u64
}
