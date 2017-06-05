package app

import (
	"fmt"
	"os"

	"github.com/PMoneda/flow"

	"bitbucket.org/mundipagg/boletoapi/api"
	"bitbucket.org/mundipagg/boletoapi/config"
	"bitbucket.org/mundipagg/boletoapi/log"
	"bitbucket.org/mundipagg/boletoapi/models"
	"bitbucket.org/mundipagg/boletoapi/robot"
	"bitbucket.org/mundipagg/boletoapi/util"
)

//Run starts boleto api Application
func Run(devMode, mockMode, disableLog bool) {
	configFlags(devMode, mockMode, disableLog)
	installflowConnectors()
	robot.GoRobots()
	installLog()
	api.InstallRestAPI()

}

func installLog() {
	err := log.Install()
	if err != nil {
		fmt.Println("Log SEQ Fails")
		os.Exit(-1)
	}
}

func installflowConnectors() {
	flow.RegisterConnector("logseq", util.SeqLogConector)
	flow.RegisterConnector("apierro", models.BoletoErrorConector)
}

func configFlags(devMode, mockMode, disableLog bool) {
	if devMode {
		os.Setenv("API_PORT", "3000")
		os.Setenv("API_VERSION", "0.0.1")
		os.Setenv("ENVIROMENT", "Development")
		os.Setenv("SEQ_URL", "http://localhost:5341")   // http://stglog.mundipagg.com/ 192.168.8.119:5341
		os.Setenv("SEQ_API_KEY", "4jZzTybZ9bUHtJiPdh6") //4jZzTybZ9bUHtJiPdh6
		os.Setenv("ENABLE_REQUEST_LOG", "false")
		os.Setenv("ENABLE_PRINT_REQUEST", "true")
		os.Setenv("URL_BB_REGISTER_BOLETO", "https://cobranca.homologa.bb.com.br:7101/registrarBoleto")
		os.Setenv("URL_BB_TOKEN", "https://oauth.hm.bb.com.br:43000/oauth/token")
		os.Setenv("URL_CAIXA", "https://des.barramento.caixa.gov.br/sibar/ManutencaoCobrancaBancaria/Boleto/Externo")
		os.Setenv("URL_CITI", "https://citigroupsoa.citigroup.com/comercioeletronico/registerboleto/RegisterBoletoSOAP")
		os.Setenv("APP_URL", "http://localhost:3000/boleto")
		os.Setenv("ELASTIC_URL", "http://localhost:9200")
		os.Setenv("MONGODB_URL", "localhost:27017")
		os.Setenv("BOLETO_JSON_STORE", "/home/philippe/boletodb/upMongo")
		if mockMode {
			os.Setenv("URL_BB_REGISTER_BOLETO", "http://localhost:4000/registrarBoleto")
			os.Setenv("URL_BB_TOKEN", "http://localhost:4000/oauth/token")
			os.Setenv("URL_CAIXA", "http://localhost:4000/caixa/registrarBoleto")
			os.Setenv("URL_CITI", "http://localhost:4000/citi/registrarBoleto")
		}
	}
	config.Install(mockMode, devMode, disableLog)
}
