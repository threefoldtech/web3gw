# Token Transfer Action

> transfers tokens to the given address.

- action name: !!eth.token.transfer
- parameters:
  - contract_address [required]
  - destination [required]
  - amount [required]

## Example

```md
  !!eth.token.transfer
      contract_address: b27a31f1b0af2946b7f582768f03239b1ec07c2c
      destination: b27a31f1b0af2946b7f58276cffc6731e42c6e1a
      amount: 100
```
