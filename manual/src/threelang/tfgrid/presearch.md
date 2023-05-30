# Presearch Namespace

- To deploy a presearch instance, use the Presearch namespace.

## Create Operation

- action name: !!tfgrid.presearch.create
- parameters:
  - name [required]
  - farm_id [optional]
  - capacity [required]
    - a string in ['small', 'medium', 'large', 'extra-large'] indicating the capacity of the presearch instance
    - small: 1 vCPU, 2GB RAM, 10GB SSD
    - medium: 2 vCPU, 4GB RAM, 50GB SSD
    - large: 4 vCPU, 8GB RAM, 240 SSD
    - extra-large: 8vCPU, 16GB RAM, 480GB SSD
  - disk_size [optional]
  - ssh_key [required]
- arguments:
  - public_ip

## Get Operation

- action name: !!tfgrid.presearch.get
- parameters:
  - name [required]

## Update Operations

- Update operations are not allowed on presearch instances.
  
## Delete Operation

- action_name: !!tfgrid.presearch.delete
- parameters:
  - name [required]
