#!/bin/bash
echo "Starting tests"

PROJECTPATH=$GOPATH/src/bitbucket.org/mundipagg/boletoapi

cd $PROJECTPATH

glide install

go test $(go list ./... | grep -v /vendor/) -v