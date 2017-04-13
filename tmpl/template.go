package tmpl

import (
	"bytes"
	"html/template"
)

// Builder e o padrao que vai transformar mensagens de entrada em mensagens de saida para os bancos
type Builder interface {
	From(interface{}) Builder
	To(string) Builder
	Transform() (string, error)
}

type msgBuilder struct {
	from     interface{}
	template string
}

func (b msgBuilder) From(obj interface{}) Builder {
	b.from = obj
	return b
}

func (b msgBuilder) To(template string) Builder {
	b.template = template
	return b
}

func (b msgBuilder) Transform() (string, error) {
	buf := bytes.NewBuffer(nil)
	t := template.Must(template.New("transform").Parse(b.template))
	err := t.ExecuteTemplate(buf, "transform", b.from)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// New cria um novo builder
func New() Builder {
	return msgBuilder{}
}
