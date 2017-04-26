package models

//Validator estrutura de validação
type Validator struct {
	Rules []Rule
}

//Rule é a regra que será adiciona a camada de validação
type Rule func(interface{}) error

//Push insere no Validator uma nova regra
func (v *Validator) Push(r Rule) {
	v.Rules = append(v.Rules, r)
}

//NewValidator retorna nova instância de validação
func NewValidator() *Validator {
	return &Validator{}
}
