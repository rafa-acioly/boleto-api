#!/bin/bash

PROJECTPATH=$GOPATH/src/bitbucket.org/mundipagg/boletoapi
echo "Changing directory to path $PROJECTPATH"
cd $PROJECTPATH
echo "Installing dependencies with glide"
glide install
echo "Starting build"
go build -v