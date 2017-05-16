package models

// Authentication autenticação para entrada na API do banco
type Authentication struct {
	Username           string `json:",omitempty"`
	Password           string `json:",omitempty"`
	AuthorizationToken string `json:",omitempty"`
}
