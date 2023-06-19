# Set Fungible Approval For All Action

> sets the fungible approval for all the given fungible token.

- action name: !!eth.fungible.set_approval_for_all
- parameters:
  - contract_address [required]
  - from [required]
  - to [require]
  - approved [required]

## Example

```md
  !!eth.fungible.set_approval_for_all
      contract_address: b27a31f1b0af2946b7f582768f03239b1ec07c2c
      from: b27a31f1b0af2946b7f58276af7fc56681767523
      to: b27a31f1b0af2946b7f58276cffc6731e42c6e1a
      approved: true
```
