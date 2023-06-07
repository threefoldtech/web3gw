# Discourse action

- To deploy a discourse instance, use the discourse action.

## Create Operation

- action name: !!tfgrid.discourse.create
- parameters:
  - name [optional]
  - farm_id [optional]
    - if 0, machines could span multiple nodes on different farms
  - capacity [required]
    - a string in ['small', 'medium', 'large', 'extra-large'] indicating the capacity of the discourse instance
    - small: 1 vCPU, 2GB RAM, 10GB SSD
    - medium: 2 vCPU, 2GB RAM, 50GB SSD
    - large: 4 vCPU, 4GB RAM, 100 SSD
    - extra-large: 4vCPU, 8GB RAM, 150GB SSD
  
  - ssh_key [required]
  - developer_email [required]

  - smtp_address [required]
  - smtp_port [required]
  - smtp_username [required]
  - smtp_password [required]
  - smtp_tls [required]

- Example:
  
  ```
  !!tfgrid.discourse.create
      name: discoursename
      capacity: large
      ssh_key: my_ssh_key
      developer_email: email@gmail.com
      smtp_hostname: myhostname
      smtp_port: 9000
      smtp_username: username1
      smtp_password: pass1
      smtp_tls: true
  ```

## Get Operation

- action name: !!tfgrid.discourse.get
- parameters:
  - name [required]

- Example:
  
  ```
  !!tfgrid.discourse.get
      name: discoursename
  ```

## Delete Operation

- action_name: !!tfgrid.discourse.delete
- parameters:
  - name [required]

- Example:
  
  ```
  !!tfgrid.discourse.delete
      name: discoursename
  ```
