# TFT Stellar go client

## Configuration

The client can be constructed with the following parameters:

- `secret` - The secret of the account that will be used to sign transactions. If empty, a keypair will be generated.
- `network` - The network to use, either `testnet` or `public`

## Transfering TFT

```go
client, err := NewClient("stellarsecret" "testnet")
if err != nil {
    log.fatal(err)
}

// Transfer 100 TFT to someDestinationAddress with someMemo
err = client.Transfer("someDestinationAddres", "someMemo", "100")
if err != nil {
    log.fatal(err)
}
```

