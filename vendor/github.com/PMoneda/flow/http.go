package flow

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

var defaultDialer = &net.Dialer{Timeout: 16 * time.Second, KeepAlive: 16 * time.Second}

var cfg = &tls.Config{
	InsecureSkipVerify: true,
}
var client = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig:     cfg,
		Dial:                defaultDialer.Dial,
		TLSHandshakeTimeout: 16 * time.Second,
		DisableCompression:  true,
		DisableKeepAlives:   true,
	},
}

func getClient(skip string) *http.Client {
	/*_skip := false
	if skip != "" {
		_skip, _ = strconv.ParseBool(skip)
	}
	cfg := &tls.Config{
		InsecureSkipVerify: _skip,
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: cfg,
		},
	}*/
	return client
}
func httpConnector(next func(), e *ExchangeMessage, out Message, u URI, params ...interface{}) error {

	newData := NewExchangeMessage()
	var skip string
	var opts map[string]string
	method := "GET"
	authMethod := ""
	username := ""
	password := ""
	if len(params) > 0 {
		opts = params[0].(map[string]string)
		skip = opts["insecureSkipVerify"]
		method = opts["method"]
		authMethod = opts["auth"]
		username = opts["username"]
		password = opts["password"]
	}
	client := getClient(skip)
	var req *http.Request
	var err error
	b := e.GetBody()
	var body io.Reader
	switch t := b.(type) {
	case string:
		body = strings.NewReader(t)
	default:
		if j, err := json.Marshal(b); err != nil {
			newData.SetHeader("error", err.Error())
			newData.SetBody(err)
			out <- newData
			next()
		} else {
			body = strings.NewReader(string(j))
		}
	}
	req, err = http.NewRequest(method, u.raw, body)

	if err != nil {
		newData.SetHeader("error", err.Error())
		newData.SetBody(err)
		out <- newData
		next()
		return err
	}

	if authMethod == "basic" {
		req.SetBasicAuth(username, password)
	}
	header := e.head
	keys := header.ListKeys()
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "pt-BR,pt;q=0.8,en-US;q=0.6,en;q=0.4")
	for _, key := range keys {
		req.Header.Add(key, header.Get(key))
	}
	req.Close = true
	resp, errResp := client.Do(req)
	if errResp != nil {
		newData.SetHeader("error", errResp.Error())
		newData.SetBody(errResp)
		out <- newData
		next()
		return errResp
	}
	defer resp.Body.Close()
	data, errResponse := ioutil.ReadAll(resp.Body)
	if errResponse != nil {
		newData.SetHeader("error", errResponse.Error())
		newData.SetBody(errResponse)
		out <- newData
		next()
		return errResponse
	}

	for k := range resp.Header {
		newData.SetHeader(k, resp.Header.Get(k))
	}
	newData.SetHeader("status", fmt.Sprintf("%d", resp.StatusCode))
	newData.SetBody(string(data))
	out <- newData
	next()
	return nil
}
