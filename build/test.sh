#!/bin/bash

if [[ "$1" != "" ]]; then
    PROJECTPATH="$1"
    echo "Changing directory to path $1"
    cd $PROJECTPATH

    echo "Starting tests"
    go test $(go list ./... | grep -v /vendor/) -v
else
    echo "[ERROR] Expecting build directory as argument"
    exit 1
fi