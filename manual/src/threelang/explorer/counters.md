# Statistics actions
Get counters from the explorer

## Get Operation
- action name: !!explorer.stats.get
- paramters: 
    - `status`: status of the nodes (`up`, `down`)

- example:
    ```bash
    !!explorer.stats.get
        status: up
    ```