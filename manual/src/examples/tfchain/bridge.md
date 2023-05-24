# Bridge examples

## Convert TFChain TFT to Stellar TFT

- `-m`: tfchain mnemonic
- `-n`: tfchain network
- `-d`: stellar address (destination address)
- `-a`: amount to transfer from tfchain to stellar

```sh
v -cg run convert_to_stellar.v -m mnemonic -d destination_stellar -n mainnet -a "100"
```

## Convert Stellar TFT to TFChain TFT

- `-s`: stellar secret
- `-n`: stellar network (testnet or public)
- `-a`: amount to transfer from stellar to tfchain
- `-t`: the twinid in tfchain (destination)


```sh
v -cg run convert_to_tft.v -s secret -n public -t 1651 -a "100"
```
