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

			var responses []model.Response
			var code int

			if r.method == "GET" {
				responses, code = r.get(srv, conf.Cut)
			} else {
				responses, code = r.post(srv, conf.Cut)
			}

			for _, resp := range responses {
				if r.expected_code != code {
					fmt.Println("=>PRUEBA: ", prueba)
					fmt.Printf("=>RESPUESTA: %v\n=>MENSAJE: %v\n=>SE ESPERABA:[%v]\n=>SE OBTUVO:[%v]\n", resp, resp.Message, r.expected_code, code)
					log.Fatal()
				}
			}

		})
	}

}
