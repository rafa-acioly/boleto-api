package util

import (
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

var defaultDialer = &net.Dialer{Timeout: 16 * time.Second, KeepAlive: 16 * time.Second}

var cfg *tls.Config = &tls.Config{
	InsecureSkipVerify: true,
}
var client *http.Client = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig:     cfg,
		Dial:                defaultDialer.Dial,
		TLSHandshakeTimeout: 16 * time.Second,
	},
}

// DefaultHTTPClient retorna um cliente http configurado para dar um skip na validação do certificado digital
func DefaultHTTPClient() *http.Client {

	return client
}

//Post faz um requisição POST para uma URL e retorna o response, status e erro
func Post(url, body string, header map[string]string) (string, int, error) {
	return doRequest("POST", url, body, header)
}

//Get faz um requisição GET para uma URL e retorna o response, status e erro
func Get(url, body string, header map[string]string) (string, int, error) {
	return doRequest("GET", url, body, header)
}

func doRequest(method, url, body string, header map[string]string) (string, int, error) {
	client := DefaultHTTPClient()
	message := strings.NewReader(body)
	req, err := http.NewRequest(method, url, message)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	if header != nil {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}
	resp, errResp := client.Do(req)
	if errResp != nil {
		return "", 0, errResp
	}
	defer resp.Body.Close()
	data, errResponse := ioutil.ReadAll(resp.Body)
	if errResponse != nil {
		return "", resp.StatusCode, errResponse
	}
	sData := string(data)
	return sData, resp.StatusCode, nil
}

//HeaderToMap converte um http Header para um dicionário string -> string
func HeaderToMap(h http.Header) map[string]string {
	m := make(map[string]string)
	for k, v := range h {
		m[k] = v[0]
	}
	return m
}
