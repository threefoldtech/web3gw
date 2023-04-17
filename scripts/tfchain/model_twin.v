module tfchain

pub struct Twin{
pub:
	id u32
	account_id string
	relay struct {
		has_value bool
		as_value string
	}
	entities []EntityProof
	pk struct {
		has_value bool
		as_value string
	}
}

pub struct EntityProof{
pub:
	entity_id u32
	signature string
}
