module main

import freeflowuniverse.crystallib.rpcwebsocket
import threefoldtech.threebot.nostr { CreateChannelMessageInput, FetchChannelMessageInput, NostrClient }
import flag { FlagParser }
import log { Logger }
import os
import time

const (
	default_server_address = 'ws://127.0.0.1:8080'
)

fn create_channel(mut fp FlagParser, mut nostr_client NostrClient, mut logger Logger) ! {
	fp.usage_example('create channel_name [-a about] [-p picture_url]')
	fp.description('Creates a new channel on the specified relay')
	fp.limit_free_args_to_exactly(1)!

	name := fp.args[0] or { return error('failed to get channel name') }
	about := fp.string('about', `d`, '', 'Channel description')
	pic_url := fp.string('picture', `p`, '', 'Channel picture URL')

	_ := fp.finalize() or {
		println(fp.usage())
		return error('${err}')
	}

	channel_id := nostr_client.create_channel(name: name, about: about, picture: pic_url)!
	logger.info('Channel ID  ${channel_id}')
}

fn list_channels(mut fp FlagParser, mut nostr_client NostrClient, mut logger Logger) ! {
	fp.usage_example('list')
	fp.description('Lists all channels on the specified relay')
	_ := fp.finalize()!

	channels := nostr_client.list_channels()!
	logger.info('channels ${channels}')
}

fn read_messages(mut fp FlagParser, mut nostr_client NostrClient, mut logger Logger) ! {
	fp.usage_example('read')
	fp.description('Reads all messages on a specific channel')

	channel_id := fp.string_opt('channel', `c`, 'Channel ID to read messages from') or {
		println(fp.usage())
		return error('${err}')
	}
	_ := fp.finalize() or {
		println(fp.usage())
		return error('${err}')
	}

	messages := nostr_client.get_channel_message(FetchChannelMessageInput{
		channel_id: channel_id
	})!
	logger.info('channel messages:\n${messages}')
}

fn send_message(mut fp FlagParser, mut nostr_client NostrClient, mut logger Logger) ! {
	fp.usage_example('send')
	fp.description('Sends a message to the specified channel')

	channel_id := fp.string_opt('channel', `c`, 'Channel ID to send message to') or {
		println(fp.usage())
		return error('${err}')
	}
	content := fp.string_opt('content', `t`, 'Text message content') or {
		println(fp.usage())
		return error('${err}')
	}
	message_id := fp.string('message', `m`, '', 'Message ID to reply to')
	public_key := fp.string('public_key', `p`, '', 'Public Key of user to reply to')
	_ := fp.finalize() or {
		println(fp.usage())
		return error('${err}')
	}

	nostr_client.create_channel_message(CreateChannelMessageInput{
		content: content
		channel_id: channel_id
		message_id: message_id
		public_key: public_key
	})!
}

fn subscribe_channel(mut fp FlagParser, mut nostr_client NostrClient, mut logger Logger) ! {
	fp.usage_example('subscribe')
	fp.description('Subscribes to the specified channel')

	channel_id := fp.string_opt('channel', `c`, 'Channel ID to subscribe to') or {
		println(fp.usage())
		return error('${err}')
	}

	_ := fp.finalize() or {
		println(fp.usage())
		return error('${err}')
	}

	nostr_client.subscribe_channel_message(id: channel_id)!

	for {
		time.sleep(2 * time.second)
		events := nostr_client.get_events()!
		if events.len == 0 {
			continue
		}
		logger.info('Events: ${events}')
	}
}

fn main() {
	mut logger := Logger(&log.Log{
		level: .info
	})

	mut fp := flag.new_flag_parser(os.args)
	fp.application('Welcome to the web3_proxy client. The web3_proxy client allows you to execute all remote procedure calls that the web3_proxy server can handle.')
	fp.description('')
	fp.skip_executable()
	fp.allow_unknown_args()

	address := fp.string('address', `a`, '${default_server_address}', 'The address of the web3_proxy server to connect to.')
	debug_log := fp.bool('debug', 0, false, 'By setting this flag the client will print debug logs too.')
	mut secret := fp.string('secret', `s`, '', 'The secret to use for nostr. if none was provided, a new secret is generated')
	relay_url := fp.string('relay_url', `u`, 'https://nostr01.grid.tf/', 'Relay URL to connect to.')
	operation := fp.string('operation', `o`, 'list', 'Command to run on nostr channels')
	remainig_args := fp.finalize() or {
		logger.error('${err}')
		exit(1)
	}

	if debug_log {
		logger.set_level(.debug)
	}

	mut myclient := rpcwebsocket.new_rpcwsclient(address, &logger) or {
		logger.error('Failed creating rpc websocket client: ${err}')
		exit(1)
	}

	_ := spawn myclient.run()

	mut nostr_client := nostr.new(mut myclient)
	if secret == '' {
		secret = nostr_client.generate_keypair()!
	}
	nostr_client.load(secret)!
	nostr_client.connect_to_relay(relay_url) or {
		logger.error('${err}')
		exit(1)
	}

	match operation {
		'create' {
			mut new_fp := flag.new_flag_parser(remainig_args)
			create_channel(mut new_fp, mut nostr_client, mut logger) or {
				logger.error('${err}')
				exit(1)
			}
		}
		'list' {
			mut new_fp := flag.new_flag_parser(remainig_args)
			list_channels(mut new_fp, mut nostr_client, mut logger) or {
				logger.error('${err}')
				exit(1)
			}
		}
		'read' {
			mut new_fp := flag.new_flag_parser(remainig_args)
			read_messages(mut new_fp, mut nostr_client, mut logger) or {
				logger.error('${err}')
				exit(1)
			}
		}
		'send' {
			mut new_fp := flag.new_flag_parser(remainig_args)
			send_message(mut new_fp, mut nostr_client, mut logger) or {
				logger.error('${err}')
				exit(1)
			}
		}
		'subscribe' {
			mut new_fp := flag.new_flag_parser(remainig_args)
			subscribe_channel(mut new_fp, mut nostr_client, mut logger) or {
				logger.error('${err}')
				exit(1)
			}
		}
		else {
			logger.error('invalid operation ${operation}')
			exit(1)
		}
	}
}
