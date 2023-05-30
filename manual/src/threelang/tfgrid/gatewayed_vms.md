# Machines with Gateways Namespace

- To deploy a number of virtual machines on the same network with a domain name refer to a running server on the machines, use the `gatewayed_vms` namespace.

## Create Operation

- action name: !!tfgrid.gatewayed_vms.create
- parameters:
  - name [required]
  - farm_id [optional]
    - if 0, machines could span multiple nodes on different farms
  - times [optional]
    - indicates how many machines to deploy
    - default is one
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
  - backend_port
        the port where the gateway should point to. default is `80`
- arguments:
  - add_wireguard_access
  - public_ips

## Get Operation

- action name: !!tfgrid.gatewayed_vms.get
- parameters:
  - name [required]

## Delete Operation

- action_name: !!tfgrid.gatewayed_vms.delete
- parameters:
  - name [required]
