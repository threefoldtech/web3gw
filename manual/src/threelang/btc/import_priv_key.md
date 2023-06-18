# Import Private Key Action

> Imports the passed private key which must be the wallet import format (WIF). The WIF string must be a base58-encoded string.

- action name: !!btc.imports.priv_key
- parameters:
  - wif [required]

## Example

```md
    !!btc.imports.priv_key
        wif: tykbn2pG
```
