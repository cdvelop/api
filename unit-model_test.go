package api_test

import (
	"github.com/cdvelop/filehandler"
	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
	"github.com/cdvelop/unixid"
)

var (
	product = ModuleProduct()
)

type module struct{}

func ModuleProduct() *model.Module {
	m := module{}

	newModule := &model.Module{
		ModuleName: "product",
		Title:      "Producto",
		Objects:    nil,
	}

	newObject := model.Object{
		ObjectName: "product",
		Table:      "product",

		Fields: []model.Field{
			{Name: "id_product", Legend: "Id", Input: unixid.InputPK()},
			{Name: "name", Legend: "Nombre", Input: input.Text()},
		},
		Module: newModule,
		BackHandler: model.BackendHandler{
			CreateApi: m,
			ReadApi:   m,
			UpdateApi: m,
			DeleteApi: m,
		},
	}

	newModule.Objects = append(newModule.Objects, &newObject)

	// fmt.Printf("TAMAÑO OBJETOS MODULO: [%v]\n", len(newModule.Objects))

	return newModule
}

func (m module) Create(u *model.User, params ...map[string]string) error {
	// fmt.Println("parámetros Create recibidos:", params)

	params[0]["id_product"] = "4"

	return nil
}

func (m module) Read(u *model.User, params ...map[string]string) ([]map[string]string, error) {
	// fmt.Println("parámetros leer todo recibidos:", params)
	for _, v := range params {
		if v["id_product"] == "1" {
			return []map[string]string{
				{"name": "manzana"},
				{"name": "peras"},
			}, nil
		}
	}
	return nil, model.Error("nada encontrado")
}

func (m module) Update(u *model.User, params ...map[string]string) error {
	// fmt.Println("parámetros Update recibidos:", params)
	return nil
}

func (m module) Delete(u *model.User, params ...map[string]string) ([]map[string]string, error) {
	// fmt.Println("parámetros Delete recibidos:", params)
	return []map[string]string{{"id_product": "1", "name": "pera"}}, nil
}

func (m module) FilePath(params map[string]string) (file_path, file_area string, err error) {
	// fmt.Println("parámetros leer archivo recibidos:", params)
	return "./README.md", "s", nil
}

func (m module) GetFileSettings() *filehandler.FileSetting {
	return &filehandler.FileSetting{
		MaximumFilesAllowed: 6,
		MaximumKbSize:       50,
	}
}

func (module) FileUpload(object_name, area_file string, file_request ...any) (out []map[string]string, err error) {

	return []map[string]string{{"id_product": "200"}}, nil
}
