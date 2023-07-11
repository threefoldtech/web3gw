# Add multisig Owner Action

> adds a new owner to the given multisig contract.

- action name: !!eth.multisig.add_owner
- parameters:
  - contract_address [required]
  - target [required]
  - threshold [required]
    - The number of required confirmations to execute a transaction

## Example

```md
  !!eth.multisig.add_owner
      contract_address: b27a31f1b0af2946b7f582768f03239b1ec07c2c
      target: b27a31f1b0af2946b7f58276af7fc56681767523
      threshold: 10
```
