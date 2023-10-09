set -ex
go env -w CGO_ENABLED="0"
go build -o ~/go/bin/griddriver
echo build ok

