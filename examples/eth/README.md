# Ethereum examples

## Prerequisites 

To run this example on Stellar testnet and Ethereum testnet you need:

- Stellar testnet account with at least 10 XLM and TFT Trustline + TFT
- Ethereum testnet account with at least 0.01 ETH on Goerli network

## Deposit to Ethereum bridge

- `-s`: stellar secret
- `-d`: destination ethereum account
- `-m`: amount of tft in string format
- `-n`: network (testnet or public)

```sh
v -cg run deposit_bridge_eth.v -m 10 -d destination_eth -n testnet -s stellar_s
```

## Withdraw from Ethereum bridge

- `-s`: ethereum secret
- `-d`: destination stellar account
- `-m`: amount in units of 10^7

```sh
v -cg run withdraw_bridge_eth.v -s secret -d destination_stellar_addrr -m 100000000
```