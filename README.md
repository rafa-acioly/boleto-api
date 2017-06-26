What is the Online "Boleto" API?
--------------

BoletoOnline is an API for boleto's online register in banks and boleto's creation for payments.

Currently, we support the banks below:
* Banco do Brasil
* Caixa(coming soon)
* Citi (coming soon)
* Santander (coming soon)
* Bradesco (coming soon)
* Itaú (coming soon)

The integration order will follow the list above but we may have changes considering our clients demands.

API Building
--------------

The API was developed using GO language and therefore it is necessary to install the language tools in case you need to compile the application from the source.

GO can be downloaded [here](https://golang.org/dl/)

Before cloning the Project, it should be created the file path inside $GOPATH

	% mkdir -p "$GOPATH/src/github.com/mundipagg"
	% cd $GOPATH/src/github.com/mundipagg 
	% git clone https://github.com/mundipagg/boleto-api


Before compiling the application, it should be installed the [Glide](http://glide.sh/), which is the application dependency manager.

After installing GO, do:

	% cd devops
	% ./build

The script build.sh will download the application dependency and install wkhtmltox, necessary to the boleto's creation in PDF format.

Running the application
-------------

To run the API with default configuration
Eg:

Linux (*NIX):

	% ./boleto-api
	
Windows:

	% boleto-api.exe

If you want to run the API in dev mode, which will reload all variables in standard environment, you should execute the application like this:

	% ./boleto-api -dev

In case you want to run the application in mock mode, not using the bank integration but a local database, you should use mock option:

	% ./boleto-api -mock

In case you want to run the application with log turned off, you should use the option -nolog:

	% ./boleto-api -nolog

You can combine all these options and, in case you want to use them altogether, you can simply use the -airplane-mode option

	% ./boleto-api -airplane-mode
	

Using Online "Boleto" API
------------------

You can use [Postman](https://chrome.google.com/webstore/detail/postman/fhbjgbiflinjbdggehcddcbncdddomop) to request the API's services or even the curl

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
API's success response
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
In case of any error in the register proccess, the API's response will follow the pattern below:
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

Installing the API
-----------------

To install the API's executable, it is only necessary to compile the application  and configure the necessary environment variables.

Edit file $HOME/.bashrc.sh
```
    export API_PORT="3000"
    export API_VERSION="0.0.1"
    export ENVIRONMENT="Development"
    export SEQ_URL="http://example.mundipagg.com"
    export SEQ_API_KEY="API_KEY"
    export ENABLE_REQUEST_LOG="false"
    export ENABLE_PRINT_REQUEST="true"
    export URL_BB_REGISTER_BOLETO="https://cobranca.desenv.bb.com.br:7101/registrarBoleto"
    export URL_BB_TOKEN="https://oauth.desenv.bb.com.br:43000/oauth/token"
    export MONGODB_URL="10.0.2.15:27017"
    export APP_URL="http://localhost:8080/boleto"
```
    % go build && mv boleto-api /usr/local/bin

Then the application will be installed locally.

Instaling the API via Docker
-----------------

Before deploying, you should open the file [docker-compose](/devops/docker-compose.yml) and configure the information which is relevant to the environment. After setting up the docker-compose, you can install the application using the file deploy.sh

    % cd devops
    % ./deploy.sh . local

The script will create the Docker's volume directories, compile the application, mount the API's and MongoDB's images and upload the containers.
For more information about docker-compose, see [doc](https://docs.docker.com/compose/). 

The parameters sent to the script show that the deploy will run locally and not via TFS. In case you don't send the argument "local", the script will use docker-compose.release.yml.

After being deployed, the application can be stoped or started.
    
    % cd devops/
    % ./stop.sh
    % ./start.sh
 
Backup & Restore
----------------- 

To backup the MongoDB datebase, run the following command:

    % cd devops
    % ./doBackup.sh

The generated backups, by default, will be stored in the diretory `$HOME/backups` with the name `bck_boleto-api-YYYY-MM-DD.tar`.
To restore a backup:
    
    % cd devops
    % ./doRestore.sh

When doing the restore, the script will aks for the restore file date and it should be informed a valid date of the backup in pattern: `YYYY-MM-DD`.
    
How to contribute
-----------------

To contribute, see [CONTRIBUTING](CONTRIBUTING.md)

Source Code Layout
---

The application root contains only the file main.go and some config and documentation files.

In the root, we have the following packages:

* `api`: Rest Controllers
* `auth`: Bank authentication
* `bank`: Boleto's register
* `boleto`: User boleto's creation
* `cache`: Database (key value) in-memory used only when the applicatio in runned in mock mode
* `config`: Application config
* `db`: Database persistency
* `devops`: Contains the upload, deploy, backup and restore files from the application
* `letters`: Integration banks layouts
* `log`: Application log
* `models`: Application's data structure
* `parser`: XML parser
* `test`: Tests utilitaries
* `tmpl`: Template utilitaries
* `util`: Miscellaneous utilitaries

For more information
-----------------

See [FAQ](./FAQ.md)
