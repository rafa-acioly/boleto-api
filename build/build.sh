#!/bin/bash
echo "O usuario que esta executando o script: $USER"
#export GOPATH="/home/mundipagg/go"
PROJECTPATH="$GOPATH/src/bitbucket.org/mundipagg/boletoapi"
echo "Changing directory to path $PROJECTPATH"
ls -la $PROJECTPATH
cd $PROJECTPATH

echo "ENV:"
env

echo "GO ENV:"
go env

echo "GO PATH:"
echo "$GOPATH"

echo "Installing dependencies with glide"
go get
go run main.go

echo "Starting build"
go build -v