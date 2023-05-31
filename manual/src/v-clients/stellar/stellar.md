# v Stellar client

## Create and load a client

Needs an [RPC websocket client to a web3proxy server](../vclients.md#rpc-websocket-client) which is assumed to be present in a variable `rpcClient`.

The `secret` parameter requires a private key of the Stellar account the client needs to use.

```v
secret := 'SB1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890ABCDEFGH'

mut stellar_client := stellar.new(mut rpcClient)
stellar_client.load(secret: secret)!

```

## Convert TFT on Stellar to TFT on Ethereum

The stellar clients provides an easy way to convert [TFT on Stellar](https://github.com/threefoldfoundation/tft-stellar) to [TFT on Ethereum](https://github.com/threefoldfoundation/tft/tree/main/ethereum) using the Stellar-Ethereum bridge.

The `amount` parameter is a string in decimal format of the number of TFT's to convert. Keep in mind that a conversion fee of 2000 TFT will be deducted so make sure the amount is larger than that.

The destination parameter is the Ethereum account that will receive the TFT's.

The following snippet will send 1000.50 TFT (3000.50 - 2000 conversion fee) to 0x65e491D7b985f77e60c85105834A0332fF3002CE.

```v
amount := '3000.50'
destination := '0x65e491D7b985f77e60c85105834A0332fF3002CE'
stellar_client.bridge_to_eth(amount: amount, destination: destination)!
```

The conversion from TFT on Ethereum to TFT on Stellar is part of the [Ethereum client](../ethereum/ethereum.md#convert-tft-on-ethereum-to-tft-on-stellar).
