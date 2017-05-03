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

//Assert aplica todas as validações no objeto passado como parâmetro
func (v *Validator) Assert(o interface{}) Errors {
	errs := NewErrors()
	for _, assert := range v.Rules {
		err := assert(o)
		switch t := err.(type) {
		case ErrorResponse:
			errs = append(errs, t)
		default:
			if err != nil {
				errs.Append(err.Error(), "Erro interno")
			}
		}
	}
	return errs
}

//NewValidator retorna nova instância de validação
func NewValidator() *Validator {
	return &Validator{}
}
