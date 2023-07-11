# Channel Action

Nostr Channel Actions allow you to:

- Create new channels
- List all channels in a relay.
- Send a message to a channel.
- Read messages in a channel.
- Subscribe to a channel.

## Create Channel

- action name: !!nostr.channel.create
- parameters:
  - name: [required]
    - new channel name
  - description: [optional]
    - channel description
  - picture: [optional]
    - channel picture URL

- Example:

```md
    !!nostr.channel.create
        name: my_new_channel_name
        description: this is my new channel
        picture: https://www.my_channel_pic_url.com
```

## List Channels

- action name: !!nostr.channel.list
- parameters:

- Example:

```md
    !!nostr.channel.list
```

## Send Channels

- action name: !!nostr.channel.send
- parameters:
  - channel [required]
    - this is the Channel ID to send the message to. a Channel ID is the event ID of the channel creation event.
  - content [required]
    - this is the content of the message.
  - reply_to [optional]
    - this is the message ID to reply to, if any.
  - public_key_author [optional]
    - this is the public key of the author of the message that you want to reply to, if any.

- Example:

```md
    !!nostr.channel.send
        channel: f27ffebc7314cbbb12ad24ff3c127ef8070f9f341b5561251c355c274984beea
        content: my message content
        reply_to: 55d4bf31efac0bb926ca1127237f729051ca563fd74f6579e61e7c0d9ca60e0b
```

## Read Channel Messages

- action name: !!nostr.channel.read
- parameters:
  - channel [required]
    - this is the Channel ID to read messages from. a Channel ID is the event ID of the channel creation event.

- Example:

```md
    !!nostr.channel.read
        channel: f27ffebc7314cbbb12ad24ff3c127ef8070f9f341b5561251c355c274984beea
```

## Subscribe to Channel

- action name: !!nostr.channel.subscribe
- parameters:
  - channel [required]
    - this is the Channel ID to subscribe to. a Channel ID is the event ID of the channel creation event.

- Example:

```md
    !!nostr.channel.subscribe
        channel: f27ffebc7314cbbb12ad24ff3c127ef8070f9f341b5561251c355c274984beea
```

## Read Channel Subscription messages

- action name: !!nostr.channel.read_sub
- parameters:
  - channel [required]
    - this is the Channel ID to read messages from. a Channel ID is the event ID of the channel creation event.
  - id [optional]
    - subscription id. if not provided, a new subscription is created and the subscription id is printed in logs.
  - count [optional]
    - number of messages to read. defaults to `10`

- Example

```md
    !!nostr.channel.read_sub
        id: "SUBSCRIPTION ID" 
        count: 30
```
