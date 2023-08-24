# Explorer Example

- This example uses the tfgrid explorer to retrieve information about the grid

```md
!!tfgrid.nodes.get
    network: dev
    count: true
    node_id: 33
    size: 100
    page: 1

!!tfgrid.farms.get
    network: test
    count: true
    free_ips: 3

!!tfgrid.contracts.get
    network: qa
    count: true
    state: Created
    size: 10
    page: 5

!!tfgrid.stats.get
    network: dev
    status: up

!!tfgrid.twins.get
    twin_id: 1234
```
