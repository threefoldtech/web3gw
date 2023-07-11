# Get Balance Action

> Returns the available balance for the specified account using the default number of minimum confirmations. You can provide * as an account to get the balance of all accounts.

- action name: !!btc.get.balance
- parameters:
  - account [required]

## Example

```md
    !!btc.get.balance
        account: acc1
```
