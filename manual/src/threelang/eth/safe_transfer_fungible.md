# Safe Transfer Fungible Action

- action name: !!eth.core.safe_transfer_fungible
- parameters:
  - contract_address [required]
  - from [required]
  - to [require]
  - token_id [required]

## Example

```md
  !!eth.core.safe_transfer_fungible
      contract_address: b27a31f1b0af2946b7f582768f03239b1ec07c2c
      from: b27a31f1b0af2946b7f58276af7fc56681767523
      to: b27a31f1b0af2946b7f58276cffc6731e42c6e1a
      token_id: 123456
```
