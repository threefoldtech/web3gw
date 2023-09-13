# 3bot

> :warning: The repository has recently modified its name from web3_proxy to 3bot. Although Github will redirect the fetches, pushes, etc to the new name they do strongly recommend to change the remote:
>
> ```git remote set-url origin git@github.com:threefoldtech/3bot.git```

The 3bot implements a json rpc 2.0 server which opens up remote procedure calls for many clients such as tfgrid, tfchain, stellar, bitcoin, etc. Each directory in [here](server/pkg/) is a client that the web3 proxy supports.

Project setup:

- [Server](server/): the implementation for the json rpc 2.0 server
- [Lib](lib/): the V library that implements the client for the json rpc 2.0 server, each of the server's methods can be called from that library
- [Examples](examples): contains some examples of V scripts for each client that the json rpc 2.0 server supports, these scripts show you how you can use the V library

## How to run the web3 proxy server (json rpc 2.0 server)

The server is implemented in go and is located in [this folder](server/). To build and run the server execute these commands:

```shell
cd server
go build && ./server --debug
```

The json rpc 2.0 server should now be up and running and should allow you to call any of the clients that the server supports.

## Installing the V library

If you wish to use the V library to contact the server you first have to install the V libary. Please execute the command shown below. The V library has a dependency on [crystallib](https://github.com/freeflowuniverse/crystallib). The bash scrips show below will do that for you unless you install it first. Installing it first allows you to be able to use different branches from crystallib as runtime code.

```sh
./install.sh
```

## Examples

You should now be able to run the examples under [examples](examples/). They should give you an idea of how you can use the V library to interact with the proxy.

## Adding new clients

Follow the steps [here](server/) to add the client to the json rpc 2.0 server and then look [here](lib/) to add the necessary code for the V library. This repository contains a pipeline that builds and tests both the server and V clients. This should be green for development at all times. Whenever we add a new client to the web3 proxy it should be build by CI to make sure that it is building at all times.

## Documentation

To generate the documentation for the project, run `v run doc.vsh`. This builds: 
- MDBook Documentation in html format from content in manual folder in `docs`. 
- OpenRPC Documents for the JSON-RPC API's at `docs/openrpc` from V clients in `lib`.
- OpenRPC Playground for the JSON-RPC API's at `docs/playground`.

Find out more about how comments in code are used to generate OpenRPC Documents for domains, and how to annotate your code accordingly [here](https://github.com/freeflowuniverse/crystallib/tree/development/openrpc)

To locally generate specific documents and not all of the aforementioned artifacts, comment out the [lines](https://github.com/threefoldtech/3bot/blob/596331a5051d15502681d200fa408ee0983debc0/doc.vsh#LL88-L91) in the doc.vsh script accordingly.

Note that running this command overwrites prebuilt content in docs if any, and is not necessary beyond testing locally as the script is run in CI workflow upon pushing / merging changes to the development branch. The docs are generated automatically and are made available on Github Pages at the [projects page](https://threefoldtech.github.io/3bot)

Links to generated documents:
- [Manual](https://threefoldtech.github.io/3bot)
- [OpenRPC Document](https://threefoldtech.github.io/3bot/openrpc/openrpc.json) for all clients
- OpenRPC Document for each client: https://threefoldtech.github.io/3bot/openrpc/<client_name>/openrpc.json
- [OpenRPC Playground](https://threefoldtech.github.io/3bot/playground/)

**While /docs is already in .gitignore, please avoid pushing generated docs.**

See [manual](/3script/manual/readmd.md) for more info on using mdbooks for documentation.
