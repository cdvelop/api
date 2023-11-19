package api_test

import (
	"github.com/cdvelop/fileserver"
	"github.com/cdvelop/testools"
)

const static_files = "frontend/built/static"

var (
	testData = map[string]testools.Request{
		"se espera creación de producto ok": {
			Method:   "POST",
			Endpoint: "create",
			Object:   product.Objects[0].ObjectName,
			Data:     []map[string]string{{"name": "manzana"}},
			Expected: []map[string]string{{"id_product": "4", "name": "manzana"}},
		},
		"se espera lectura json de producto ok": {
			Method:   "POST",
			Endpoint: "read",
			Object:   product.Objects[0].ObjectName,
			Data:     []map[string]string{{"id_product": "1"}},
			Expected: []map[string]string{{"name": "manzana"}, {"name": "peras"}},
		},

		"se espera error id 2 no esta en el controlador": {
			Method:   "POST",
			Endpoint: "read",
			Object:   product.Objects[0].ObjectName,
			Data:     []map[string]string{{"id_product": "2"}},
			Expected: "nada encontrado",
		},

		"se espera actualización de producto ok": {
			Method:   "POST",
			Endpoint: "update",
			Object:   product.Objects[0].ObjectName,
			Data:     []map[string]string{{"id_product": "1", "name": "pera"}},
			Expected: []map[string]string{{"id_product": "1", "name": "pera"}},
		},

		"se espera eliminación de producto ok": {
			Method:   "POST",
			Endpoint: "delete",
			Object:   product.Objects[0].ObjectName,
			Data:     []map[string]string{{"id_product": "1"}},
			Expected: []map[string]string{{"id_product": "1", "name": "pera"}},
		},

		"se espera lectura fichero de productos ok": {
			Method:   "GET",
			Endpoint: "file?id=1",
			Object:   "file",
			Data:     nil,
			Expected: []map[string]string{{"file": string(fileserver.GetFile("./README.md"))}},
		},

		"se espera lectura fichero estático  ok": {
			Method:   "GET",
			Endpoint: "static/dino-test.png",
			Object:   "",
			Expected: []map[string]string{{"file": string(fileserver.GetFile(static_files + "/dino-test.png"))}},
		},

		"se espera bad request no existe controlador": {
			Method:   "POST",
			Endpoint: "create",
			Object:   "dino",
			Expected: "error isHandlerOk objeto: dino no encontrado",
		},
		"se espera subida de archivo ok": {
			Method:   "POST",
			Endpoint: "upload",
			Object:   product.Objects[0].ObjectName,
			Data:     []map[string]string{{"name": "1"}},
			Expected: []map[string]string{{"id_product": "200"}},
		},
	}
)
