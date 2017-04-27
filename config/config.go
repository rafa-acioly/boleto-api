package config

import "os"
import "sync"

//Config é a estrutura que tem todas as configurações da aplicação
type Config struct {
	APIPort             string
	Version             string
	SEQUrl              string
	SEQAPIKey           string
	EnableRequestLog    bool
	EnablePrintRequest  bool
	Environment         string
	SEQDomain           string
	ApplicationName     string
	URLBBRegisterBoleto string
	URLBBToken          string
	MockMode            bool
	AppURL              string
}

var cnf Config
var scnf sync.Once

//Get retorna o objeto de configurações da aplicação
func Get() Config {
	return cnf
}
func Install(mockMode bool) {
	cnf = Config{
		APIPort:             ":" + os.Getenv("API_PORT"),
		Version:             os.Getenv("API_VERSION"),
		SEQUrl:              os.Getenv("SEQ_URL"),                        //Pegar o SEQ de dev
		SEQAPIKey:           os.Getenv("SEQ_API_KEY"),                    //Staging Key:
		EnableRequestLog:    os.Getenv("ENABLE_REQUEST_LOG") == "true",   // Log a cada request no SEQ
		EnablePrintRequest:  os.Getenv("ENABLE_PRINT_REQUEST") == "true", // Imprime algumas informacoes da request no console
		Environment:         os.Getenv("ENVIRONMENT"),
		SEQDomain:           "One",
		ApplicationName:     "BoletoOnline",
		URLBBRegisterBoleto: os.Getenv("URL_BB_REGISTER_BOLETO"),
		URLBBToken:          os.Getenv("URL_BB_TOKEN"),
		MockMode:            mockMode,
		AppURL:              os.Getenv("APP_URL"),
	}
}
