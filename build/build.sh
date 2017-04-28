#!/bin/bash

if [[ "$1" != "" ]]; then
    PROJECTPATH="$1"
    echo "Changing directory to path $1"
    cd $PROJECTPATH

    echo "Installing dependencies with glide"
    pwd
    glide install

    echo "Starting build"
    go build -v
else
    echo "[ERROR] Expecting build directory as argument"
    exit 1
fi