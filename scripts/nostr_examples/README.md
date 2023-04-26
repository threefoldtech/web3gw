# Nostr examples

## Chat example

### Relay

Install a local relay first:

```
git clone git@github.com:scsibug/nostr-rs-relay.git
cd nostr-rs-relay
cargo build -r
```

Open the `config.toml` in `nostr-rs-relay` and change the port from `8080` to `8081`.

Now run the relay:

```
./target/release/nostr-rs-relay
```

### Chat

To run the chat example, open 2 terminal windows. In one window run the consumer (this will start streaming for direct messages on a nostr relay):

Use following secret (b4fc308f04cb3dc80c2caf18dadc42ba4a7dbdbc1471a2e40fa091ac0e96d711) or generate your own

```
v -cg run main_nostr_chat_consumer.v -s b4fc308f04cb3dc80c2caf18dadc42ba4a7dbdbc1471a2e40fa091ac0e96d711
```

In a second pane, run the publisher:

Receiver: (2bd6ab7f8a9c8c1a337611786aa06a8ab9be0a03bd0ab9417d190109be9cc9a7)
Secret: (fb2184cac1bfa5694977d289b698afefc1d012d978e72de6b433dba1cd54ec3d)

```
 v -cg run main_nostr_chat_publisher.v -r 2bd6ab7f8a9c8c1a337611786aa06a8ab9be0a03bd0ab9417d190109be9cc9a7 -s fb2184cac1bfa5694977d289b698afefc1d012d978e72de6b433dba1cd54ec3d
```

You should now see a message in the consumer pane.