package config

//Config é a estrutura que tem todas as configurações da aplicação
type Config struct {
	APIPort            string
	Version            string
	SEQUrl             string
	SEQAPIKey          string
	EnableRequestLog   bool
	EnablePrintRequest bool
}

//Get retorna o objeto de configurações da aplicação
func Get() Config {
	cnf := Config{
		APIPort:            ":3000",
		Version:            "0.0.1",
		SEQUrl:             "http://stglog.mundipagg.com/", //Pegar o SEQ de dev
		SEQAPIKey:          "4jZzTybZ9bUHtJiPdh6",
		EnableRequestLog:   true,
		EnablePrintRequest: true,
	}
	return cnf
}
