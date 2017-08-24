package models

// Authentication autenticação para entrada na API do banco
type Authentication struct {
	Username           string `json:"username,omitempty"`
	Password           string `json:"password,omitempty"`
	AuthorizationToken string `json:"authentication_token,omitempty"`
}
