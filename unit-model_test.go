package api_test

import (
	"fmt"

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
			CreateApi:  m,
			ReadOneApi: m,
			ReadAllApi: m,
			FileApi:    m,
			UpdateApi:  m,
			DeleteApi:  m,
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

func (m module) Create(params *[]map[string]string) (*[]map[string]string, error) {
	fmt.Println("parámetros Create recibidos:", params)
	return &[]map[string]string{
		{"id": "2"},
	}, nil
}

func (m module) Update(params *[]map[string]string) error {
	fmt.Println("parámetros Update recibidos:", params)
	return nil
}

func (m module) Delete(params *[]map[string]string) error {
	fmt.Println("parámetros Delete recibidos:", params)
	return nil
}
func (m module) ReadOne(params *map[string]string) (*map[string]string, error) {
	fmt.Println("parámetros leer uno recibidos:", params)
	return &map[string]string{
		"name": "manzana",
	}, nil
}
func (m module) ReadAll(params *map[string]string) (*[]map[string]string, error) {
	fmt.Println("parámetros leer todo recibidos:", params)
	return &[]map[string]string{
		{"name": "manzana"},
		{"name": "peras"},
	}, nil
}

func (m module) FilePath(params *map[string]string) (string, error) {
	fmt.Println("parámetros leer archivo recibidos:", params)
	return "./README.md", nil
}
