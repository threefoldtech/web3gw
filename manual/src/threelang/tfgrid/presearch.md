# Presearch actions

- To deploy a presearch instance, use the Presearch actions.

## Create Operation

- action name: !!tfgrid.presearch.create
- parameters:
  - name [required]
  - farm_id [optional]
  - disk_size [optional]
  - ssh_key [required]
  - public_ip
    - yes or no to add a public ip to the presearch instance

- Example:
  
  ```
  !!tfgrid.presearch.create
      name: mypresearch
      farm_id: 3
      disk_size: 10GB
      public_ip: yes
  ```

## Get Operation

- action name: !!tfgrid.presearch.get
- parameters:
  - name [required]

- Example:
  
  ```
  !!tfgrid.presearch.get
      name: mypresearch
  ```

## Update Operations

- Update operations are not allowed on presearch instances.
  
## Delete Operation

- action_name: !!tfgrid.presearch.delete
- parameters:
  - name [required]

- Example:
  
  ```
  !!tfgrid.presearch.delete
      name: mypresearch
  ```
