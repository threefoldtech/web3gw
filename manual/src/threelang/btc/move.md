# Move Action

> Moves specified amount from one account in your wallet to another. Only funds with the default number of minimum confirmations will be used. A comment can also be added to the transaction.

- action name: !!btc.wallet.move
- parameters:
  - from_account [required]
  - to_account [required]
  - amount [required]
  - min_confirmations [required]
  - comment [optional]

## Example

```md
    !!btc.wallet.move
        from_account: acc1
        to_account: acc2
        amount: 100
        min_confirmations: 5
        comment: my comment
```
