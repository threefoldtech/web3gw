# Gateway FQDN Namespace

- To deploy a gateway fqdn workload, use the Gateway FQDN namespace

## Create Operation

- action name: !!tfgrid.gateway_fqdn.create
- parameters:
  - name [required]
  - node_id [required]
  - fqdn [required]
  - backend [required]
    - the URL that the gateway will pass traffic to.
- arguments:
  - tls_passthrough

## Get Operation

- action name: !!tfgrid.gateway_fqdn.get
- parameters:
  - name [required]

## Update Operations

- Update operations are not allowed on gateway fqdn.

## Delete Operation

- action_name: !!tfgrid.gateway_fqdn.delete
- parameters:
  - name [required]
