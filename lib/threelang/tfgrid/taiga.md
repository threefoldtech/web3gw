# Taiga Namespace

- To deploy a taiga instance, use the Taiga namespace.

## Create Operation

- action name: !!tfgrid.taiga.create
- parameters:
  - model_name [required]
  - farm_id [optional]
  - capacity [required]
    - a string in ['small', 'medium', 'large'] indicating the capacity of the vm
    - small: 2 vCPU, 4GB RAM, 10GB SSD
    - medium: 4 vCPU, 8GB RAM, 20GB SSD
    - large: 8 vCPU, 16GB RAM, 30GB SSD
  - disk_size [optional]
  - ssh_key [required]
- arguments:
  - public_ip

## Read Operation

- action name: !!tfgrid.taiga.read
- parameters:
  - model_name [required]

## Update Operations

- Update operations are not allowed on taiga instances.
  
## Delete Operation

- action_name: !!tfgrid.taiga.delete
- parameters:
  - model_name [required]
