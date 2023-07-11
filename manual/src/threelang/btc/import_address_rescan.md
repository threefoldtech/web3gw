# Import Address Rescan Action

> Imports the passed public address. When rescan is true, the block history is scanned for transactions addressed to provided address.

- action name: !!btc.imports.address_rescan
- parameters:
  - address [required]
  - account [required]
  - rescan [optional]
    - defaults to `false`

## Example

```md
    !!btc.imports.address_rescan
        address: b27a31f1b0af2946b7f582768f03239b1ec07c2c
        account: account1
        rescan: true
```
