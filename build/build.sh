#!/bin/bash
GG="/home/mundipagg/go"
PROJECTPATH="$GG/src/bitbucket.org/mundipagg/boletoapi"
echo "Changing directory to path $PROJECTPATH"
ls -la $PROJECTPATH
cd $PROJECTPATH
echo "Installing dependencies with glide"
go get
go run main.go
echo "Starting build"
go build -v