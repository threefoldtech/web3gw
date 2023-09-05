# Kubernetes Example

- This example deployes, gets, updates, and deletes a kubernetes cluster on the tfgrid.

```md
!!tfgrid.core.login
    mnemonic: 'YOUR MNEMONIC'
    network: dev

!!tfgrid.sshkeys.new
    name: default
    ssh_key: 'YOUR SSH KEY'

!!tfgrid.kubernetes.create
    name: myk8s
    farm_id: 1
    workers: 2
    capacity: small

!!tfgrid.kubernetes.get
    name: myk8s

!!tfgrid.kubernetes.add
    name: myk8s
    farm_id: 1
    capacity: small
    ssh_key: default

!!tfgrid.kubernetes.remove
    name: myk8s
    worker_name: w2

!!tfgrid.kubernetes.delete
    name: myk8s
```
