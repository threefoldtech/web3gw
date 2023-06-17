# Remove multisig Owner Action

- action name: !!eth.core.remove_multisig_owner
- parameters:
  - contract_address [required]
  - target [required]
  - threshold [required]

## Example

```md
  !!eth.core.remove_multisig_owner
      contract_address: b27a31f1b0af2946b7f582768f03239b1ec07c2c
      target: b27a31f1b0af2946b7f58276af7fc56681767523
      threshold: 123456789
```
