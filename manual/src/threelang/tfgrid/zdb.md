# ZDB Namespace

- To deploy a ZDB, use the ZDB namespace

## Create Operation

- action name: !!tfgrid.zdb.create
- parameters:
  - model_name [required]
  - farm_id [optional]
  - password [required]
  - size [required]
    - size of the ZDB in GB
- arguments:
  - public
  - user_mode

## Read Operation

- action name: !!tfgrid.zdb.read
- parameters:
  - model_name [required]

## Update Operations

- Update operations are not allowed on ZDBs.

## Delete Operation

- action_name: !!tfgrid.zdb.delete
- parameters:
  - model_name [required]
