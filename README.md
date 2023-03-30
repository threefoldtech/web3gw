# Web 3 Proxy

Project setup:

- Server (the proxy server)
- Clients (the web3 clients)
- Scripts (client scripts to invoke server commands)

## Documentation

To generate the documentation for the project, run `bash doc.sh`. This builds mdbook documentation in html format from content in manual folder in /docs. This overwrites prebuilt content in docs if any. To save your changes to the generated documentation, simply commit your changes in the manual/src folder. The CI workflow will regenerate the documentation on github pages upon pushing / merging changes to the development branch.

**While /docs is already in .gitignore, please avoid pushing generated docs.**

See [manual](/manual/readmd.md) for more info on using mdbooks for documentation.
