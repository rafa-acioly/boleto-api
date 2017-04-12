package config

//Config é a estrutura que tem todas as configurações da aplicação
type Config struct {
	APIPort string
}

//GetConfig retorna o objeto de configurações da aplicação
func GetConfig() Config {
	cnf := Config{
		APIPort: ":3000",
	}
	return cnf
}
