# TF Chain Actions
In a simple format you can do almost all the extrinsic and queries on the chian. List of the namespaces

## Name spaces and their actions
- [TF Chain Actions](#tf-chain-actions)
  - [Name spaces and their actions](#name-spaces-and-their-actions)
  - [Account](#account)
  - [Balance](#balance)
  - [Twins](#twins)
  - [Nodes](#nodes)
  - [Farms](#farms)
  - [Contracts](#contracts)
  - [Service Contract](#service-contract)
  - [Metadata](#metadata)


## Account
- Create new account on the chain by Generating mnemonic
    ```md
    !!tfchain.account.create
        network:devnet 
    ```
    - `network`: is the tfchain network, should be one of (mainnet, testnet, qanet, devnet)
- Load the account to do chain extrinsic 
    ```md
    !!tfchain.client.load
        network:devnet
        mnemonic:\'secret words as your mnemonic\'
    ```
    - `network`: is the tfchain network, should be one of (mainnet, testnet, qanet, devnet)
    - `mnemonic`: twin mnemonic on the chain

## Balance
- Transfer balance between tfchain twins
    ```md
    !!tfchain.balance.transfer
        amount:100
        destination:5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY
    ```
    - `amount`: the amount of sent TFTs
    - `destination`: tfchain address in SS58 format
- Swap balance to stellar account
    ```md
    !!tfchain.balance.swap
        amount:100
        destination:GCCVPYFOHY7ZB7557JKENAX62LUAPLMGIWNZJAFV2MITK6T32V37KEJU
    ```
    - `amount`: the amount of sent TFTs
    - `destination`: stellar address
- Get current balance for tfchain account
    ```md
    !!tfchain.balance.get
        address:5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY   
    ```
    - `address`: tfchain address in SS58 format
  
## Twins
- Get a twin info with different parameters
    ```
    !!tfchain.twins.get
        id:29
    ```
    ```
    !!tfchain.twins.get
        pubkey:5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY
    ```
    You can get the twin info by only providing one of the following:
    - `id`: twin id
    - `pubkey`: SS58 address for the twin

- Create new twin on the chain for your account.
    ```md
    !!tfchain.twins.create
        network:devnet
        pubkey:5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY
    ```
    - `network`: is the tfchain network, should be one of (mainnet, testnet, qanet, devnet)
    - `pubkey`: SS58 address for the twin
    

## Nodes
- Get nodes info from the chain
    ```md
    !!tfchain.nodes.get
        farm_id:1
    ```
    This will return all the nodes for this farm.
    ```
    !!tfchain.nodes.get
        node_id:11 
    ```
    This will only return the node with the same id.

## Farms
- Get farms info
    ```
    !!tfchain.farms.get
        id:1
    ```
    ```
    !!tfchain.farms.get
        name:freefarm
    ```
    - `id`: is the farm id
    - `name`: the full farm name "case-sensitive"
- Create new farm on the chain
    ```md
    !!tfchain.farms.create
        name:newfarm
    ```
    ```
    !!tfchain.farms.create
        name:newfarmwithips
        public_ips: 185.206.122.152/16
        gateways: 185.206.122.152
    ```
    It is optional to add ips to the farm
    - `public_ips`: is a list of IP addresses in CIDR format xxx.xxx.xxx.xxx/xx. separated by comma.
    - `gateways`: is a list of Gateways for the IP in ipv4 format. separated by comma.
    Values on both public_ips/gateways lists are mapped by their indices. so for example gateways[0] is the gateway for the ip public_ips[0] and so on. 

## Contracts
- Get the contract info
    ```
    !!tfchain.contracts.get
        id:6142
    ```
    Get specific contract by it's id.
    ```
    !!tfchain.contracts.get
        node_id:11
    ```
    Get all contracts deployed on specific node.
    ```
    !!tfchain.contracts.get
        node_id:11
        hash:ff38eba6dbe9c5d2a1eb829bc48d77a6
    ```
    Get specific contract by it's hash and the node where it is deployed.

- Create new Name contract
    ```
    !!tfchain.contracts.create_name
        name:foo
    ```

- Create new Node contract
    ```
    !!tfchain.contracts.create_node
        node_id:11
        body: `string`
        hash: `string`
        public_ips:1
        solution_provider_id:102
    ```
    - `node_id`: the node where you want to deploy the contract
    - `body`: the deployment itself
    - `hash`: hash of the deployment.
    - `public_ips`: number of public ips attached to the contract
    - `solution_provider_id`: `optional` solution provider twin id
- Create new Rent contract
    ```md
    !!tfchain.contracts.create_rent
        node_id:29
        solution_provider_id:102
    ```
    - `node_id`: the wanted to rent node.
    - `solution_provider_id`: `optional` solution provider twin id

- Cancel contracts
    ```
    !!tfchain.contracts.cancel
        contract_id:6142
    ```
    Cancel single contract by its id.
    ```
    !!tfchain.contracts.cancel
        contract_ids:6142,6143,6144
    ```
    Cancel batch of contracts with ids separated with comma.

## Service Contract
- Create new service contract
    ```
    !!tfchain.service_contract.create
        service:  `string`
        consumer: `string`    
    ```
    - `service`: SS58 address
    - `consumer`: SS58 address

- Approve service contract
    ```
    !!tfchain.service_contract.approve
        contract_id: 2015
    ```
- Reject service contract
    ```
    !!tfchain.service_contract.reject
        contract_id:2015
    ```

- Bill service contract
    ```
    !!tfchain.service_contract.bill
        contract_id:2015     
        variable_amount:100
    ```
- Cancel service contract 
    ```
    !!tfchain.service_contract.cancel
        contract_id:2015
    ```
- Update service contract
    ```
    !!tfchain.service_contract.set
        contract_id:2015
        base_fee:      `u64`
        variable_fee:  `u64`
    ```
    Set fees for service contract
    - `base_fee`: amount of base fee
    - `variable_fee`: amount of variable fee

    ```
    !!tfchain.service_contract.set
        metadata: string
    ```
    Set metadata for service contract

## Metadata
- Get zos version
    ```
    !!tfchain.metadata.zos_version
    ```
- Get the current height of the chain
    ```
    !!tfchain.metadata.chain_height
    ```
