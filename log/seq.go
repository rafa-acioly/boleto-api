package log

import (
	"bitbucket.org/mundipagg/boletoapi/config"
	"github.com/mundipagg/goseq"
)

var logger *goseq.Logger
var messageType string
var Operation string
var Recipient string

//Install instala o "servico" de log do SEQ
func Install() error {
	_logger, err := goseq.GetLogger(config.Get().SEQUrl, config.Get().SEQAPIKey)
	if err != nil {
		return err
	}
	_logger.SetDefaultProperties(map[string]string{
		"Application": config.Get().ApplicationName,
		"Environment": config.Get().Environment,
		"Domain":      config.Get().SEQDomain,
	})
	logger = _logger
	return nil
}

func formatter(message string) string {
	return "[" + config.Get().ApplicationName + ": " + Operation + "] - " + messageType + " " + message
}

func Request(content string) {
	messageType = "Request"
	props := goseq.NewProperties()
	props.AddProperty("MessageType", messageType)
	props.AddProperty("Content", content)

	msg := formatter("to BancoDoBrasil (http://example.com)")

	logger.Information(msg, props)
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
