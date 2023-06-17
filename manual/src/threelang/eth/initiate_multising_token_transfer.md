# Initiate multisig Token Transfer Action

- action name: !!eth.core.initiate_multisig_token_transfer
- parameters:
  - contract_address [required]
  - token_address [required]
  - destination [required]
  - amount [require]

## Example

```md
  !!eth.core.initiate_multisig_token_transfer
      contract_address: b27a31f1b0af2946b7f582768f03239b1ec07c2c
      token_address: b27a31f1b0af2946b7f58276af7fc56681767523
      destination: b27a31f1b0af2946b7f58276cffc6731e42c6e1a
      amount: 100
```
