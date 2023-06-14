# Twins

## Get twin
Get a twin info with different parameters. You can get the twin info by only providing one of the parameters:
- action name: !!tfchain.twins.get
- parameters: 
    - `id`: [optional] twin id
    - `pubkey`: [optional] SS58 address for the twin
- example:
    ```md
    !!tfchain.twins.get
        id:29
    ```
    ```md
    !!tfchain.twins.get
        pubkey:5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY
    ```

## Create twin
Create new twin on the chain for your account.
- action name: !!tfchain.twins.create
- parameters:
    - `network`: [optional] is the tfchain network, should be one of (mainnet, testnet, qanet, devnet)
    - `pubkey`: [optional] SS58 address for the twin
- example:
    ```md
    !!tfchain.twins.create
        network:devnet
        pubkey:5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY
    ```