# Contracts Action
Query and filter contracts on the chain.

## Filter contracts
- action name: !!explor.contracts.filter
- parameters:
	- `contract_id`: contract id
	- `twin_id`: twin id of the contract creator
	- `node_id`: node id of the contract
	- `type`: type of the contract
	- `state`: state of the contract
	- `name`: contract name
	- `number_of_public_ips`: number of public ips in the contract
	- `deployment_data`: deployment metadata
	- `deployment_hash`: deployment hash

    - `size`: size of the returned batch of the contracts. default is 50
    - `page`: offset of the returned batch of the contracts. default is 1
    - `randomize`: if true, the returned batch of contracts will be random. default is false
    - `count`: if true, will return the number of contracts that match the filter even if the size is set. default is false.

- examples:
    - get specific contract by it's id
        ```bash
        !!explor.contracts.filter
            contract_id: 2014
        ```
    - get all contracts on a node
        ```bash
        !!explor.contracts.filter
            node_id: 11
        ```
