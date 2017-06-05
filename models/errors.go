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

// IHttpNotFound interface para implementar Error
type IHttpNotFound interface {
	Error() string
	Message() string
}

// IGatewayTimeout interface para timeout
type IGatewayTimeout interface {
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

// HttpNotFound objeto para erros 404 da aplicação: ex boleto nÃ£o encontrado
type HttpNotFound struct {
	Err string
	Msg string
}

// Message retorna a mensagem final para o usuário
func (e HttpNotFound) Message() string {
	return e.Msg
}

// Error retorna o erro original
func (e HttpNotFound) Error() string {
	return e.Err
}

//NewHTTPNotFound cria um novo objeto NewHttpNotFound a partir de uma mensagem original e final
func NewHTTPNotFound(err, msg string) HttpNotFound {
	return HttpNotFound{Err: err, Msg: msg}
}

// GatewayTimeout objeto para erros 404 da aplicação: ex boleto nÃ£o encontrado
type GatewayTimeout struct {
	Err string
	Msg string
}

// Message retorna a mensagem final para o usuário
func (e GatewayTimeout) Message() string {
	return e.Msg
}

// Error retorna o erro original
func (e GatewayTimeout) Error() string {
	return e.Err
}

//NewGatewayTimeout cria um novo objeto NewGatewayTimeout a partir de uma mensagem original e final
func NewGatewayTimeout(err, msg string) GatewayTimeout {
	return GatewayTimeout{Err: err, Msg: msg}
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
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
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
