echo "Rodando os testes"
PROJECTPATH=$GOPATH/src/bitbucket.org/mundipagg/boletoapi
cd $PROJECTPATH
go test $(go list ./... | grep -v /vendor/) -v