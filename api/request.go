package api

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/edocm/huecli/config"
	"github.com/spf13/viper"
)

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func Request(method string, url string, body []byte) ([]byte, error) {
	var req *http.Request

	switch method {
	case http.MethodGet:
		req = createRequest(http.MethodGet, url, nil)
	case http.MethodPut:
		req = createRequest(http.MethodPut, url, body)
	case http.MethodPost:
		req = createRequest(http.MethodPost, url, body)
	case http.MethodDelete:
		req = createRequest(http.MethodDelete, url, body)
	default:
		return nil, fmt.Errorf("can not create a request with the given request method: %v", method)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while sending http request: %v", err)
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error while parsing response body: %v", err)
	}
	return resBody, nil
}

func createRequest(method string, url string, body []byte) *http.Request {
	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		log.Fatal(err)
	}
	if config.Exists {
		req.Header.Set("hue-application-key", viper.GetString("username"))
	}
	return req
}
