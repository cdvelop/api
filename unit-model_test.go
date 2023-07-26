package api_test

import (
	"fmt"
	"net/http"

	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
)

var (
	product = ModuleProduct()
)

type module struct {
}

func ModuleProduct() *model.Module {
	m := module{}

	newModule := model.Module{
		Name:  "product",
		Title: "Producto",
	}

	newObject := model.Object{
		Name: "product",
		BackendRequest: model.BackendRequest{
			CreateApi: m,
			ReadApi:   m,
			UpdateApi: m,
			DeleteApi: m,
			FileApi:   m,
		},

		Fields: []model.Field{
			{Name: "id_product", Legend: "Id", Input: input.Pk()},
			{Name: "name", Legend: "Nombre", Input: input.Text()},
		},
	}

	newObject.AddModule(&newModule)

	fmt.Printf("TAMAÑO OBJETOS MODULO: [%v]\n", len(newModule.Objects))

	return &newModule
}

func (m module) Create(params ...map[string]string) error {
	fmt.Println("parámetros Create recibidos:", params)

	params[0]["id"] = "2"

	return nil
}

func (m module) Read(params ...map[string]string) ([]map[string]string, error) {
	fmt.Println("parámetros leer todo recibidos:", params)
	return []map[string]string{
		{"name": "manzana"},
		{"name": "peras"},
	}, nil
}

func (m module) Update(params ...map[string]string) ([]map[string]string, error) {
	fmt.Println("parámetros Update recibidos:", params)
	return []map[string]string{}, nil
}

func (m module) Delete(params ...map[string]string) ([]map[string]string, error) {
	fmt.Println("parámetros Delete recibidos:", params)
	return []map[string]string{}, nil
}

func (m module) FilePath(params map[string]string) (string, error) {
	fmt.Println("parámetros leer archivo recibidos:", params)
	return "./README.md", nil
}

func (m module) UploadFile(r *http.Request, params map[string]string) ([]map[string]string, error) {
	fmt.Println("Upload File:", r)

	return []map[string]string{{"file": "./README.md"}}, nil
}

func (m module) MaximumFileSize() int64 {
	return 200
}
