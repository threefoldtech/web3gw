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
  - actions parameters
  - action arguments
- an action name is the string following "!!"
- action names consist of three parts separated by ".", and in this order:
  - module name
  - namespace
  - operation
- action parameters are all the key value pairs that follow an action name
- action arguments are all the single values that follow an action name
- parameters and arguments could be mixed toghether and do not have a particular order
- each model's actions are processed in order

### Module actions

- [TFGrid Actions](./tfgrid/grid_actions.md)

## Example

- if a user wants to deploy 4 machines, it could go as follows:
  
```md
    !!tfgrid.machines.create
        name: 'project1'
        network: private
        ssh_key: 'ssh_key'
        number_of_machines: 4
        capacity: medium
```

- this would deploy 4 vms, with medium capacity (cru, mru, sru, hru) on the same network.
- this is fairly simple, as we decide most of the specs for the user.
- the problem comes when the user wants to have customizable workloads, as the md parser does not recognize nested structures. this requires some helper actions to assist in executing main actions.

## Example for helper actions

- if a user wants to deploy 4 machines, each of which is different from the other, it could go as follows:

```md
    !!tfgrid.main.customizable_machines.deploy
        name: 'project1'
        network: private
        
    !!tfgrid.helper.construct.vm
        project: 'project1'
        name: 'vm1'
        memory: 2048
        cru: 4
        rootfs_size: 4096

    !!tfgrid.helper.construct.disk
        project: 'project1'
        vm_name: 'vm1'
        size: 10
        mountpoint: '/disk1'

    !!tfgrid.helper.construct.vm
        project: 'project1'
        name: 'vm2'
        memory: 1024
        cru: 2
        rootfs_size: 2048

    !!tfgrid.helper.construct.zlogs
        project: 'project1'
        vm_name: 'vm2'
        output: 'http://1.1.1.1:9000'
```

- this way, users could customize their deployments as they want, and the 3bot parser will know what to do.
- the actions have a specific format that the user has to follow:
  - first part is the module name, this will let the 3bot parser choose which module it will parse for.
  - second part is either 'main' or 'helper':
    - main actions are the objectives that the user wants acheived.
    - helper actions only help in building main actions and giving them enough information in a simple & easy way.
  - helper actions should have proper references to which main actions they belong to. in our example a disk should have a reference to which project it belongs to, and which vm it should be attached to.
  - later parts of the action may be specified by each module's parser
- after the 3bot parser is done with parsing, a user could execute all it's main actions (or maybe specific ones)
