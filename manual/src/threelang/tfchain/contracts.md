# Contracts

## Get 
Get the contract info. by one of the following examples.
- action name: !!chain.contracts.get
- parameters:
    - `id`: [optional] contract id
    - `node_id`: [optional] the node where the contract is deployed
    - `hash`: [optional] the contract hash
- example:

    - Get specific contract by it's id.
        ```
        !!chain.contracts.get
            id:6142
        ```
    - Get all contracts deployed on specific node.
        ```
        !!chain.contracts.get
            node_id:11
        ```
    - Get specific contract by it's hash and the node where it is deployed.
        ```
        !!chain.contracts.get
            node_id:11
            hash:829f2c6b1a654ffe4d3ddcd89e161fe1
        ```

## Create
Create new contract on the chain. one of three types (name, node, rent).
- action name: !!chain.contracts.create
- paramters: 
    - `type`: [required] contract type, should be one of (name, node, rent)
    - `name`: [required] contract name, should be unique.
    - `node_id`: [required] the node where you want to deploy the contract
    - `body`: [optional] metadata for the contract. default is empty string.
    - `hash`: [required] hash of the deployment.
    - `public_ips`: [optional] number of public ips attached to the contract. default is 0.
    - `solution_provider_id`: [optional] solution provider twin id

- example:

    - Create new Name contract
        ```
        !!chain.contracts.create
            type:name
            name:foo
        ```

    - Create new Node contract
        ```
        !!chain.contracts.create
            type:node
            node_id:11
            body:'metadata'
            hash:829f2c6b1a654ffe4d3ddcd89e161fe1
            public_ips:1
            solution_provider_id:102
        ```

    - Create new Rent contract
        ```md
        !!chain.contracts.create
            type:rent
            node_id:29
            solution_provider_id:102
        ```

## Cancel
Cancel contract by it's id or batch of contracts by their ids.
- action name: !!chain.contracts.cancel
- parameters:
    - `contract_id`: [optional] contract id
    - `contract_ids`: [optional] batch of contract ids separated with comma.
- examples:
    ```
    !!chain.contracts.cancel
        contract_id:6142
    ```
    ```
    !!chain.contracts.cancel
        contract_ids:6142,6143,6144
    ```
