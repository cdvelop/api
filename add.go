package api

import (
	"github.com/cdvelop/model"
	out "github.com/cdvelop/output"
)

// options:
// static cache duración de archivos estáticos en el navegador
// ej: "cache:year" (un año), week (semana), month (mes) default day. NOTE: modo dev = no-cache
// ej: authAdapter = GetUser(r *http.Request) *model.User. nil case default dev user
func Add(h *model.Handlers, options ...string) (c *config, err string) {

	c = &config{
		AuthAdapter:    h,
		ObjectsHandler: h,
		ModuleHandler:  h,
		DataConverter:  h,
		FileApi:        h,
		Logger:         h,

		static_cache: "public, max-age=86400", // Configurar el encabezado de caché para 1 día
	}
	h.BackendBootDataUser = c

	c.production_mode = h.ProductionMode

	// fmt.Println("**TAMAÑO OBJETOS:", m.ModuleName, len(m.Objects))

	c.processOptions(options...)

	if c.production_mode {
		out.PrintOK("*** Api en Modo Producción ***\n")
	} else {
		out.PrintWarning("*** Api en Modo Desarrollo ***\n")
		c.static_cache = "no-cache"
	}

	err = h.CheckInterfaces("api config", *c)
	if err != "" {
		return nil, err
	}

	return c, ""
}
