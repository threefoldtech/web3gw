!!tfgrid.login
 	mnemonic: 'YOUR MNEMONIC'
 	network: dev

!!tfgrid.k8s.create
    name:test_3lang_k8s
    farm_id:1
    replica:2
    capacity:small
    add_wireguard_access
    ssh_key:'YOUR PUBLIC SSH KEY'

!!tfgrid.k8s.get
    name:test_3lang_k8s

!!tfgrid.k8s.add
    name:test_3lang_k8s
    farm_id:1
    capacity:small
    ssh_key:'YOUR PUBLIC SSH KEY'

!!tfgrid.k8s.remove
    name:test_3lang_k8s
    worker_name:w2

!!tfgrid.k8s.delete
    name:test_3lang_k8s
