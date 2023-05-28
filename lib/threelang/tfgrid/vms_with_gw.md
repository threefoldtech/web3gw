# Machines with Gateways Namespace

- To deploy a number of virtual machines on the same network with a domain name refer to a running server on the machines, use the `vms_with_gws` namespace.

## Create Operation

- action name: !!tfgrid.vms_with_gws.create
- parameters:
  - model_name [required]
  - farm_id [optional]
    - if 0, machines could span multiple nodes on different farms
  - number_of_machines [required]
    - a number in the range [1, 252]
  - capacity [required]
    - a string in ['small', 'medium', 'large'] indicating the capacity of the machines
    - small: 1 vCPU, 2GB RAM, 10GB SSD
    - medium: 2 vCPU, 4GB RAM, 20GB SSD
    - large: 4 vCPU, 8GB RAM, 50GB SSD
  - disk_size [optional]
    - size in GB of disk to be mounted on each machine at "/mnt/disk"
  - ssh_key [required]
  - backend_port 
        the port where the gateway should point to. default is `80`
- arguments:
  - public_network
  - public_ips

## Read Operation

- action name: !!tfgrid.vms_with_gws.read
- parameters:
  - model_name [required]

## Delete Operation

- action_name: !!tfgrid.vms_with_gws.delete
- parameters:
  - model_name [required]
