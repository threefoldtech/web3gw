# Import Private Key Rescan Action

> Imports the passed private key which must be the wallet import format (WIF). It sets the account label to the one provided. When rescan is true, the block history is scanned for transactions addressed to provided privKey. The WIF string must be a base58-encoded string.

- action name: !!btc.imports.priv_key_rescan
- parameters:
  - wif [required]
  - label [required]
  - rescan [optional]
    - defaults to `false`

## Example

```md
    !!btc.imports.priv_key_rescan
        wif: tykbn2pG
        label: label1
        rescan: true
```
