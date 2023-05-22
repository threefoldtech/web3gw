# Swap examples

## Swap Eth to Ethereum TFT

- `-s`: stellar secret
- `-m`: amount of tft in string format

```sh
v -cg run swap_eth_for_tft.v -m "0.00001" -s ethereum_s
```

## Swap Ethereum TFT to Eth

- `-s`: ethereum secret
- `-m`: amount of tft in string format (can be with decimals: "0.1")
- `-e`: ethereum node url

```sh
v -cg run swap_tft_for_eth.v -s secret -d destination_stellar_addrr -m "100.50" -e https://goerli.infura.io/v3/your_infura_key
```
