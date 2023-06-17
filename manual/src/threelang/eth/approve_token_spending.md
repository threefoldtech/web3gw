# Approve Token Spending Action

Approves spending from a token contract with a limit

- action name: !!eth.core.approve_token_spending
- parameters:
  - contract_address [required]
    - contract address
  - spender [required]
    - spender address
  - amount [required]
    - spending amount

## Example

```md
  !!eth.core.approve_token_spending
      contract_address: b27a31f1b0af2946b7f582768f03239b1ec07c2c
      spender: b27a31f1b0af2946b7f58276af7fc56681767523
```
