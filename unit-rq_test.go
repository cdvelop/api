package api_test

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/model"
)

type request struct {
	endpoint      string //ej: upload/files download/files
	method        string //ej: "PUT","GET"
	data          map[string]string
	object        *model.Object
	expected_code int
}

func (r request) post(s *httptest.Server, c *cutkey.Cut) ([]model.Response, int) {
	body, err := cutkey.Encode(r.object, r.data)
	if err != nil {
		log.Fatal(err)
	}

	return r.sendRequest(s, c, body)
}

func (r *request) get(s *httptest.Server, c *cutkey.Cut) ([]model.Response, int) {
	data := url.Values{}
	for key, value := range r.data {
		data.Set(key, value)
	}
	params := data.Encode()

	r.endpoint = fmt.Sprintf("%s?%s", r.endpoint, params)

	return r.sendRequest(s, c, nil)
}

func (r request) sendRequest(s *httptest.Server, c *cutkey.Cut, body []byte) ([]model.Response, int) {

	req, err := http.NewRequest(r.method, s.URL+r.endpoint, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return c.DecodeResponses(resp), res.StatusCode
}
