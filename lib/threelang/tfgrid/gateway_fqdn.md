# Gateway FQDN Namespace

- To deploy a gateway fqdn workload, use the Gateway FQDN namespace

## Create Operation

- action name: !!tfgrid.gateway_fqdn.create
- parameters:
  - model_name [required]
  - node_id [required]
  - fqdn [required]
  - backend [required]
    - the URL that the gateway will pass traffic to.
- arguments:
  - tls_passthrough

## Read Operation

- action name: !!tfgrid.gateway_fqdn.read
- parameters:
  - model_name [required]

## Update Operations

- Update operations are not allowed on gateway fqdn.

## Delete Operation

- action_name: !!tfgrid.gateway_fqdn.delete
- parameters:
  - model_name [required]
