# Balance

## Transfer 
Transfer balance between tfchain twins
- action name: !!tfchain.balance.transfer
- parameters:
  - `amount`: [required] the amount of sent TFTs
  - `destination`: [required] tfchain address in SS58 format

- example:
  ```md
  !!tfchain.balance.transfer
      amount:100
      destination:5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY
  ```

## Swap
Swap balance to stellar account
- action name: !!tfchain.balance.swap
- parameters:
  - `amount`: [required] the amount of sent TFTs
  - `destination`: [required] stellar address

- example:
  ```md
  !!tfchain.balance.swap
      amount:100
      destination:GCCVPYFOHY7ZB7557JKENAX62LUAPLMGIWNZJAFV2MITK6T32V37KEJU
  ```
## Get
Get current balance for tfchain account
- action name: !!tfchain.balance.get
- parameters:
  - `address`: [optional] tfchain address in SS58 format. in case no address found. it will get the balance of the loaded key.
  
- example:
  ```md
  !!tfchain.balance.get
      address:5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY  
  ```
