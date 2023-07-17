module tfchain

[params]
pub struct Load {
pub:
	network string
	mnemonic string
}

[params]
pub struct Transfer {
pub:
	amount      u64
	destination string
}

[params]
pub struct CreateTwin {
pub:
	relay string
	pk    []byte
}

[params]
pub struct AcceptTermsAndConditions {
	link string
	hash string
}

[params]
pub struct CreateNodeContract {
	node_id              u32
	body                 string
	hash                 string
	public_ips           u32
	solution_provider_id ?u64
}

[params]
pub struct CreateRentContract {
	node_id              u32
	solution_provider_id ?u64
}

[params]
pub struct ServiceContractCreate {
	service  string
	consumer string
}

[params]
pub struct ServiceContractBill {
	contract_id     u64
	variable_amount u64
	metadata        string
}

[params]
pub struct SetServiceContractFees {
	contract_id  u64
	base_fee     u64
	variable_fee u64
}

[params]
pub struct ServiceContractSetMetadata {
	contract_id u64
	metadata    string
}

pub struct PublicIPInput {
	ip      string
	gateway string
}

[params]
pub struct CreateFarm {
	name       string
	public_ips []PublicIPInput
}

[params]
pub struct SwapToStellar {
	target_stellar_address string 
	amount u64
}

[params]
pub struct AwaitBridgedFromStellar {
pub:
	memo                            string // the memo to look for
	timeout                         int    // the timeout after which we abandon the search
	height int    // at what blockheight to start looking for the transaction
}
