# Await Transaction on Eth Bridge Action

> Await till a transaction is processed on ethereum bridge that contains a specific memo

- action name: !!stellar.bridge.await_transaction_on_eth_bridge
- parameters:
  - memo [required]

## Example

```md
    !!stellar.bridge.await_transaction_on_eth_bridge
        memo: mymemo
```
