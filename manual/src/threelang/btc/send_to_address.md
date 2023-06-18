# Send To Address Action

> Sends the passed amount to the given address with a comment if provided and returns the hash of the transaction

- action name: !!btc.account.send_to_address
- parameters:
  - address [required]
  - amount [required]
  - comment [optional]
    - is intended to be used for the purpose of the transaction.
  - comment_to [optional]
    - is intended to be used for who the transaction is being sent to.

## Example

```md
    !!btc.account.send_to_address
        address: b27a31f1b0af2946b7f582768f03239b1ec07c2c
        amount: 100
        comment: comment1
```
