PROJECTPATH=$GOPATH/src/bitbucket.org/mundipagg/boletoapi

rm -rf PROJECTPATH -v

mkdir $PROJECTPATH

mv -v ~/myagent/_work/1/s/* -t $PROJECTPATH 

mv $PROJECTPATH/s $PROJECTPATH/boletoapi -v

cd $PROJECTPATH

glide install

go build -v

go test $(go list ./... | grep -v /vendor/) -v