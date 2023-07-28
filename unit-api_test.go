package api_test

import (
	"fmt"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/cdvelop/api"
	"github.com/cdvelop/model"
)

func Test_Api(t *testing.T) {

	conf := api.Add([]*model.Module{product})

	mux := conf.ServeMuxAndRoutes()

	srv := httptest.NewServer(mux)
	defer srv.Close()

	for prueba, r := range testData {
		t.Run((prueba), func(t *testing.T) {

			r.Cut = conf.Cut
			r.Server = srv

			var responses []model.Response
			var code int

			if r.Method == "GET" {
				responses, code = r.Get()
			} else {
				responses, code = r.Post()
			}

			for _, resp := range responses {
				if r.ExpectedCode != code {
					fmt.Println("=>PRUEBA: ", prueba)
					fmt.Printf("=>RESPUESTA: %v\n=>MENSAJE: %v\n=>SE ESPERABA:[%v]\n=>SE OBTUVO:[%v]\n", resp, resp.Message, r.ExpectedCode, code)
					log.Fatal()
				}
			}

		})
	}
}
