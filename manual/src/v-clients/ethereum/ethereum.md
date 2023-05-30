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
