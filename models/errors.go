package models

// IErrorResponse interface para implementar Error
type IErrorResponse interface {
	Error() string
	ErrorCode() string
}

// IServerError interface para implementar Error
type IServerError interface {
	Error() string
	Message() string
}

// IFormatError interface para implementar Error
type IFormatError interface {
	Error() string
}

// FormatError objeto para erros de input no request da API
type FormatError struct {
	Err string
}

func (e FormatError) Error() string {
	return e.Err
}

// InternalServerError objeto para erros internos da aplicação: ex banco de dados
type InternalServerError struct {
	Err string
	Msg string
}

// Message retorna a mensagem final para o usuário
func (e InternalServerError) Message() string {
	return e.Msg
}

// Error retorna o erro original
func (e InternalServerError) Error() string {
	return e.Err
}

//NewInternalServerError cria um novo objeto InternalServerError a partir de uma mensagem original e final
func NewInternalServerError(err, msg string) InternalServerError {
	return InternalServerError{Err: err, Msg: msg}
}

//NewErrorResponse cria um novo objeto de ErrorReponse com código e mensagem
func NewErrorResponse(code, msg string) ErrorResponse {
	return ErrorResponse{Code: code, Message: msg}
}

//NewFormatError cria um novo objeto de FormatError com descrição do erro
func NewFormatError(e string) FormatError {
	return FormatError{Err: e}
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
	return []ErrorResponse{errorResponse}
}

// NewSingleErrorCollection cria nova coleção de erros com 1 item
func NewSingleErrorCollection(code, msg string) Errors {
	return NewErrorCollection(NewErrorResponse(code, msg))
}

// NewErrors cria nova coleção de erros vazia
func NewErrors() Errors {
	return []ErrorResponse{}
}

// Append adiciona mais um erro na coleção
func (e *Errors) Append(code, message string) {
	*e = append(*e, ErrorResponse{Code: code, Message: message})
}
