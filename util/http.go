package util

import (
	"crypto/tls"
	"net/http"
)

// DefaultHTTPClient retorna um cliente http configurado para dar um skip na validação do certificado digital
func DefaultHTTPClient() *http.Client {
	cfg := &tls.Config{
		InsecureSkipVerify: true,
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: cfg,
		},
	}
	return client
}
