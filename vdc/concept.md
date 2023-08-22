# VDC technical concept

## definition 3script

A VDC is defined through 3script. It's a declarative language to describe a VDC and is parsed into an in-memory model.

**What** instead of **How-to**.

## query 3script

3script that queries the deployed state on the grid (reality). The output can be json for machine parsing or markdown for human interpretation.

## action 3script

A 3script with the actions needed to bring the reality to the desired state defined in the definition 3script.
These are 3script actions executed through the web3gw.

## from definition to reality

1. The definition 3script is parsed to an in memory model.
2. A query 3script is generated to query the reality.
3. The query 3script is executed and an action 3script is generated (the web3gw sal is not called directly to align relity to the model)
4. The action 3script is executed to bring reality in the desired state
