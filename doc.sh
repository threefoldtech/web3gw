set -ex

rm -rf docs/_docs
rm -rf docs/_docs/v

pushd manual
bash run.sh
popd

# compile openrpc cli binary
v ~/.vmodules/freeflowuniverse/crystallib/openrpc/cli
OPENRPC_CLI=~/.vmodules/freeflowuniverse/crystallib/openrpc/cli/cli

# generate doc for entire api
$OPENRPC_CLI docgen -t "Web3Proxy JSON-RPC API" -p -o server lib

for file in lib/*
do
    # generate individual docs per domain
    name=${file##*/}
    $OPENRPC_CLI docgen -t "$name JSON-RPC API" -p -o server/pkg/$name $file
done

# v fmt -w .
# v doc -m -f html . -readme -comments -no-timestamp

# mv _docs docs/v/