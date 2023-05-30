# Discourse Namespace

- To deploy a discourse instance, use the discourse namespace.

## Create Operation

- action name: !!tfgrid.discourse.create
- parameters:
  - name [required]
  - farm_id [optional]
    - if 0, machines could span multiple nodes on different farms
  - capacity [required]
    - a string in ['small', 'medium', 'large', 'extra-large'] indicating the capacity of the discourse instance
    - small: 1 vCPU, 2GB RAM, 10GB SSD
    - medium: 2 vCPU, 4GB RAM, 50GB SSD
    - large: 4 vCPU, 8GB RAM, 240 SSD
    - extra-large: 8vCPU, 16GB RAM, 480GB SSD
  
  - ssh_key [required]
  - developer_email [required]

  - smtp_hostname [required]
  - smtp_port [required]
  - smtp_username [required]
  - smtp_password [required]
  - smtp_tls [required]

## Get Operation

- action name: !!tfgrid.discourse.get
- parameters:
  - name [required]

## Delete Operation

- action_name: !!tfgrid.discourse.delete
- parameters:
  - name [required]
