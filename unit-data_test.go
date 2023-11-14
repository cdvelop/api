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
		// "se espera lectura json de producto ok": {
		// 	Endpoint:       "read",
		// 	Object:     product.Objects[0].Name,
		// 	Data:         []map[string]string{{"id_product": "1"}},
		// 	Expected: 200,
		// },
		// "se espera actualización de producto ok": {
		// 	Endpoint:       "update",
		// 	Object:    product.Objects[0].Name,
		// 	Data:         []map[string]string{{"id_product": "1", "name": "pera"}},
		// 	Expected: 200,
		// },
		// "se espera eliminación de producto ok": {
		// 	Endpoint:       "delete",
		// 	Object:     product.Objects[0].Name,
		// 	Data:         []map[string]string{{"id_product": "1"}},
		// 	Expected: 200,
		// },

		// "se espera lectura fichero de productos ok": {
		// 	Endpoint:       "file",
		// 	Object:     "/file?id=1",
		// 	Data:         []map[string]string{{}},
		// 	Expected: 200,
		// },

		// "se espera lectura fichero estático  ok": {
		// 	Endpoint:       "",
		// 	Object:     "/static/dino-test.png",
		// 	Expected: 200,
		// },

		// "se espera error no existe controlador": {
		// 	Endpoint:       "create",
		// 	Object:     "dino",
		// 	Expected: 400,
		// },
	}
)
