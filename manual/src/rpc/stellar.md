
# Stellar

TODO: intro

## Remote Procedure Calls

### Load

- secret = The secret of the account that will be used to sign transactions. If empty, a keypair will be generated.
- network = The network to use, either `testnet` or `public`

****Request****

```
{
    "jsonrpc": "2.0",
    "method": "stellar.Load",
    "params": {
        "secret": string,
        "network": string
    },
    "id": "<GUID>"
}
```

**Response**

```
{
    "jsonrpc": "2.0",
    "result": "",
    "id": "<GUID>"
}
```

### Transfer

****Request****

```
{
    "jsonrpc": "2.0",
    "method": "stellar.Transfer",
    "params": {
        "destination": string,
        "memo": string,
        "amount": string
    },
    "id": "<GUID>"
}
```

**Response**

```
{
    "jsonrpc": "2.0",
    "result": "",
    "id": "<GUID>"
}
```

### GetBalance

Get the TFT balance for an account, if the passed address is an empty string, the balance of the account of the client is returned.

**Request**

```
{
    "jsonrpc": "2.0",
    "method": "stellar.Balance",
    "params": "<address>",
    "id": "<GUID>"
}
```

**Response**

```
{
    "jsonrpc": "2.0",
    "result": i64,
    "id": "<GUID>"
}
```
