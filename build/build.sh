#!/bin/bash

# PROJECTPATH=$GOPATH/src/bitbucket.org/mundipagg/boletoapi

# echo "Mudando para o diretório no repositório no GOPATH"
# cd $PROJECTPATH

pwd

echo "Instalando dependências com o glide"
glide install

echo "Fazendo o build do projeto"
go build -v