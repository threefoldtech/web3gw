# Add multisig Owner Action

Adds an owner with some threshold to a wallet.

- action name: !!eth.core.add_multisig_owner
- parameters:
  - contract_address [required]
  - target [required]
  - threshold [required]
    - The number of required confirmations to execute a transaction

## Example

```md
  !!eth.core.add_multisig_owner
      contract_address: b27a31f1b0af2946b7f582768f03239b1ec07c2c
      target: b27a31f1b0af2946b7f58276af7fc56681767523
      threshold: 10
```
