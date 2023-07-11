# Balance

## Send to another tfchain account 
Transfer balance between tfchain accounts
- action name: !!chain.money.send
- parameters:
  - `channel`: [optional] channel name to use for the transfer. default is `tfchain`
  - `currency`: [optional] currency to use for the transfer. default is `tft`
  - `to`: [required] tfchain address in SS58 format
  - `amount`: [required] the amount of sent TFTs

- example:
  ```md
  !!chain.money.send 
    channel:tfchain
    currency:tft
    to:5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY
    amount:100
  ```

## Send to stellar account
Swap balance to stellar account
- action name: !!chain.balance.swap
- parameters:
  - `channel`: [optional] channel name to use for the transfer. default is `tfchain`
  - `currency`: [optional] currency to use for the transfer. default is `tft`
  - `to`: [required] tfchain address in SS58 format
  - `amount`: [required] the amount of sent TFTs

- example:
  ```md
  !!chain.money.send
    channel:stellar
    currency:tft
    to:GCCVPYFOHY7ZB7557JKENAX62LUAPLMGIWNZJAFV2MITK6T32V37KEJU
    amount:100
  ```
## Get balance
Get current balance for tfchain account
- action name: !!chain.money.balance
- parameters:
  - `channel`: [optional] channel name to use for the transfer. default is `tfchain`
  - `currency`: [optional] currency to use for the transfer. default is `tft``
  - `address`: [optional] tfchain address in SS58 format. in case no address found. default is the address of the current user.
  
- example:
  ```md
  !!chain.money.balance
    channel:tfchain
    currency:tft
    address:5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY  
  ```
