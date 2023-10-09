set -ex
cd ~/code/github/threefoldtech/3bot/web3gw
go env -w CGO_ENABLED="0"
go build -o ~/go/bin/web3gw
echo "build done"