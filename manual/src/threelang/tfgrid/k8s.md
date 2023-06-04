# Kubernetes actions

- To deploy a kubernetes cluster, use the kubernetes actions

## Create Operation

- action name: !!tfgrid.kubernetes.create
- parameters:
  - name [required]
  - farm_id [optional]
    - if 0, cluster nodes could span multiple nodes on different farms
  - workers [required]
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
  - add_wireguard_access
    - yes or no to add wireguard access point
  - add_public_ip_to_master
    - yes or no to add a public ip to the master node
  - add_public_ips_to_workers
    - yes or no to add public ips to workers

- Example:
  
  ```
  !!tfgrid.kubernetes.create
      name: myk8s
      farm_id: 1
      workers: 4
      capacity: small
      disk_size: 5GB
      add_public_ip_to_master: yes
  ```

## Get Operation

- action name: !!tfgrid.kubernetes.get
- parameters:
  - name [required]

- Example:
  
  ```
  !!tfgrid.kubernetes.get
      name: myk8s
  ```

## Update Operations

### Add Operation

- action_name: !!tfgrid.kubernetes.add
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
  - add_public_ip
    - yes or no to add a public ip to the new worker

- Example:
  
  ```
  !!tfgrid.kubernetes.add
      name: myk8s
      farm_id: 2
      capacity: small
      disk_size: 10GB
  ```

### Remove Operation

- action_name: !!tfgrid.kubernetes.remove
- parameters:
  - name [required]
  - worker_name [required]

- Example:
  
  ```
  !!tfgrid.kubernetes.remove
      name: myk8s
      worker_name: worker1
  ```

## Delete Operation

- action_name: !!tfgrid.kubernetes.delete
- parameters:
  - name [required]

- Example:
  
  ```
  !!tfgrid.kubernetes.delete
      name: myk8s
  ```
