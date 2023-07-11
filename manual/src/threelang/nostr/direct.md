# Direct Action

Nostr Direct Actions allow you to:

- Send a direct message
- Receive direct messages

## Send Direct Message

- action name: !!nostr.direct.send
- parameters:
  - receiver [required]
    - this is the public key of the user you are sending the message to.
  - content [required]
    - this is the content of the direct message

- Example:

```md
    !!nostr.direct.send
        receiver: 8484e72068a9e8d2145f87d11db30030f62ba6914227e8ab9be260515360ce30
        content: content of my direct message
```

## Receive Direct Messages

- action name: !!nostr.direct.read
- paramters:
  - id [optional]
    - subscription id. if not provided, a new subscription is created and the subscription id is printed in logs.
  - count [optional]
    - count of messages to read. defaults to `10`

- Example:

```md
    !!nostr.direct.read
        count: 20
```
