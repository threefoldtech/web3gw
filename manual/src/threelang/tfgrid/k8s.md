# K8s Namespace

- To deploy a kubernetes cluster, use the K8s namespace

## Create Operation

- action name: !!tfgrid.k8s.create
- parameters:
  - name [required]
  - farm_id [optional]
    - if 0, cluster nodes could span multiple nodes on different farms
  - number_of_workers [required]
    - a number in the range [1, 252]
  - capacity [required]
    - a string in ['small', 'medium', 'large', 'extra-large'] indicating the capacity of the cluster nodes
    - small: 1 vCPU, 2GB RAM, 10GB SSD
    - medium: 2 vCPU, 4GB RAM, 50GB SSD
    - large: 4 vCPU, 8GB RAM, 240 SSD
    - extra-large: 8vCPU, 16GB RAM, 480GB SSD
  - disk_size [optional]
    - size in GB of disk to be mounted on each node at "/mnt/disk"
  - ssh_key [required]
- arguments:
  - add_wireguard_access
  - add_public_ip_to_master
  - add_public_ips_to_workers

## Get Operation

- action name: !!tfgrid.k8s.get
- parameters:
  - name [required]
  - master_name [required]

## Update Operations

### Add Operation

- action_name: !!tfgrid.k8s.add
- parameters:
  - name [required]
  - farm_id [optional]
    - if 0, cluster nodes could span multiple nodes on different farms
  - capacity [required]
    - a string in ['small', 'medium', 'large', 'extra-large'] indicating the capacity of the worker
    - small: 1 vCPU, 2GB RAM, 10GB SSD
    - medium: 2 vCPU, 4GB RAM, 50GB SSD
    - large: 4 vCPU, 8GB RAM, 240 SSD
    - extra-large: 8vCPU, 16GB RAM, 480GB SSD
  - disk_size [optional]
    - size in GB of disk to be mounted on each cluster node at "/mnt/disk"
  - ssh_key [required]
- arguments:
  - public_ip

### Remove Operation

- action_name: !!tfgrid.k8s.remove
- parameters:
  - name [required]
  - master_name [required]
  - worker_name [required]

## Delete Operation

- action_name: !!tfgrid.k8s.delete
- parameters:
  - name [required]
