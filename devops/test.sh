#!/bin/bash

PROJECTPATH=$GOPATH/src/github.com/mundipagg/boletoapi
echo "Changing directory to path $PROJECTPATH"
cd $PROJECTPATH
echo "Starting tests"
go test $(go list ./... | grep -v /vendor/) -v
