# Create Wallet Action

> Creates a new wallet account taken into account the provided arguments.

- action name: !!btc.wallet.create
- parameters:
  - name [required]
  - disable_private_keys [optional]
    - defaults to `false`
  - create_blank_wallet [optional]
    - defaults to `false`
  - passphrase [required]
  - avoid_reuse [optional]
    - defaults to `false`

## Example

```md
    !!btc.wallet.create
        name: name1
        disable_private_keys: true
        create_blank_wallet: true
        passphrase: this is my phrase
        avoid_reuse: false
```
