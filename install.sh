#!/usr/bin/env bash
set -ex
SOURCE=${BASH_SOURCE[0]}
DIR_OF_THIS_SCRIPT="$( dirname "$SOURCE" )"
ABS_DIR_OF_SCRIPT="$( realpath $DIR_OF_THIS_SCRIPT )"
mkdir -p ~/.vmodules/threefoldtech
ln -s $ABS_DIR_OF_SCRIPT/web3gw/client ~/.vmodules/threefoldtech/web3gw

# install crystallib
if !(v list | grep -q 'freeflowuniverse.crystallib'); then
    git clone https://github.com/freeflowuniverse/crystallib.git ~/.vmodules/freeflowuniverse/crystallib
fi