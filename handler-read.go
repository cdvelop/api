package api

import (
	"net/http"

	"github.com/cdvelop/model"
)

func (c config) read(p *petition) {

	params, err := c.decodeStringMapData(p)
	if err != nil {
		c.error(p, err)
		return
	}

	err = p.o.ValidateData(false, true, params...)
	if err != nil {
		c.error(p, err)
		return
	}

	// fmt.Printf("params read(p *petition) %s\n", params)

	data, err := p.o.BackHandler.Read(p.u, params...)
	if err != nil {
		c.error(p, err)
		return
	}

	c.success(p, "read", "ok", data...)
}

func (c config) readFile(p *petition) {
	// retorna objeto estático ej imagen.jpg
	params := make(map[string]string)

	gerUrlParams(p.r, params)

	// fmt.Printf("Estás en la página de lectura archivo %s\n", params)

	file_path, file_area, err := c.FilePath(params)
	if err != nil {
		c.error(p, err, http.StatusBadRequest)
		return
	}

	// fmt.Println("AREA USUARIO", u.Area)

	if file_area != p.u.Area {
		c.error(p, model.Error("no autorizado para leer archivo"), http.StatusUnauthorized)
		return
	}

	http.ServeFile(p.w, p.r, file_path)
}
