# Transfer Fungible Action

> transfers the given fungible token.

- action name: !!eth.fungible.transfer
- parameters:
  - contract_address [required]
  - from [required]
  - to [required]
  - token_id [rqeuired]

## Example

```md
  !!eth.fungible.transfer
      contract_address: b27a31f1b0af2946b7f582768f03239b1ec07c2c
      from: b27a31f1b0af2946b7f58276af7fc56681767523
      to: b27a31f1b0af2946b7f58276cffc6731e42c6e1a
      token_id: 123456
```
