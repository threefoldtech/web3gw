# K8s Namespace

- To deploy a kubernetes cluster, use the K8s namespace

## Create Operation

- action name: !!tfgrid.k8s.create
- parameters:
  - model_name [required]
  - farm_id [optional]
    - if 0, cluster nodes could span multiple nodes on different farms
  - number_of_workers [required]
    - a number in the range [1, 252]
  - capacity [required]
    - a string in ['small', 'medium', 'large'] indicating the capacity of each cluster node
    - small: 2 vCPU, 4GB RAM, 10GB SSD
    - medium: 4 vCPU, 8GB RAM, 20GB SSD
    - large: 8 vCPU, 16GB RAM, 30GB SSD
  - disk_size [optional]
    - size in GB of disk to be mounted on each node at "/mnt/disk"
  - ssh_key [required]
- arguments:
  - public_network
  - master_public_ip
  - workers_public_ips

## Read Operation

- action name: !!tfgrid.k8s.read
- parameters:
  - model_name [required]
  - master_name [required]

## Update Operations

### Add Operation

- action_name: !!tfgrid.k8s.add
- parameters:
  - model_name [required]
  - farm_id [optional]
    - if 0, cluster nodes could span multiple nodes on different farms
  - capacity [required]
    - a string in ['small', 'medium', 'large'] indicating the capacity of each cluster node
    - small: 2 vCPU, 4GB RAM, 10GB SSD
    - medium: 4 vCPU, 8GB RAM, 20GB SSD
    - large: 8 vCPU, 16GB RAM, 30GB SSD
  - disk_size [optional]
    - size in GB of disk to be mounted on each cluster node at "/mnt/disk"
  - ssh_key [required]
- arguments:
  - public_ip

### Remove Operation

- action_name: !!tfgrid.k8s.remove
- parameters:
  - model_name [required]
  - master_name [required]
  - worker_name [required]

## Delete Operation

- action_name: !!tfgrid.k8s.delete
- parameters:
  - model_name [required]
