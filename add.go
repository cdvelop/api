package api

import (
	"github.com/cdvelop/model"
	out "github.com/cdvelop/output"
)

// options:
// static cache duración de archivos estáticos en el navegador
// ej: "cache:year" (un año), week (semana), month (mes) default day. NOTE: modo dev = no-cache
// ej: authAdapter = GetUser(r *http.Request) *model.User. nil case default dev user
func Add(h *model.Handlers, options ...string) (*config, error) {

	err := h.CheckInterfaces("api config", config{})
	if err != nil {
		return nil, err
	}

	c := &config{
		LoginUser:      h,
		ObjectsHandler: h,
		DataConverter:  h,
		FileApi:        h,
		Logger:         h,

		bootHandlers: []*model.Object{},

		static_cache: "public, max-age=86400", // Configurar el encabezado de caché para 1 día
	}

	// fmt.Println("**TAMAÑO OBJETOS:", m.ModuleName, len(m.Objects))
	for _, o := range h.GetObjects() {

		if o.BootResponse != nil {
			c.bootHandlers = append(c.bootHandlers, o)
		}

	}

	c.processOptions(options...)

	if c.LoginUser == nil {
		c.developer_mode = true
	}

	if !c.developer_mode {
		out.PrintOK("*** Api en Modo Producción ***\n")
	} else {
		out.PrintWarning("*** Api en Modo Desarrollo ***\n")
		c.static_cache = "no-cache"
	}

	return c, nil
}
