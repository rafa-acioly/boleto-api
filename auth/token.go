package auth

// Token representa um token de autenticacao
type Token struct {
	Status           int
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	ExpiresIn        uint   `json:"expires_in"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
