# Load Action

> Load a client, connecting to the rpc endpoint at the given network and loading a keypair from the given secret.

- action name: !!stellar.core.load
- parameters:
  - secret [require]
  - network [optional]
    - if not provided, defaults to 'public'

## Example

```md
    !!stellar.core.load
        secret: my secret
        network: public
```
