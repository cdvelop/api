package api_test

import (
	"testing"

	"github.com/cdvelop/model"
	"github.com/cdvelop/testools"
)

func Test_Api(t *testing.T) {

	objects := ModuleProduct().Objects

	h := &model.Handlers{
		FileApi: module{},
	}

	app, err := testools.NewApiTestDefault(t, h, objects...)
	if err != nil {
		t.Fatal(err)
	}

	defer app.Close()

	for prueba, r := range testData {
		t.Run((prueba), func(t *testing.T) {
			r.ApiTest = app

			// fmt.Println("ENDPOINT:", endpoint)
			app.SendOneRequest(r.Method, app.BuildEndPoint(r), r.Object, r.Data, func(resp []map[string]string, err error) {
				var response any
				if err != nil {
					response = err.Error()
				} else {
					response = resp
				}

				// fmt.Println("RESPUESTA:", response)

				testools.CheckTest(prueba, r.Expected, response, t)
			})

		})
	}
}
