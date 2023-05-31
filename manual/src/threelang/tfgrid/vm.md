# VM Namespace

- To deploy a number of virtual machines on the same network, use the vm namespace.

## Create Operation

- action name: !!tfgrid.vm.create
- parameters:
  - network [required]
  - farm_id [optional]
    - if 0, the grid nodes will be chosen at random and the deployed vms might span multiple farms.
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
  - gateway
    - to add a gateway point to this machine on port 9000

## Get Operation

- action name: !!tfgrid.vm.get
- parameters:
  - network [required]

## Update Operations

### Add Operation

- to add a number of vms to a previously deployed network, use the Create operation above.

### Remove Operation

- action_name: !!tfgrid.vm.remove
- parameters:
  - network_name [required]
  - machine_name [required]

## Delete Operation

- action_name: !!tfgrid.vm.delete
- parameters:
  - network_name [required]
