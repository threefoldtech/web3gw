# Machines Actions

- To deploy a number of virtual machines on the same network, use the machines actions.

## Create Operation

- action name: !!tfgrid.machines.create
- parameters:
  - name [optional]
    - vm name
  - network [optional]
    - if not provided, a random network name will be generated.
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
  - ssh_key [optional]
    - the name of the ssh key defined with an sshkey action.
    - if not provided, the sshkey with the name 'default' will be used.
  - add_wireguard_access [optional]
    - yes or no to add a wireguard access point to the network
  - add_public_ips [optional]
    - yes or no to add public ips to the machines
  - gateway [optional]
    - yes or no to add a gateway point to this machine on port 9000

- Example:
  
  ```md
  !!tfgrid.machines.create
      network: skynet
      sshkey: my_ssh_key
      capacity: small
      times: 2
      gateway: yes
      add_wireguard_access: yes
  ```

## Get Operation

- action name: !!tfgrid.machines.get
- parameters:
  - network [required]

- Example:
  
  ```md
  !!tfgrid.machines.get
      name: skynet
  ```

## Update Operations

### Add Operation

- to add a number of vms to a previously deployed network, use the Create operation above.

### Remove Operation

- action_name: !!tfgrid.machines.remove
- parameters:
  - network_name [required]
  - machine_name [required]

- Example:
  
  ```md
  !!tfgrid.machines.remove
      network_name: skynet
      machine_name: vm1
  ```

## Delete Operation

- action_name: !!tfgrid.machines.delete
- parameters:
  - network_name [required]

- Example:
  
  ```md
  !!tfgrid.machines.delete
      network_name: skynet
  ```
