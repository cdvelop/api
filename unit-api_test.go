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
	if err != "" {
		t.Fatal(err)
		return
	}

	defer app.Close()

	for prueba, r := range testData {
		t.Run((prueba), func(t *testing.T) {
			r.ApiTest = app
			r.TestName = prueba

			// fmt.Println("ENDPOINT:", endpoint)
			app.SendOneRequest(r.Method, app.BuildEndPoint(r), r.Object, r.Data, func(resp []map[string]string, err string) {
				var response any
				if err != "" {
					response = err
				} else {
					response = resp
				}
				// fmt.Println("RESPUESTA:", response)
				if !r.CheckTest(r.Expected, response) {
					t.Fatal()
					return
				}
			})

		})
	}

}
