# Contracts

## Get 
Get the contract info. by one of the following examples.
- action name: !!tfchain.contracts.get
- parameters:
    - `id`: [optional] contract id
    - `node_id`: [optional] the node where the contract is deployed
    - `hash`: [optional] the contract hash
- example:

    - Get specific contract by it's id.
        ```
        !!tfchain.contracts.get
            id:6142
        ```
    - Get all contracts deployed on specific node.
        ```
        !!tfchain.contracts.get
            node_id:11
        ```
    - Get specific contract by it's hash and the node where it is deployed.
        ```
        !!tfchain.contracts.get
            node_id:11
            hash:ff38eba6dbe9c5d2a1eb829bc48d77a6
        ```

## Create
Create new contract on the chain. one of three types (name, node, rent).
- action name: !!tfchain.contracts.create
- paramters: 
    - `type`: [required] contract type, should be one of (name, node, rent)
    - `name`: [required] contract name, should be unique.
    - `node_id`: [required] the node where you want to deploy the contract
    - `body`: [required] the deployment itself
    - `hash`: [required] hash of the deployment.
    - `public_ips`: [required] number of public ips attached to the contract
    - `solution_provider_id`: [optional] solution provider twin id

- example:

    - Create new Name contract
        ```
        !!tfchain.contracts.create
            type:name
            name:foo
        ```

    - Create new Node contract
        ```
        !!tfchain.contracts.create
            type:node
            node_id:11
            body: `string`
            hash: `string`
            public_ips:1
            solution_provider_id:102
        ```

    - Create new Rent contract
        ```md
        !!tfchain.contracts.create
            type:rent
            node_id:29
            solution_provider_id:102
        ```

## Cancel
Cancel contract by it's id or batch of contracts by their ids.
- action name: !!tfchain.contracts.cancel
- parameters:
    - `contract_id`: [optional] contract id
    - `contract_ids`: [optional] batch of contract ids separated with comma.
- examples:
    ```
    !!tfchain.contracts.cancel
        contract_id:6142
    ```
    ```
    !!tfchain.contracts.cancel
        contract_ids:6142,6143,6144
    ```
