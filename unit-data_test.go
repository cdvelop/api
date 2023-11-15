package api_test

import (
	"github.com/cdvelop/testools"
)

var (
	testData = map[string]testools.Request{
		"se espera creación de producto ok": {
			Method:   "POST",
			Endpoint: "create",
			Object:   product.Objects[0].Name,
			Data:     []map[string]string{{"name": "manzana"}},
			Expected: []map[string]string{{"id_product": "4", "name": "manzana"}},
		},
		"se espera lectura json de producto ok": {
			Method:   "GET",
			Endpoint: "read",
			Object:   product.Objects[0].Name,
			Data:     []map[string]string{{"id_product": "1"}},
			Expected: []map[string]string{{"name": "manzana"}, {"name": "peras"}},
		},
		"se espera error lectura no se puede con método post": {
			Method:   "POST",
			Endpoint: "read",
			Object:   product.Objects[0].Name,
			Data:     []map[string]string{{"id_product": "1"}},
			Expected: []map[string]string{{"error": "acción: read no permitida con método POST"}},
		},
		"se espera actualización de producto ok": {
			Method:   "POST",
			Endpoint: "update",
			Object:   product.Objects[0].Name,
			Data:     []map[string]string{{"id_product": "1", "name": "pera"}},
			Expected: []map[string]string{{"id_product": "1", "name": "pera"}},
		},

		"se espera eliminación de producto ok": {
			Method:   "POST",
			Endpoint: "delete",
			Object:   product.Objects[0].Name,
			Data:     []map[string]string{{"id_product": "1"}},
			Expected: []map[string]string{{"id_product": "1", "name": "pera"}},
		},

		"se espera lectura fichero de productos ok": {
			Method:   "GET",
			Endpoint: "file?id=1",
			Object:   "",
			Data:     nil,
			Expected: []map[string]string{{"status": "200 OK"}},
		},

		"se espera lectura fichero estático  ok": {
			Method:   "GET",
			Endpoint: "static/dino-test.png",
			Object:   "",
			Expected: []map[string]string{{"status": "200 OK"}},
		},

		"se espera bad request no existe controlador": {
			Method:   "POST",
			Endpoint: "create",
			Object:   "dino",
			Expected: []map[string]string{{"status": "400 Bad Request"}},
		},
	}
)
