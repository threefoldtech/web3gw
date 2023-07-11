# Client actions

## Load client
- action name: !!chain.client.load
- parameters:
    - `network`: is the tfchain network, should be one of (main, test, qa, dev)
    - `mnemonic`: twin mnemonic on the chain

- example:
  ```md
  !!chain.client.load
      network:dev 
      mnemonic:'YOUR MNEMONIC'
  ```