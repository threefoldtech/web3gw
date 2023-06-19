# Channels Example

The channels example exposes the following functionality:

- Listing all public channels in a relay
- Reading all messages sent to a channel
- Sending a message to a channel, with the ability to mark the message as a reply for another message/user.
  
## CLI Arguments

The [channels cli](../../../../examples/nostr/channels.v) has the following arguments:

- secret: this is your secret for nostr. if none was provided, a new secret will be generated for you.
- realy_url: this is the relay URL to connect to. this defaults to `https://nostr01.grid.tf/`
- operation: this is the operation that you want to perform. must be one of `list`, `read`, or `send`.

### List Operation Arguments

There are no extra arguments for the list operation.

```sh
    v run channels.v -s "YOUR SECRET" -o list
```

### Read Operation Arguments

- channel: this is the Channel ID to read messages from. a Channel ID is the event ID of the channel creation event.

```sh
    v run channels.v -s "YOUR SECRET" -o list -channel "f27ffebc7314cbbb12ad24ff3c127ef8070f9f341b5561251c355c274984beea"
```

### Send Operation Arguments

- channel: this is the Channel ID to send the message to. a Channel ID is the event ID of the channel creation event.
- content: this is the content of the message.
- message: this is the message ID to reply to, if any.
- public_key: this is the public key of the author of the message that you want to reply to, if any.

```sh
    v run channels.v -s "YOUR SECRET" -o send -channel "f27ffebc7314cbbb12ad24ff3c127ef8070f9f341b5561251c355c274984beea" -content "Message content" -message "55d4bf31efac0bb926ca1127237f729051ca563fd74f6579e61e7c0d9ca60e0b"
```
