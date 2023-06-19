# Twins

## Get twin
Get a twin info with different parameters. You can get the twin info by only providing one of the parameters:
- action name: !!chain.twins.get
- parameters: 
    - `id`: [optional] twin id
    - `pubkey`: [optional] SS58 address for the twin
- example:
    ```md
    !!chain.twins.get
        id:29
    ```
    ```md
    !!chain.twins.get
        pubkey:5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY
    ```

## Create twin
Create new twin on the chain for your account.
- action name: !!chain.twins.create
- parameters:
    - `network`: [optional] is the tfchain network, should be one of (main, test, qa, dev)
    - `pubkey`: [optional] SS58 address for the twin
- example:
    ```md
    !!chain.twins.create
        network:dev
        pubkey:5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY
    ```