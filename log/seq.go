package log

import (
	"net/http"

	"bitbucket.org/mundipagg/boletoapi/config"
	"github.com/mundipagg/goseq"
)

var logger *goseq.Logger
var messageType string

// Operation a operacao usada na API
var Operation string

// Recipient o nome do banco
var Recipient string

//Install instala o "servico" de log do SEQ
func Install() error {
	_logger, err := goseq.GetLogger(config.Get().SEQUrl, config.Get().SEQAPIKey)
	if err != nil {
		return err
	}
	_logger.SetDefaultProperties(map[string]interface{}{
		"Application": config.Get().ApplicationName,
		"Environment": config.Get().Environment,
		"Domain":      config.Get().SEQDomain,
	})
	logger = _logger
	return nil
}

func formatter(message string) string {
	return "[{Application}: {Operation}] - {MessageType} " + message
}

// Request loga o request para algum banco
func Request(content interface{}, url string, headers http.Header) {
	go (func() {
		messageType = "Request"
		props := goseq.NewProperties()
		props.AddProperty("MessageType", messageType)
		props.AddProperty("Content", content)
		props.AddProperty("Recipient", Recipient)
		props.AddProperty("Headers", headers)
		props.AddProperty("Operation", Operation)
		props.AddProperty("URL", url)

		msg := formatter("to {Recipient} ({URL})")

		logger.Information(msg, props)
	})()
}

// Response loga o response para algum banco
func Response(content interface{}) {
	go (func() {
		messageType = "Response"
		props := goseq.NewProperties()
		props.AddProperty("MessageType", messageType)
		props.AddProperty("Content", content)
		props.AddProperty("Recipient", Recipient)
		props.AddProperty("Operation", Operation)

		msg := formatter("from {Recipient}")

		logger.Information(msg, props)
	})()
}

//Info loga mensagem do level INFO
func Info(msg string) {
	go logger.Information(msg, goseq.NewProperties())
}

//Warn loga mensagem do leve Warning
func Warn(msg string) {
	go logger.Warning(msg, goseq.NewProperties())
}

//Fatal loga mensagem do leve fatal
func Fatal(msg string) {
	go logger.Fatal(msg, goseq.NewProperties())
}

//Close fecha a conexao com o SEQ
func Close() {
	logger.Close()
}
