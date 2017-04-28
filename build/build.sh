echo "Deletando arquivos do repositório no GOPATH"
PROJECTPATH=$GOPATH/src/bitbucket.org/mundipagg/boletoapi
rm -rdfv PROJECTPATH/*

echo "Criando diretório do repositório no GOPATH"
mkdir -p $PROJECTPATH

echo "Movendo arquivos do repositório do workspace do agente para o GOPATH"
mv -v ~/myagent/_work/1/s/* -t $PROJECTPATH 
mv $PROJECTPATH/s $PROJECTPATH/boletoapi -v

echo "Mudando para o diretório no repositório no GOPATH"
cd $PROJECTPATH

echo "Instalando dependências com o glide"
glide install

echo "Fazendo o build do projeto"
go build -v