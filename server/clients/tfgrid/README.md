# TFGrid

- TFGrid package provides the access point to operations on the tfgrid.
- operations on the tfgrid are exposed by the Runner struct.

## Runner

- Internally, a Runner holds a tfgrid client instance.
- It provides a layer on top of the grid client to make queries to the grid a bit more manageable and easier.
- It abstracts deployments/contracts operations from users into project operations.
- A user is able to deploy, read, and/or delete multiple types of projects.
- A user could deploy all types of tfgrid workloads using a combination of projects.

## changes

- rename functions to convert from grid-client types to web3proxy types
- rename functions to convert from web3proxy types to grid-client types
- add AddWGAccess flag to k8s cluster (maybe also iprange)
- think of a better naming for contractsInfo struct
- make a Contains function instead of using the one in workloads package
- in k8s remove worker, check if there is no other worker on workerNodeID before updating network
- hide names generate (like generateProjectName, ....)

## deploy

- make sure there is no contracts with project name
- convert from web3proxy types to grid-client types
- deploy
- load from state
- convert from grid client types to web3proxy types
- return

## delete

- generate project name
- delete by project name

## get

- load model contracts
- load from state (set state, load mnha)
- convert to web3proxy types
- return

## add (k8s, machines)

- load model contracts
- load from state (set state, load mnha)
- convert new workload from web3proxy types to grid-client types
- update deployment (network included)
  - if there is a new node, add to network
- load updated model
- return

## remove (k8s, machines)

- load model contracts
- load from state(set state, load mnha)
- update deployment (network included)
  - if there is another deployment on the same node of the removed vm, eshta
  - else, delete node from network
- load updated model
- return


## notes

- solutino type should be more indicative of solution type, find somewhere else for project name


## refactoring #15

- build two getters:
  - one works with graphql
  - the other works without graphql
- local state on web proxy: project to grid state (map[uint32]contractIDs) (map[projectName]gridState)
- 