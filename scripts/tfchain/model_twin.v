module tfchain

pub struct Twin{
pub:
	id u32
	account string
	relay ?string
	entities []EntityProof
	pk ?string
}

pub struct EntityProof{
pub:
	entity_id u32
	signature string
}