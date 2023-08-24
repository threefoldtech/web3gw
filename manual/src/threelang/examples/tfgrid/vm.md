# Machines Example

- This example deployes, gets, udpates, and deletes a network of machines on the tfgrid.

```md
!!tfgrid.core.login
 mnemonic: 'YOUR MNEMONICS'
 network: dev

!!tfgrid.sshkeys.new
 name: default
 ssh_key: 'YOUR SSH KEY'

!!tfgrid.machine.create
 network: skynet
 capacity: small
 times: 2
 gateway: yes
 add_wireguard_access: yes
 disk_size: 10GB

!!tfgrid.machine.create
 network: skynet
 capacity: medium

!!tfgrid.machine.remove
 network: skynet
 machine: ewbjpuqe

!!tfgrid.machine.get
 network: skynet

!!tfgrid.machine.delete
 network: skynet
```
