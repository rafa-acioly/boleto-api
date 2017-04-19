package log

import (
	"net/http"

	"fmt"

	"bitbucket.org/mundipagg/boletoapi/config"
	"bitbucket.org/mundipagg/boletoapi/models"
	"github.com/mundipagg/goseq"
)

var logger *goseq.Logger
var messageType string

// Operation a operacao usada na API
var Operation string

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
	return "[{Application}: {Operation}] - {MessageType} " + message
}

// Request loga o request para algum banco
func Request(content, url string, headers http.Header, bankNumber models.BankNumber) {
	messageType = "Request"
	props := goseq.NewProperties()
	props.AddProperty("MessageType", messageType)
	props.AddProperty("Content", content)
	props.AddProperty("BankName", bankNumber.BankName())
	props.AddProperty("Headers", formatHeaders(headers))
	props.AddProperty("Operation", Operation)
	props.AddProperty("URL", url)

	msg := formatter("to {BankName} ({URL})")

	logger.Information(msg, props)
}

// Response loga o response para algum banco
func Response(content string, bankNumber models.BankNumber) {
	messageType = "Response"
	props := goseq.NewProperties()
	props.AddProperty("MessageType", messageType)
	props.AddProperty("Content", content)
	props.AddProperty("BankName", bankNumber.BankName())
	props.AddProperty("Operation", Operation)

	msg := formatter("from {BankName}")

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

func formatHeaders(headers http.Header) string {
	s := ""
	for name, headers := range headers {
		for _, h := range headers {
			s += fmt.Sprintf("%v: %v\n", name, h)
		}
	}
	return s
}
