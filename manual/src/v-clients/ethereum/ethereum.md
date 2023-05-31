# v Ethereum client

## Create and load a client

Needs an [RPC websocket client to a web3proxy server](../vclients.md#rpc-websocket-client) which is assumed to be present in a variable `rpcClient`.

The `url` parameter needs the url to an Ethereum full node and the `secret` parameter a private key in hex format of the Ethereum account the client needs to use.

```v
eth_url := 'ws://185.69.167.224:8546'
secret := '1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef'

mut eth_client := eth.new(mut rpcClient)
eth_client.load(url: eth_url, secret: secret)!
```

## Convert TFT on Ethereum to TFT on Stellar

The Ethereum  clients provides an easy way to convert [TFT on Ethereum](https://github.com/threefoldfoundation/tft/tree/main/ethereum) to  [TFT on Stellar](https://github.com/threefoldfoundation/tft-stellar) using the Stellar-Ethereum bridge.

The `amount` parameter is a string in decimal format of the number of TFT's to convert. Keep in mind that a conversion fee of 1 TFT will be deducted so make sure the amount is larger than that.

The destination parameter is the Stellar account that will receive the TFT's.

The following snippet will send 49.12 TFT (50.12 - 1 conversion fee) to GBN4RY5FDSY5MJJKD3G4QYXLQ73H6MXYPUXT4YMV3JXWA2HCXAJTFOZ2.

```v
amount := '50.12'
destination := 'GBN4RY5FDSY5MJJKD3G4QYXLQ73H6MXYPUXT4YMV3JXWA2HCXAJTFOZ2'

eth_client.withdraw_eth_tft_to_stellar(destination: destination, amount: amount)!
```

The conversion from TFT on Stellar to TFT on Ethereum  is part of the [Stellar client](../stellar/stellar.md#convert-tft-on-stellar-to-tft-on-ethereum).
