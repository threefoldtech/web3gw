# ZDB Namespace

- To deploy a ZDB, use the ZDB namespace

## Create Operation

- action name: !!tfgrid.zdb.create
- parameters:
  - name [required]
  - farm_id [optional]
  - password [required]
  - size [required]
    - size of the ZDB in GB
- arguments:
  - public
  - user_mode

## Get Operation

- action name: !!tfgrid.zdb.get
- parameters:
  - name [required]

## Update Operations

- Update operations are not allowed on ZDBs.

## Delete Operation

- action_name: !!tfgrid.zdb.delete
- parameters:
  - name [required]
