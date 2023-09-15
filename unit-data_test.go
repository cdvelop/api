package api_test

import (
	"github.com/cdvelop/testools"
)

var (
	testData = map[string]testools.Request{
		"se espera creación de producto ok": {
			Endpoint:     "/create/" + product.Objects[0].Name,
			Method:       "POST",
			Data:         []map[string]string{{"name": "manzana"}},
			Object:       product.Objects[0],
			ExpectedCode: 200,
		},
		"se espera lectura json de producto ok": {
			Endpoint:     "/read/" + product.Objects[0].Name,
			Method:       "GET",
			Data:         []map[string]string{{"id_product": "1"}},
			Object:       product.Objects[0],
			ExpectedCode: 200,
		},
		"se espera actualización de producto ok": {
			Endpoint:     "/update/" + product.Objects[0].Name,
			Method:       "POST",
			Data:         []map[string]string{{"id_product": "1", "name": "pera"}},
			Object:       product.Objects[0],
			ExpectedCode: 200,
		},
		"se espera eliminación de producto ok": {
			Endpoint:     "/delete/" + product.Objects[0].Name,
			Method:       "POST",
			Data:         []map[string]string{{"id_product": "1"}},
			Object:       product.Objects[0],
			ExpectedCode: 200,
		},

		"se espera lectura fichero de productos ok": {
			Endpoint:     "/file/" + product.Objects[0].Name,
			Method:       "GET",
			Data:         []map[string]string{{"id_product": "1"}},
			Object:       product.Objects[0],
			ExpectedCode: 200,
		},

		"se espera lectura fichero estático  ok": {
			Endpoint:     "/static/dino-test.png",
			Method:       "GET",
			ExpectedCode: 200,
		},

		"se espera error no existe controlador": {
			Endpoint:     "/create/dino",
			Method:       "POST",
			ExpectedCode: 400,
		},
	}
)
