# Explorer Example

- This example uses the tfgrid explorer to retrieve information about the grid

!!tfgrid.explorer.nodes
    network: dev
    count: true
    node_id: 33
    size: 100
    page: 1

!!tfgrid.explorer.farms
    network: test
    count: true
    free_ips: 3

!!tfgrid.explorer.contracts
    network: qa
    count: true
    state: Created
    size: 10
    page: 5

!!tfgrid.explorer.node
    node_id: 1234

!!tfgrid.explorer.stats
    network: dev
    status: up

!!tfgrid.explorer.twins
    twin_id: 1234
