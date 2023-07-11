# Is Approved Action

> returns true if the given hash is approved for the given multisig contract.

- action name: !!eth.multisig.is_approved
- parameters:
  - contract_address [required]
  - hash [required]

## Example

```md
  !!eth.multisig.is_approved
      contract_address: b27a31f1b0af2946b7f582768f03239b1ec07c2c
      hash: 0x00000000000000000000000000000000
```
