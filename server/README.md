# web3_proxy

## Server

### Build:

```
go build .
```

### Run:

`./server`

Server will now listen on `ws://localhost:8080`

### Examples with websocat

1. Load eth: `{"jsonrpc":"2.0", "id": 1, "method": "eth.Load", "params":["http://[2a04:7700:1003:1:4883:fff:fe19:e118]:8545","abcdefabcdefabcdefabcdefabcdefababcdefabcdefabcdefabcdefabcdefab"]}`

2. Get eth balance: `{"jsonrpc":"2.0", "id": 1, "method": "eth.Balance", "params":["0x49E02993791d762EbD2E4ac2FcA80CbAD6029be0"]}`

3. Get eth gnosis multisig contract owners: `{"jsonrpc":"2.0", "id": 1, "method": "eth.GetMultisigOwners", "params":["0x575862299B321053469f4941619393349E98E760"]}`

4. Get eth gnosis multisig threshold: `{"jsonrpc":"2.0", "id": 1, "method": "eth.GetMultisigThreshold", "params":["0x575862299B321053469f4941619393349E98E760"]}`
