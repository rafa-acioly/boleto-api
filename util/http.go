package util

import (
	"crypto/tls"
	"net/http"
)

var cfg *tls.Config = &tls.Config{
	InsecureSkipVerify: true,
}
var client *http.Client = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: cfg,
	},
}

// DefaultHTTPClient retorna um cliente http configurado para dar um skip na validação do certificado digital
func DefaultHTTPClient() *http.Client {

	return client
}
