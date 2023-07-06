module nostr

import freeflowuniverse.crystallib.actionsparser { Action }
import time

fn (mut n NostrHandler) channel(action Action) ! {
	match action.name {
		'create' {
			// create a new channel
			name := action.params.get('name')!
			about := action.params.get_default('about', '')!
			pic_url := action.params.get_default('pic_url', '')!

			channel_id := n.client.create_channel(name: name, about: about, picture: pic_url)!
			n.logger.info('Channel ID ${channel_id}')
		}
		'send' {
			// send message to channel
			channel_id := action.params.get('channel_id')!
			content := action.params.get('content')!
			message_id := action.params.get_default('reply_msg_id', '')!
			public_key := action.params.get_default('reply_usr_pk', '')!

			n.client.create_channel_message(
				channel_id: channel_id
				content: content
				message_id: message_id
				public_key: public_key
			)!
		}
		'subscribe' {
			// subscribe to channel
			channel_id := action.params.get('channel_id')!

			n.client.subscribe_channel_message(id: channel_id)!

			for {
				time.sleep(2 * time.second)
				events := n.client.get_events()!
				if events.len == 0 {
					continue
				}
				n.logger.info('Message Events: ${events}')
			}
		}
		'read' {
			// read channel messages
			channel_id := action.params.get('channel_id')!

			messages := n.client.get_channel_message(channel_id: channel_id)!
			n.logger.info('Channel Messages: ${messages}')
		}
		'list' {
			// list all channels on relay
			channels := n.client.list_channels()!
			n.logger.info('Channels: ${channels}')
		}
		else {
			return error('operation ${action.name} is not supported on nostr groups')
		}
	}
}
