set -ex

#!/bin/bash

# Get the current Go version
current_version=$(go version | awk '{print $3}' | sed 's/go//')

# Define the minimum required version
min_version="1.21.3"

# Function to compare versions
version_gt() { test "$(printf '%s\n' "$@" | sort -V | head -n 1)" != "$1"; }

# Check if the current version is greater than or equal to the minimum version
if version_gt "$min_version" "$current_version"; then
    echo "Your Go version is older than $min_version. Please update it."
    exit 1
else
    echo "Your Go version is $current_version, which is up to date."
fi


go env -w CGO_ENABLED="0"
go build -o ~/go/bin/griddriver
echo build ok

sudo cp ~/go/bin/griddriver /usr/local/bin