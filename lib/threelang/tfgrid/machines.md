# Machines Namespace

- To deploy a number of virtual machines on the same network, use the machines namespace.

## Create Operation

- action name: !!tfgrid.machines.create
- parameters:
  - model_name [required]
  - farm_id [optional]
    - if 0, machines could span multiple nodes on different farms
  - number_of_machines [required]
    - a number in the range [1, 252]
  - capacity [required]
    - a string in ['small', 'medium', 'large'] indicating the capacity of the machines
    - small: 2 vCPU, 4GB RAM, 10GB SSD
    - medium: 4 vCPU, 8GB RAM, 20GB SSD
    - large: 8 vCPU, 16GB RAM, 30GB SSD
  - disk_size [optional]
    - size in GB of disk to be mounted on each machine at "/mnt/disk"
  - ssh_key [required]
- arguments:
  - public_network
  - public_ips

## Read Operation

- action name: !!tfgrid.machines.read
- parameters:
  - model_name [required]

## Update Operations

### Add Operation

- action_name: !!tfgrid.machines.add
- parameters:
  - model_name [required]
  - farm_id [optional]
    - if 0, machines could span multiple nodes on different farms
  - capacity [required]
    - a string in ['small', 'medium', 'large'] indicating the capacity of the machines
    - small: 2 vCPU, 4GB RAM, 10GB SSD
    - medium: 4 vCPU, 8GB RAM, 20GB SSD
    - large: 8 vCPU, 16GB RAM, 30GB SSD
  - disk_size [optional]
    - size in GB of disk to be mounted on each machine at "/mnt/disk"
  - ssh_key [required]
- arguments:
  - public_ips

### Remove Operation

- action_name: !!tfgrid.machines.remove
- parameters:
  - model_name [required]
  - machine_name [required]

## Delete Operation

- action_name: !!tfgrid.machines.delete
- parameters:
  - model_name [required]
