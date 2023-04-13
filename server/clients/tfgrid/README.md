# TFGrid

- TFGrid package provides the access point to operations on the tfgrid.
- operations on the tfgrid are exposed by the Runner struct.

## Runner

- Internally, a Runner holds a tfgrid client instance.
- It provides a layer on top of the grid client to make queries to the grid a bit more manageable and easier.
- It abstracts deployments/contracts operations from users into project operations.
- A user is able to deploy, read, and/or delete multiple types of projects.
- A user could deploy all types of tfgrid workloads using a combination of projects.
