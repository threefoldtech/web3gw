# TFGRID CLI
A command line interface for the ThreeFold Grid.

## Usage
```bash
v run . [module] [operation] [flags]
```
To get manual for each module and avilable operations

- get available modules:
    `v run .`
- get available operations for specific module:
    `v run . vms`
- get available flags for operation
    `v run . vms create -help`

## Example
- deploy three vms on the same network with gateways
    ```bash
    v run . vms create \
        -grid dev \
        -mnemonic 'mom picnic deliver again rug night rabbit music motion hole lion where' \
        -network test_cli \
        -capacity small \
        -times 3 \
        -gateway true
    ```