# Import Public Key Rescan Action

> Imports the passed public key. When rescan is true, the block history is scanned for transactions addressed to provided pubkey.

- action name: !!btc.imports.pub_key_rescan
- parameters:
  - pub_key [required]
  - rescan [optional]

## Example

```md
    !!btc.imports.pub_key_rescan
        pub_key: b27a31f1b0af2946b7f582768f03239b1ec07c2c
        rescan: false
```
