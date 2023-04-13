module stellar

import freeflowuniverse.crystallib.rpcwebsocket { RpcWsClient }

const (
	default_timeout = 500000
)

pub fn load(mut client RpcWsClient, secret string, relayUrl string) ! {
	_ := client.send_json_rpc[[]string, string]('nostr.Load', [secret, relayUrl], default_timeout)!
}

pub fn connect_to_relay(mut client RpcWsClient, relayUrl string) ! {
	_ := client.send_json_rpc[[]string, string]('nostr.ConnectToRelay', [relayUrl], default_timeout)!
}

pub fn publish_to_relays(mut client RpcWsClient, tags []string, content string) ! {
	_ := client.send_json_rpc[[]string, string]('nostr.PublishEventToRelays', [tags, content], default_timeout)!
}