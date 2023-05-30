# Peertube Namespace

- To deploy a peertube instance, use the Peertube namespace.

## Create Operation

- action name: !!tfgrid.peertube.create
- parameters:
  - name [required]
  - farm_id [optional]
    - if 0, machines could span multiple nodes on different farms
  - capacity [required]
    - a string in ['small', 'medium', 'large', 'extra-large'] indicating the capacity of the peertube instance
    - small: 1 vCPU, 2GB RAM, 10GB SSD
    - medium: 2 vCPU, 4GB RAM, 50GB SSD
    - large: 4 vCPU, 8GB RAM, 240 SSD
    - extra-large: 8vCPU, 16GB RAM, 480GB SSD
  - ssh_key [required]
  - db_username [required]
  - db_password [required]
  - admin_email [required]
  - smtp_hostname [required]
  - smtp_username [required]
  - smtp_password [required]

## Get Operation

- action name: !!tfgrid.peertube.get
- parameters:
  - name [required]

## Update Operations

- Update operations are not allowed on gateway names.
  
## Delete Operation

- action_name: !!tfgrid.peertube.delete
- parameters:
  - name [required]
