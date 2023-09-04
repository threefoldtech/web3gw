# Machines Example

- This example deployes, gets, udpates, and deletes a network of machines on the tfgrid.

```md
!!tfgrid.core.login
mnemonic: '<MNEMONICS>'
network: dev

!!tfgrid.sshkeys.new
name: SSH_KEY
ssh_key: ''

!!tfgrid.network.create
network: skynet
capacity: small
times: 2
gateway: yes
add_wireguard_access: yes
disk_size: 10GB

!!tfgrid.network.create
network: skynet
capacity: medium

!!tfgrid.network.remove
network: skynet
machine: ewbjpuqe

!!tfgrid.network.get
network: skynet

!!tfgrid.network.delete
network: skynet
```
