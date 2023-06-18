# Transfer Action

> Transfer an amount of TFT from the loaded account to the destination.

- action name: !!stellar.account.transfer
- parameters:
  - amount [required]
  - destination [required]
  - memo [optional]

## Example

```md
    !!stellar.account.transfer
        amount: 100
        destination: b27a31f1b0af2946b7f582768f03239b1ec07c2c
        memo: my memo
```
