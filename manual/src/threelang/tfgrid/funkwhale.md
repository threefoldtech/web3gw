# Funkwhale Namespace

- To deploy a funkwhale instance, use the funkwhale namespace.

## Create Operation

- action name: !!tfgrid.funkwhale.create
- parameters:
  - name [required]
  - farm_id [optional]
    - if 0, machines could span multiple nodes on different farms
  - capacity [required]
    - a string in ['small', 'medium', 'large', 'extra-large'] indicating the capacity of the funkwhale instance
    - small: 1 vCPU, 2GB RAM, 10GB SSD
    - medium: 2 vCPU, 4GB RAM, 50GB SSD
    - large: 4 vCPU, 8GB RAM, 240 SSD
    - extra-large: 8vCPU, 16GB RAM, 480GB SSD
  
  - ssh_key [required]
  - admin_email [required]
  - admin_username [required]
  - admin_password [required]

## Get Operation

- action name: !!tfgrid.funkwhale.get
- parameters:
  - name [required]

## Delete Operation

- action_name: !!tfgrid.funkwhale.delete
- parameters:
  - name [required]
