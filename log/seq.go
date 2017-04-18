package log

import (
	"bitbucket.org/mundipagg/boletoapi/config"
	"github.com/mundipagg/goseq"
)

var logger *goseq.Logger

//Install instala o "servico" de log do SEQ
func Install() error {
	_logger, err := goseq.GetLogger(config.Get().SEQUrl, config.Get().SEQAPIKey)
	if err != nil {
		return err
	}
	_logger.SetDefaultProperties(map[string]string{
		"Application": "BoletoOnline",
	})
	logger = _logger
	return nil
}

//Info loga mensagem do level INFO
func Info(msg string) {
	logger.Information(msg, goseq.NewProperties())
}

//Warn loga mensagem do leve Warning
func Warn(msg string) {
	logger.Warning(msg, goseq.NewProperties())
}

//Fatal loga mensagem do leve fatal
func Fatal(msg string) {
	logger.Fatal(msg, goseq.NewProperties())
}

//Close fecha a conexao com o SEQ
func Close() {
	logger.Close()
}
