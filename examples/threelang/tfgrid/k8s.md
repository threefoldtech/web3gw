# Kubernetes Example

- This example deployes, gets, updates, and deletes a kubernetes cluster on the tfgrid.

!!tfgrid.core.login
    mnemonic: 'YOUR MNEMONIC'
    network: dev

!!tfgrid.sshkeys.new
    name: default
    ssh_key: 'YOUR SSH KEY'

!!tfgrid.kubernetes.create
    name:test_3lang_k8s
    farm_id:1
    capacity:small
    add_wireguard_access
    ssh_key:default

!!tfgrid.kubernetes.get
    name:test_3lang_k8s

!!tfgrid.kubernetes.add
    name:test_3lang_k8s
    farm_id:1
    capacity:small
    ssh_key:default

!!tfgrid.kubernetes.remove
    name:test_3lang_k8s
    worker_name:w2

!!tfgrid.kubernetes.delete
    name:test_3lang_k8s
