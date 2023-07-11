# Nodes

## Get
Get nodes info from the chain
- action name: !!chain.nodes.get
- parameters:
    - `farm_id`: [optional] this will return all the nodes for this farm.
    - `node_id`: [optional] this will only return the node with the same id.

- example:
    ```md
    !!chain.nodes.get
        farm_id:1
    ```
    ```
    !!chain.nodes.get
        node_id:11 
    ```