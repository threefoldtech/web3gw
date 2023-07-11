# Import Private Key Lable Action

> Imports the passed private key which must be the wallet import format (WIF). It sets the account label to the one provided. The WIF string must be a base58-encoded string.

- action name: !!btc.imports.priv_key_label
- parameters:
  - wif [required]
  - label [required]

## Example

```md
    !!btc.imports.priv_key_label
        wif: tykbn2pG
        label: label1
```
