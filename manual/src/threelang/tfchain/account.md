# Account actions

## Create account
- action name: !!chain.account.create
- parameters:
  - `network`: is the tfchain network, should be one of (main, test, qa, dev)

- example:
  ```md
  !!chain.account.create
      network:dev 
  ```
  
## Get address
Get the address for the loaded account.
- action name: !!chain.account.address

- example:
  ```md
  !!chain.account.address
  ```