package api

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

type Request struct {
	Endpoint string //ej: upload/files download/files
	Method   string //ej: "PUT","GET"
	Data     []map[string]string
	*model.Object
	ExpectedCode int

	*cutkey.Cut
	*httptest.Server
}

func (r Request) Post() ([]model.Response, int) {
	body, err := cutkey.Encode(r.Object, r.Data...)
	if err != nil {
		log.Fatal(err)
	}

	return r.SendRequest(body)
}

func (r *Request) Get() ([]model.Response, int) {
	url_values := url.Values{}

	for _, data := range r.Data {
		for key, value := range data {
			// Agregar cada clave-valor al url_values
			url_values.Add(key, value)
		}
	}

	params := url_values.Encode()

	r.Endpoint = fmt.Sprintf("%s?%s", r.Endpoint, params)

	return r.SendRequest(nil)
}

func (r Request) SendRequest(body []byte) ([]model.Response, int) {

	req, err := http.NewRequest(r.Method, r.Server.URL+r.Endpoint, bytes.NewBuffer(body))
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

	return r.DecodeResponses(resp), res.StatusCode
}
