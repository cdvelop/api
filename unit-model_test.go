package api_test

import (
	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
	"github.com/cdvelop/unixid"
)

var (
	product = ModuleProduct()
)

type module struct {
}

func ModuleProduct() *model.Module {
	m := module{}

	newModule := &model.Module{
		ModuleName: "product",
		Title:      "Producto",
		Objects:    nil,
	}

	newObject := model.Object{
		Name:  "product",
		Table: "product",
		BackendHandler: model.BackendHandler{
			CreateApi:   m,
			ReadApi:     m,
			UpdateApi:   m,
			DeleteApi:   m,
			FileHandler: m,
		},

		Fields: []model.Field{
			{Name: "id_product", Legend: "Id", Input: unixid.InputPK()},
			{Name: "name", Legend: "Nombre", Input: input.Text()},
		},
		Module: newModule,
	}

	newModule.Objects = append(newModule.Objects, &newObject)

	// fmt.Printf("TAMAÑO OBJETOS MODULO: [%v]\n", len(newModule.Objects))

	return newModule
}

func (m module) Create(u *model.User, params ...map[string]string) error {
	// fmt.Println("parámetros Create recibidos:", params)

	params[0]["id"] = "2"

	return nil
}

func (m module) Read(u *model.User, params ...map[string]string) ([]map[string]string, error) {
	// fmt.Println("parámetros leer todo recibidos:", params)
	return []map[string]string{
		{"name": "manzana"},
		{"name": "peras"},
	}, nil
}

func (m module) Update(u *model.User, params ...map[string]string) error {
	// fmt.Println("parámetros Update recibidos:", params)
	return nil
}

func (m module) Delete(u *model.User, params ...map[string]string) ([]map[string]string, error) {
	// fmt.Println("parámetros Delete recibidos:", params)
	return []map[string]string{}, nil
}

func (m module) GetFilePathByID(params map[string]string) (file_path, file_area string, err error) {
	// fmt.Println("parámetros leer archivo recibidos:", params)
	return "./README.md", "s", nil
}

func (m module) RegisterNewFile(new *model.FileNewToStore, form_data map[string]string) (map[string]string, error) {
	// fmt.Println("Upload File:", r)

	return map[string]string{"file": "./README.md"}, nil
}

func (m module) ConfigFile() *model.FileConfig {
	return &model.FileConfig{
		MaximumFilesAllowed: 6,
		InputNameWithFiles:  "files",
		MaximumFileSize:     0,
		MaximumKbSize:       50,
		AllowedExtensions:   ".jpg, .png, .jpeg",
	}
}

func (module) GenerateFileNameOnDisk() string {
	return "123"
}

func (module) UploadFolderPath(form_data map[string]string) string {
	return "/"
}
