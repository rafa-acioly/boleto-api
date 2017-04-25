curl https://glide.sh/get | sh

cd boletoapi

glide install

go build -v

go test $(go list ./... | grep -v /vendor/)