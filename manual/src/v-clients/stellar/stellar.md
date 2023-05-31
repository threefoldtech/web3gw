# v Stellar client

## Create and load a client

Needs an [RPC websocket client to a web3proxy server](../vclients.md#rpc-websocket-client) which is assumed to be present in a variable `rpcClient`.

The `secret` parameter requires a private key of the Stellar account the client needs to use.

```v
secret := 'SB1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890ABCDEFGH'

mut stellar_client := stellar.new(mut rpcClient)
stellar_client.load(secret: secret)!

```
