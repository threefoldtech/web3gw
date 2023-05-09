# Web 3 Proxy

Project setup:

- Server (the proxy server)
- Lib (the clients in V lang)
- Examples (example scripts)

## Installing

```sh
./install.sh
```

This creates a symbolic link from ~/.vmodules/threefoldtech/threebot to the lib folder of this repository.

## How to run the web3 proxy

First start the server:

```sh
cd server
go build
./server --debug
```

Then you can go through the documentation under [scripts](scripts/)

## Examples

See `./examples` folder for examples of how to interact with the proxy.

## Documentation

To generate the documentation for the project, run `bash doc.sh`. This builds mdbook documentation in html format from content in manual folder in /docs. This overwrites prebuilt content in docs if any. To save your changes to the generated documentation, simply commit your changes in the manual/src folder. The CI workflow will regenerate the documentation on github pages upon pushing / merging changes to the development branch.

**While /docs is already in .gitignore, please avoid pushing generated docs.**

See [manual](/manual/readmd.md) for more info on using mdbooks for documentation.
