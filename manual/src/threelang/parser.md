# 3bot parser

- the objective of this parser is to be able to parse and execute actions from md files on different tf modules
- we already have md file parser that could extract actions and their parameters.
- these actions could be passed to the 3bot parser to parse and execute the user's requests
- the 3bot parser should have smaller inner processors for each tf module

## Actions

- an action is indicated in an md file by a line starting with !!
- actions are delimited by new lines.
- actions consist of:
  - action names
  - action parameters
  - action arguments
- an action name is the string following the "!!"
- action names consist of three parts separated by a ".", in this order:
  - module name
  - namespace
  - operation
- action parameters are all the key value pairs that follow an action name
- action arguments are all the single values that follow an action name
- parameters and arguments could be mixed toghether and do not have a particular order

### Module actions

- [TFGrid Actions](./tfgrid/grid_actions.md)

## Example

- if a user wants to deploy a group of 4 machines on the same network:
  
```md
    !!tfgrid.machine.create
        name: 'my machines'
        ssh_key: 'ssh_key'
        times: 4
        capacity: medium
```

- this would deploy 4 vms, with medium capacity (cru, mru, sru) on the same network.
