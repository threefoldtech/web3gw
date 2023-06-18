# Load Action

> Connects to the bitcoin node. This should be the first call to execute.

- action name: !!btc.core.load
- parameters:
  - host [required]
  - user [required]
  - pass [required]

## Example

```md
    !!btc.core.load
        host: http://1.1.1.1
        user: user1
        pass: pass1
```
