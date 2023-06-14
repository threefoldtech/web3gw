# Machines Example

- This example deployes, gets, udpates, and deletes a network of machines on the tfgrid.

!!tfgrid.core.login
	mnemonic: 'YOUR MNEMONICS'
	network: dev

!!tfgrid.sshkeys.new
	name: default
	ssh_key: 'YOUR SSH KEY'

!!tfgrid.machines.create
	network: skynet
	capacity: small
	times: 2
	gateway: yes
	add_wireguard_access: yes
	disk_size: 10GB

!!tfgrid.machines.create
	network: skynet
	capacity: medium

!!tfgrid.machines.remove
	network: skynet
	machine: ewbjpuqe

!!tfgrid.machines.get
	network: skynet

!!tfgrid.machines.delete
	network: skynet