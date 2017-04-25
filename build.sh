curl https://glide.sh/get | sh

git clone --branch="master" --depth 50 https://phsantiago@bitbucket.org/mundipagg/boletoapi.git

cd boletoapi

glide install

go build -v

go test $(go list ./... | grep -v /vendor/)