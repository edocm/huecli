package api

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func GET(url string) []byte {
	req := createRequest(http.MethodGet, url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return resBody
}

func PUT(url string, body []byte) []byte {
	req := createRequest(http.MethodPut, url, body)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return resBody
}

func POST(url string, body []byte) []byte {
	req := createRequest(http.MethodPost, url, body)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return resBody
}

func DELETE(url string, body []byte) []byte {
	req := createRequest(http.MethodDelete, url, body)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return resBody
}

func createRequest(method string, url string, body []byte) *http.Request {
	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		log.Fatal(err)
	}
	return req
}
