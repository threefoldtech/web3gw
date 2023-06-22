# Client
An actor that loads all the needed clients to interact with the supported chains (`tfchain`, `stellar`, `etherum`, `bitcoin`, ...).

## Actions
- action name: `!!web3gw.client.load`
- parameters:
    
    TFChain client configuration:
    - `tfc_mnemonic`: required, mnemonic to load the tfchain client with
    - `tfc_network`: default to `dev`, tfchain network to connect to
    
    Bitcoin client configuration:
    - `btc_host`
    - `btc_user`
    - `btc_pass`
    
    Ethereum client configuration:
    - `eth_url`
    - `eth_secret`
    
    Stellar client configuration:
    - `str_network`: default to `public`, stellar network to connect to
    - `str_secret`

- examples:
    Load handler that holds tfchain, stellar clients:
    ```md
    !!web3gw.client.load
        tfc_mnemonic:'candy maple cake sugar pudding cream honey rich smooth crumble sweet treat'
        tfc_network:dev
        str_network:public
        str_secret:'SAKRB7EE6H23EF733WFU76RPIYOPEWVOMBBUXDQYQ3OF4NF6ZY6B6VLW'