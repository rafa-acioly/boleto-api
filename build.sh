curl https://glide.sh/get | sh

PROJECTPATH=$GOPATH/src/bitbucket.org/mundipagg

mkdir $PROJECTPATH

mv ~/myagent/_work/1/s -t $PROJECTPATH

mv $PROJECTPATH/s $PROJECTPATH/boletoapi

cd $PROJECTPATH/boletoapi

glide install

go build -v

go test $(go list ./... | grep -v /vendor/)