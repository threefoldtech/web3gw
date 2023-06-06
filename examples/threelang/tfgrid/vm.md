# Machines Example

- This example deployes, gets, udpates, and deletes a network of machines on the tfgrid.

!!tfgrid.core.login
 	mnemonic: 'YOUR MNEMONICS'
	network: dev

!!tfgrid.sshkeys.new
	name: default
	ssh_key: 'YOUR SSH KEY'

!!tfgrid.machines.create
	sshkey: default
	network: skynet
	capacity: small
	times: 2
	gateway: yes
	add_wireguard_access: yes

!!tfgrid.machines.create
	network: skynet
	capacity: medium
	sshkey: default

!!tfgrid.machines.remove
	network: skynet
	machine: ewbjpuqe

!!tfgrid.machines.get
	network: skynet

!!tfgrid.machines.delete
	network: skynet