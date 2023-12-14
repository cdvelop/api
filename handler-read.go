package api

import (
	"net/http"
)

func (c config) read(p *petition) {
	const this = "api read error "
	params, err := c.decodeStringMapData(p)
	if err != "" {
		c.error(p, this+err)
		return
	}

	err = p.o.ValidateData(false, false, params...)
	if err != "" {
		c.error(p, this+err)
		return
	}

	// fmt.Printf("params read(p *petition) %s\n", params)

	data, err := p.o.BackHandler.Read(p.u, params...)
	if err != "" {
		c.error(p, this+err)
		return
	}

	c.success(p, "read", "ok", data...)
}

func (c config) readFile(p *petition) {
	const this = "api readFile error "
	// retorna objeto estático ej imagen.jpg
	params := make(map[string]string)

	gerUrlParams(p.r, params)

	// fmt.Printf("Estás en la página de lectura archivo %s\n", params)

	file_path, file_area, err := c.FilePath(params)
	if err != "" {
		c.error(p, this+err, http.StatusBadRequest)
		return
	}

	// fmt.Println("AREA USUARIO", u.Area)

	if file_area != p.u.Area {
		c.error(p, this+"no autorizado para leer archivo", http.StatusUnauthorized)
		return
	}

	http.ServeFile(p.w, p.r, file_path)
}
