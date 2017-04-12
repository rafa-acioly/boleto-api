package config

//Config é a estrutura que tem todas as configurações da aplicação
type Config struct {
	APIPort string
	Version string
}

//GetConfig retorna o objeto de configurações da aplicação
func GetConfig() Config {
	cnf := Config{
		APIPort: ":3000",
		Version: "0.0.1",
	}
	return cnf
}
