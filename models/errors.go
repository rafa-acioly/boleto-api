package models

// IErrorResponse interface para implementar Error
type IErrorResponse interface {
	Error() string
	ErrorCode() string
}

//IErrorHTTP interface para retorno de status code
type IErrorHTTP interface {
	Error() string
	StatusCode() int
}

//ErrorStatusHTTP tipo de erro para forçar o status code http
type ErrorStatusHTTP struct {
	Code    int
	Message string
}

func (e ErrorStatusHTTP) Error() string {
	return e.Message
}

//StatusCode retorna o status code para forçar na volta da requisição
func (e ErrorStatusHTTP) StatusCode() int {
	return e.Code
}

// ErrorResponse objeto de erro
type ErrorResponse struct {
	Code    string
	Message string
}

func (e ErrorResponse) Error() string {
	return e.Message
}

// ErrorCode retorna código do erro
func (e ErrorResponse) ErrorCode() string {
	return e.Code
}

// Errors coleção de erros
type Errors []ErrorResponse

// NewErrorCollection cria nova coleção de erros
func NewErrorCollection(errorResponse ErrorResponse) Errors {
	return []ErrorResponse{
		ErrorResponse{
			Code:    errorResponse.Code,
			Message: errorResponse.Message,
		},
	}
}

// NewEmptyErrorCollection cria nova coleção de erros vazia
func NewEmptyErrorCollection() Errors {
	return []ErrorResponse{}
}

//NewErrorResponse retorna um ErrorResponse
func NewErrorResponse(code, message string) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: message,
	}
}

//NewSingleErrorCollection retorna colecao com um erro apenas
func NewSingleErrorCollection(code, message string) Errors {
	return NewErrorCollection(NewErrorResponse(code, message))
}

// Append adiciona mais um erro na coleção
func (e *Errors) Append(code, message string) {
	er := NewErrorResponse(code, message)
	*e = append(*e, er)
}
