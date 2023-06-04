!!tfgrid.core.login
 	mnemonic: 'route visual hundred rabbit wet crunch ice castle milk model inherit outside'
 	network: dev

!!tfgrid.sshkeys.new
    name: default
	ssh_key: 'ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCQi7Qp0fs4WowSBQJonYHNWNJ5q776XbFO8HnUggyGse1Z4CFZyVpUdWaIzpFkQdivAACSKmqfE6jAunX7HuujTQhLhVgs/iCQ3ALQfQ118Jh1g2dz7S3/klHJs6JqfXLKtwDHzq2DuEDjls5PPoD6SVipjQo+kFO2tvKUYOrXryGW5VNPSUKtXZJX4kxtLzCANqENMSqZIBiJhXj7+JQq8Kp6E117dkLxh4BmPJmGS4stSAfgSdmEWgm0MgAbHkc2X+fLsWrcEBYaXE1b+n70bVXGDVEfeuMNZjBlsgULVR0DXY5zxegciOSNr1b7yF/ZdoALN0gmQg+AywPy92+oeY7EXLabDoDUKcE+42EHscXEkTHlhCieF/W9worCzGqpMwJuBDNvDu5kP1y/vB+ZfPVTlZ1kS77/OuDTr/zssQI/SgSszVXTyVSFIFIbXLGuUDscnmPHVPV4PnmeOa2aeF1cgX0o/JErQ8+iu2wqQKueZT4QAUFyknIgXloSBVs= mariocs@mario-codescalers'

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
