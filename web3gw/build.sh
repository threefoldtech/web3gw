set -ex
go env -w CGO_ENABLED="0"
go build -o ~/go/bin/web3gw-server
echo "build done"