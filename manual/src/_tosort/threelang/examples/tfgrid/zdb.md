# ZDB Example

- This example deploys, gets, and deletes a ZDB workload on the tfgrid.

```md
!!tfgrid.core.login
 mnemonic: 'YOUR MNEMONICS'
 network: dev

!!tfgrid.zdbs.create
 name: hamadazdb
 size: 10GB
 password: pass1

!!tfgrid.zdbs.get
 name: hamadazdb

!!tfgrid.zdbs.delete
 name: hamadazdb
```md
