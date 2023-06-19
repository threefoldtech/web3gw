# Approve Token Spending Action

> approves token spending for the given address.

- action name: !!eth.token.approve_token_spending
- parameters:
  - contract_address [required]
    - contract address
  - spender [required]
    - spender address
  - amount [required]
    - spending amount

## Example

```md
  !!eth.token.approve_token_spending
      contract_address: b27a31f1b0af2946b7f582768f03239b1ec07c2c
      spender: b27a31f1b0af2946b7f58276af7fc56681767523
```
