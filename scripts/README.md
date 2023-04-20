# Scripts

## Clients

For each client the web3 proxy supports there is a module in this directory. These modules all implement their own client that has one RpcWsClient (see below). These clients are the ones opening specific functionality of the clients. They all have functions that resembles the function below. 

A couple of important rules to follow when creating a new function. The function should use the send_json_rpc function to send data to the server. That functions takes two generic types. The first one is object that will be the params field in the json rpc 2.0 request. The second is the object it expects when receiving a json rpc 2.0 response (the object inside the result field of the json result).

The server only supports sending by order for now. This means we have to send a list of objects to the server. In V you can of course only send a list of objects of the same type. Therefore we agreed to send a list of only one object, which is the arguments that the corresponding function on the server side requires. The first generic will therefore always be a list of a specific type of object. The result (the second generic) can be any type. If you expect no result just use string and ignore the result. Always add documentation for each of the calls, this documentation is used when generating the documentation for openrpc. Documenting the structs required to send data is highly appreciated!

```
[noinit]
pub struct TfChainClient {
mut:
	client &RpcWsClient
}

pub fn new(mut client RpcWsClient) TfChainClient {
	return TfChainClient{
		client: &client
	}
}

// Load your mnemonic with this call. Choose the network while doing so. The network should be one of:
// mainnet, testnet, qanet, devnet 
pub fn (mut t TfChainClient) load(args Load) ! {
	_ := t.client.send_json_rpc[[]Load, string]('tfchain.Load', [args],
		tfchain.default_timeout)!
}

...

// Get a twin by id. 
pub fn (mut t TfChainClient) get_twin(id u32) !Twin {
	return t.client.send_json_rpc[[]u32, Twin]('tfchain.GetTwin', [id], tfchain.default_timeout)!
}

```

## Adding scripts to CI

Please always add the scripts to the build V part on the CI so that we ensure that the V code always builds.

## Creating your own script

Steps to run client:
1) Make sure you have the repo [crystallib](https://github.com/freeflowuniverse/crystallib) on branch development
2) Add the calls you want to execute in main.v in function execute_rpcs and then run:

```
v run main.v -m "<YOUR MNEMONIC COMES HERE>"
```
You can add the --debug flag to see the json data being send to the server and the response it receives.

## Testing existing functionality

For every client we have a file called *main_\<client\>.v*. You can add your test code there.

