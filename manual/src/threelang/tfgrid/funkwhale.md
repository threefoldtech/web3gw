# Funkwhale action

- To deploy a funkwhale instance, use the funkwhale action.

## Create Operation

- action name: !!tfgrid.funkwhale.create
- parameters:
  - name [required]
  - farm_id [optional]
    - if 0, machines could span multiple nodes on different farms
  - capacity [required]
    - a string in ['small', 'medium', 'large', 'extra-large'] indicating the capacity of the funkwhale instance
    - small: 2 vCPU, 2GB RAM, 10GB SSD
    - medium: 2 vCPU, 4GB RAM, 50GB SSD
    - large: 2 vCPU, 8GB RAM, 240 SSD
    - extra-large: 8vCPU, 16GB RAM, 480GB SSD
  
  - ssh_key [required]
  - admin_email [required]
  - admin_username [required]
  - admin_password [required]

- Example:
  
  ```
  !!tfgrid.funkwhale.create
      name: funkwhale_instance
      farm_id: 4
      capacity: medium
      ssh_key: my_ssh_key
      admin_email: email@gmail.com
      admin_username: username1
      admin_password: pass1
  ```

## Get Operation

- action name: !!tfgrid.funkwhale.get
- parameters:
  - name [required]

- Example:
  
  ```
  !!tfgrid.funkwhale.get
      name: funkwhale_instance
  ```

## Delete Operation

- action_name: !!tfgrid.funkwhale.delete
- parameters:
  - name [required]

- Example:
  
  ```
  !!tfgrid.funkwhale.delete
      name: funkwhale_instance
  ```
