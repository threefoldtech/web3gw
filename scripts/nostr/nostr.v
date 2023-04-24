module nostr

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

const (
	default_timeout = 500000
)

[params]
pub struct RelayMessage {
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

pub fn (mut n NostrClient) load(secret string) ! {
	_ := n.client.send_json_rpc[[]string, string]('nostr.Load', [secret], default_timeout)!
}

pub fn (mut n NostrClient) connect_to_relay(relay_url string) ! {
	_ := n.client.send_json_rpc[[]string, string]('nostr.ConnectRelay', [relay_url], default_timeout)!
}

pub fn (mut n NostrClient) connect_to_auth_relay(relay_url string) ! {
	_ := n.client.send_json_rpc[[]string, string]('nostr.ConnectAuthRelay', [relay_url], default_timeout)!
}

pub fn (mut n NostrClient) generate_keypair() !string {
	return n.client.send_json_rpc[[]string, string]('nostr.GenerateKeyPair', []string{}, default_timeout)!
}

pub fn (mut n NostrClient) publish_to_relays(args RelayMessage) ! {
	_ := n.client.send_json_rpc[[]RelayMessage, string]('nostr.PublishEventToRelays', [args], default_timeout)!
}

pub fn (mut n NostrClient) subscribe_to_relays() ! {
	_ := n.client.send_json_rpc[[]string, string]('nostr.SubscribeRelays', []string{}, default_timeout)!
}

pub fn (mut n NostrClient) get_events() ![]Event {
	return n.client.send_json_rpc[[]string, []Event]('nostr.GetEvents', []string{}, default_timeout)!
}

pub fn (mut n NostrClient) close_subscription(id string) ! {
	_ := n.client.send_json_rpc[[]string, string]('nostr.CloseSubscription', [id], default_timeout)!
}

pub fn (mut n NostrClient) get_subscription_ids() ![]string {
	return n.client.send_json_rpc[[]string, []string]('nostr.GetSubscriptionIds', []string{}, default_timeout)!
}