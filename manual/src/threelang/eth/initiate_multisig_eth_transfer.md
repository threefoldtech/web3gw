# Initiate multisig Eth Transfer Action

> initiates a multisig eth transfer.

- action name: !!eth.multisig.initiate_eth_transfer
- parameters:
  - contract_address [required]
  - destination [required]
  - amount [require]

## Example

```md
  !!eth.multisig.initiate_eth_transfer
      contract_address: b27a31f1b0af2946b7f582768f03239b1ec07c2c
      destination: b27a31f1b0af2946b7f58276af7fc56681767523
      amount: 100
```
