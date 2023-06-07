# Taiga action

- To deploy a taiga instance, use the Taiga action.

## Create Operation

- action name: !!tfgrid.taiga.create
- parameters:
  - name [required]
  - farm_id [optional]
  - capacity [required]
    - a string in ['small', 'medium', 'large', 'extra-large'] indicating the capacity of the taiga instance
    - small: 2 vCPU, 2GB RAM, 100GB SSD
    - medium: 2 vCPU, 4GB RAM, 150GB SSD
    - large: 4 vCPU, 4GB RAM, 250 SSD
    - extra-large: 4vCPU, 8GB RAM, 400GB SSD
  - disk_size [optional]
  - ssh_key [required]
  - admin_username [required]
  - admin_password [required]
  - admin_email [required]
  - public_ip
    - yes or no to add a public ip to the taiga instance

- Example:
  
  ```
  !!tfgrid.taiga.create
      name: hamadataiga
      capacity: medium
      size: 10GB
      ssh_key: my_taiga_ssh_key
      admin_username: user1
      admin_password: pass1
      admin_email: email@gmail.com
  ```

## Get Operation

- action name: !!tfgrid.taiga.get
- parameters:
  - name [required]

- Example:
  
  ```
  !!tfgrid.taiga.get
      name: hamadataiga
  ```

## Update Operations

- Update operations are not allowed on taiga instances.
  
## Delete Operation

- action_name: !!tfgrid.taiga.delete
- parameters:
  - name [required]

- Example:
  
  ```
  !!tfgrid.taiga.delete
      name: hamadataiga
  ```
