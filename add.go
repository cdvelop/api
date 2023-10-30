package api

import (
	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/model"
	out "github.com/cdvelop/output"
)

// options:
// static cache duración de archivos estáticos en el navegador
// ej: "cache:year" (un año), week (semana), month (mes) default day. NOTE: modo dev = no-cache
// ej: authAdapter = GetUser(r *http.Request) *model.User. nil case default dev user
func Add(modules []*model.Module, a authAdapter, options ...string) *config {

	out.SetupLogsToFile("app")

	c := config{
		Cut:            nil,
		bootHandlers:   []*model.Object{},
		createHandlers: []*model.Object{},
		readHandlers:   []*model.Object{},
		updateHandlers: []*model.Object{},
		deleteHandlers: []*model.Object{},
		fileHandlers:   []*model.Object{},

		static_cache: "public, max-age=86400", // Configurar el encabezado de caché para 1 día
		auth:         a,
	}

	var registered = make(map[string]struct{})

	var module_objects []*model.Object

	for _, m := range modules {

		// fmt.Println("**TAMAÑO OBJETOS:", m.ModuleName, len(m.Objects))

		for _, o := range m.Objects {
			if o != nil {

				if _, exist := registered[o.Name]; !exist {

					if o.BootResponse != nil {
						c.bootHandlers = append(c.bootHandlers, o)
					}

					if o.CreateApi != nil {
						// fmt.Println("createHandlers ", o.Name)
						c.createHandlers = append(c.createHandlers, o)
					}

					if o.ReadApi != nil {
						// fmt.Println("readHandlers ", o.Name)
						c.readHandlers = append(c.readHandlers, o)
					}

					if o.UpdateApi != nil {
						// fmt.Println("updateHandlers ", o.Name)
						c.updateHandlers = append(c.updateHandlers, o)
					}

					if o.DeleteApi != nil {
						// fmt.Println("deleteHandlers ", o.Name)
						c.deleteHandlers = append(c.deleteHandlers, o)
					}

					if o.FileHandler != nil {
						// fmt.Println("fileHandlers ", o.Name)
						c.fileHandlers = append(c.fileHandlers, o)

						// agregamos file api al sistema solo si es nil
						if c.fileApi == nil {
							c.fileApi = o.FileHandler
						}

					}

					registered[o.Name] = struct{}{}

					module_objects = append(module_objects, o)

				}
			}
		}
	}

	c.Cut = cutkey.Add(module_objects...)

	c.processOptions(options...)

	if c.auth == nil {
		c.auth = auth{}
		c.developer_mode = true
	}

	if !c.developer_mode {
		out.PrintOK("*** Api en Modo Producción ***\n")
	} else {
		out.PrintWarning("*** Api en Modo Desarrollo ***\n")
		c.static_cache = "no-cache"
	}

	return &c
}
