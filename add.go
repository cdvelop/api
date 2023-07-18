package api

import (
	"fmt"
	"strings"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/gotools"
	"github.com/cdvelop/model"
)

// options:
// dev (modo desarrollador)
// static cache duración de archivos estáticos en el navegador
// ej: "cache:year" (un año) default day. modo dev = no-cache
func Add(modules []*model.Module, options ...string) *config {

	c := config{
		createHandlers:   []*model.Object{},
		pathFileHandlers: []*model.Object{},
		readOneHandlers:  []*model.Object{},
		readAllHandlers:  []*model.Object{},
		updateHandlers:   []*model.Object{},
		deleteHandlers:   []*model.Object{},
		static_cache:     "public, max-age=86400", // Configurar el encabezado de caché para 1 día
	}

	var registered = make(map[string]struct{})

	var module_objects []*model.Object

	for _, m := range modules {

		for _, o := range m.Objects {
			if o != nil {

				if _, exist := registered[o.Name]; !exist {

					if o.CreateApi != nil {
						// fmt.Println("createHandlers ", o.Name)
						c.createHandlers = append(c.createHandlers, o)
					}

					if o.ReadOneApi != nil {
						fmt.Println("readOneHandlers ", o.Name)
						c.readOneHandlers = append(c.readOneHandlers, o)
					}

					if o.ReadAllApi != nil {
						fmt.Println("readAllHandlers ", o.Name)
						c.readAllHandlers = append(c.readAllHandlers, o)
					}

					if o.FileApi != nil {
						fmt.Println("pathFileHandlers ", o.Name)
						c.pathFileHandlers = append(c.pathFileHandlers, o)
					}

					if o.UpdateApi != nil {
						// fmt.Println("updateHandlers ", o.Name)
						c.updateHandlers = append(c.updateHandlers, o)
					}

					if o.DeleteApi != nil {
						// fmt.Println("deleteHandlers ", o.Name)
						c.deleteHandlers = append(c.deleteHandlers, o)
					}

					registered[o.Name] = struct{}{}

					if o.Module.Name == m.Name {
						module_objects = append(module_objects, o)
					}

				}
			}
		}
	}

	c.Cut = cutkey.Add(module_objects...)

	for _, option := range options {

		switch {

		case strings.Contains(option, "cache:"):
			var cache_option string
			err := gotools.ExtractTwoPointArgument(option, &cache_option)
			if err == nil {

				var seconds string

				switch cache_option {

				case "week":
					seconds = "604800" // 7 días

				case "month":
					seconds = "2592000" // 1 mes

				case "year":
					seconds = "31536000" // 1 año
				}

				c.static_cache = "public, max-age=" + seconds
			}

		case option == "dev":
			c.developer_mode = true

		}
	}

	if c.developer_mode {
		c.static_cache = "no-cache"
	}

	return &c
}
