package api

import (
	"net/http"

	"github.com/cdvelop/model"
)

func (c config) read(p *petition) {

	// fmt.Printf("Estás en el Manejador de lectura de data de %s\n", o.Name)

	params, err := paramsCheckIn(p, false, false)
	if err != nil {
		c.error(p, err)
		return
	}

	data, err := p.o.Read(p.u, params)
	if err != nil {
		c.error(p, err)
		return
	}

	// fmt.Printf("Manejador de lectura RESPUESTA %s\n", data)

	c.success(p, "read", "ok", data...)
}

func (c config) readFile(p *petition) {
	// retorna objeto estático ej imagen.jpg
	params := make(map[string]string)

	gerUrlParams(p.r, params)

	// fmt.Printf("Estás en la página de lectura archivo %s\n", params)

	file_path, file_area, err := c.FilePath(params)
	if err != nil {
		errorHttp(p, err, http.StatusBadRequest)
		return
	}

	// fmt.Println("AREA USUARIO", u.Area)

	if file_area != p.u.Area {
		errorHttp(p, model.Error("no autorizado para leer archivo"), http.StatusUnauthorized)
		return
	}

	http.ServeFile(p.w, p.r, file_path)
}
