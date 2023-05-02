module nostr

pub struct Event {
	id        string
	pubkey    string
	created_at u64
	kind      int
	tags      []string
	content   string
	sig       string

	// extra map[string]any
}

[params]
pub struct TextNote {
	tags []string
	content string
}

[params]
pub struct Metadata {
	tags []string
	metadata NostrMetadata
}

pub struct NostrMetadata {
	name string
	about string
	picture string
}

[params]
pub struct DirectMessage {
	receiver string
	tags []string
	content string
}

pub struct Stall {
	id string
	name string
	description string
	currency string
	shipping []Shipping
}

pub struct StallCreateInput {
	tags []string
	stall Stall
}

pub struct Shipping {
	id string
	name string
	cost f64
	countries []string
}

pub struct Product {
	id string
	stall_id string
	name string
	description string
	images []string
	currency string
	price f64
	quantity int
	specs [][]string
}

pub struct ProductCreateInput {
	tags []string
	product Product
}