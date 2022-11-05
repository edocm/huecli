package api

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/edocm/huecli/config"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
)

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func Request(method string, url string, body []byte) ([]byte, error) {
	var req *http.Request

	switch method {
	case "GET":
		req = createRequest(http.MethodGet, url, nil)
	case "PUT":
		req = createRequest(http.MethodPut, url, body)
	case "POST":
		req = createRequest(http.MethodPost, url, body)
	case "DELETE":
		req = createRequest(http.MethodDelete, url, body)
	default:
		return nil, fmt.Errorf("can not create a request with the given request method: %v", method)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while sending http request: %v", err)
	}
	resBody, err := ioutil.ReadAll(res.Body)
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
