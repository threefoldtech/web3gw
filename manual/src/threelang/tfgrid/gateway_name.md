# Gateway Name Actions

- To deploy a gateway name workload, use the Gateway Name actions

## Create Operation

- action name: !!tfgrid.gateway_name.create
- parameters:
  - name [required]
  - farm_id [optional]
  - backend [required]
    - the URL that the gateway will pass traffic to.
  - tls_passthrough [optional]
    - yes or no

- Example:
  
  ```
  !!tfgrid.gateway_name.create 
      name: hamadagateway
      backend: http://1.1.1.1:9000
  ```

## Get Operation

- action name: !!tfgrid.gateway_name.get
- parameters:
  - name [required]

- Example:
  
  ```
  !!tfgrid.gateway_name.get
      name: hamadagateway
  ```

## Update Operations

- Update operations are not allowed on gateway names.

## Delete Operation

- action_name: !!tfgrid.gateway_name.delete
- parameters:
  - name [required]

- Example:
  
  ```
  !!tfgrid.gateway_name.delete
      name: hamadagateway
  ```