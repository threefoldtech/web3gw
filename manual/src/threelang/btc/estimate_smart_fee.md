# Estimate Smart Fee Action

> Provides a more accurate estimated fee given an estimation mode.

- action name: !!btc.estimate.smart_fee
- parameters:
  - conf_target [required]
    - confirmation target in blocks
  - mode [optional]
    - defines the different fee estimation modes, should be one of UNSET, ECONOMICAL or CONSERVATIVE. defaults to `CONSERVATIVE`

## Example

```md
    !!btc.estimate.smart_fee
        conf_target: 1
        mode: UNSET
```
