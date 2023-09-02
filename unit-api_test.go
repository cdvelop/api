package api_test

import (
	"log"
	"net/http/httptest"
	"testing"

	"github.com/cdvelop/api"
	"github.com/cdvelop/model"
	"github.com/cdvelop/testools"
)

func Test_Api(t *testing.T) {

	conf := api.Add([]*model.Module{product}, nil)

	mux := conf.ServeMuxAndRoutes()

	srv := httptest.NewServer(mux)
	defer srv.Close()

	for prueba, r := range testData {
		t.Run((prueba), func(t *testing.T) {

			r.Cut = conf.Cut
			r.Server = srv

			var responses []model.Response
			var code int
			var err error

			if r.Method == "GET" {
				responses, code, err = r.Get(r.Data...)
			} else {
				responses, code, err = r.CutPost()
			}

			if err != nil {
				log.Fatal(err)
			}

			for _, resp := range responses {
				testools.CheckTest(prueba, r.ExpectedCode, code, resp)
			}
		})
	}
}
