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
	}

	newObject := &model.Object{
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

	newModule.AddObjectsToModule(newObject)

	// fmt.Printf("TAMAÑO OBJETOS MODULO: [%v]\n", len(newModule.Objects))

	return newModule
}

func (m module) Create(u *model.User, params ...map[string]string) (err string) {
	// fmt.Println("parámetros Create recibidos:", params)

	params[0]["id_product"] = "4"

	return ""
}

func (m module) Read(u *model.User, params ...map[string]string) (out []map[string]string, err string) {
	// fmt.Println("parámetros leer todo recibidos:", params)
	for _, v := range params {
		if v["id_product"] == "1" {
			return []map[string]string{
				{"name": "manzana"},
				{"name": "peras"},
			}, ""
		}
	}
	return nil, "nada encontrado"
}

func (m module) Update(u *model.User, params ...map[string]string) (err string) {
	// fmt.Println("parámetros Update recibidos:", params)
	return ""
}

func (m module) Delete(u *model.User, params ...map[string]string) (err string) {

	for i := 0; i < len(params); i++ {
		if params[i]["id_product"] == "1" {
			params[i]["name"] = "pera"
			return
		}
	}

	// fmt.Println("parámetros Delete recibidos:", params)
	return "error nada encontrado para eliminar"
}

func (m module) FilePath(params map[string]string) (file_path, file_area, err string) {
	// fmt.Println("parámetros leer archivo recibidos:", params)
	return "./README.md", "s", ""
}

func (m module) GetFileSettings() *filehandler.FileSetting {
	return &filehandler.FileSetting{
		MaximumFilesAllowed: 6,
		MaximumKbSize:       50,
	}
}

func (module) FileUpload(object_name, area_file string, file_request ...any) (out []map[string]string, err string) {

	return []map[string]string{{"id_product": "200"}}, ""
}
