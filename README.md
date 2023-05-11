# Web 3 Proxy

The web3 proxy implements a json rpc 2.0 server which opens up remote procedure calls for many clients such as tfgrid, tfchain, stellar, bitcoin, etc. Each directory in [here](server/pkg/) is a client that the web3 proxy supports.

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

To generate the documentation for the project, run `bash doc.sh`. This builds mdbook documentation in html format from content in manual folder in /docs. This overwrites prebuilt content in docs if any. To save your changes to the generated documentation, simply commit your changes in the manual/src folder. The CI workflow will regenerate the documentation on github pages upon pushing / merging changes to the development branch.

**While /docs is already in .gitignore, please avoid pushing generated docs.**

See [manual](/manual/readmd.md) for more info on using mdbooks for documentation.
