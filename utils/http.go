package utils

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// http methods
const (
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodPatch  = "PATCH"
	MethodDelete = "DELETE"
)

var client = &http.Client{}

// CallAPI calls HTTP API
func CallAPI(method, url, body string, headers map[string]string, cookies map[string]string) (data []byte, err error) {
	req, err := http.NewRequest(method, url, strings.NewReader(body))

	for k, v := range headers {
		req.Header.Add(k, v)
	}
	for k, v := range cookies {
		req.AddCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)

	return
}
