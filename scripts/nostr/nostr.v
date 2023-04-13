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