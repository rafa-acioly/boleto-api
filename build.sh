PROJECTPATH=$GOPATH/src/bitbucket.org/mundipagg/boletoapi

rm -rf PROJECTPATH

mkdir $PROJECTPATH

mv ~/myagent/_work/1/s -t $PROJECTPATH -v 

mv $PROJECTPATH/s $PROJECTPATH

cd $PROJECTPATH

glide install

go build -v

go test $(go list ./... | grep -v /vendor/)