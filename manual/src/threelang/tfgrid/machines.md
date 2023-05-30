# Machines Namespace

- To deploy a number of virtual machines on the same network, use the machines namespace.

## Create Operation

- action name: !!tfgrid.machine.create
- parameters:
  - network_name [required]
  - farm_id [optional]
    - if 0, machines could span multiple nodes on different farms
  - times [required]
    - indicates how many machines to deploy
    - a number in the range [1, 252]
  - capacity [required]
    - a string in ['small', 'medium', 'large', 'extra-large'] indicating the capacity of the machines
    - small: 1 vCPU, 2GB RAM, 10GB SSD
    - medium: 2 vCPU, 4GB RAM, 50GB SSD
    - large: 4 vCPU, 8GB RAM, 240 SSD
    - extra-large: 8vCPU, 16GB RAM, 480GB SSD
  - disk_size [optional]
    - size in GB of disk to be mounted on each machine at "/mnt/disk"
  - ssh_key [required]
- arguments:
  - add_wireguard_access
  - public_ips

## Get Operation

- action name: !!tfgrid.machines.get
- parameters:
  - network_name [required]

## Update Operations

### Add Operation

- action_name: !!tfgrid.machines.add
- parameters:
  - network_name [required]
  - farm_id [optional]
    - if 0, machines could span multiple nodes on different farms
  - capacity [required]
    - a string in ['small', 'medium', 'large', 'extra-large'] indicating the capacity of the machines
    - small: 1 vCPU, 2GB RAM, 10GB SSD
    - medium: 2 vCPU, 4GB RAM, 50GB SSD
    - large: 4 vCPU, 8GB RAM, 240 SSD
    - extra-large: 8vCPU, 16GB RAM, 480GB SSD
  - disk_size [optional]
    - size in GB of disk to be mounted on each machine at "/mnt/disk"
  - ssh_key [required]
- arguments:
  - public_ips

### Remove Operation

- action_name: !!tfgrid.machines.remove
- parameters:
  - network_name [required]
  - machine_name [required]

## Delete Operation

- action_name: !!tfgrid.machines.delete
- parameters:
  - network_name [required]
