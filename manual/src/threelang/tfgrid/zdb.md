# ZDB Actions

- To deploy a ZDB, use the ZDB actions

## Create Operation

- action name: !!tfgrid.zdbs.create
- parameters:
  - name [required]
  - farm_id [optional]
  - password [required]
  - size [required]
    - size of the ZDB in GB
  - public
    - yes or no
  - user_mode
    - 'seq' or 'user'. defaults to 'user'

- Example:
  
  ```
  !!tfgrid.zdbs.create 
      name: hamadazdb
      size: 10GB
      password: pass1
  ```

## Get Operation

- action name: !!tfgrid.zdbs.get
- parameters:
  - name [required]

- Example:
  
  ```
  !!tfgrid.zdbs.get
      name: hamadazdb
  ```

## Update Operations

- Update operations are not allowed on ZDBs.

## Delete Operation

- action_name: !!tfgrid.zdbs.delete
- parameters:
  - name [required]

- Example:
  
  ```
  !!tfgrid.zdbs.delete
      name: hamadazdb
  ```
