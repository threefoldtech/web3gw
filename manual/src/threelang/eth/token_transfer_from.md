# Token Transfer From Action

> transfers tokens from the given address.

- action name: !!eth.token.transfer_from
- parameters:
  - contract_address [required]
  - from [required]
  - destination [required]
  - amount [required]

## Example

```md
  !!eth.token.transfer_from
      contract_address: b27a31f1b0af2946b7f582768f03239b1ec07c2c
      from: b27a31f1b0af2946b7f58276af7fc56681767523
      destination: b27a31f1b0af2946b7f58276cffc6731e42c6e1a
      amount: 100
```
