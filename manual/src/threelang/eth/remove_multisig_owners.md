# Remove multisig Owner Action

> removes an owner from the given multisig contract.

- action name: !!eth.multisig.remove_owner
- parameters:
  - contract_address [required]
  - target [required]
  - threshold [required]

## Example

```md
  !!eth.multisig.remove_owner
      contract_address: b27a31f1b0af2946b7f582768f03239b1ec07c2c
      target: b27a31f1b0af2946b7f58276af7fc56681767523
      threshold: 123456789
```
