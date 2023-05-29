# Funkwhale Namespace

- To deploy a funkwhale instance, use the funkwhale namespace.

## Create Operation

- action name: !!tfgrid.funkwhale.create
- parameters:
  - model_name [required]
  - farm_id [optional]
    - if 0, machines could span multiple nodes on different farms
  - capacity [required]
    - a string in ['small', 'medium', 'large'] indicating the capacity of the machines
    - small: 2 vCPU, 1GB RAM, 50GB SSD
    - medium: 2 vCPU, 4GB RAM, 100GB SSD
    - large: 4 vCPU, 8GB RAM, 250GB SSD
  
  - ssh_key [required]
  - admin_email [required]
  - admin_username [required]
  - admin_password [required]

## Read Operation

- action name: !!tfgrid.funkwhale.read
- parameters:
  - model_name [required]

## Delete Operation

- action_name: !!tfgrid.funkwhale.delete
- parameters:
  - model_name [required]
