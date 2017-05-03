package log

import (
	"fmt"
	"net/http"

	"bitbucket.org/mundipagg/boletoapi/config"
	"github.com/mundipagg/goseq"
)

var logger *goseq.Logger

// Operation a operacao usada na API
var Operation string

// Recipient o nome do banco
var Recipient string

// Log struct com os elemtos do log
type Log struct {
	Operation   string
	Recipient   string
	NossoNumero int
	logger      *goseq.Logger
}

//Install instala o "servico" de log do SEQ
func Install() error {
	_logger, err := goseq.GetLogger(config.Get().SEQUrl, config.Get().SEQAPIKey, 150)
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

//CreateLog cria uma nova instancia do Log
func CreateLog() *Log {
	return &Log{
		logger: logger,
	}
}

// Request loga o request para algum banco
func (l Log) Request(content interface{}, url string, headers http.Header) {
	go (func() {
		props := goseq.NewProperties()
		props.AddProperty("MessageType", "Request")
		props.AddProperty("Content", content)
		props.AddProperty("Recipient", l.Recipient)
		props.AddProperty("Headers", headers)
		props.AddProperty("Operation", l.Operation)
		props.AddProperty("NossoNumero", l.NossoNumero)
		props.AddProperty("URL", url)

		msg := formatter("to {Recipient} ({URL})")

		l.logger.Information(msg, props)
	})()
}

// Response loga o response para algum banco
func (l Log) Response(content interface{}, url string) {
	go (func() {
		props := goseq.NewProperties()
		props.AddProperty("MessageType", "Response")
		props.AddProperty("Content", content)
		props.AddProperty("Recipient", l.Recipient)
		props.AddProperty("Operation", l.Operation)
		props.AddProperty("NossoNumero", l.NossoNumero)
		props.AddProperty("URL", url)
		msg := formatter("from {Recipient} ({URL})")

		l.logger.Information(msg, props)
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

// Fatal loga erros da aplicação
func (l Log) Fatal(content interface{}, m string) {
	go (func() {
		props := goseq.NewProperties()
		props.AddProperty("MessageType", "Error")
		props.AddProperty("Content", content)
		props.AddProperty("Recipient", l.Recipient)
		props.AddProperty("Operation", l.Operation)
		props.AddProperty("NossoNumero", l.NossoNumero)
		msg := formatter(m)

		l.logger.Fatal(msg, props)
	})()
}

//Close fecha a conexao com o SEQ
func Close() {
	fmt.Println("Closing SEQ Connection")
	logger.Close()
}
