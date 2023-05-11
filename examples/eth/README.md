# Ethereum examples

## Deposit to Ethereum bridge

```sh
v -cg run deposit_bridge_eth.v -s "stellar_secret" -n "(testnet/public)" -a "amount" -d "destination_stellar_addr"
```

## Withdraw from Ethereum bridge

```sh
v -cg run withdraw_bridge_eth.v -s "eth_secret" -d "destination_stellar_addr" -e "eth_node_url"  
```