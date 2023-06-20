# Farms Action
Query and filter farms on the chain.

## Filter farms
- action name: !!explorer.farms.filter
- parameters:
    - `free_ips`: number of free public IPs on the farm
    - `total_ips`: number of total public IPs on the farm
    - `stellar_address`: farm stellar address
    - `pricing_policy_id`: farm pricing policy id 
    - `farm_id`: farm id 
    - `twin_id`: twin id for the farm
    - `name`: full name of the farm (case-sensitive)
    - `name_contains`: substring of the farm name (case-insensitive)
    - `certification_type`: farm certification type (`DIY`, `Gold`)
    - `dedicated`: true if the farm is dedicated


    - `size`: size of the returned batch of the farms. default is 50
    - `page`: offset of the returned batch of the farms. default is 1
    - `randomize`: if true, the returned batch of farms will be random. default is false
    - `count`: if true, will return the number of farms that match the filter even if the size is set. default is false.

- examples:
    - get specific farm by it's id
        ```bash
        !!explorer.farms.filter
            farm_id: 1
        ```
    - get all farms that marked as dedicated
        ```bash
        !!explorer.farms.filter
            dedicated: true
        ```
    - filter farms based on capacity
        ```bash
        !!explorer.farms.filter
            free_ips: 2
        ```