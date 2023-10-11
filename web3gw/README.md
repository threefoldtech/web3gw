# web3_proxy

## Server

### Build and Run

Inside the web3gw directory, run:

```sh
cd web3gw
./build.sh &&  ~/go/bin/web3gw-server --debug 
```

The server will now listen on `ws://localhost:8080`

### Remote procedure calls

Each client is located inside its own folder. We usually create a client.go file inside the client's folder. In that file we provide remote procedure calls that look like the code snippet below.

A couple of things to notice:

- We always have one function (usually called Load) which creates the client object and saves it in the state of the session for future calls.
- Functions other than the "Load" function can retrieve the client from the conState object (which is passed to the function when calling it).
- We use only one argument (next to the context and the connection state) that represents the required argument(s) for that remote procedure call (this is the decoded parameters from the params field of the incoming json 2.0 rpc). If a rpc requires more then one argument a new struct should be created to collect those arguments. You can look at the existing clients for inspiration.
- Each function returns either error or (<SOME_OBJECT>, error). If the latter is used the returned object will be serialized and put into the result field of the json 2.0 rpc response.

```go
func (c *Client) ServiceContractSetMetadata(ctx context.Context, conState jsonrpc.State, args ServiceContractSetMetadata) error {
 log.Debug().Msgf("Tfchain: setting metadata for service contract %s", args.ContractID)

 state := State(conState)
 if state.client == nil {
  return false, pkg.ErrClientNotConnected{}
 }

 return state.client.ServiceContractSetMetadata(*state.identity, args.ContractID, args.Metadata)
}
```

### Adding more clients

1) Create a folder for the client and add a client.go file in that folder
2) Add the functions following the guidelines mentioned above
3) Register the client in main.go
4) Add the V client in [the V library](../lib)

### Examples with websocat

1. Load eth: `{"jsonrpc":"2.0", "id": 1, "method": "eth.Load", "params":["http://[2a04:7700:1003:1:4883:fff:fe19:e118]:8545","abcdefabcdefabcdefabcdefabcdefababcdefabcdefabcdefabcdefabcdefab"]}`

2. Get eth balance: `{"jsonrpc":"2.0", "id": 1, "method": "eth.Balance", "params":["0x49E02993791d762EbD2E4ac2FcA80CbAD6029be0"]}`

3. Get eth gnosis multisig contract owners: `{"jsonrpc":"2.0", "id": 1, "method": "eth.GetMultisigOwners", "params":["0xa1c47964b774A977CAda6EFC80a14d833630ac38"]}`

4. Get eth gnosis multisig threshold: `{"jsonrpc":"2.0", "id": 1, "method": "eth.GetMultisigThreshold", "params":["0xa1c47964b774A977CAda6EFC80a14d833630ac38"]}`

5. Get eth token balance for custom token: `{"jsonrpc":"2.0", "id": 1, "method": "eth.GetTokenBalance", "params":["0x403351d9a97b48B290bCE1bF1d8797812Ae527DF", "0xbD330A6F55518b5dc6B984c01dd7f023775fbe7d"]}`
