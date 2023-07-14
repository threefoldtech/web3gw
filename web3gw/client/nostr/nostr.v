module nostr

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

const (
	default_timeout = 500000
)

[openrpc: exclude]
[noinit]
pub struct NostrClient {
mut:
	client &RpcWsClient
}

[openrpc: exclude]
pub fn new(mut client RpcWsClient) NostrClient {
	return NostrClient{
		client: &client
	}
}

// Loads the nostr secret
pub fn (mut n NostrClient) load(secret string) ! {
	_ := n.client.send_json_rpc[[]string, string]('nostr.Load', [secret], nostr.default_timeout)!
}

// Gets the nostr encoded id
pub fn (mut n NostrClient) get_id() !string {
	return n.client.send_json_rpc[[]string, string]('nostr.GetId', []string{}, nostr.default_timeout)!
}

// Gets the nostr public key
pub fn (mut n NostrClient) get_public_key() !string {
	return n.client.send_json_rpc[[]string, string]('nostr.GetPublicKey', []string{},
		nostr.default_timeout)!
}

// Connects to the provided authenticated relay
pub fn (mut n NostrClient) connect_to_auth_relay(relay_url string) ! {
	_ := n.client.send_json_rpc[[]string, string]('nostr.ConnectToAuthRelay', [
		relay_url,
	], nostr.default_timeout)!
}

// Connects to the given relay url
pub fn (mut n NostrClient) connect_to_relay(relay_url string) ! {
	_ := n.client.send_json_rpc[[]string, string]('nostr.ConnectToRelay', [relay_url],
		nostr.default_timeout)!
}

// Generates a new keypair and return the secret
pub fn (mut n NostrClient) generate_keypair() !string {
	return n.client.send_json_rpc[[]string, string]('nostr.GenerateKeyPair', []string{},
		nostr.default_timeout)!
}

// Publishes a text note to the relay
pub fn (mut n NostrClient) publish_text_note(args TextNote) ! {
	_ := n.client.send_json_rpc[[]TextNote, string]('nostr.PublishTextNote', [args], nostr.default_timeout)!
}

// Publishes metadata to the relay
pub fn (mut n NostrClient) publish_metadata(args Metadata) ! {
	_ := n.client.send_json_rpc[[]Metadata, string]('nostr.PublishMetadata', [args], nostr.default_timeout)!
}

// Publishes a direct message to a receiver
pub fn (mut n NostrClient) publish_direct_message(args DirectMessage) ! {
	_ := n.client.send_json_rpc[[]DirectMessage, string]('nostr.PublishDirectMessage',
		[args], nostr.default_timeout)!
}

// Subscribes to text notes on all relays
pub fn (mut n NostrClient) subscribe_text_notes() ! {
	_ := n.client.send_json_rpc[[]string, string]('nostr.SubscribeTextNotes', []string{},
		nostr.default_timeout)!
}

// Subscribes to direct messages on all relays and decrypts them
pub fn (mut n NostrClient) subscribe_direct_messages() ! {
	_ := n.client.send_json_rpc[[]string, string]('nostr.SubscribeDirectMessages', []string{},
		nostr.default_timeout)!
}

// Subscribes to stall creation on all relays
pub fn (mut n NostrClient) subscribe_stall_creation() ! {
	_ := n.client.send_json_rpc[[]string, string]('nostr.SubscribeStallCreation', []string{},
		nostr.default_timeout)!
}

// Subscribes to relays for product creation
pub fn (mut n NostrClient) subscribe_product_creation() ! {
	_ := n.client.send_json_rpc[[]string, string]('nostr.SubscribeProductCreation', []string{},
		nostr.default_timeout)!
}

// Closes the subscription with the given id
pub fn (mut n NostrClient) close_subscription(id string) ! {
	_ := n.client.send_json_rpc[[]string, string]('nostr.CloseSubscription', [id], nostr.default_timeout)!
}

// Returns all subscription ids
pub fn (mut n NostrClient) get_subscription_ids() ![]string {
	return n.client.send_json_rpc[[]string, []string]('nostr.GetSubscriptionIds', []string{},
		nostr.default_timeout)!
}

// Returns all the events for all subscriptions
pub fn (mut n NostrClient) get_events() ![]Event {
	return n.client.send_json_rpc[[]string, []Event]('nostr.GetEvents', []string{}, nostr.default_timeout)!
}

// Publishes a new stall to the relay
pub fn (mut n NostrClient) publish_stall(args StallCreateInput) ! {
	_ := n.client.send_json_rpc[[]StallCreateInput, string]('nostr.PublishStall', [
		args,
	], nostr.default_timeout)!
}

// Publishes a new product to the relay
pub fn (mut n NostrClient) publish_product(args ProductCreateInput) ! {
	_ := n.client.send_json_rpc[[]ProductCreateInput, string]('nostr.PublishProduct',
		[args], nostr.default_timeout)!
}

// Creates a new channel
pub fn (mut n NostrClient) create_channel(args CreateChannelInput) !string {
	return n.client.send_json_rpc[[]CreateChannelInput, string]('nostr.CreateChannel',
		[
		args,
	], nostr.default_timeout)!
}

// Subsribes to channel creation events
pub fn (mut n NostrClient) subscribe_channel_creation() !string {
	return n.client.send_json_rpc[[]string, string]('nostr.SubscribeChannelCreation',
		[]string{}, nostr.default_timeout)!
}

// Creates a new channel message event
pub fn (mut n NostrClient) create_channel_message(args CreateChannelMessageInput) ! {
	_ := n.client.send_json_rpc[[]CreateChannelMessageInput, string]('nostr.CreateChannelMessage',
		[
		args,
	], nostr.default_timeout)!
}

// Subscribes to channel messages or message replies, depending on the id provided
pub fn (mut n NostrClient) subscribe_channel_message(args SubscribeChannelMessageInput) !string {
	return n.client.send_json_rpc[[]SubscribeChannelMessageInput, string]('nostr.SubscribeChannelMessage',
		[
		args,
	], nostr.default_timeout)!
}

// Lists all channels on the connected relays
pub fn (mut n NostrClient) list_channels() ![]RelayChannel {
	return n.client.send_json_rpc[[]string, []RelayChannel]('nostr.ListChannels', []string{},
		nostr.default_timeout)!
}

// Returns all messages sent to a channel
pub fn (mut n NostrClient) get_channel_message(args FetchChannelMessageInput) ![]RelayChannelMessage {
	return n.client.send_json_rpc[[]FetchChannelMessageInput, []RelayChannelMessage]('nostr.GetChannelMessages',
		[
		args,
	], nostr.default_timeout)!
}
