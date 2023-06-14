# Account actions

## Create account
- action name: !!tfchain.account.create
- parameters:
  - `network`: is the tfchain network, should be one of (mainnet, testnet, qanet, devnet)

- example:
  ```md
  !!tfchain.account.create
      network:devnet 
  ```
  