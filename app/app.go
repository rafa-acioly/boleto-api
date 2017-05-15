package app

import (
	"fmt"
	"os"

	"bitbucket.org/mundipagg/boletoapi/api"
	"bitbucket.org/mundipagg/boletoapi/config"
	"bitbucket.org/mundipagg/boletoapi/log"
	"bitbucket.org/mundipagg/boletoapi/robot"
)

//Run starts boleto api Application
func Run(devMode, mockMode, disableLog bool) {
	configFlags(devMode, mockMode, disableLog)
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

func configFlags(devMode, mockMode, disableLog bool) {
	if devMode {
		os.Setenv("API_PORT", "3000")
		os.Setenv("API_VERSION", "0.0.1")
		os.Setenv("ENVIROMENT", "Development")
		os.Setenv("SEQ_URL", "http://192.168.8.119:5341") // http://stglog.mundipagg.com/ 192.168.8.119:5341
		os.Setenv("SEQ_API_KEY", "4jZzTybZ9bUHtJiPdh6")   //4jZzTybZ9bUHtJiPdh6
		os.Setenv("ENABLE_REQUEST_LOG", "false")
		os.Setenv("ENABLE_PRINT_REQUEST", "true")
		os.Setenv("URL_BB_REGISTER_BOLETO", "https://cobranca.homologa.bb.com.br:7101/registrarBoleto")
		os.Setenv("URL_BB_TOKEN", "https://oauth.hm.bb.com.br:43000/oauth/token")
		os.Setenv("APP_URL", "http://localhost:3000/boleto")
		os.Setenv("ELASTIC_URL", "http://localhost:9200")
		os.Setenv("MONGODB_URL", "localhost:27017")
		os.Setenv("BOLETO_JSON_STORE", "/home/philippe/boletodb/upMongo")
		if mockMode {
			os.Setenv("URL_BB_REGISTER_BOLETO", "http://localhost:4000/registrarBoleto")
			os.Setenv("URL_BB_TOKEN", "http://localhost:4000/oauth/token")
		}
	}
	config.Install(mockMode, devMode, disableLog)
}
