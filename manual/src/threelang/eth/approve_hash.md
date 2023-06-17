# Approve Hash Action

Approves a transaction hash

- action name: !!eth.core.approve_hash
- parameters:
  - contract_address [required]
    - contract address
  - hash [required]
    - transaction hash

## Example

```md
  !!eth.core.approve_hash
      contract_address: b27a31f1b0af2946b7f582768f03239b1ec07c2c
      hash: 0x00000000000000000000000000000000
```
