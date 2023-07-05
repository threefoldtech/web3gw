module nostr

pub struct Event {
	id         string
	pubkey     string
	created_at u64
	kind       int
	tags       [][]string
	content    string
	sig        string
	// extra map[string]any
}

[params]
pub struct TextNote {
	tags    []string
	content string
}

[params]
pub struct Metadata {
	tags     []string
	metadata NostrMetadata
}

pub struct NostrMetadata {
	name    string
	about   string
	picture string
}

[params]
pub struct DirectMessage {
	receiver string
	tags     []string
	content  string
}

pub struct Stall {
	id          string
	name        string
	description string
	currency    string
	shipping    []Shipping
}

pub struct StallCreateInput {
	tags  []string
	stall Stall
}

pub struct Shipping {
	id        string
	name      string
	cost      f64
	countries []string
}

pub struct Product {
	id          string
	stall_id    string
	name        string
	description string
	images      []string
	currency    string
	price       f64
	quantity    int
	specs       [][]string
}

pub struct ProductCreateInput {
	tags    []string
	product Product
}

[params]
pub struct CreateChannelInput {
	tags    []string
	name    string
	about   string
	picture string
}

pub struct CreateChannelMessageInput {
	content    string
	channel_id string
	// Message ID is used for replies
	message_id string
	// Public Key of author to reply to
	public_key string
}

[params]
pub struct SubscribeChannelMessageInput {
	// Id of the channel or message to make a subscription for
	id string
}

pub struct Channel {
	name    string
	about   string
	picture string
}

pub struct RelayChannel {
	name    string
	about   string
	picture string
	relay   string
	id      string
}

pub struct FetchChannelMessageInput {
	channel_id string
}

pub struct RelayChannelMessage {
	content string
	tags    [][]string
	relay   string
	id      string
}
