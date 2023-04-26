module nostr

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

const (
	default_timeout = 500000
)

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

[noinit]
pub struct NostrClient {
mut:
	client &RpcWsClient
}

pub fn new(mut client RpcWsClient) NostrClient {
	return NostrClient{
		client: &client
	}
}

// load the nostr client with a secret
pub fn (mut n NostrClient) load(secret string) ! {
	_ := n.client.send_json_rpc[[]string, string]('nostr.Load', [secret], default_timeout)!
}

// connect to a relay given a url
pub fn (mut n NostrClient) connect_to_relay(relay_url string) ! {
	_ := n.client.send_json_rpc[[]string, string]('nostr.ConnectRelay', [relay_url], default_timeout)!
}

// connect to the authenticated relay
pub fn (mut n NostrClient) connect_to_auth_relay(relay_url string) ! {
	_ := n.client.send_json_rpc[[]string, string]('nostr.ConnectAuthRelay', [relay_url], default_timeout)!
}

// generate a keypair
pub fn (mut n NostrClient) generate_keypair() !string {
	return n.client.send_json_rpc[[]string, string]('nostr.GenerateKeyPair', []string{}, default_timeout)!
}

// get the nostr encoded id
pub fn (mut n NostrClient) get_id() !string {
	return n.client.send_json_rpc[[]string, string]('nostr.GetId', []string{}, default_timeout)!
}

// publish a text note to the relay
pub fn (mut n NostrClient) publish_text_note(args TextNote) ! {
	_ := n.client.send_json_rpc[[]TextNote, string]('nostr.PublishTextNote', [args], default_timeout)!
}

// publish metadata to the relay
pub fn (mut n NostrClient) publish_metadata(args Metadata) ! {
	_ := n.client.send_json_rpc[[]Metadata, string]('nostr.PublishMetadata', [args], default_timeout)!
}

// publish a direct message to the relay given a receiver
pub fn (mut n NostrClient) publish_direct_message(args DirectMessage) ! {
	_ := n.client.send_json_rpc[[]DirectMessage, string]('nostr.PublishDirectMessage', [args], default_timeout)!
}

// subscribe to the relay for messages and events
pub fn (mut n NostrClient) subscribe_to_relays() ! {
	_ := n.client.send_json_rpc[[]string, string]('nostr.SubscribeRelays', []string{}, default_timeout)!
}

// get all the events for the subscriptions
pub fn (mut n NostrClient) get_events() ![]Event {
	return n.client.send_json_rpc[[]string, []Event]('nostr.GetEvents', []string{}, default_timeout)!
}

// close a subscription given an id
pub fn (mut n NostrClient) close_subscription(id string) ! {
	_ := n.client.send_json_rpc[[]string, string]('nostr.CloseSubscription', [id], default_timeout)!
}

// get all the subscription ids
pub fn (mut n NostrClient) get_subscription_ids() ![]string {
	return n.client.send_json_rpc[[]string, []string]('nostr.GetSubscriptionIds', []string{}, default_timeout)!
}