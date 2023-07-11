# Approve Hash Action

> approves a hash for the given multisig contract.

- action name: !!eth.multisig.approve_hash
- parameters:
  - contract_address [required]
    - contract address
  - hash [required]
    - transaction hash

## Example

```md
  !!eth.multisig.approve_hash
      contract_address: b27a31f1b0af2946b7f582768f03239b1ec07c2c
      hash: 0x00000000000000000000000000000000
```
