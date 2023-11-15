package api_test

import (
	"testing"

	"github.com/cdvelop/testools"
)

func Test_Api(t *testing.T) {

	objects := ModuleProduct().Objects

	app, err := testools.NewApiTestDefault(t, module{}, objects...)
	if err != nil {
		t.Fatal(err)
	}

	defer app.Close()

	for prueba, r := range testData {
		t.Run((prueba), func(t *testing.T) {

			r.Server = app.Server

			endpoint := app.Server.URL

			if r.Endpoint != "" {
				endpoint += "/" + r.Endpoint
			}
			if r.Object != "" {
				endpoint += "/" + r.Object
			}
			// fmt.Println("ENDPOINT:", endpoint)
			// var err error
			app.SendOneRequest(r.Method, endpoint, r.Object, r.Data, func(resp []map[string]string, err error) {

				if err != nil {
					t.Fatal(err)
					return
				}

				testools.CheckTest(prueba, r.Expected, resp, t)
			})

		})
	}
}
