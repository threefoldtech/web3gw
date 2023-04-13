module nostr

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

const (
	default_timeout = 500000
)

pub fn load(mut client RpcWsClient, secret string) ! {
	_ := client.send_json_rpc[[]string, string]('nostr.Load', [secret], default_timeout)!
}

pub fn connect_to_relay(mut client RpcWsClient, relayUrl string) ! {
	_ := client.send_json_rpc[[]string, string]('nostr.ConnectRelay', [relayUrl], default_timeout)!
}

pub fn connect_to_auth_relay(mut client RpcWsClient, relayUrl string) ! {
	_ := client.send_json_rpc[[]string, string]('nostr.ConnectAuthRelay', [relayUrl], default_timeout)!
}

pub fn generate_keypair(mut client RpcWsClient) !string {
	return client.send_json_rpc[[]string, string]('nostr.GenerateKeyPair', []string{}, default_timeout)!
}

pub struct RelayMessage {
	tags []string
	content string
}

pub fn publish_to_relays(mut client RpcWsClient, tags []string, content string) ! {
	_ := client.send_json_rpc[[]RelayMessage, string]('nostr.PublishEventToRelays', [RelayMessage { tags: tags, content: content }], default_timeout)!
}

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

pub fn subscribe_to_relays(mut client RpcWsClient) ! {
	_ := client.send_json_rpc[[]string, string]('nostr.SubscribeRelays', []string{}, default_timeout)!
}

pub fn get_events(mut client RpcWsClient) ![]Event {
	return client.send_json_rpc[[]string, []Event]('nostr.GetEvents', []string{}, default_timeout)!
}