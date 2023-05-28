# Peertube Namespace

- To deploy a peertube instance, use the Peertube namespace.

## Create Operation

- action name: !!tfgrid.peertube.create
- parameters:
  - model_name [required]
  - farm_id [optional]
    - if 0, machines could span multiple nodes on different farms
  - capacity [required]
    - a string in ['small', 'medium', 'large'] indicating the capacity of the machines
    - small: 2 vCPU, 4GB RAM, 10GB SSD
    - medium: 4 vCPU, 8GB RAM, 20GB SSD
    - large: 8 vCPU, 16GB RAM, 30GB SSD
  - ssh_key [required]
  - db_username [required]
  - db_password [required]
  - admin_email [required]
  - smtp_hostname [required]
  - smtp_username [required]
  - smtp_password [required]

## Read Operation

- action name: !!tfgrid.peertube.read
- parameters:
  - model_name [required]

## Update Operations

- Update operations are not allowed on gateway names.
  
## Delete Operation

- action_name: !!tfgrid.peertube.delete
- parameters:
  - model_name [required]
