module tfchain

pub struct Contract {
pub mut:
	version              u32
	state                ContractState
	contract_id          u64
	twin_id              u32
	contract_type        ContractType
	solution_provider_id OptionU64
}

pub struct ContractState {
pub mut:
	is_created bool
	is_deleted bool
	as_deleted struct {
		is_canceled_by_user bool
		is_out_of_funds     bool
	}

	is_grace_period             bool
	as_grace_period_blocknumber u64
}

pub struct NodeContract {
pub mut:
	node            u32
	deployment_hash []byte
	deployment_data string
	public_ips      []PublicIP
}

pub struct NameContract {
pub mut:
	name string
}

pub struct RentContract {
pub mut:
	node u32
}

pub struct ContractType {
pub mut:
	is_node_contract bool
	node_contract    NodeContract
	is_name_contract bool
	name_contract    NameContract
	is_rent_contract bool
	rent_contract    RentContract
}
