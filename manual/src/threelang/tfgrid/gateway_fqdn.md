# Gateway FQDN Actions

- To deploy a gateway fqdn workload, use the Gateway FQDN actions

## Create Operation

- action name: !!tfgrid.gateway_fqdn.create
- parameters:
  - name [required]
  - node_id [required]
  - fqdn [required]
  - backend [required]
    - the URL that the gateway will pass traffic to.
  - tls_passthrough [optional]
    - yes or no

- Example:
  
  ```
  !!tfgrid.gateway_fqdn.create
      name: hamadafqdn
      node_id: 11
      backend: http://1.1.1.1:9000
      fqdn: hamada1.3x0.me
  ```


## Get Operation

- action name: !!tfgrid.gateway_fqdn.get
- parameters:
  - name [required]

- Example:
  
  ```
  !!tfgrid.gateway_fqdn.get
      name: hamadafqdn
  ```

## Update Operations

- Update operations are not allowed on gateway fqdn.

## Delete Operation

- action_name: !!tfgrid.gateway_fqdn.delete
- parameters:
  - name [required]


- Example:
  
  ```
  !!tfgrid.gateway_fqdn.delete
      name: hamadafqdn
  ```
