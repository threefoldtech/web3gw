# Client actions

## Load client
- action name: !!tfchain.client.load
- parameters:
    - `network`: is the tfchain network, should be one of (mainnet, testnet, qanet, devnet)
    - `mnemonic`: twin mnemonic on the chain

- example:
  ```md
  !!tfchain.client.load
      network:devnet 
      mnemonic:'YOUR MNEMONIC'
  ```