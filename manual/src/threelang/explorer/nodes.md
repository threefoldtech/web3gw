# Nodes Action
Query and filter nodes on the chain.

## Filter nodes
- action name: !!explorer.nodes.filter
- parameters:
    - `status`: node status (`up`, `down`)
    - `free_mru`: free memory
    - `free_hru`: free HDD space
    - `free_sru`: free SSD space
    - `total_mru`: total memory
    - `total_hru`: total HDD space 
    - `total_sru`: total SSD space 
    - `total_cru`: total CPU cores 
    - `country`: country where the node is located
    - `country_contains`: substring of country where the node is located
    - `city`: city where the node is located
    - `city_contains`: substring of city where the node is located
    - `farm_name`: full farm name (case-sensitive)
    - `farm_name_contains`: substring of farm name (case-insensitive)
    - `farm_id`: farm id where the node is registered
    - `free_ips`: number of free IPs
    - `gateway`: true if the node is a gateway
    - `dedicated`: true if the node farm is dedicated
    - `rentable`: true if the node is available for rent
    - `rented`: true if the node is rented
    - `rented_by`: twin id for the twin that rented the node
    - `available_for`: twin id for the twin that can use the node to deploy on
    - `node_id`: node id
    - `twin_id`: twin id for the node


    - `size`: size of the returned batch of the nodes. default is 50
    - `page`: offset of the returned batch of the nodes. default is 1
    - `randomize`: if true, the returned batch of nodes will be random. default is false
    - `count`: if true, will return the number of nodes that match the filter even if the size is set. default is false.

- examples:
    - get specific node by it's id
        ```bash
        !!explorer.nodes.filter
            node_id: 11
        ```
    - get all nodes that has public access
        ```bash
        !!explorer.nodes.filter
            gateway: true
        ```
    - filter nodes based on capacity
        ```bash
        !!explorer.nodes.filter
            free_mru: 2GB
            free_hru: 100GB
            free_sru: 50GB
        ```

## Get node
- action name: !!explorer.nodes.get
- paramters: 
    - `node_id`: node id

- example:
    ```bash
    !!explorer.nodes.get
        node_id: 11
    ```

## Get node status
- action name: !!explorer.nodes.status
- paramters: 
    - `node_id`: node id

- example:
    ```bash
    !!explorer.nodes.status
        node_id: 11
    ```