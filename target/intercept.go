package target

import (
	"net/http"
	"net/http/httptest"
	"strings"
)

type Intercept interface {
	ServeHTTP(rw *httptest.ResponseRecorder, r *http.Request)
}

func CallAPI(intercepter Intercept, method, url, body string, extraHeaders map[string]string) (
	rw *httptest.ResponseRecorder) {

	req, _ := http.NewRequest(method, url, strings.NewReader(body))
	for k, v := range extraHeaders {
		req.Header.Set(k, v)
	}

	rw = httptest.NewRecorder()
	intercepter.ServeHTTP(rw, req)

	return
}
