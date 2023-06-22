# Money
Is an inter-chain centralized actor that can do all money related actions on supported chains (`tfchain`, `stellar`, `etherum`, `bitcoin`, ...).

## Actions

### Send tokens between accounts
Send tokens from one account to another account on the same chain.
- supported sends:
    - btc on bitcoin -> btc on bitcoin
    - eth on ethereum -> eth on ethereum
    - token on ethereum -> token on ethereum
    - tft on stellar -> tft on stellar
    - tft on tfchain -> tft on tfchain

- action name: `!!web3gw.money.send`
- parameters:
    - `channel`: the chain where the transaction will be done. (`tfchain`, `stellar`, `ethereum`, `bitcoin`, ...)
    - `currency`: the currency to send (`tft`, `btc`, `eth`, `xlm`, ...). If the channel or the target_channel is `tfchain` then the currency is `tft` by default.
    - `from`: the source address or twin_id if the channel is `tfchain`
    - `to`: the destination address or twin_id if the target_channel is `tfchain`
    - `amount`: the amount to send

- examples:
    ```md
    !!web3gw.money.send
        channel:tfchain
        from:29
        to:28
        currency:tft
        amount:100
    ```

### Swap tokens on the same chain
Swap from token to another token on the same chain.
- supported swaps:
    - eth on ethereum -> tft on ethereum
    - tft on ethereum -> eth on ethereum
    - tft on stellar -> xlm on stellar

- action name: `!!web3gw.money.swap`
- parameters: 
    - `channel`: the chain where the transaction will be done. (`stellar`, `ethereum`, ...)
    - `from`: the source address or twin_id if the channel is `tfchain`
    - `to`: the destination address or twin_id if the target_channel is `tfchain`
    - `currency`: the currency to send (`tft`, `eth`, `xlm`, ...).
    - `target_currency`: the target currency to swap to (`tft`, `eth`, `xlm`, ...).
    - `amount`: the amount to swap
- examples:
    ```md
    !!web3gw.money.swap
        channel:stellar
        currency:tft
        target_currency:xlm
        from:GCCVPYFOHY7ZB7557JKENAX62LUAPLMGIWNZJAFV2MITK6T32V37KEJU
        to:GCCVPYFOHY7ZB7557JKENAX62LUAPLMGIWNZJAFV2MITK6T32V37KEJU
        amount:100
    ```
### Bridge TFT between chains
Transfer TFT from one chain to another chain.
- supported bridges:
    - tft on ethereum -> tft on stellar
    - tft on stellar -> tft on ethereum
    - tft on stellar -> tft on tfchain

- action name: `!!web3gw.money.bridge`
- parameters:
    - `channel`: the source chain to send the TFT from. (`stellar`, `ethereum`, ...)
    - `target_channel`: the destination chain to send the TFT to. (`stellar`, `ethereum`, `tfchain`, ...)
    - `from`: the source address or twin_id if the channel is `tfchain`
    - `to`: the destination address or twin_id if the target_channel is `tfchain`
    - `amount`: the amount to send
- examples:
    ```md
    !!web3gw.money.bridge
        channel:stellar
        target_channel:tfchain
        from:GCCVPYFOHY7ZB7557JKENAX62LUAPLMGIWNZJAFV2MITK6T32V37KEJU
        to:28
        amount:100
    ```