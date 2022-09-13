package api

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func GET(url string) string {
	req := createRequest(http.MethodGet, url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(resBody)
}

func PUT(url string, body io.Reader) string {
	req := createRequest(http.MethodPut, url, body)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(resBody)
}

func POST(url string, body io.Reader) string {
	req := createRequest(http.MethodPost, url, body)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(resBody)
}

func DELETE(url string, body io.Reader) string {
	req := createRequest(http.MethodDelete, url, body)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(resBody)
}

func createRequest(method string, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal(err)
	}
	return req
}
