module model

import json

type OptionU64 = EmptyOption | u64
type OptionBool = EmptyOption | bool

[params]
pub struct FarmFilter {
pub mut:
	page               OptionU64  = EmptyOption{}
	size               OptionU64  = EmptyOption{}
	ret_count          OptionBool = EmptyOption{}
	free_ips           OptionU64  = EmptyOption{}
	total_ips          OptionU64  = EmptyOption{}
	stellar_address    string
	pricing_policy_id  OptionU64 = EmptyOption{}
	farm_id            OptionU64 = EmptyOption{}
	twin_id            OptionU64 = EmptyOption{}
	name               string
	name_contains      string
	certification_type string
	dedicated          OptionBool = EmptyOption{}
}

// serialize FarmFilter to map
pub fn (f &FarmFilter) to_map() map[string]string {
	mut m := map[string]string{}

	match f.page {
		EmptyOption {}
		u64 {
			m['page'] = f.page.str()
		}
	}
	match f.size {
		EmptyOption {}
		u64 {
			m['size'] = f.size.str()
		}
	}
	match f.ret_count {
		EmptyOption {}
		bool {
			m['ret_count'] = f.ret_count.str()
		}
	}
	match f.free_ips {
		EmptyOption {}
		u64 {
			m['free_ips'] = f.free_ips.str()
		}
	}
	match f.total_ips {
		EmptyOption {}
		u64 {
			m['total_ips'] = f.total_ips.str()
		}
	}
	if f.stellar_address != '' {
		m['stellar_address'] = f.stellar_address
	}
	match f.pricing_policy_id {
		EmptyOption {}
		u64 {
			m['pricing_policy_id'] = f.pricing_policy_id.str()
		}
	}
	match f.farm_id {
		EmptyOption {}
		u64 {
			m['farm_id'] = f.farm_id.str()
		}
	}
	match f.twin_id {
		EmptyOption {}
		u64 {
			m['twin_id'] = f.twin_id.str()
		}
	}

	if f.name != '' {
		m['name'] = f.name
	}
	if f.name_contains != '' {
		m['name_contains'] = f.name_contains
	}
	if f.certification_type != '' {
		m['certification_type'] = f.certification_type
	}
	match f.dedicated {
		EmptyOption {}
		bool {
			m['dedicated'] = f.dedicated.str()
		}
	}
	return m
}

[params]
pub struct ContractFilter {
pub mut:
	page                 OptionU64  = EmptyOption{}
	size                 OptionU64  = EmptyOption{}
	ret_count            OptionBool = EmptyOption{}
	contract_id          OptionU64  = EmptyOption{}
	twin_id              OptionU64  = EmptyOption{}
	node_id              OptionU64  = EmptyOption{}
	contract_type        string
	state                string
	name                 string
	number_of_public_ips OptionU64 = EmptyOption{}
	deployment_data      string
	deployment_hash      string
}

// serialize ContractFilter to map
pub fn (f &ContractFilter) to_map() map[string]string {
	mut m := map[string]string{}
	match f.page {
		EmptyOption {}
		u64 {
			m['page'] = f.page.str()
		}
	}
	match f.size {
		EmptyOption {}
		u64 {
			m['size'] = f.size.str()
		}
	}
	match f.ret_count {
		EmptyOption {}
		bool {
			m['ret_count'] = f.ret_count.str()
		}
	}
	match f.contract_id {
		EmptyOption {}
		u64 {
			m['contract_id'] = f.contract_id.str()
		}
	}
	match f.twin_id {
		EmptyOption {}
		u64 {
			m['twin_id'] = f.twin_id.str()
		}
	}
	match f.node_id {
		EmptyOption {}
		u64 {
			m['node_id'] = f.node_id.str()
		}
	}
	if f.contract_type != '' {
		m['contract_type'] = f.contract_type
	}
	if f.state != '' {
		m['state'] = f.state
	}
	if f.name != '' {
		m['name'] = f.name
	}
	match f.number_of_public_ips {
		EmptyOption {}
		u64 {
			m['number_of_public_ips'] = f.number_of_public_ips.str()
		}
	}
	if f.deployment_data != '' {
		m['deployment_data'] = f.deployment_data
	}
	if f.deployment_hash != '' {
		m['deployment_hash'] = f.deployment_hash
	}
	return m
}

[params]
pub struct NodeFilter {
pub mut:
	page          OptionU64  = EmptyOption{}
	size          OptionU64  = EmptyOption{}
	ret_count     OptionBool = EmptyOption{}
	free_mru      OptionU64  = EmptyOption{}
	free_sru      OptionU64  = EmptyOption{}
	free_hru      OptionU64  = EmptyOption{}
	free_ips      OptionU64  = EmptyOption{}
	city          string
	country       string
	farm_name     string
	ipv4          string
	ipv6          string
	domain        string
	status        string
	dedicated     OptionBool = EmptyOption{}
	rentable      OptionBool = EmptyOption{}
	rented_by     OptionU64  = EmptyOption{}
	available_for OptionU64  = EmptyOption{}
	farm_ids      []u64
}

// serialize NodeFilter to map
pub fn (p &NodeFilter) to_map() map[string]string {
	mut m := map[string]string{}
	match p.page {
		EmptyOption {}
		u64 {
			m['page'] = p.page.str()
		}
	}
	match p.size {
		EmptyOption {}
		u64 {
			m['size'] = p.size.str()
		}
	}
	match p.ret_count {
		EmptyOption {}
		bool {
			m['ret_count'] = p.ret_count.str()
		}
	}
	match p.free_mru {
		EmptyOption {}
		u64 {
			m['free_mru'] = p.free_mru.str()
		}
	}
	match p.free_sru {
		EmptyOption {}
		u64 {
			m['free_sru'] = p.free_sru.str()
		}
	}
	match p.free_hru {
		EmptyOption {}
		u64 {
			m['free_hru'] = p.free_hru.str()
		}
	}
	match p.free_ips {
		EmptyOption {}
		u64 {
			m['free_ips'] = p.free_ips.str()
		}
	}
	if p.status != '' {
		m['status'] = p.status
	}
	if p.city != '' {
		m['city'] = p.city
	}
	if p.country != '' {
		m['country'] = p.country
	}
	if p.farm_name != '' {
		m['farm_name'] = p.farm_name
	}
	if p.ipv4 != '' {
		m['ipv4'] = p.ipv4
	}
	if p.ipv6 != '' {
		m['ipv6'] = p.ipv6
	}
	if p.domain != '' {
		m['domain'] = p.domain
	}
	match p.dedicated {
		EmptyOption {}
		bool {
			m['dedicated'] = p.dedicated.str()
		}
	}
	match p.rentable {
		EmptyOption {}
		bool {
			m['rentable'] = p.rentable.str()
		}
	}
	match p.rented_by {
		EmptyOption {}
		u64 {
			m['rented_by'] = p.rented_by.str()
		}
	}
	match p.available_for {
		EmptyOption {}
		u64 {
			m['available_for'] = p.available_for.str()
		}
	}

	if p.farm_ids.len > 0 {
		m['farm_ids'] = json.encode(p.farm_ids).all_after('[').all_before(']')
	}
	return m
}

pub enum NodeStatus {
	all
	online
}

[params]
pub struct ResourceFilter {
pub mut:
	free_mru_gb u64
	free_sru_gb u64
	free_hru_gb u64
	free_ips    u64
}

[params]
pub struct StatFilter {
pub mut:
	status NodeStatus
}

[params]
pub struct TwinFilter {
pub mut:
	page       OptionU64  = EmptyOption{}
	size       OptionU64  = EmptyOption{}
	ret_count  OptionBool = EmptyOption{}
	twin_id    OptionU64  = EmptyOption{}
	account_id string
}

// serialize TwinFilter to map
pub fn (p &TwinFilter) to_map() map[string]string {
	mut m := map[string]string{}
	match p.page {
		EmptyOption {}
		u64 {
			m['page'] = p.page.str()
		}
	}
	match p.size {
		EmptyOption {}
		u64 {
			m['size'] = p.size.str()
		}
	}
	match p.ret_count {
		EmptyOption {}
		bool {
			m['ret_count'] = p.ret_count.str()
		}
	}
	match p.twin_id {
		EmptyOption {}
		u64 {
			m['twin_id'] = p.twin_id.str()
		}
	}
	if p.account_id != '' {
		m['account_id'] = p.account_id
	}
	return m
}
