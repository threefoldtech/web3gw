module nostr

import freeflowuniverse.crystallib.actionsparser { Action }
import time

fn (mut n NostrHandler) direct(action Action) ! {
	match action.name {
		'send' {
			// send direct message
			receiver := action.params.get('receiver')!
			content := action.params.get('content')!

			n.client.publish_direct_message(
				receiver: receiver
				content: content
			)!
		}
		'read' {
			// reads and subscribes to direct messages
			n.client.subscribe_to_direct_messages()!

			for {
				time.sleep(2 * time.second)
				events := n.client.get_events()!
				if events.len == 0 {
					continue
				}
				n.logger.info('Direct Message Events: ${events}')
			}
		}
		else {
			return error('operation ${action.name} is not supported on nostr direct messages')
		}
	}
}
