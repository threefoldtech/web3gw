# Ethereum examples

## Prerequisites

To run this example on Stellar testnet and Ethereum testnet you need:

- Stellar testnet account with at least 10 XLM and TFT Trustline + TFT
- Ethereum testnet account with at least 0.01 ETH on Goerli network

## Convert Stellar TFT to Ethereum TFT

- `-s`: stellar secret
- `-d`: destination ethereum account
- `-m`: amount of tft in string format
- `-n`: network (testnet or public)

```sh
v -cg run convert_to_eth.v -m 10 -d destination_eth -n testnet -s stellar_s
```

## Convert Ethereum TFT to Stellar TFT

- `-s`: ethereum secret
- `-d`: destination stellar account
- `-m`: amount in units of 10^7

```sh
v -cg run convert_to_stellar.v -s secret -d destination_stellar_addrr -m 100000000
```
