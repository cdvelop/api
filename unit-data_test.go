package api_test

var (
	testData = map[string]request{
		"se espera creación de producto ok": {
			endpoint:      "/create/" + product.Objects[0].Api(),
			method:        "POST",
			data:          map[string]string{"name": "manzana"},
			object:        product.Objects[0],
			expected_code: 200,
		},
		"se espera actualización de producto ok": {
			endpoint:      "/update/" + product.Objects[0].Api(),
			method:        "POST",
			data:          map[string]string{"id_product": "1", "name": "pera"},
			object:        product.Objects[0],
			expected_code: 200,
		},
		"se espera eliminación de producto ok": {
			endpoint:      "/delete/" + product.Objects[0].Api(),
			method:        "POST",
			data:          map[string]string{"id_product": "1"},
			object:        product.Objects[0],
			expected_code: 200,
		},
		"se espera lectura json de producto ok": {
			endpoint:      "/readone/" + product.Objects[0].Api(),
			method:        "GET",
			data:          map[string]string{"id_product": "1"},
			object:        product.Objects[0],
			expected_code: 200,
		},
		"se espera lectura json de productos ok": {
			endpoint:      "/readall/" + product.Objects[0].Api(),
			method:        "GET",
			data:          map[string]string{"name": "frutas"},
			object:        product.Objects[0],
			expected_code: 200,
		},
		"se espera lectura fichero de productos ok": {
			endpoint:      "/file/" + product.Objects[0].Api(),
			method:        "GET",
			data:          map[string]string{"id_product": "1"},
			object:        product.Objects[0],
			expected_code: 200,
		},

		"se espera lectura fichero estático  ok": {
			endpoint:      "/static/dino-test.png",
			method:        "GET",
			expected_code: 200,
		},
	}
)
