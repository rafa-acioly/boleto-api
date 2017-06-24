O que é a API de Registro de Boleto Online?
--------------

BoletoOnline é uma API para registro online de boletos junto ao banco e geração de boletos para pagamento.


Atualmente os bancos suportados são:

* Banco do Brasil
* Caixa(em breve)
* Citi (em breve)
* Santander (em breve)
* Bradesco (em breve)
* Itau (em breve)

A ordem de integração seguirá a lista acima mas poderá haver modificações na prioridade dependendo do cliente.

Construindo a API
--------------

A API foi desenvolvida utilizando a linguagem GO e portanto é necessário instalar as ferramentas da linguagem caso queira compilar a aplicação a partir do fonte.

O Go pode ser baixado [aqui](https://golang.org/dl/)

Antes de fazer o clone do Projeto, deve ser criado o caminho de pastas dentro do $GOPATH

	% mkdir -p "$GOPATH/src/bitbucket.org/mundipagg"
	% cd $GOPATH/src/bitbucket.org/mundipagg 
	% git clone https://bitbucket.org/mundipagg/boletoapi

Antes de compilar a aplicação deve-se instalar o [Glide](http://glide.sh/) que é o gerenciador de dependências da aplicação.

Após instalar o GO, faça:

	% cd devops
	% ./build

O script build.sh irá baixar todas as dependências da aplicação e instalar o wkhtmltox, necessário para a geração do boleto em PDF.

Executando a aplicação
-------------

Para executar a API com as configurações default

Ex: 

Linux (*NIX):

	% ./boletoapi
	
Windows:

	% boletoapi.exe

Se você quiser rodar a API em modo dev, que irá carregar todas as variáveis de ambiente padrão, você deve executar a aplicação da seguinte forma:

	% ./boletoapi -dev

Caso queira executar a aplicação em modo mock, para não realizar diretamente a integração com o banco e usar uma base de dados em memória, deve-se usar a opção mock:

	% ./boletoapi -mock

Caso queira executar a aplicação com o log desligado, deve-se usar a opção -nolog:

	% ./boletoapi -nolog

Você pode combinar essas opções como quiser e caso queira usar todas elas juntas, basta usar a opção -airplane-mode

	% ./boletoapi -airplane-mode
	

Usando a API de boleto online
------------------

Pode-ser usar o [Postman](https://chrome.google.com/webstore/detail/postman/fhbjgbiflinjbdggehcddcbncdddomop) para criar chamar os serviços da API ou mesmo o curl

```
% curl -X POST \
  http://localhost:3000/v1/boleto/register \
  -d '{
    "Authentication" : {
        "Username":"user",
        "Password":"pass"
    },
    "Agreement":{
        "AgreementNumber":11111,
        "WalletVariation":19,
        "Wallet":17,
        "Agency":"123",
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
}

```
Resposta de sucesso da API
```
{
  "Url": "http://localhost:3000/boleto?fmt=html&id=g8HXWatft9oMLdTMAqzxbnPYFv3sqgV_KD0W7j8Cy9nkCLZMIK1WH2p9JwP1Jzz4ZtohmQ==",
  "DigitableLine": "00190000090101405100500066673179971340000010000",
  "BarCodeNumber": "00199713400000100000000001014051000006667317",
  "Links": [
    {
      "href": "http://localhost:3000/boleto?fmt=html&id=wOKZh6K_moLwXTW0Xr3oelh9YkYWXdl3VyURiQ-bu6TcuDzxdZI52BnQnuzNpGeh4TapUA==",
      "rel": "html",
      "method": "GET"
    },
    {
      "href": "http://localhost:3000/boleto?fmt=pdf&id=wOKZh6K_moLwXTW0Xr3oelh9YkYWXdl3VyURiQ-bu6TcuDzxdZI52BnQnuzNpGeh4TapUA==",
      "rel": "pdf",
      "method": "GET"
    }
  ]
}

```
Caso aconteça algum erro no processo de registro online a resposta entregue pela API seguirá o seguinte padrão.
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

Para instalar o executável da API precisa-se apenas compilar a aplicação e configurar as variáveis de ambiente necessárias.

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

Desta forma a aplicação será instalada de forma local na máquina.

Instalando a API via Docker
-----------------

Antes de fazer o deploy deve-se abrir o arquivo [docker-compose](/devops/docker-compose.yml) e configurar as informações que sejam pertinentes ao ambiente. Após ajustar o docker-compose pode-se instalar a aplicação usando o arquivo deploy.sh

    % cd devops
    % ./deploy.sh . local

O script irá criar os diretórios de volume do Docker, compilar a aplicação, montar as imagens da API e do MongoDB e subir os containers. Para mais informações sobre docker-compose consulte a [doc](https://docs.docker.com/compose/). 
Os parâmetros passados para o script dizem que o deploy será feito de forma local e não via TFS, caso não passe o argumento "local", o script irá utilizar o docker-compose.release.yml.

Após levantada, aplicação poderá ser parada ou iniciada.
    
    % cd devops/
    % ./stop.sh
    % ./start.sh
 
Backup e Restore
----------------- 

Para realizar o backup da base do MongoDB execute o seguinte comando:

    % cd devops
    % ./doBackup.sh

Os backups gerados, por padrão, serão armazenados no diretório `$HOME/backups` com o nome `bck_boletoapi-YYYY-MM-DD.tar`.
Para restaurar um backup:
    
    % cd devops
    % ./doRestore.sh

Quando fizer o restore, o script irá solicitar a data do arquivo de restore e deverá ser informada uma data válida do backup no padrão: `YYYY-MM-DD`.
    
Como contrubuir
-----------------

Para contrubuir dê uma olhada no arquivo [CONTRIBUTING](CONTRIBUTING.md)

Layout do Código Fonte
---

A Raiz da aplicação contém apenas o arquivo main.go e alguns arquivos de configuração e documentação.

Dentro da raiz temos alguns pacotes:

* `api`: Controladores Rest
* `auth`: Autenticação com os bancos
* `bank`: Registro de boletos
* `boleto`: Geração do boleto para o usuário
* `cache`: Banco de dados (chave-valor) in-memory utilizado apenas quando roda a aplicação em modo mock
* `config`: Configuração da aplicação
* `db`: Persistência de dados
* `devops`: Contém os arquivos de subida, deploy, backup e restore da aplicação
* `letters`: Layouts de integração com os bancos
* `log`: Log da Aplicação
* `models`: Estruturas de dados da aplicação
* `parser`: XML parser
* `test`: Utilitários de testes
* `tmpl`: Utilitário de template
* `util`: Utilitários de forma geral

Para mais informações
-----------------

Consulte o [FAQ](./FAQ.md)