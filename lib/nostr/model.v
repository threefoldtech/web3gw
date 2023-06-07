module nostr

pub struct Event {
	id         string
	pubkey     string
	created_at u64
	kind       int
	tags       []string
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

pub struct CreateChannelInput {
	tags  []string
	name  string
	about string
}

pub struct CreateChannelMessageInput {
	tags    []string
	content string
	// reply_to is either the channel ID for root messages, or a message ID for replies
	reply_to string
}

pub struct SubscribeChannelMessageInput {
	// Id of the channel or message for which the reply is intended
	message_id string
}

pub struct Channel {
	name    string
	about   string
	picture string
}

pub struct RelayChannel {
	Channel
	relay string
	id    string
}

pub struct FetchChannelMessageInput {
	channel_id string
}

pub struct ChannelMessage {
	// Content of the message
	content string
	// ReplyTo is either the ID of a message to reply to, or the ID of the channel create message of the channel to post in
	// if this is a root message in the channel
	reply_to string
}

pub struct RelayChannelMessage {
	ChannelMessage
	relay string
	id    string
}
