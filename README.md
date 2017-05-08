# Boleto API

Consulte o [FAQ](./FAQ.md)


O que é a Api de BoletoOnline?
--------------

BoletoOnline é uma api para registro online de boleto junto ao banco e a geração do boleto para pagamento


Bancos suportados:

* Banco do Brasil
* Caixa(em breve)
* Citi (em breve)
* Santander (em breve)
* Bradesco (em breve)
* Itau (em breve)

A ordem de integração seguirá a lista superior mas pode haver modificação na prioridade dependendo do cliente

Construindo a Api
--------------

A Api foi desenvolvida utilizando a linguagem GO e portanto é necessário instalar as ferramentas da linguagem caso queira compilar a aplicação
O Go pode ser baixado [aqui](https://golang.org/dl/)
Antes de fazer o clone do Projeto você deve criar o caminho de pastas dentro do $GOPATH

	% mkdir -p "$GOPATH/src/bitbucket.org/mundipagg"
	% cd $GOPATH/src/bitbucket.org/mundipagg 
	% git clone https://bitbucket.org/mundipagg/boletoapi

Antes de compilar a aplicação você deve instalar o [Glide](http://glide.sh/). O Glide é o gerenciador de dependências da aplicação.
Após a instalaçãoi do Glide você pode executar o comando para baixar todas as dependências da aplicação

	% glide install

Para compilar a aplicação você deve executar o comando

	% go build

Rodando a aplicação
-------------

Para rodar a api com as configurações default
Ex: 
Linux (*NIX):

	% ./boletoapi
	
Windows:

	% boletoapi.exe

Se você quiser rodar a api em modo dev, que irá carregar as variáveis de ambiente padrão você deve executar a aplicação da seguinte forma:

	% ./boletoapi -dev

Se você quiser rodar a aplicação em modo mock para não realizar diretamente a integração com o banco e usar uma base de dados em memória você deve usar a opção mock:

	% ./boletoapi -mock

Você também pode combinar as duas opções:

	% ./boletoapi -dev -mock
	

Usando a API de boleto online
------------------

Você pode usar o [Postman](https://chrome.google.com/webstore/detail/postman/fhbjgbiflinjbdggehcddcbncdddomop) para criar chamar os serviços da api ou mesmo o curl

```
% curl -X POST \
  http://localhost:3000/v1/boleto/register \
  -d '{
    "Authentication" : {
        "Username":"eyJpZCI6IjgwNDNiNTMtZjQ5Mi00YyIsImNvZGlnb1B1YmxpY2Fkb3IiOjEwOSwiY29kaWdvU29mdHdhcmUiOjEsInNlcXVlbmNpYWxJbnN0YWxhY2FvIjoxfQ",
        "Password":"eyJpZCI6ImY1NzViYjgtYjBiNy00YSIsImNvZGlnb1B1YmxpY2Fkb3IiOjEwOSwiY29kaWdvU29mdHdhcmUiOjEsInNlcXVlbmNpYWxJbnN0YWxhY2FvIjoxLCJzZXF1ZW5jaWFsQ3JlZGVuY2lhbCI6MX0"
    },
    "Agreement":{
        "AgreementNumber":1014051,
        "WalletVariation":19,
        "Wallet":17,
        "Agency":"1233",
        "AgencyDigit":"2",
        "Account":"1231231",
        "AccountDigit":"3"
    },
    "Title":{
      "ExpireDate": "2017-05-25",
        "AmountInCents":30000,
        "OurNumber":101405187,
        "Instructions":"Instruções"
    },
    "Buyer":{
        "Name":"BoletoOnlione",
        "Document": {
            "Type":"CNPJ",
            "Number":"73400584000166"
        },
        "Address":{
            "Street":"Rua Teste",
            "Number": "11",
            "Complement":"",
            "ZipCode":"12345678",
            "City":"Rio de Janeiro",
            "District":"Melhor bairro",
            "StateCode":"RJ"
        }
    },
    "Recipient":{
        "Name":"Nome do Recebedor",
        "Document": {
            "Type":"CNPJ",
            "Number":"12312312312366"
        },
        "Address":{
            "Street":"Rua do Recebedor",
            "Number": "322",
            "Complement":"2º Piso loja 404",
            "ZipCode":"112312342",
            "City":"Rio de Janeiro",
            "District":"Outro bairro",
            "StateCode":"RJ"
        }
    },
    "BankNumber":1

}'

```
Resposta de sucesso da API
```
{
  "Url": "http://localhost:3000/boleto?fmt=html&id=g8HXWatft9oMLdTMAqzxbnPYFv3sqgV_KD0W7j8Cy9nkCLZMIK1WH2p9JwP1Jzz4ZtohmQ==",
  "DigitableLine": "00190000090101405100500066673179971340000010000",
  "BarCodeNumber": "00199713400000100000000001014051000006667317"
}

```
Caso aconteça algum erro no processo de registro online a resposta entregue pela api seguirá o seguinte padrão
```
{
  "Errors": [
    {
      "Code": "MPExpireDate",
      "Message": "Data de expiração não pode ser menor que a data de hoje"
    }
  ]
}
```

Instalando a API
-----------------

Para instalar o executável da API você precisa apenas compilar a aplicação e configurar as variáveis de ambiente necessárias
para o seu ambiente.
Edite o arquivo $HOME/.bashrc.sh
```
    export API_PORT="3000"
    export API_VERSION="0.0.1"
    export ENVIRONMENT="Development"
    export SEQ_URL="http://192.168.8.119:5341"
    export SEQ_API_KEY="4jZzTybZ9bUHtJiPdh6"
    export ENABLE_REQUEST_LOG="false"
    export ENABLE_PRINT_REQUEST="true"
    export URL_BB_REGISTER_BOLETO="https://cobranca.desenv.bb.com.br:7101/registrarBoleto"
    export URL_BB_TOKEN="https://oauth.desenv.bb.com.br:43000/oauth/token"
    export MONGODB_URL="10.0.2.15:27017"
    export APP_URL="http://localhost:8080/boleto"
```
    % go build && mv boletoapi /usr/local/bin

Desta forma você irá instalar de forma local o executável da API

Instalando a API via Docker
-----------------

Antes de fazer o deploy você deve abrir o arquivo [docker-compose](/devops/docker-compose.yml) e configurar com as informações que sejam pertinentes ao seu ambiente. Após ajustar o docker-compose você pode instalar a aplicação usando o arquivo deploy.sh

    % cd devops
    % ./deploy.sh

O script irá criar os diretórios de volume do Docker, buildar a aplicação montar as imagens da API e do MongoDB e subir os containers. Para mais informações sobre docker-compose consulte a [doc](https://docs.docker.com/compose/)
Após levantada a apĺicação você pode parar ou iniciar os containers.
    
    % cd devops/
    % ./stop.sh
    % ./start.sh
 
Backup e Restore
----------------- 

Para realizar o backup da base do MongoDB usa executa o seguinte comando:

    % cd devops
    % ./doBackup.sh

Os backups gerados, por padrão, serão armazenados no diretório `$HOME/backups` com o nome `bck_boletoapi-2017-05-08.tar`.
Para restaurar um backup:
    
    % cd devops
    % ./doRestore.sh

Quando fizer o restore o script irá solicitar a data do arquivo de restore e deverá ser informada uma data válida do backup no padrão: `2017-05-08`.
    
Como contrubuir
-----------------

Para contrubuir dê uma olhada no arquivo [CONTRIBUTING](CONTRIBUTING.md)

Layout do Código Fonte
---

A Raiz da aplicação apenas contém o arquivo main.go e alguns arquivos de configuração e documentação
Dentro da raiz temos alguns pacotes:
* `api`: Pacote de controladores Rest
* `auth`: Pacote de autenticação dos bancos
* `bank`: Pacote de registro de boleto de cada banco
* `boleto`: Pacote responsável pela geração do boleto para o usuário
* `cache`: Pacote de Banco de dados (chave-valor) in-memory utilizado apenas quando roda a aplicação em modo mock
* `config`: Pacote de gerenciamento de configuração da aplicação
* `db`: Pacote de persistência de dados
* `devops`: Contém os arquivos de subida, deploy, backup e restore da aplicação
* `letters`: Pacote que contém os layouts de integração com os bancos
* `log`: Pacote de Log da Aplicação
* `models`: Pacote com as estruturas de dados da aplpicação
* `parser`: Pacote de XML
* `test`: Pacote utilitário de testes unitários
* `tmpl`: Pacote utilitário de template
* `util`: Pacote com utilitários de forma geral

Para mais informações
-----------------

Consulte o [FAQ](./FAQ.md)
