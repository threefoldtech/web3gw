module nostr

pub struct Event {
	id         string     // 32-bytes lowercase hex-encoded sha256 of the serialized event data
	pubkey     string     // 32-bytes lowercase hex-encoded public key of the event creator
	created_at u64        // unix timestamp in seconds
	kind       int        // event kind
	tags       [][]string // event tags
	content    string     // arbitrary string
	sig        string     // 64-bytes hex of the signature of the sha256 hash of the serialized event data, which is the same as the "id" field
}

[params]
pub struct TextNote {
	tags    []string // event tags
	content string // text note content
}

[params]
pub struct Metadata {
	tags     []string // event tags
	metadata NostrMetadata // metadata details
}

pub struct NostrMetadata {
	name    string // user name
	about   string // user about
	picture string  // user picture URL
}

[params]
pub struct DirectMessage {
	receiver string // 32-bytes lowercase hex-encoded public key of the message receiver
	tags     []string // event tags
	content  string // message content
}

pub struct Stall {
	id          string // UUID generated by the merchant.
	name        string // stall name
	description string // stall description
	currency    string // currency used in the stall
	shipping    []Shipping // shipping details for the stall
}

pub struct StallCreateInput {
	tags  []string // event tags
	stall Stall // stall details
}

pub struct Shipping {
	id        string // UUID of the shipping zone, generated by the merchant
	name      string // zone name
	cost      f64 // cost for shipping. The currency is defined at the stall struct
	countries []string // countries included in this zone
}

pub struct Product {
	id          string // UUID generated by the merchant.
	stall_id    string // UUID of the stall to which this product belong to
	name        string // product name
	description string // product description
	images      []string // array of image URLs,
	currency    string // currency used
	price       f64 // cost of product
	quantity    int // available items
	specs       [][]string // key value pairs of product specs
}

pub struct ProductCreateInput {
	tags    []string // event tags
	product Product // product details
}

[params]
pub struct CreateChannelInput {
	tags    []string // event tags
	name    string // new channel name
	about   string // channel description
	picture string // channel picture URL
}

pub struct RelayChannel {
	name    string // channel name
	about   string // channel description
	picture string // channel picture URL
	relay   string // relay URL
	id      string // event ID of channel creation
}

[params]
pub struct CreateChannelMessageInput {
	content    string // message content
	channel_id string // event id of the channel to send the message to
	message_id string // event id of message to reply to
	public_key string // Public Key of author to reply to
}

[params]
pub struct SubscribeChannelMessageInput {
	id string // Id of the channel or message to make a subscription for
}

[params]
pub struct FetchChannelMessageInput {
	channel_id string // channel ID to fetch messages from
}

pub struct RelayChannelMessage {
	content string // message content
	tags    [][]string // event tags
	relay   string // relay URL
	id      string // message event ID
}
