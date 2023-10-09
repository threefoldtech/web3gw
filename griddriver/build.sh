set -ex
cd ~/code/github/threefoldtech/3bot/griddriver
go env -w CGO_ENABLED="0"
go build -o griddriver